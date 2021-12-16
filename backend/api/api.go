package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	sessionManager *SessionManager

	http.Server
}

func New() *APIServer {
	apiServer := &APIServer{
		Server: http.Server{
			Addr: ":5000",
		},
	}

	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, req)
		})
	})
	router.HandleFunc("/login", apiServer.LoginHandler).Methods(http.MethodGet)

	apiServer.Handler = router
	return apiServer
}
