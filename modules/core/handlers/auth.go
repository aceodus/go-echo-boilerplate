package handlers

import (
	"net/http"

	"github.com/aceodus/go-echo-boilerplate/config"
	"github.com/aceodus/go-echo-boilerplate/modules/core/usecases"
	"github.com/aceodus/go-echo-boilerplate/pkg/middlewares"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	UserUC  usecases.UserUsecase
	Configs *config.AppConfig
}

func NewAuthHandler(g *echo.Group, middManager *middlewares.MiddlewareManager, userUsecase usecases.UserUsecase, appCfgs *config.AppConfig) {
	handler := &AuthHandler{
		UserUC:  userUsecase,
		Configs: appCfgs,
	}

	apiV1 := g.Group("auth")
	apiV1.GET("/login", handler.LoginForm)
	apiV1.GET("/success", handler.LoginSuccess)
}

func (h *AuthHandler) LoginForm(c echo.Context) error {
	data := map[string]interface{}{
		"config":      h.Configs.FirebaseAuthCreds,
		"success_url": h.Configs.BaseURL + "auth/success",
	}
	return c.Render(http.StatusOK, "login.go.tpl", data)
}

func (h *AuthHandler) LoginSuccess(c echo.Context) error {
	data := map[string]interface{}{
		"config": h.Configs.FirebaseAuthCreds,
	}
	return c.Render(http.StatusOK, "success.go.tpl", data)
}
