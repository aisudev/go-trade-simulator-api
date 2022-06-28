package controllers

import (
	"net/http"
	"trade_simulator/managers"
	"trade_simulator/models"
	"trade_simulator/utils"

	"github.com/labstack/echo/v4"
)

type transactionHandle struct {
	sm *managers.ServiceManager
}

func NewTransactionController(e *echo.Group, sm *managers.ServiceManager) *transactionHandle {
	h := transactionHandle{sm: sm}

	e.POST("/open", h.Open)
	e.POST("/close", h.Close)
	e.POST("/close_all", h.CloseAll)

	return &h
}

func (h *transactionHandle) Open(c echo.Context) error {
	id := c.Get("uid")

	form := models.TransactionForm{}
	if err := c.Bind(&form); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response(false, "Unexpected Entity", nil, err))
	}

	tx, err := h.sm.TransactionService.Open(id.(string), form.HistoricalID, form.Amount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, "", tx, nil))
}

func (h *transactionHandle) Close(c echo.Context) error {
	id := c.Get("uid")

	form := models.TransactionForm{}
	if err := c.Bind(&form); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response(false, "Unexpected Entity", nil, err))
	}

	tx, err := h.sm.TransactionService.Close(id.(string), form.HistoricalID, form.Amount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, "", tx, nil))
}

func (h *transactionHandle) CloseAll(c echo.Context) error {
	id := c.Get("uid")

	form := models.TransactionForm{}
	if err := c.Bind(&form); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response(false, "Unexpected Entity", nil, err))
	}

	tx, err := h.sm.TransactionService.CloseAll(id.(string), form.HistoricalID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, "", tx, nil))
}
