package service

import (
	"mutants/helper"
	"mutants/models"
	"mutants/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"errors"
)

type MutantService struct {
	repository interfaces.IMutantRepository
}

func New(repository interfaces.IMutantRepository) interfaces.IMutantService {
	return &MutantService{
		repository: repository,
	}
}

func (service *MutantService) IsMutant(ctx *gin.Context) (bool,error) {
	validate := validator.New()
	var err error
	var dna models.DNA
	ctx.BindJSON(&dna)
	if err = validate.Struct(dna); err != nil{
		return false, errors.New("La longitud de Dna es incorrecto")
	}
	dna.IsMutant, err = helper.VerifyDNA(&dna)
	if err != nil {
		return false, err
	}
	go service.repository.SaveDna(dna, ctx)
	return dna.IsMutant, err
}

func (service *MutantService) GetStatistics(ctx *gin.Context) (models.Statistics, error) {
	result, err := service.repository.GetStatistics(ctx)
	return result, err
}
