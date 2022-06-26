package controllers

import (
	"net/http"
	"trade_simulator/managers"
	"trade_simulator/models"
	"trade_simulator/utils"

	"github.com/labstack/echo/v4"
)

type handle struct {
	sm *managers.ServiceManager
}

func NewAuthController(e *echo.Group, sm *managers.ServiceManager) *handle {
	h := handle{sm: sm}
	return &h
}

func (h *handle) SignUp(c echo.Context) error {
	form := models.SignUpForm{}
	if err := c.Bind(&form); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response(false, "unexpected entity", nil, err))
	}

	if err := h.sm.UserService.Create(&form); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "fail to create an new account", nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, "", nil, nil))
}
