package handler

import (
	"github.com/mimaken3/ShareIT/domain/service"
	"github.com/mimaken3/ShareIT/infrastructure"
)

var userService service.UserServiceInterface

func DI() {
	userRepo := infrastructure.NewUserDB(infrastructure.DB)
	userService = service.NewUserService(userRepo)
}
