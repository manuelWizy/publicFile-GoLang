package main

import (
	"log"
	"net/http"
)

type APIServer struct{
	addr string
}

func NewApiServer(addr string) *APIServer{
	return &APIServer{
		addr: addr,
	}
}
func (s *APIServer) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("GET /files/{fileId}", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		fileId := r.PathValue("fileId")
		namespace := q.Get("namespace")
		w.Write([]byte("USER ID: " + fileId + " NAMESPACE: " + namespace))
	})

	server := http.Server {
		Addr: s.addr,
		Handler: RequestLoggerMiddleware(router),
	}

	return server.ListenAndServe()
}

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method %s, path %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r )
	}
}