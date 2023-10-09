package main
import (
		"go-api/pkg/api"
		"log"
)

func main(){

	router := api.InitRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}