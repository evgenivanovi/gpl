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

	grpcfw "github.com/evgenivanovi/gpl/server"
	slogx "github.com/evgenivanovi/gpl/stdx/log/slog"
)

// RunServer
// When we receive a SIGINT or SIGTERM signal,
// we instruct our server to stop accepting any new HTTP(s) requests,
// and give any in-flight requests a ‘grace period’ to complete before the application is terminated.
func RunServer(cfg *Configuration) error {

	var httpServer *http.Server = nil
	if cfg.App.Settings.HttpEnabled() {
		httpServer = &http.Server{
			Addr:    cfg.App.Settings.HttpAddress(),
			Handler: cfg.HTTPHandler,
		}
	}

	var httpsServer *http.Server = nil
	if cfg.App.Settings.HttpsEnabled() {
		httpsServer = &http.Server{
			Addr:      cfg.App.Settings.HttpsAddress(),
			Handler:   cfg.HTTPHandler,
			TLSConfig: cfg.App.Settings.tls.config,
		}
	}

	var grpcServer *grpcfw.GRPCServer = nil
	if cfg.App.Settings.GrpcEnabled() {
		grpcServer = grpcfw.NewGrpcServer(
			grpcfw.WithGrpcServerConfig(
				*grpcfw.NewGRPCServerConfig(cfg.App.Settings.grpcPort),
			),
			grpcfw.WithGrpcReflection(cfg.GRPCReflection),
			grpcfw.WithServices(cfg.GRPCServices...),
			grpcfw.WithUnaryInterceptors(cfg.GRPCUnaryMWs...),
			grpcfw.WithStreamInterceptors(cfg.GRPCStreamMWs...),
		)
	}

	// Create a startErrorCh channel.
	// We will use this to receive any errors returned by the Startup() function.
	startErrorCh := make(chan error)
	// Create a shutdownErrorCh channel.
	// We will use this to receive any errors returned by the Shutdown() function.
	shutdownErrorCh := make(chan error)

	go shutdown(cfg.App, shutdownErrorCh, httpServer, httpsServer, grpcServer)
	go startup(cfg.App, startErrorCh, httpServer, httpsServer, grpcServer)

	// Calling Shutdown() on our HTTP servers will cause ListenAndServe() or ListenAndServerTLS()
	// to immediately return a http.ErrServerClosed error.
	// So if we see this error, it is actually a good thing and an indication that the graceful executor has started.
	// So we check specifically for this, only returning the error if it is NOT http.ErrServerClosed.
	// Otherwise, we wait to receive the return value from Shutdown() on the shutdownError channel.
	// If return value is an error, we know that there was a problem with the graceful executor, and we return the error.
	select {
	case err := <-startErrorCh:
		{
			if err != nil {
				slogx.Log().Debug(
					fmt.Sprintf("Could not start service due to (error: %s)", err),
				)
				return err
			}
			break
		}
	case err := <-shutdownErrorCh:
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

func startup(
	app *Application,
	errs chan error,
	httpServer *http.Server,
	httpsServer *http.Server,
	grpcServer *grpcfw.GRPCServer,
) {

	app.Start()

	if app.Settings.httpEnabled && httpServer != nil {
		go startHttpServer(
			app,
			errs,
			httpServer,
		)
	}

	if app.Settings.httpsEnabled && httpsServer != nil {
		go startHttpsServer(
			app,
			errs,
			httpsServer,
		)
	}

	if app.Settings.grpcEnabled && grpcServer != nil {
		go startGrpcServer(
			app,
			errs,
			grpcServer,
		)
	}

}

func startHttpServer(
	app *Application,
	errs chan error,
	server *http.Server,
) {

	if !app.Settings.httpEnabled {
		return
	}

	slogx.Log().Debug(
		"HTTP server has been run on address: '" + server.Addr + "'",
	)

	err := server.ListenAndServe()

	// Calling Shutdown() on our HTTP servers will cause ListenAndServe() or ListenAndServerTLS()
	// to immediately return a http.ErrServerClosed error.
	// So if we see this error, it is actually a good thing and an indication that the graceful executor has started.
	// So we check specifically for this, only returning the error if it is NOT http.ErrServerClosed.
	if !errors.Is(err, http.ErrServerClosed) {
		errs <- err
		return
	}

}

func startHttpsServer(
	app *Application,
	errs chan error,
	server *http.Server,
) {

	if !app.Settings.httpsEnabled {
		return
	}

	slogx.Log().Debug(
		"HTTPs server has been run on address: '" + server.Addr + "'",
	)

	err := server.ListenAndServeTLS(
		app.Settings.tls.cert,
		app.Settings.tls.key,
	)

	// Calling Shutdown() on our HTTP servers will cause ListenAndServe() or ListenAndServerTLS()
	// to immediately return a http.ErrServerClosed error.
	// So if we see this error, it is actually a good thing and an indication that the graceful executor has started.
	// So we check specifically for this, only returning the error if it is NOT http.ErrServerClosed.
	if !errors.Is(err, http.ErrServerClosed) {
		errs <- err
		return
	}

}

func startGrpcServer(
	app *Application,
	errs chan error,
	server *grpcfw.GRPCServer,
) {

	if !app.Settings.grpcEnabled {
		return
	}

	slogx.Log().Debug(
		"gRPC server has been run on address: '" + app.Settings.GrpcAddress() + "'",
	)

	server.StartChannable(errs)

}

func shutdown(
	app *Application,
	errs chan error,
	httpServer *http.Server,
	httpsServer *http.Server,
	grpcServer *grpcfw.GRPCServer,
) {

	httpServers := make([]*http.Server, 0)

	if app.Settings.httpEnabled && httpServer != nil {
		httpServers = append(httpServers, httpServer)
	}

	if app.Settings.httpsEnabled && httpsServer != nil {
		httpServers = append(httpServers, httpsServer)
	}

	grpcServers := make([]*grpcfw.GRPCServer, 0)

	if app.Settings.grpcEnabled && grpcServer != nil {
		grpcServers = append(grpcServers, grpcServer)
	}

	doShutdown(app, errs, grpcServers, httpServers)

}

func doShutdown(
	app *Application,
	errs chan error,
	grpcServers []*grpcfw.GRPCServer,
	httpServers []*http.Server,
) {

	// Create a quit channel which carries os.Signal values.
	quit := make(chan os.Signal, 1)

	// Use signal.Notify() to listen for incoming syscall.SIGINT and syscall.SIGTERM signals and
	// relay them to the quitChannel channel.
	// Any other signals will not be caught by signal.Notify() and
	// will retain their default behavior.
	//
	// syscall.SIGKILL signals are not catchable
	// (and will always cause the application to terminate immediately),
	// and we’ll leave SIGQUIT with its default behavior
	// (as it’s handy if you want to execute a non-graceful shutdown via a keyboard shortcut).
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Read the signal from the quitChannel channel.
	// This code will block until a signal is received.
	sig := <-quit

	// Log a message to say that the signal has been caught.
	// Notice that we also call the String() method on the signal to get the signal name and
	// include it in the log entry properties.
	slogx.Log().Debug(
		"Shutting down server",
		slog.String("os.signal", sig.String()),
	)

	// Create a context with a timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Call Shutdown() on our server, passing in the context we just made.
	// Shutdown() will return nil if the graceful executor was successful, or an
	// error (which may happen because of a problem closing the listeners, or
	// because the executor didn't complete before the context deadline is hit).
	// We relay this return value to the shutdownError channel.
	for _, server := range httpServers {
		err := server.Shutdown(ctx)
		if err != nil {
			errs <- err
		}
	}

	for _, server := range grpcServers {
		server.Stop()
	}

	slogx.Log().Debug("Completing tasks")
	app.Close()

	// Then we return nil on the shutdownError channel,
	// to indicate that the executor was completed without any issues.
	errs <- nil

}
