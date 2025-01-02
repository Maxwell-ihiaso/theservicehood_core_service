package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/maxwell-ihiaso/theservicehood_core_service/pkg/utils/keycloak"
)

func init() {
	err := godotenv.Load("../../configs/.env")

	if err != nil {
		panic(err)
	}

	fmt.Println("Loaded env")
}

func main() {

	s := newServer(os.Getenv("HOST"), os.Getenv("PORT"), keycloak.NewKeycloak())
	s.listen()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

}
