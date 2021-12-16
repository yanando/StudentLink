package main

import (
	"log"

	"github.com/yanando/StudentLink/api"
)

func main() {
	log.Println("Initializing API")
	apiServer := api.New()
	go func() {
		err := apiServer.ListenAndServe()
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Println("API started")
	}()

}
