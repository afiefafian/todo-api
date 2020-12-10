package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-redis/redis/v8"

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
	memDB := initMemBD()
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Init router
	r := httpDelivery.NewRouter(db, memDB)

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

func initMemBD() *redis.Client {
	address := fmt.Sprintf("%s:%s", viper.GetString(`redis.host`), viper.GetString(`redis.port`))
	username := viper.GetString(`redis.username`)
	password := viper.GetString(`redis.password`)
	db, _ := strconv.Atoi(viper.GetString(`redis.db`))

	redis := redis.NewClient(&redis.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})
	return redis
}
