package controllers

import (
	"net/http"
	"trade_simulator/managers"
	"trade_simulator/models"
	"trade_simulator/utils"

	"github.com/labstack/echo/v4"
)

type handleUser struct {
	sm *managers.ServiceManager
}

func NewUserController(e *echo.Group, sm *managers.ServiceManager) *handleUser {
	h := handleUser{sm: sm}

	e.GET("/info", h.Info)
	e.POST("/withdraw", h.Withdraw)
	e.POST("/deposit", h.Deposit)

	return &h
}

func (h *handleUser) Info(c echo.Context) error {
	id := c.Get("uid")
	user, err := h.sm.UserService.OneByID(id.(string))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.Response(false, "", nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, "", user, nil))
}

func (h *handleUser) Withdraw(c echo.Context) error {
	id := c.Get("uid")
	form := models.BalanceForm{}
	if err := c.Bind(&form); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response(false, "unexpected entity", nil, err))
	}

	user, err := h.sm.UserService.Withdraw(id.(string), form.Amount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "withdraw failed", nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, "", user, nil))
}

func (h *handleUser) Deposit(c echo.Context) error {
	id := c.Get("uid")
	form := models.BalanceForm{}
	if err := c.Bind(&form); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response(false, "unexpected entity", nil, err))
	}

	user, err := h.sm.UserService.Deposit(id.(string), form.Amount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "deposit failed", nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, "", user, nil))
}
