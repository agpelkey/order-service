package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func (app *application) server() error {
    
    // declare server settings
    srv := http.Server{
        Addr: fmt.Sprintf(":%d", app.config.port),
        Handler: app.routes(),
        IdleTimeout: time.Minute,
        ReadTimeout: 10*time.Second,
        WriteTimeout: 30*time.Second,
    }

    // channel to later receive any errors returned by the graceful shutdown
    shutdownError := make(chan error)

    // start a background go routine to run for the lifetime of the application
    go func() {
    
        // create a channel to catch all signals
        quit := make(chan os.Signal, 1)

        // use signal.Notify to listen for incoming signals and send them to the quit channel
        signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

        // Read out the signals from the quit channel
        s := <-quit

        // log out the signal message from the channel
        log.Println(s.String())

        // create a context with 5 second timeout
        ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
        defer cancel()

        // Call shutdown() on our server, and relay the return value from the function to the
        // shutdownErr channel
        shutdownError <- srv.Shutdown(ctx)
    }()

    fmt.Println("starting server")

    // Start the server. Note that if Shutdown() is executed successfully then ListenAndServe *will*
    // return an http.ErrServerClosed. So, we can check for any error that is *not* what we expect. 
    err := srv.ListenAndServe()
    if !errors.Is(err, http.ErrServerClosed) {
        return err
    }

    // Otherwise, we wait for the return value from our call to Shutdown(). 
    // If return value is an error, then we now know our graceful shutdown 
    // encountered an issue
    err = <- shutdownError
    if err != nil {
        return err
    }

    log.Println("stopping server")

    return nil
}
