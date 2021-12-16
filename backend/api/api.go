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

func New(dataManager datamanager.Datamanager) *APIServer {
	apiServer := &APIServer{
		dataManager: dataManager,
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
	router.HandleFunc("/login", apiServer.LoginHandler).Methods(http.MethodPost)
	router.HandleFunc("/user", apiServer.GetUserHandler).Methods(http.MethodGet)
	router.HandleFunc("/user", apiServer.GetUserHandler).Methods(http.MethodPatch)

	apiServer.Handler = router
	return apiServer
}
