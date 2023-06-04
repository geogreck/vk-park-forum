// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package api

import (
	"fmt"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/gofiber/fiber/v2"
)

// Nickname defines model for Nickname.
type Nickname = string

// User Информация о пользователе.
type User struct {
	// About Описание пользователя.
	About *string `json:"about,omitempty"`

	// Email Почтовый адрес пользователя (уникальное поле).
	Email openapi_types.Email `json:"email"`

	// Fullname Полное имя пользователя.
	Fullname string `json:"fullname"`

	// Nickname Имя пользователя (уникальное поле).
	// Данное поле допускает только латиницу, цифры и знак подчеркивания.
	// Сравнение имени регистронезависимо.
	Nickname *string `json:"nickname,omitempty"`
}

// UserUpdate Информация о пользователе.
type UserUpdate struct {
	// About Описание пользователя.
	About *string `json:"about,omitempty"`

	// Email Почтовый адрес пользователя (уникальное поле).
	Email *openapi_types.Email `json:"email,omitempty"`

	// Fullname Полное имя пользователя.
	Fullname *string `json:"fullname,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Создание нового пользователя
	// (POST /user/{nickname}/create)
	UserCreate(c *fiber.Ctx, nickname Nickname) error
	// Получение информации о пользователе
	// (GET /user/{nickname}/profile)
	UserGetOne(c *fiber.Ctx, nickname Nickname) error
	// Изменение данных о пользователе
	// (POST /user/{nickname}/profile)
	UserUpdate(c *fiber.Ctx, nickname Nickname) error
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

type MiddlewareFunc fiber.Handler

// UserCreate operation middleware
func (siw *ServerInterfaceWrapper) UserCreate(c *fiber.Ctx) error {

	var err error

	// ------------- Path parameter "nickname" -------------
	var nickname Nickname

	err = runtime.BindStyledParameter("simple", false, "nickname", c.Params("nickname"), &nickname)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter nickname: %w", err).Error())
	}

	return siw.Handler.UserCreate(c, nickname)
}

// UserGetOne operation middleware
func (siw *ServerInterfaceWrapper) UserGetOne(c *fiber.Ctx) error {

	var err error

	// ------------- Path parameter "nickname" -------------
	var nickname Nickname

	err = runtime.BindStyledParameter("simple", false, "nickname", c.Params("nickname"), &nickname)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter nickname: %w", err).Error())
	}

	return siw.Handler.UserGetOne(c, nickname)
}

// UserUpdate operation middleware
func (siw *ServerInterfaceWrapper) UserUpdate(c *fiber.Ctx) error {

	var err error

	// ------------- Path parameter "nickname" -------------
	var nickname Nickname

	err = runtime.BindStyledParameter("simple", false, "nickname", c.Params("nickname"), &nickname)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter nickname: %w", err).Error())
	}

	return siw.Handler.UserUpdate(c, nickname)
}

// FiberServerOptions provides options for the Fiber server.
type FiberServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router fiber.Router, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, FiberServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router fiber.Router, si ServerInterface, options FiberServerOptions) {
	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	for _, m := range options.Middlewares {
		router.Use(m)
	}

	router.Post(options.BaseURL+"/user/:nickname/create", wrapper.UserCreate)

	router.Get(options.BaseURL+"/user/:nickname/profile", wrapper.UserGetOne)

	router.Post(options.BaseURL+"/user/:nickname/profile", wrapper.UserUpdate)

}
