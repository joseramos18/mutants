package interfaces

import (
	"mutants/models"
	"context"
)

type IMutantRepository interface {
	SaveDna(models.DNA, context.Context)
	GetStatistics(context.Context) models.Statistics
}
