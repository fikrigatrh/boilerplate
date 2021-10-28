package router

import (
	"boilerplate/config/db"
	"boilerplate/config/env"
	"boilerplate/config/log"
	"boilerplate/controller"
	"boilerplate/middlewares"
	"boilerplate/repository/example1"
	error2 "boilerplate/usecase/error"
	example12 "boilerplate/usecase/example1"
	"github.com/Saucon/errcntrct"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	router := gin.New() // CALL LIBRARY GIN GONIC FOR ROUTER

	DB := db.DB // CALL FUNCTION DB
	sqlDb := db.SqlDB

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	logger := log.NewLogCustom(env.Config)

	if err := errcntrct.InitContract(env.Config.JSONPathFile); err != nil {
		logger.Fatal(err, "router : init contract", nil)
	}

	// CALL DEPENDENCY REPOSITORY IN ABOVE
	exsmpleRepo := example1.NewExample1RepoImpl(DB, sqlDb)

	// CALL DEPENDENCY USECASE IN ABOVE
	exampleUsecase := example12.NewUserUsecaseImpl(exsmpleRepo, logger)

	// create router group
	routerGroup := router.Group("/example")

	//error handler
	errorHandlerUsecase := error2.NewErrorHandlerUsecase()

	//middleware
	middlewares.NewTraceIDHandler(routerGroup)
	middlewares.NewErrorHandler(routerGroup, errorHandlerUsecase)

	// CALL DEPENDENCY CONTROLLER IN ABOVE
	controller.NewExampleControllerImpl(routerGroup, exampleUsecase, logger, errorHandlerUsecase)

	// RETURN ROUTER
	return router
}