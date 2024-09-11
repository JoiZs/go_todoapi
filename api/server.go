package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/JoiZs/go_todoapi/db"
	"github.com/JoiZs/go_todoapi/model"
	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type apiErr struct {
	Error string
}

type APIServer struct {
	serverAddr string
}

func NewAPI(host string) *APIServer {
	return &APIServer{
		serverAddr: host,
	}
}

func changeHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusOK, apiErr{"NoErr"})
		}
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	dbI := db.DbConnect()

	var todo model.Todo

	dbI.First(&todo)

	fmt.Print(todo)

	router.HandleFunc("/", changeHandler(s.getTodoList))

	log.Println("Server at> localhost", s.serverAddr)
	http.ListenAndServe(s.serverAddr, router)
}

func (s *APIServer) getTodoList(w http.ResponseWriter, r *http.Request) error {
	return nil
}
