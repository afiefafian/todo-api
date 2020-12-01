package main

import (
	"fmt"
	"log"
	"net/http"

	httpDelivery "todo_api/src/delivery/http"

	"github.com/go-pg/pg/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
)

func init() {
	config()
}

func main() {
	// Init DB
	db := initDB()
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Init router
	router := httprouter.New()

	// Set http routing
	httpDelivery.InitRoute(router, db)

	fmt.Println("Served and listen at", viper.GetString(`port`))
	log.Fatal(http.ListenAndServe(":"+viper.GetString("port"), router))
}

func config() {
	viper.SetConfigFile("config/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool("debug") {
		log.Println("Service RUN on DEBUG mode")
	}
}

func initDB() *pg.DB {
	dbUser := viper.GetString(`database.user`)
	dbName := viper.GetString(`database.name`)

	return pg.Connect(&pg.Options{User: dbUser, Database: dbName})
}
