package handler

import (
	"github.com/mimaken3/ShareIT-api/domain/service"
	"github.com/mimaken3/ShareIT-api/infrastructure"
)

var userService service.UserServiceInterface

func DI() {
	userRepo := infrastructure.NewUserDB(infrastructure.DB)
	userService = service.NewUserService(userRepo)
}
