// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"about-vaccine/src/base/dao"
	"about-vaccine/src/controller"
	"about-vaccine/src/repo"
	"about-vaccine/src/router"
	"about-vaccine/src/service"
)

// Injectors from wire.go:

func InitApplication(dsn string) (*router.APIRouter, error) {
	engine, err := dao.NewEngine(dsn)
	if err != nil {
		return nil, err
	}
	db := dao.NewDB(engine)
	userRepo := repo.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	vaersRepo := repo.NewVaersRepo(db)
	vaersResultRepo := repo.NewVaersResultRepo(db)
	vaersSymptomRepo := repo.NewVaersSymptomRepo(db)
	vaersSymptomTermRepo := repo.NewVaersSymptomTermRepo(db)
	vaersVaxRepo := repo.NewVaersVaxRepo(db)
	vaersVaxTermRepo := repo.NewVaersVaxTermRepo(db)
	vaersService := service.NewVaersService(vaersRepo, vaersResultRepo, vaersSymptomRepo, vaersSymptomTermRepo, vaersVaxRepo, vaersVaxTermRepo)
	vaersController := controller.NewVaersController(vaersService)
	vaccineRepo := repo.NewVaccineRepo(db)
	vaccineTypeRepo := repo.NewVaccineTypeRepo(db)
	vaccineService := service.NewVaccineService(vaccineRepo, vaccineTypeRepo)
	vaccineController := controller.NewVaccineController(vaccineService)
	oaeTermRepo := repo.NewOAETermRepo(db)
	oaeTermService := service.NewOaeTermService(oaeTermRepo)
	oaeTermController := controller.NewOAETermController(oaeTermService)
	adverseEventRepo := repo.NewAdverseEventRepo(db)
	adverseSymptomRepo := repo.NewAdverseSymptomRepo(db)
	adverseVaccineRepo := repo.NewAdverseVaccineRepo(db)
	adverseEventService := service.NewAdverseEventService(adverseEventRepo, adverseSymptomRepo, adverseVaccineRepo, vaccineService)
	adverseEventController := controller.NewAdverseEventController(adverseEventService)
	apiRouter := router.NewAPIRouter(userController, vaersController, vaccineController, oaeTermController, adverseEventController)
	return apiRouter, nil
}
