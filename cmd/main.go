package main

import (
	"fmt"
	"log"

	"github.com/billowdev/document-system-field-manager/internal/adapters/database"
	"github.com/billowdev/document-system-field-manager/pkg/configs"
)

func main() {
	params := configs.NewFiberHttpServiceParams()
	fiberHTTP := configs.NewFiberHTTPService(params)
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal("Failed to start Database:", err)
	}
	_ = db
	portString := fmt.Sprintf(":%v", params.Port)

	err = fiberHTTP.Listen(portString)

	if err != nil {
		log.Fatal("Failed to start golang Fiber server:", err)
	}
}
