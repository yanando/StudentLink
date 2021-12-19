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

	dataManager.AddChatMessage(datamanager.Message{
		ID:          1,
		AuthorID:    1,
		RecipientID: 0,
		Content:     "Test",
	})

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
