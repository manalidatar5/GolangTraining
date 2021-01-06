package main


import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/you/hello/api"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	api_router := r.PathPrefix("/api").Subrouter()
	api_router.HandleFunc("/login", api.CreateTokenEndpoint).Methods("POST")
	api_router.HandleFunc("/create", api.ValidateMiddleware(api.UserCreate)).Methods("POST")
	api_router.HandleFunc("/delete", api.ValidateMiddleware(api.Userdelete)).Methods("DELETE")
	api_router.HandleFunc("/update", api.ValidateMiddleware(api.Userupdate)).Methods("PUT")
	api_router.HandleFunc("/users", api.ValidateMiddleware(api.Getuser)).Methods("GET")
	fmt.Printf("Starting server at port 8091\n")
	http.ListenAndServe("localhost:8091", r)

}


