package controller
import (
	"mutants/interfaces"
	"mutants/repository"
	"mutants/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetControllers(ctx *gin.Engine, client *mongo.Client) { 
	
	var mutantRepository interfaces.IMutantRepository
	if client != nil {
		mutantRepository = repository.New(client)
	} else {
		mutantRepository = repository.NewMockRepository()
	}
	var (
	mutantService  interfaces.IMutantService = service.New(mutantRepository)
	)
	ctx.POST("mutant", func(ctx *gin.Context) {
		isMutant, err := mutantService.IsMutant(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		}else {
			if isMutant {
				ctx.JSON(http.StatusOK, true)
			} else {
				ctx.JSON(http.StatusForbidden, false)
			}
		}
	})
	ctx.GET("stats", func(ctx *gin.Context){
		response,_ :=mutantService.GetStatistics(ctx)
		ctx.JSON(http.StatusOK, response)
	})
}