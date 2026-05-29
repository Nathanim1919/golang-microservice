package main

import (
	"fmt"
     "context"
     "os/signal"
     "os"
     "go-microservice/application"
)


func main(){
   app := application.New()

   ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
   defer cancel()

   err := app.Start(ctx)

   if err != nil {
     fmt.Printf("Failed to start app: %v\n", err)
   }
}
