package utils

import (
	"github.com/omrfrkazt/golang-generic-repository-pattern/internal/entities"
	"github.com/omrfrkazt/golang-generic-repository-pattern/internal/models"
)

func UserToEntity(model *models.UserModel) *entities.User {
	var entity = &entities.User{}
	entity.ID = model.ID
	entity.Name = model.Name
	entity.Surname = model.Surname
	entity.Email = model.Email
	entity.Password = model.Password
	entity.Username = model.Username
	return entity
}

func UserToModel(entity *entities.User) *models.UserModel {
	var model = &models.UserModel{}
	model.ID = entity.ID
	model.Name = entity.Name
	model.Surname = entity.Surname
	model.Email = entity.Email
	model.Password = entity.Password
	model.Username = entity.Username
	return model
}

func UsersToModel(entityList *[]entities.User) *[]models.UserModel {
	var modelList []models.UserModel
	for _, entity := range *entityList {
		modelList = append(modelList, *UserToModel(&entity))
	}
	return &modelList
}

func UsersToEntity(modelList *[]models.UserModel) *[]entities.User {
	var entityList []entities.User
	for _, model := range *modelList {
		entityList = append(entityList, *UserToEntity(&model))
	}
	return &entityList
}
