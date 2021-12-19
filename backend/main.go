package main

import (
	"log"

	"github.com/yanando/StudentLink/api"
	"github.com/yanando/StudentLink/datamanager"
)

func main() {
	dataManager := &datamanager.StudentLinkDatabase{}
	if err := dataManager.Start(); err != nil {
		log.Fatal(err)
	}
	log.Println("Started datamanager")
	defer dataManager.Close()

	sessionManager := api.NewSessionManager()
	log.Println("Initializing API")
	apiServer := api.NewAPIServer(dataManager, sessionManager)
	go func() {
		err := apiServer.ListenAndServe()
		if err != nil {
			log.Fatal(err)
			return
		}
	}()

	select {}
}
