package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"strconv"

	httpDelivery "github.com/afiefafian/todo-api/src/delivery/http"
	httpMiddleware "github.com/afiefafian/todo-api/src/delivery/http/middleware"

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

	initMessage()
	log.Fatal(http.ListenAndServe(":"+viper.GetString("port"), httpMiddleware.HTTPLogger(&r.Router)))
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

	redisClient := redis.NewClient(&redis.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})
	return redisClient
}

func initMessage() {
	fmt.Println("\n   _____  __ ____  __ _____  __  __ ____\n  / ___/ / //_/\\ \\/ // ___/ / / / //  _/\n  \\__ \\ / ,<    \\  / \\__ \\ / /_/ / / /  \n ___/ // /| |   / / ___/ // __  /_/ /   \n/____//_/ |_|  /_/ /____//_/ /_//___/ ")
	fmt.Println("@2020 Powered by: HTTPRouter\n")
	log.Printf("Listening on : %v\n", viper.GetString(`port`))
	log.Println("--------------------------------")
}
