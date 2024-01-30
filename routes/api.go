package routes

import (
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Get("/users/{id}", userController.Show)

	webHookController := controllers.NewWebhookController()
	facades.Route().Get("api/", webHookController.Index)
	facades.Route().Post("api/", webHookController.Index)
}
