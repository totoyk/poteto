package repository

import (
	"github.com/totoyk/echo/internal/domain/model"
	"gorm.io/gorm"
)

type PetRepository struct {
	Conn *gorm.DB
}

func NewPetRepository(dbconn *gorm.DB) *PetRepository {
	return &PetRepository{
		Conn: dbconn,
	}
}

func (r *PetRepository) FindAll() (pets []model.Pet, err error) {
	err = r.Conn.Find(&pets).Error
	return
}

func (r *PetRepository) FindById(id uint) (pet model.Pet, err error) {
	err = r.Conn.First(&pet, id).Error
	return
}

func (r *PetRepository) Create(pet *model.Pet) (err error) {
	err = r.Conn.Create(pet).Error
	return
}

func (r *PetRepository) Update(pet *model.Pet) (err error) {
	err = r.Conn.Save(pet).Error
	return
}

func (r *PetRepository) Delete(pet *model.Pet) (err error) {
	err = r.Conn.Delete(pet).Error
	return
}
