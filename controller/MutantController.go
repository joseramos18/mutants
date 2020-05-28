package controller
import (
	"github.com/gin-gonic/gin"
	"mutants/repository"
	"mutants/service"
	"mutants/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func SetControllers(ctx *gin.Engine, client *mongo.Client) { 
	var ( 
	mutantRepository interfaces.IMutantRepository = repository.New(client)
	mutantService  interfaces.IMutantService = service.New(mutantRepository)
	)
	ctx.POST("mutant", func(ctx *gin.Context) {
		if mutantService.IsMutant(ctx) {
			ctx.JSON(http.StatusOK, true)
		} else {
			ctx.JSON(http.StatusForbidden, false)
		}
	})
	ctx.GET("stats", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK, mutantService.GetStatistics(ctx))
	})
}