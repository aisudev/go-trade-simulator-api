package controllers

import (
	"net/http"
	"trade_simulator/managers"
	"trade_simulator/models"
	"trade_simulator/utils"

	"github.com/labstack/echo/v4"
)

type assetHandle struct {
	sm *managers.ServiceManager
}

func NewAssetController(e *echo.Group, sm *managers.ServiceManager) *assetHandle {
	h := assetHandle{sm: sm}
	e.POST("", h.CreateAsset)
	return &h
}

func (h *assetHandle) CreateAsset(c echo.Context) error {
	form := models.AssetForm{}
	if err := c.Bind(&form); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response(false, "", nil, err))
	}

	if err := h.sm.AssetService.Create(&form); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, "created", nil, nil))
}
