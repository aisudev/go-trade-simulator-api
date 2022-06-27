package controllers

import "trade_simulator/managers"

type transactionHandle struct {
	sm *managers.ServiceManager
}

func NewTransactionController(sm *managers.ServiceManager) *transactionHandle {
	h := transactionHandle{sm: sm}

	return &h
}
