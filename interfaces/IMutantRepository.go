package interfaces

import (
	"mutants/models"
	"context"
)

type IMutantRepository interface {
	SaveDna(models.DNA, context.Context) error
	GetStatistics(context.Context) (models.Statistics,error)
}
