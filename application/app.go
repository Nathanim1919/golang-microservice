package application

import (
	"fmt"
	"context"
	"net/http"
	"time"
	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
}


func New() *App{
  app :=  &App{
      router: loadRoutes(),
      	rdb:    redis.NewClient(&redis.Options{}),

  }
  return app
}


func (a *App) Start(ctx context.Context) error{
	server := &http.Server {
		Addr: ":3000",
		Handler: a.router,
	}


	err := a.rdb.Ping(ctx).Err()

	if err != nil {
		return fmt.Errorf("error while connecting to redis: %w", err)
	}


	defer func(){
		if err := a.rdb.Close(); err != nil {
			fmt.Println("Failed to close redis", err)
		}
	}()

	fmt.Println("Starting server..")
    cha := make(chan error, 1)
	go func (){
		err = server.ListenAndServe()

		if err != nil {
			cha <- fmt.Errorf("failed to start server: %w", err)
		}

		close(cha)
	}()

select {
	case err = <- cha:
		return err
	case<-ctx.Done():
	    shutdownCtx,cancel := context.WithTimeout(context.Background(), 5 * time.Second)
					defer cancel()
		return server.Shutdown(shutdownCtx)
}

	return nil
}
