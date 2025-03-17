package main

import (
	"log"
	"server/db"
	"server/internal/user"
	"server/router"
)

func main() {
    dbConn, err := db.NewDatabase()
    if err != nil {
        log.Fatalf("could not initialize db connection: %s", err)
    }

    userRep := user.NewRepository(dbConn.GetDB())
    userSvc := user.NewService(userRep)
    userHandler := user.NewHandler(userSvc)

    router.InitRouter(userHandler)
    router.Start("localhost:8080")
}
