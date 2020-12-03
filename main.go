package main

import (
	"log"
	"net/http"

	httpDelivery "todo_api/src/delivery/http"

	"github.com/go-pg/pg/v10"
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
	r := httpDelivery.NewRouter(db)

	log.Println("SKYSHI @2020 ----")
	log.Printf("Listening on : %v", viper.GetString(`port`))
	log.Fatal(http.ListenAndServe(":"+viper.GetString("port"), &r.Router))
}

func config() {
	viper.SetConfigFile("config/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initDB() *pg.DB {
	dbUser := viper.GetString(`database.user`)
	dbName := viper.GetString(`database.name`)

	return pg.Connect(&pg.Options{User: dbUser, Database: dbName})
}
