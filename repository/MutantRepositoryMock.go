package repository


import (
	"errors"
	"mutants/interfaces"
	"mutants/models"
)

type MutantRepositoryMock struct {

}

func NewMockRepository() interfaces.IMutantRepository {
	return &MutantRepository{}
}

func (repository *MutantRepositoryMock) SaveDna() error {
	return errors.New("hol")
}

func (repository *MutantRepositoryMock) GetStatistics() (models.Statistics,  error){
	return models.Statistics{}, nil
}


