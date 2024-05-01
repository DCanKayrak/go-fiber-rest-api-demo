package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang-rest-api-demo/internal/application/controller/request"
	"golang-rest-api-demo/internal/application/controller/response"
	"golang-rest-api-demo/internal/application/handler/user"
	"golang-rest-api-demo/internal/application/query"
	"net/http"
)

type IUserController interface {
	Save(ctx *fiber.Ctx) error
	GetUserById(ctx *fiber.Ctx) error
	GetUser(ctx *fiber.Ctx) error
}

type userController struct {
	userQueryService   query.IUserQueryService
	userCommandHandler user.ICommandHandler
}

func NewUserController(userQueryService query.IUserQueryService, userCommandHandler user.ICommandHandler) IUserController {
	return &userController{
		userQueryService:   userQueryService,
		userCommandHandler: userCommandHandler,
	}
}

func (c *userController) Save(ctx *fiber.Ctx) error {
	var req request.UserCreateRequest
	err := ctx.BodyParser(&req)

	if err != nil {
		fmt.Printf("userController.Save ERROR -> There was an error while binding json - ERROR: %v\n", err.Error())
		return err
	}

	fmt.Printf("userController.Save STARTED with request: %#v\n", req)

	if err = c.userCommandHandler.Save(ctx.UserContext(), req.ToCommand()); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON("User Created Successfully")
}

func (c *userController) GetUserById(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")

	if userId == "" {
		return ctx.Status(http.StatusBadRequest).JSON("userId cannot be empty")
	}

	fmt.Printf("userController.GetUserById STARTED with userId: %s\n", userId)

	user, err := c.userQueryService.GetById(ctx.UserContext(), userId)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(response.ToUserResponse(user))
}

func (c *userController) GetUser(ctx *fiber.Ctx) error {
	fmt.Printf("userController.GetUser INFO - Started \n")

	users, err := c.userQueryService.Get(ctx.UserContext())

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(response.ToUserResponseList(users))
}
