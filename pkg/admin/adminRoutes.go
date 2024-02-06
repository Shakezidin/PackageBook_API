package admin

import (
	"log"
	"net/http"

	"github.com/Shakezidin/middleware"
	"github.com/Shakezidin/pkg/admin/handler"
	pb "github.com/Shakezidin/pkg/admin/pb"
	"github.com/Shakezidin/pkg/config"
	"github.com/gin-gonic/gin"
)

type Admin struct {
	cfg    *config.Configure
	client pb.AdminClient
}

func NewAdminRoutes(c *gin.Engine, cfg config.Configure) {
	client, err := ClientDial(cfg)
	if err != nil {
		log.Fatalf("error Not connected with gRPC server, %v", err.Error())
	}

	adminHandler := &Admin{
		cfg:    &cfg,
		client: client,
	}

	apiVersion := c.Group("/api/v1")

	admin := apiVersion.Group("/admin")
	{
		admin.POST("/login", adminHandler.AdminLogin)

		admin.POST("/category/add", adminHandler.AdminAuthenticate, adminHandler.AddCategory)
		admin.GET("/catagory/view", adminHandler.AdminAuthenticate, adminHandler.ViewCatagories)

		admin.GET("/pacakages/view", adminHandler.AdminAuthenticate, adminHandler.ViewPackages)
		admin.GET("/pacakage/view", adminHandler.AdminAuthenticate, adminHandler.ViewPackage)
		admin.GET("/package/status", adminHandler.AdminAuthenticate, adminHandler.PackageStatus)

		admin.GET("/destination/view", adminHandler.AdminAuthenticate, adminHandler.ViewDestination)
		admin.GET("/activity/view", adminHandler.AdminAuthenticate, adminHandler.ViewActivity)

		admin.GET("/users/view", adminHandler.AdminAuthenticate, adminHandler.ViewUsers)
		// admin.GET("/user/view",adminHandler.AdminAuthenticate,adminHandler.ViewUser)
		// admin.GET("/user/block",adminHandler.AdminAuthenticate,adminHandler.BlockUser)
		// admin.GET("/user/view/blocked",adminHandler.AdminAuthenticate,adminHandler.ViewBlockedUsers)
		// admin.GET("/user/view/unblocked",adminHandler.AdminAuthenticate,adminHandler.ViewUnBlockedUsers)

		admin.GET("coordinator/view", adminHandler.AdminAuthenticate, adminHandler.ViewCoordinators)
		// admin.GET("/coordinator/view",adminHandler.AdminAuthenticate,adminHandler.ViewCoordinator)
		// admin.GET("/coordinator/block",adminHandler.AdminAuthenticate,adminHandler.BlockCoordinator)
		// admin.GET("/coordinator/view/blocked",adminHandler.AdminAuthenticate,adminHandler.ViewBlockCoordinator)
		// admin.GET("/coordinator/view/unblocked",adminHandler.AdminAuthenticate,adminHandler.ViewUnBlockedCoordinator)

		admin.GET("bookings/view", adminHandler.AdminAuthenticate, adminHandler.ViewBookings)
		// admin.GET("booking/view", adminHandler.AdminAuthenticate, adminHandler.ViewBooking)

		// admin.GET("/banners", adminHandler.AdminAuthenticate, adminHandler.ViewCoordinators)
		// admin.GET("/banner/details", adminHandler.AdminAuthenticate, adminHandler.ViewCoordinators)
		// admin.GET("/banner/activate", adminHandler.AdminAuthenticate, adminHandler.ViewCoordinators)
		// admin.DELETE("/banner/delete", adminHandler.AdminAuthenticate, adminHandler.ViewCoordinators)
	}
}

func (a *Admin) AdminAuthenticate(ctx *gin.Context) {
	email, _, err := middleware.ValidateToken(ctx, "admin")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":  err.Error(),
			"status": http.StatusUnauthorized,
		})
		return
	}
	ctx.Set("registered_email", email)
	ctx.Next()
}

func (a *Admin) AdminLogin(ctx *gin.Context) {
	handler.AdminLoginHandler(ctx, a.client, "admin")
}

func (a *Admin) AddCategory(ctx *gin.Context) {
	handler.AddCategory(ctx, a.client)
}

func (a *Admin) ViewCatagories(ctx *gin.Context) {
	handler.ViewCatagories(ctx, a.client)
}

func (a *Admin) ViewPackages(ctx *gin.Context) {
	handler.ViewPackages(ctx, a.client)
}

func (a *Admin) ViewPackage(ctx *gin.Context) {
	handler.ViewPackage(ctx, a.client)
}

func (a *Admin) PackageStatus(ctx *gin.Context) {
	handler.PackageStatus(ctx, a.client)
}

func (a *Admin) ViewDestination(ctx *gin.Context) {
	handler.ViewDestination(ctx, a.client)
}

func (a *Admin) ViewActivity(ctx *gin.Context) {
	handler.ViewActivity(ctx, a.client)
}

func (a *Admin) ViewUsers(ctx *gin.Context) {
	// handler.ViewUser(ctx, a.client)
}

func (a *Admin) ViewCoordinators(ctx *gin.Context) {
	// handler.ViewCoordinators(ctx, a.client)
}

func (a *Admin) ViewBookings(ctx *gin.Context) {
	// handler.ViewBookings(ctx, a.client)
}
