package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("#", err)
	}
}
func main() {
	portString := os.Getenv("PORT")
	if portString == "" {
		panic("error")
	}
	fmt.Println(portString)
	createServer(portString)
}
func createServer(port string) {

	router := chi.NewRouter()

	router.Use(
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"*"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}))

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	fmt.Printf("Server Runnig On Port : " + port)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
