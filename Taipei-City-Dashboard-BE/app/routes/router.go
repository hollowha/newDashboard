// Package routes stores all the routes for the Gin router.
package routes

import (
	"TaipeiCityDashboardBE/app/controllers"
	"TaipeiCityDashboardBE/app/middleware"
	"TaipeiCityDashboardBE/global"

	"github.com/gin-gonic/gin"
)

// router.go configures all API routes

var (
	Router      *gin.Engine
	RouterGroup *gin.RouterGroup
)

// ConfigureRoutes configures all routes for the API and sets version router groups.
func ConfigureRoutes() {


	

	Router.Use(middleware.ValidateJWT)
	// API routers
	RouterGroup = Router.Group("/api/" + global.VERSION)
	configureAuthRoutes()
	configureUserRoutes()
	configureComponentRoutes()
	configureDashboardRoutes()
	configureIssueRoutes()
	configureLikeComponent()
	configureFollowComponent()
	configureCommentComponent()
	configureTestRoutes()

	// test routes
}

// configureTestRoutes configures all test routes.
func configureTestRoutes() {
	testRoutes := RouterGroup.Group("/test")
	testRoutes.GET("/data", controllers.GetData) // 使用 GetData 控制器
}

func configureAuthRoutes() {
	// auth routers
	authRoutes := RouterGroup.Group("/auth")
	authRoutes.Use(middleware.LimitAPIRequests(global.AuthLimitAPIRequestsTimes, global.LimitRequestsDuration))
	authRoutes.Use(middleware.LimitTotalRequests(global.AuthLimitTotalRequestsTimes, global.TokenExpirationDuration))
	authRoutes.POST("/login", controllers.Login)
	// taipeipass login callback
	authRoutes.GET("/callback", controllers.ExecIssoAuth)
	authRoutes.POST("/logout", controllers.IssoLogOut)
}

func configureUserRoutes() {
	userRoutes := RouterGroup.Group("/user")
	userRoutes.Use(middleware.LimitAPIRequests(global.UserLimitAPIRequestsTimes, global.LimitRequestsDuration))
	userRoutes.Use(middleware.LimitTotalRequests(global.UserLimitTotalRequestsTimes, global.TokenExpirationDuration))
	userRoutes.Use(middleware.IsLoggedIn())
	{
		userRoutes.GET("/me", controllers.GetUserInfo)
		userRoutes.PATCH("/me", controllers.EditUserInfo)
	}
	userRoutes.Use(middleware.IsSysAdm())
	{
		userRoutes.GET("/", controllers.GetAllUsers)
		userRoutes.PATCH("/:id", controllers.UpdateUserByID)
	}
}

// configureComponentRoutes configures all component routes.
func configureComponentRoutes() {
	componentRoutes := RouterGroup.Group("/component")

	componentRoutes.Use(middleware.LimitAPIRequests(global.ComponentLimitAPIRequestsTimes, global.LimitRequestsDuration))
	componentRoutes.Use(middleware.LimitTotalRequests(global.ComponentLimitTotalRequestsTimes, global.TokenExpirationDuration))
	{
		componentRoutes.GET("/", controllers.GetAllComponents)
		componentRoutes.
			GET("/:id", controllers.GetComponentByID)
		componentRoutes.
			GET("/:id/chart", controllers.GetComponentChartData)
		componentRoutes.GET("/:id/history", controllers.GetComponentHistoryData)
	}
	componentRoutes.Use(middleware.IsSysAdm())
	{
		componentRoutes.
			PATCH("/:id", controllers.UpdateComponent).
			DELETE("/:id", controllers.DeleteComponent)
		componentRoutes.
			PATCH("/:id/chart", controllers.UpdateComponentChartConfig)
		componentRoutes.PATCH("/:id/map", controllers.UpdateComponentMapConfig)
	}
}

func configureDashboardRoutes() {
	dashboardRoutes := RouterGroup.Group("/dashboard")
	dashboardRoutes.Use(middleware.LimitAPIRequests(global.DashboardLimitAPIRequestsTimes, global.LimitRequestsDuration))
	dashboardRoutes.Use(middleware.LimitTotalRequests(global.DashboardLimitTotalRequestsTimes, global.LimitRequestsDuration))
	{
		dashboardRoutes.
			GET("/", controllers.GetAllDashboards)
		dashboardRoutes.
			GET("/:index", controllers.GetDashboardByIndex)
	}
	dashboardRoutes.Use(middleware.IsLoggedIn())
	{
		dashboardRoutes.POST("/", controllers.CreatePersonalDashboard)
		dashboardRoutes.
			PATCH("/:index", controllers.UpdateDashboard).
			DELETE("/:index", controllers.DeleteDashboard)
	}
	dashboardRoutes.Use(middleware.IsSysAdm())
	{
		dashboardRoutes.POST("/public", controllers.CreatePublicDashboard)
		dashboardRoutes.GET("/check-index/:index", controllers.CheckDashboardIndex)
	}
}

func configureIssueRoutes() {
	issueRoutes := RouterGroup.Group("/issue")
	issueRoutes.Use(middleware.LimitAPIRequests(global.IssueLimitAPIRequestsTimes, global.LimitRequestsDuration))
	issueRoutes.Use(middleware.LimitTotalRequests(global.IssueLimitTotalRequestsTimes, global.LimitRequestsDuration))
	issueRoutes.Use(middleware.IsLoggedIn())
	{
		issueRoutes.
			POST("/", controllers.CreateIssue)
	}
	issueRoutes.Use(middleware.IsSysAdm())
	{
		issueRoutes.
			GET("/", controllers.GetAllIssues)
		issueRoutes.
			PATCH("/:id", controllers.UpdateIssueByID)
	}
}

func configureLikeComponent() {
	likeRoutes := RouterGroup.Group("/like")
	likeRoutes.Use(middleware.LimitAPIRequests(global.IssueLimitAPIRequestsTimes, global.LimitRequestsDuration))
	likeRoutes.Use(middleware.LimitTotalRequests(global.IssueLimitTotalRequestsTimes, global.LimitRequestsDuration))
	likeRoutes.GET("/:componentid", controllers.LikeComponentByID)
	likeRoutes.GET("/order-by-likes", controllers.GetPostsOrderByLikes)
	likeRoutes.Use(middleware.IsLoggedIn())
	{
		likeRoutes.
			POST("/", controllers.LikeComponentByID)
	}
}
func configureFollowComponent() {
	followRoutes := RouterGroup.Group("/follow")
	followRoutes.Use(middleware.LimitAPIRequests(global.IssueLimitAPIRequestsTimes, global.LimitRequestsDuration))
	followRoutes.Use(middleware.LimitTotalRequests(global.IssueLimitTotalRequestsTimes, global.LimitRequestsDuration))
	followRoutes.Use(middleware.IsLoggedIn())
	{
		followRoutes.
			POST("/", controllers.FollowComponentByID)
		followRoutes.GET("/", controllers.GetFollowComponentListByUserID)
	}
}
func configureCommentComponent() {
	commentRoutes := RouterGroup.Group("/comment")
	commentRoutes.Use(middleware.LimitAPIRequests(global.IssueLimitAPIRequestsTimes, global.LimitRequestsDuration))
	commentRoutes.Use(middleware.LimitTotalRequests(global.IssueLimitTotalRequestsTimes, global.LimitRequestsDuration))
	commentRoutes.GET("/:componentid", controllers.GetCommentComponentByID)
	commentRoutes.Use(middleware.IsLoggedIn())
	{
		commentRoutes.
			POST("/", controllers.CommentComponentByID)

	}
}
