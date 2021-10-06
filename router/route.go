package router

import (
	"boilerplate/config/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	router := gin.New() // CALL LIBRARY GIN GONIC FOR ROUTER

	DB := db.DB // CALL FUNCTION DB
	fmt.Println(DB)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// CALL DEPENDENCY REPOSITORY IN ABOVE

	// CALL DEPENDENCY USECASE IN ABOVE

	// create router group
	//newRoute := router.Group("/example")

	// CALL DEPENDENCY CONTROLLER IN ABOVE


	// RETURN ROUTER
	return router
}