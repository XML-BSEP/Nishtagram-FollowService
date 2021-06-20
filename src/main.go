package main

import (
	"FollowService/http/router"
	"FollowService/infrastructure/grpc/follow_service"
	"FollowService/infrastructure/mongo"
	"FollowService/infrastructure/seeder"
	"FollowService/interactor"
	"fmt"
	"github.com/gin-gonic/gin"
	logger "github.com/jelena-vlajkov/logger/logger"
	"google.golang.org/grpc"
	"log"
	"net"
	"context"
	"os"
)
func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return lis
}

func main() {
	logger := logger.InitializeLogger("follow-service", context.Background())

	mongoCli, ctx := mongo.NewMongoClient()
	db := mongo.GetDbName()

	seeder.SeedData(db, mongoCli, ctx)

	i := interactor.NewInteractor(mongoCli, logger)
	appHandler := i.NewAppHandler()

	g := router.NewRouter(appHandler)
	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	list := getNetListener(8077)
	grpcServer := grpc.NewServer()
	s := i.NewFollowServiceImpl()
	follow_service.RegisterFollowServiceServer(grpcServer, s)

	go func() {
		log.Fatalln(grpcServer.Serve(list))
	}()

	//g.Run(":8089")
	if os.Getenv("DOCKER_ENV") == "" {
		err := g.RunTLS(":8089", "certificate/cert.pem", "certificate/key.pem")
		if err != nil {
			return 
		}
		
	} else {
		err := g.Run(":8089")
		if err != nil {
			return 
		}
	}

	g.Run("127.0.0.1:8089")
}
