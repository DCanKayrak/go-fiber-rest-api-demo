package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"golang-rest-api-demo/configuration"
	"golang-rest-api-demo/internal/application/controller"
	"golang-rest-api-demo/internal/application/handler/user"
	"golang-rest-api-demo/internal/application/query"
	"golang-rest-api-demo/internal/application/repository"
	"golang-rest-api-demo/internal/application/web"
	"golang-rest-api-demo/internal/pkg/server"
)
import _ "golang-rest-api-demo/docs"

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
		app.Get("/swagger/*", swagger.HandlerDefault)

		app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
			URL:         "http://example.com/doc.json",
			DeepLinking: false,
			// Expand ("list") or Collapse ("none") tag groups by default
			DocExpansion: "none",
			// Prefill OAuth ClientId on Authorize popup
			OAuth: &swagger.OAuthConfig{
				AppName:  "OAuth Provider",
				ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
			},
			// Ability to change OAuth2 redirect uri location
			OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
		}))

		app.Get("/", func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusMovedPermanently).Redirect("/swagger/index.html")
		})
	}
}
