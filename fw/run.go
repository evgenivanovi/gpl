package fw

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	slogx "github.com/evgenivanovi/gpl/stdx/log/slog"
)

// RunServer
// When we receive a SIGINT or SIGTERM signal,
// we instruct our server to stop accepting any new HTTP(s) requests,
// and give any in-flight requests a ‘grace period’ to complete before the application is terminated.
func RunServer(app *Application, mux http.Handler) error {

	var httpServer *http.Server = nil
	if app.Settings.httpEnabled {
		httpServer = &http.Server{
			Addr:    app.Settings.HttpAddress(),
			Handler: mux,
		}
	}

	var httpsServer *http.Server = nil
	if app.Settings.httpsEnabled {
		httpsServer = &http.Server{
			Addr:      app.Settings.HttpsAddress(),
			Handler:   mux,
			TLSConfig: app.Settings.tls.config,
		}
	}

	// Create a startError channel.
	// We will use this to receive any errors returned by the Startup() function.
	startErrorChannel := make(chan error)
	// Create a shutdownError channel.
	// We will use this to receive any errors returned by the Shutdown() function.
	shutdownErrorChannel := make(chan error)

	go shutdown(app, shutdownErrorChannel, httpServer, httpsServer)
	go startup(app, startErrorChannel, httpServer, httpsServer)

	// Calling Shutdown() on our server will cause ListenAndServe() or ListenAndServerTLS()
	// to immediately return a http.ErrServerClosed error.
	// So if we see this error, it is actually a good thing and an indication that the graceful executor has started.
	// So we check specifically for this, only returning the error if it is NOT http.ErrServerClosed.
	// Otherwise, we wait to receive the return value from Shutdown() on the shutdownError channel.
	// If return value is an error, we know that there was a problem with the graceful executor, and we return the error.
	select {
	case err := <-startErrorChannel:
		{
			if err != nil {
				slogx.Log().Debug(
					fmt.Sprintf("Could not start service due to (error: %s)", err),
				)
				return err
			}
			break
		}
	case err := <-shutdownErrorChannel:
		{
			if err != nil {
				slogx.Log().Debug(
					fmt.Sprintf("Could not shutdown service due to (error: %s)", err),
				)
				return err
			}
			break
		}
	}

	// At this point, we know that the graceful executor completed successfully.
	slogx.Log().Debug("Stopped server")
	return nil

}

/* __________________________________________________ */

func startup(
	app *Application,
	startErrorChannel chan error,
	httpServer *http.Server,
	httpsServer *http.Server,
) {

	app.Start()

	if app.Settings.httpEnabled && httpServer != nil {
		go startHttpServer(
			app,
			startErrorChannel,
			httpServer,
		)
	}

	if app.Settings.httpsEnabled && httpsServer != nil {
		go startHttpsServer(
			app,
			startErrorChannel,
			httpsServer,
		)
	}

}

func startHttpServer(
	app *Application,
	startErrorChannel chan error,
	server *http.Server,
) {

	if !app.Settings.httpEnabled {
		return
	}

	slogx.Log().Debug(
		"HTTP Server has been run on address: '" + server.Addr + "'",
	)

	err := server.ListenAndServe()

	if !errors.Is(err, http.ErrServerClosed) {
		startErrorChannel <- err
	}

}

func startHttpsServer(
	app *Application,
	startErrorChannel chan error,
	server *http.Server,
) {

	if !app.Settings.httpsEnabled {
		return
	}

	slogx.Log().Debug(
		"HTTPs Server has been run on address: '" + server.Addr + "'",
	)

	err := server.ListenAndServeTLS(
		app.Settings.tls.cert,
		app.Settings.tls.key,
	)

	if !errors.Is(err, http.ErrServerClosed) {
		startErrorChannel <- err
	}

}

/* __________________________________________________ */

func shutdown(
	app *Application,
	shutdownErrorChannel chan error,
	httpServer *http.Server,
	httpsServer *http.Server,
) {

	servers := make([]*http.Server, 0)

	if app.Settings.httpEnabled && httpServer != nil {
		servers = append(servers, httpServer)
	}

	if app.Settings.httpsEnabled && httpsServer != nil {
		servers = append(servers, httpsServer)
	}

	shutdownServers(app, shutdownErrorChannel, servers...)

}

func shutdownServers(
	app *Application,
	shutdownChannel chan error,
	servers ...*http.Server,
) {

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
	for _, server := range servers {
		err := server.Shutdown(ctx)
		if err != nil {
			shutdownChannel <- err
		}
	}

	slogx.Log().Debug("Completing tasks")
	app.Close()

	// Then we return nil on the shutdownError channel,
	// to indicate that the executor was completed without any issues.
	shutdownChannel <- nil

}
