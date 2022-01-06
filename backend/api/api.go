package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yanando/StudentLink/datamanager"
)

type APIServer struct {
	dataManager    datamanager.Datamanager
	sessionManager *SessionManager

	http.Server
}

func NewAPIServer(dataManager datamanager.Datamanager, sessionManager *SessionManager) *APIServer {
	apiServer := &APIServer{
		sessionManager: sessionManager,
		dataManager:    dataManager,
		Server: http.Server{
			Addr: ":8080",
		},
	}

	router := mux.NewRouter().PathPrefix("/api").Subrouter()
	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			if req.Method == http.MethodOptions {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "*")
				w.WriteHeader(200)
				return
			}

			next.ServeHTTP(w, req)
		})
	})

	router.HandleFunc("/login", apiServer.LoginHandler).Methods(http.MethodPost, http.MethodOptions)

	router.HandleFunc("/user", apiServer.GetUserHandler).Methods(http.MethodGet)
	router.HandleFunc("/user", apiServer.UpdateUserHandler).Methods(http.MethodPatch)

	router.HandleFunc("/messages", apiServer.AddMessageHandler).Methods(http.MethodPost)
	router.HandleFunc("/messages", apiServer.GetMessagesHandler).Methods(http.MethodGet)

	apiServer.Handler = router
	return apiServer
}
