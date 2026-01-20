package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"project-app-bioskop-golang-fathoni/internal/wire"
	
	"syscall"
)

func ApiServer(app *wire.App) {
	fmt.Println("Server running on port 8080")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: app.Route,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal("can't run service")
		}
	}()

	// gracefully shutdown ------------------------------------------------------------------------
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	close(app.Stop)
	app.WG.Wait()

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// if err := srv.Shutdown(ctx); err != nil {
	// 	log.Println("can't shutdown service")
	// }

	log.Println("server shutdown cleanly")
}
