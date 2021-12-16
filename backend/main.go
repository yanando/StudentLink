package main

import (
	"log"

	"github.com/yanando/StudentLink/api"
	"github.com/yanando/StudentLink/datamanager"
)

func main() {
	sessionManager := api.NewSessionManager()
	log.Println("Initializing API")
	apiServer := api.NewAPIServer(&datamanager.StudentLinkDatabase{}, sessionManager)
	go func() {
		err := apiServer.ListenAndServe()
		if err != nil {
			log.Fatal(err)
			return
		}
	}()

	select {}
}
