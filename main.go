package main

import (
	"go-sample/app/controllers"
	"go-sample/app/db"
	"go-sample/app/repositories"
	"go-sample/app/router"
	"go-sample/app/services"
	"go-sample/app/utils/logic"
	"go-sample/app/utils/validation"
)

func main() {
	// DB接続
	db := db.Init()

	// logic層
	authLogic := logic.NewAuthLogic()
	userLogic := logic.NewUserLogic()
	todoLogic := logic.NewTodoLogic()
	responseLogic := logic.NewResponseLogic()
	jwtLogic := logic.NewJWTLogic()

	// validation層
	authValidate := validation.NewAuthValidation()
	todoValidate := validation.NewTodoValidation()

	// repository層
	userRepo := repositories.NewUserRepository(db)
	todoRepo := repositories.NewTodoRepository(db)

	// service層
	authService := services.NewAuthService(userRepo, authLogic, userLogic, responseLogic, jwtLogic, authValidate)
	todoService := services.NewTodoService(todoRepo, todoLogic, responseLogic, todoValidate)

	// controller層
	appController := controllers.NewAppController()
	authController := controllers.NewAuthController(authService)
	todoContoroller := controllers.NewTodoController(todoService, authService)

	// router設定
	appRouter := router.NewAppRouter(appController)
	authRouter := router.NewAuthRouter(authController)
	todoRouter := router.NewTodoRouter(todoContoroller)
	mainRouter := router.NewMainRouter(appRouter, authRouter, todoRouter)

	// API起動
	mainRouter.StartWebServer()
}
