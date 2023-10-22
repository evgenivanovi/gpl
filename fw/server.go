package fw

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/evgenivanovi/gpl/std"
	slogx "github.com/evgenivanovi/gpl/stdx/log/slog"
	"github.com/gookit/goutil/strutil"
)

/* __________________________________________________ */

const ServerDefaultHost = "localhost"
const ServerDefaultPort = 80

/* __________________________________________________ */

type ServerOp func(*ServerSettings)

func (o ServerOp) Join(op ServerOp) ServerOp {
	return func(opts *ServerSettings) {
		o(opts)
		op(opts)
	}
}

func (o ServerOp) And(ops ...ServerOp) ServerOp {
	return func(opts *ServerSettings) {
		o(opts)
		for _, fn := range ops {
			fn(opts)
		}
	}
}

/* __________________________________________________ */

type ServerSettings struct {
	host string
	port int
}

func (op ServerSettings) Address() string {
	result := strings.Builder{}
	result.WriteString(op.host)
	result.WriteString(std.Colon)
	result.WriteString(strconv.Itoa(op.port))
	return result.String()
}

func WithHost(host string) ServerOp {
	return func(opts *ServerSettings) {
		opts.host = host
	}
}

func WithHostFn(fn func() string) ServerOp {
	return WithHost(fn())
}

func WithPort(port int) ServerOp {
	return func(opts *ServerSettings) {
		opts.port = port
	}
}

func WithPortFn(fn func() int) ServerOp {
	return WithPort(fn())
}

func WithAddress(address string) ServerOp {
	array := strings.Split(address, std.Colon)
	host := array[0]
	port, _ := strutil.ToInt(array[1])
	return WithHost(host).Join(WithPort(port))
}

func WithAddressFn(fn func() string) ServerOp {
	return WithAddress(fn())
}

func NewServerOpts(opts ...ServerOp) *ServerSettings {
	opt := defaultServerOpts()
	for _, fn := range opts {
		fn(&opt)
	}
	return &opt
}

func defaultServerOpts() ServerSettings {
	return ServerSettings{
		host: ServerDefaultHost,
		port: ServerDefaultPort,
	}
}

/* __________________________________________________ */

// ConfigureServer
// When we receive a SIGINT or SIGTERM signal,
// we instruct our server to stop accepting any new HTTP requests,
// and give any in-flight requests a ‘grace period’ to complete before the application is terminated.
func ConfigureServer(
	application *Application,
	handler http.Handler,
) error {

	server := &http.Server{
		Addr:    application.ServerOpts.Address(),
		Handler: handler,
	}

	// Create a shutdownError channel.
	// We will use this to receive any errors returned by the graceful Shutdown() function.
	shutdownErrorChannel := make(chan error)

	// Start a background goroutine.
	go func() {

		// Create a quitChannel channel which carries os.Signal values.
		quitChannel := make(chan os.Signal, 1)

		// Use signal.Notify() to listen for incoming SIGINT and SIGTERM signals and
		// relay them to the quitChannel channel.
		// Any other signals will not be caught by signal.Notify() and
		// will retain their default behavior.
		signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)

		// Read the signal from the quitChannel channel.
		// This code will block until a signal is received.
		sig := <-quitChannel

		// Log a message to say that the signal has been caught.
		// Notice that we also call the String() method on the signal to get the signal name and
		// include it in the log entry properties.
		slogx.Log().Debug(
			"Caught OS signal.",
			slog.String("signal", sig.String()),
		)

		slogx.Log().Debug(
			"Shutting down server",
			slog.String("signal", sig.String()),
		)

		// Create a context with a timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Call Shutdown() on our server, passing in the context we just made.
		// Shutdown() will return nil if the graceful executor was successful, or an
		// error (which may happen because of a problem closing the listeners, or
		// because the executor didn't complete before the context deadline is hit).
		// We relay this return value to the shutdownError channel.
		err := server.Shutdown(ctx)
		if err != nil {
			shutdownErrorChannel <- err
		}

		slogx.Log().Debug("Completing tasks")
		application.Close()

		// Then we return nil on the shutdownError channel,
		// to indicate that the executor was completed without any issues.
		shutdownErrorChannel <- nil

	}()

	application.Start()

	// Calling Shutdown() on our server will cause ListenAndServe() to immediately return a http.ErrServerClosed error.
	// So if we see this error, it is actually a good thing and an indication that the graceful executor has started.
	// So we check specifically for this, only returning the error if it is NOT http.ErrServerClosed.
	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	// Otherwise, we wait to receive the return value from Shutdown() on the shutdownError channel.
	// If return value is an error, we know that there was a problem with the graceful executor, and we return the error.
	err = <-shutdownErrorChannel
	if err != nil {
		return err
	}

	// At this point, we know that the graceful executor completed successfully.
	slogx.Log().Debug("Stopped server")
	return nil

}
