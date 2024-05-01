package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"golang-rest-api-demo/configuration"
	"golang-rest-api-demo/internal/application/controller"
	"golang-rest-api-demo/internal/application/handler/user"
	"golang-rest-api-demo/internal/application/query"
	"golang-rest-api-demo/internal/application/repository"
	"golang-rest-api-demo/internal/application/web"
	"golang-rest-api-demo/internal/pkg/server"
)

func main() {
	app := fiber.New()

	app.Use(recover.New())

	configureSwaggerUi(app)

	userRepository := repository.NewUserRepository()
	userQueryService := query.NewUserQueryService(userRepository)
	userCommandHandler := user.NewCommandHandler(userRepository)
	userController := controller.NewUserController(userQueryService, userCommandHandler)

	// Router initializing
	web.InitRouter(app, userController)

	//customValidator := validation.NewCustomValidator(validator.New())

	// Start server
	server.NewServer(app).StartHttpServer()
}

func configureSwaggerUi(app *fiber.App) {
	if configuration.Env != "prod" {
		//app.Get("/swagger/*", swagger.HandlerDefault)

		app.Get("/", func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusMovedPermanently).Redirect("/swagger/index.html")
		})
	}
}
