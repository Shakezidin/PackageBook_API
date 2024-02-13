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

type AdminAPI struct {
	cfg    *config.Configure
	client pb.AdminClient
}

func NewAdminRoutes(c *gin.Engine, cfg config.Configure) {
	client, err := ClientDial(cfg)
	if err != nil {
		log.Fatalf("error: could not connect to gRPC server: %v", err)
	}

	adminAPI := &AdminAPI{
		cfg:    &cfg,
		client: client,
	}

	apiVersion := c.Group("/api/v1")

	admin := apiVersion.Group("/admin")
	{
		admin.POST("/login", adminAPI.Login)
		admin.GET("/login/dashbord", adminAPI.Authenticate, adminAPI.ViewDashboard)

		admin.POST("/category/add", adminAPI.Authenticate, adminAPI.AddCategory)
		admin.GET("/catagory/view", adminAPI.Authenticate, adminAPI.ViewCategories)

		admin.GET("/pacakages/view", adminAPI.Authenticate, adminAPI.ViewPackages)
		admin.GET("/pacakage/view", adminAPI.Authenticate, adminAPI.ViewPackage)
		admin.GET("/package/status", adminAPI.Authenticate, adminAPI.PackageStatus)

		admin.GET("/destination/view", adminAPI.Authenticate, adminAPI.ViewDestination)
		admin.GET("/activity/view", adminAPI.Authenticate, adminAPI.ViewActivity)

		admin.GET("/users/view", adminAPI.Authenticate, adminAPI.ViewUsers)
		admin.GET("/user/view",adminAPI.Authenticate,adminAPI.ViewUser)
		// admin.GET("/user/block",adminHandler.AdminAuthenticate,adminHandler.BlockUser)
		// admin.GET("/user/view/blocked",adminHandler.AdminAuthenticate,adminHandler.ViewBlockedUsers)
		// admin.GET("/user/view/unblocked",adminHandler.AdminAuthenticate,adminHandler.ViewUnBlockedUsers)

		admin.GET("/coordinators/view", adminAPI.Authenticate, adminAPI.ViewCoordinators)
		admin.GET("/coordinator/view",adminAPI.Authenticate,adminAPI.ViewCoordinator)
		// admin.GET("/coordinator/block",adminHandler.AdminAuthenticate,adminHandler.BlockCoordinator)
		// admin.GET("/coordinator/view/blocked",adminHandler.AdminAuthenticate,adminHandler.ViewBlockCoordinator)
		// admin.GET("/coordinator/view/unblocked",adminHandler.AdminAuthenticate,adminHandler.ViewUnBlockedCoordinator)

		admin.GET("/bookings/view", adminAPI.Authenticate, adminAPI.ViewBookings)
		admin.POST("/bookings/view/filter", adminAPI.Authenticate, adminAPI.FilterBookings)
		admin.GET("/booking/view", adminAPI.Authenticate, adminAPI.ViewBooking)

		// admin.GET("/banners", adminHandler.AdminAuthenticate, adminHandler.ViewCoordinators)
		// admin.GET("/banner/details", adminHandler.AdminAuthenticate, adminHandler.ViewCoordinators)
		// admin.GET("/banner/activate", adminHandler.AdminAuthenticate, adminHandler.ViewCoordinators)
		// admin.DELETE("/banner/delete", adminHandler.AdminAuthenticate, adminHandler.ViewCoordinators)
	}
}

func (a *AdminAPI) Authenticate(ctx *gin.Context) {
	email, _, err := middleware.ValidateToken(ctx, "admin")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"Error":  "Unauthorized",
			"Status": http.StatusUnauthorized,
		})
		return
	}
	ctx.Set("registered_email", email)
	ctx.Next()
}

func (a *AdminAPI) Login(ctx *gin.Context) {
	handler.AdminLoginHandler(ctx, a.client, "admin")
}

func (a *AdminAPI) ViewDashboard(ctx *gin.Context) {
	handler.ViewDashboard(ctx, a.client)
}

func (a *AdminAPI) AddCategory(ctx *gin.Context) {
	handler.AddCategory(ctx, a.client)
}

func (a *AdminAPI) ViewCategories(ctx *gin.Context) {
	handler.ViewCategories(ctx, a.client)
}

func (a *AdminAPI) ViewPackages(ctx *gin.Context) {
	handler.ViewPackages(ctx, a.client)
}

func (a *AdminAPI) ViewPackage(ctx *gin.Context) {
	handler.ViewPackage(ctx, a.client)
}

func (a *AdminAPI) PackageStatus(ctx *gin.Context) {
	handler.PackageStatus(ctx, a.client)
}

func (a *AdminAPI) ViewDestination(ctx *gin.Context) {
	handler.ViewDestination(ctx, a.client)
}

func (a *AdminAPI) ViewActivity(ctx *gin.Context) {
	handler.ViewActivity(ctx, a.client)
}

func (a *AdminAPI) ViewCoordinators(ctx *gin.Context) {
	handler.ViewCoordinators(ctx, a.client)
}

func (a *AdminAPI) ViewBookings(ctx *gin.Context) {
	handler.ViewBookings(ctx, a.client)
}

func (a *AdminAPI) ViewBooking(ctx *gin.Context) {
	handler.ViewBooking(ctx, a.client)
}

func (a *AdminAPI) FilterBookings(ctx *gin.Context) {
	handler.FilterBookings(ctx, a.client)
}

func (a *AdminAPI)ViewUsers(ctx *gin.Context){
	handler.ViewUsers(ctx,a.client)
}

func (a *AdminAPI)ViewUser(ctx *gin.Context){
	handler.ViewUser(ctx,a.client)
}

func (a *AdminAPI)ViewCoordinator(ctx *gin.Context){
	
}