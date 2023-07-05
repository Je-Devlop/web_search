package main

import (
	"Je-devlop/web_site/search"
	"Je-devlop/web_site/store"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Load enviroment file error! %s", err.Error())
	}

	r := gin.Default()

	store, err := store.NewElasticClient(os.Getenv("ELASTIC_ENDPOINT"))
	if err != nil {
		log.Panic(err.Error())
	}

	searchHandler := search.NewHandler(store)
	r.GET("/search", searchHandler.SearchContent)

	log.Fatal(r.Run(":" + os.Getenv("PORT")))
}
