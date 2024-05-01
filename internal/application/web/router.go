package web

import (
	"github.com/gofiber/fiber/v2"
	"golang-rest-api-demo/internal/application/controller"
	"net/http"
)

func InitRouter(app *fiber.App, userController controller.IUserController) {

	app.Get("/healthcheck", func(context *fiber.Ctx) error {
		return context.SendStatus(http.StatusOK)
	})

	dcankayrakRouteGroup := app.Group("/api/v1/")

	dcankayrakRouteGroup.Get("/user", userController.GetUser)
	dcankayrakRouteGroup.Post("/user", userController.Save)
	dcankayrakRouteGroup.Get("/user/:userId", userController.GetUserById)
}
