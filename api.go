package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
  w.WriteHeader(status)
  w.Header().Set("Content-Type", "application/json")
  return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
  Error string
}

// Decorates our handler functions to make the signature as required by the router 
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    if err := f(w, r); err != nil {
      // handle the error here
      WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
    }
  }
}

type APIServer struct {
	listenAddr string
}

func newApiServer(listenAddr string) *APIServer {
	return &APIServer{listenAddr}
}

func (s *APIServer) Run() {
  router := mux.NewRouter()
  log.Println("JSON API server is running on port: ", s.listenAddr)
  http.ListenAndServe(s.listenAddr, router)

  // we wrap our handle function to process the error
  // the second argument shall be Route.HandlerFunc which returns nothing but we return error
  router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))

}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
