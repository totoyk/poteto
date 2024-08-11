package interfaces

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/totoyk/echo/internal/interfaces/oas"
	"github.com/totoyk/echo/internal/interfaces/response"
	"github.com/totoyk/echo/internal/usecase"
)

type PetHandler struct {
	UsecaseReceiver usecase.PetReceiver
}

func NewPetHandler(petReceiver usecase.PetReceiver) *PetHandler {
	return &PetHandler{
		UsecaseReceiver: petReceiver,
	}
}

// (GET /pets)
func (h *PetHandler) ListPets(ctx echo.Context) error {
	result, err := h.UsecaseReceiver.ListPets(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest,
			oas.Error{Code: http.StatusBadRequest, Message: err.Error()},
		)
	}
	res := response.NewPets(ctx, result)
	return ctx.JSON(http.StatusOK, res)
}

// (GET /pets/{id})
func (h *PetHandler) FindPetById(ctx echo.Context, id uint) error {
	result, err := h.UsecaseReceiver.FindPetById(ctx, id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest,
			oas.Error{Code: http.StatusBadRequest, Message: err.Error()},
		)
	}
	res := response.NewPet(ctx, result)
	return ctx.JSON(http.StatusOK, res)
}

// (POST /pets)
func (h *PetHandler) CreatePets(ctx echo.Context) error {
	var params oas.PetRequestBody
	if err := ctx.Bind(&params); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			oas.Error{Code: http.StatusBadRequest, Message: err.Error()},
		)
	}
	err := h.UsecaseReceiver.CreatePets(ctx, params)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest,
			oas.Error{Code: http.StatusBadRequest, Message: err.Error()},
		)
	}
	return ctx.JSON(http.StatusCreated, nil)
}

// (PUT /pets/{id})
func (h *PetHandler) UpdatePet(ctx echo.Context, id uint) error {
	var params oas.PetRequestBody
	if err := ctx.Bind(&params); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			oas.Error{Code: http.StatusBadRequest, Message: err.Error()},
		)
	}
	err := h.UsecaseReceiver.UpdatePet(ctx, id, params)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest,
			oas.Error{Code: http.StatusBadRequest, Message: err.Error()},
		)
	}
	return ctx.JSON(http.StatusOK, nil)
}

// (DELETE /pets/{id})
func (h *PetHandler) DeletePet(ctx echo.Context, id uint) error {
	err := h.UsecaseReceiver.DeletePet(ctx, id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest,
			oas.Error{Code: http.StatusBadRequest, Message: err.Error()},
		)
	}
	return ctx.JSON(http.StatusNoContent, nil)
}
