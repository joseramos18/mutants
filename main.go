package main
import (
	"fmt"
	"mutants/controller"
	"mutants/infrastructures"
	"os"
	"github.com/gin-gonic/gin"
)
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := run(port); err != nil {
		fmt.Println("error")
	}
}

func run(port string) error {
	client := infrastructures.GetMongoClient()
	defer client.Disconnect(infrastructures.GetConnectionContext())
	router := gin.Default()
	controller.SetControllers(router, client)
	return router.Run(":" + port)
}