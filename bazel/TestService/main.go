package main

import (
	handler "TestService/handlers"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	//gh := handler.NewGoodBye(l)
	//hh := handler.NewHello(l)
	productHandler := handler.NewProducts(l)
	//serveMux := http.NewServeMux()
	serveMux := mux.NewRouter()
	//serveMux.Handle("/", hh)
	//serveMux.Handle("/goodBye", gh)
	//serveMux.Handle("/products", productHandler)

	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", productHandler.GetProducts)
	//getRouter.Use(productHandler.ProductMiddleware)

	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", productHandler.AddProducts)
	postRouter.Use(productHandler.ProductMiddleware)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received this := ", sig)
	tc, _ := context.WithTimeout(context.Background(), 25*time.Second)
	server.Shutdown(tc)
}
