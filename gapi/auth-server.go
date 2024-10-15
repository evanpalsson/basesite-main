package gapi

import (
	"github.com/evanpalsson/basesite/config"
	"github.com/evanpalsson/basesite/pb"
	"github.com/evanpalsson/basesite/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	config         config.Config
	authService    services.AuthService
	userService    services.UserService
	userCollection *mongo.Collection
}

func NewGrpcAuthServer(config config.Config, authService services.AuthService,
	userService services.UserService, userCollection *mongo.Collection) (*AuthServer, error) {

	authServer := &AuthServer{
		config:         config,
		authService:    authService,
		userService:    userService,
		userCollection: userCollection,
	}

	return authServer, nil
}
