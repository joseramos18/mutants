package service

import (
	"mutants/helper"
	"mutants/models"
	"mutants/interfaces"
	"github.com/gin-gonic/gin"
)

type MutantService struct {
	repository interfaces.IMutantRepository
}

func New(repository interfaces.IMutantRepository) interfaces.IMutantService {
	return &MutantService{
		repository: repository,
	}
}

func (service *MutantService) IsMutant(ctx *gin.Context) bool {
	var dna models.DNA
	ctx.BindJSON(&dna)
	dna.IsMutant = helper.VerifyDNA(&dna)
	go service.repository.SaveDna(dna, ctx)
	return dna.IsMutant
}

func (service *MutantService) GetStatistics(ctx *gin.Context) models.Statistics {
	result := service.repository.GetStatistics(ctx)
	return result
}
