package main

import (
	"os"

	"github.com/alllga/pixSim/application/grpc"
	"github.com/alllga/pixSim/infrastructure/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main()  {
	database = db.ConnectDB(os.Getenv("env"))

	grpc.StartGrpcServer(database, 50051)
	
}