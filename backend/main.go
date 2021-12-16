package main

import (
	"log"

	"github.com/yanando/StudentLink/api"
	"github.com/yanando/StudentLink/datamanager"
)

func main() {
	log.Println("Initializing API")
	apiServer := api.New(&datamanager.StudentLinkDatabase{})
	go func() {
		err := apiServer.ListenAndServe()
		if err != nil {
			log.Fatal(err)
			return
		}
	}()

	select {}
}
