package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	helperHandler "github.com/michaelyusak/go-helper/handler"
	"github.com/michaelyusak/go-helper/middleware"
	"github.com/michaelyusak/kredit-plus-xyz/adaptor"
	"github.com/michaelyusak/kredit-plus-xyz/config"
	"github.com/michaelyusak/kredit-plus-xyz/handler"
	"github.com/michaelyusak/kredit-plus-xyz/repository"
	"github.com/michaelyusak/kredit-plus-xyz/service"
	"github.com/sirupsen/logrus"
)

type routerOpts struct {
	common      *helperHandler.CommonHandler
	user        *handler.UserHandler
	transaction *handler.TransactionHandler
}

func createRouter(config config.ServiceConfig, log *logrus.Logger) *gin.Engine {
	postgres := adaptor.ConnectPostgres(config.Postgres, log)

	userRepo := repository.NewUserRepositoryPostgres(postgres)
	transactionRepo := repository.NewTransactionRepositoryPostgres(postgres)

	userService := service.NewUserService(userRepo)
	transactionService := service.NewTransactionService(transactionRepo)

	commonHandler := &helperHandler.CommonHandler{}
	userHandler := handler.NewUserHandler(userService, config.ContextTimeout)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	opt := routerOpts{
		common:      commonHandler,
		user:        userHandler,
		transaction: transactionHandler,
	}

	router := newRouter(opt, log)

	return router
}

func newRouter(routerOpts routerOpts, log *logrus.Logger) *gin.Engine {
	router := gin.New()

	corsConfig := cors.DefaultConfig()

	router.ContextWithFallback = true

	router.Use(
		middleware.Logger(log),
		middleware.RequestIdHandlerMiddleware,
		middleware.ErrorHandlerMiddleware,
		gin.Recovery(),
	)

	corsRouting(router, corsConfig)
	commonRouting(router, routerOpts.common)
	userRouting(router, routerOpts.user)

	return router
}

func corsRouting(router *gin.Engine, configCors cors.Config) {
	configCors.AllowOrigins = []string{"localhost"}
	configCors.AllowMethods = []string{"POST", "GET", "PUT", "PATCH", "DELETE"}
	configCors.AllowHeaders = []string{"Origin", "Authorization", "Content-Type", "Accept", "User-Agent", "Cache-Control"}
	configCors.ExposeHeaders = []string{"Content-Length"}
	configCors.AllowCredentials = true
	router.Use(cors.New(configCors))
}

func commonRouting(router *gin.Engine, common *helperHandler.CommonHandler) {
	router.GET("/ping", common.Ping)
	router.NoRoute(common.NoRoute)
}

func userRouting(router *gin.Engine, user *handler.UserHandler) {
	router.POST("/api/v1/register", user.Register)
}
