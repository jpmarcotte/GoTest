package employee

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	store Store
}

func NewHandler() (*Handler, error) {
	//return &Handler{store: NewMemoryStore()}, nil
	store, err := NewSQLiteStore()
	return &Handler{store: store}, err
}

func (h *Handler) GetAllEmployees(ctx echo.Context) error {
	employees, err := h.store.GetAllEmployees()
	if err != nil {
		return ctx.JSONPretty(http.StatusInternalServerError, struct{
			Message string `json:"message"`
		}{Message:err.Error()}, "  ")
	}
	return ctx.JSONPretty(http.StatusOK, employees, "  ")
}
