package api

import "net/http"

type APIServer struct {
	http.Server
}

func New() APIServer {
	server := &http.Server{
		Addr: ":5000",
	}
	return APIServer{
		Server: *server,
	}
}
