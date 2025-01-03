package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/maxwell-ihiaso/theservicehood_core_service/pkg/utils/keycloak"
)

func init() {
	cwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}
	envPath := filepath.Join(cwd, "configs/.env")
	err = godotenv.Load(envPath)

	if err != nil {
		panic(err)
	}

	fmt.Println("Loaded env")
}

func main() {

	s := newServer(os.Getenv("HOST"), os.Getenv("PORT"), keycloak.NewKeycloak())
	s.listen()

	fmt.Printf("server running on %v:%v", os.Getenv("HOST"), os.Getenv("PORT"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

}
