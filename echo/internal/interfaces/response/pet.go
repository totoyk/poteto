package response

import (
	"github.com/labstack/echo/v4"
	"github.com/totoyk/echo/internal/domain/model"
	"github.com/totoyk/echo/internal/interfaces/oas"
	"github.com/totoyk/echo/internal/util/typecast"
)

func NewPet(ctx echo.Context, m *model.Pet) oas.Pet {
	return oas.Pet{
		Id:          m.ID,
		Name:        m.Name,
		Tag:         typecast.NullStringToPtr(m.Tag),
		DateOfBirth: typecast.NullTimeToPtr(m.DateOfBirth),
	}
}

func NewPets(ctx echo.Context, ms *[]model.Pet) []oas.Pet {
	var pets []oas.Pet
	for _, m := range *ms {
		pets = append(pets, NewPet(ctx, &m))
	}
	return pets
}
