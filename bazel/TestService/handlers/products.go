package handlers

import (
	"TestService/data"
	"context"
	"fmt"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func (p *Products) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	//writer.Write([]byte("response"))
	if request.Method == http.MethodGet {
		p.GetProducts(writer, request)
		return
	}
	// update handler
	if request.Method == http.MethodPut {
		writer.WriteHeader(http.StatusAccepted)
	}
	// Handling Post Request

	//if request.Method == http.MethodPost {
	//	p.AddProducts(writer, request)
	//}
	writer.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) GetProducts(res http.ResponseWriter, req *http.Request) {
	productList := data.GetProducts()
	err := productList.ToJSON(res)
	if err != nil {
		http.Error(res, "Unable to get the List", http.StatusInternalServerError)
		return
	}
}

func (p *Products) AddProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handling POST Request")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Unable to process", http.StatusBadRequest)
	}
	prodd := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Println(prod)
	p.l.Println(prodd)
	data.AddProduct(prod)
}

type KeyProduct struct{}

func (p *Products) ProductMiddleware(next http.Handler) http.Handler {
	fmt.Println("jjjjjj")
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		prod := &data.Product{}
		err := prod.FromJSON(request.Body)
		if err != nil {
			p.l.Println("Error in converting product", err)
			http.Error(writer, "Error in reading product", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(request.Context(), KeyProduct{}, prod)
		req := request.WithContext(ctx)
		next.ServeHTTP(writer, req)
	})
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}
