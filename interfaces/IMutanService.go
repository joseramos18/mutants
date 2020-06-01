package interfaces
import (
	"github.com/gin-gonic/gin"
	"mutants/models"
)

type IMutantService interface {
	IsMutant(*gin.Context) (bool,error)
	GetStatistics(*gin.Context) (models.Statistics, error)
} 