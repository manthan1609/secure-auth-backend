package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/manthan1609/secure-auth-backend/db"
	"github.com/manthan1609/secure-auth-backend/server"
)

func init() {
	dbError := db.InitDB()

	if dbError != nil {
		log.Println("error while initializing the database")
		log.Fatalln(dbError)
	}

	jwtErr := myJwl.InitJWT()

	if jwtErr != nil {
		log.Println("error initializing the jwt")
		log.Fatalln(jwtErr)
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("error initializing the .env")
		log.Fatalln(err)
	}

}

func main() {
	var host = os.Getenv("HOST")
	var port = os.Getenv("PORT")

	serverErr := server.StartServer(host, port)

	if serverErr != nil {
		log.Println("error while starting the server")
		log.Fatalln(serverErr)
	}
}
