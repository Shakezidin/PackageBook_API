package coordinator

import (
	"log"
	"net/http"

	"github.com/Shakezidin/middleware"
	"github.com/Shakezidin/pkg/config"
	"github.com/Shakezidin/pkg/coordinator/handler"
	pb "github.com/Shakezidin/pkg/coordinator/pb"
	"github.com/gin-gonic/gin"
)

type Coordinator struct {
	cfg    *config.Configure
	client pb.CoordinatorClient
}

func NewCoordinatorRoute(c *gin.Engine, cfg config.Configure) {
	client, err := ClientDial(cfg)
	if err != nil {
		log.Fatalf("error Not connected with gRPC server, %v", err.Error())
	}

	CoordinatorHandler := &Coordinator{
		cfg:    &cfg,
		client: client,
	}

	apiVersion := c.Group("/api/v1")

	coordinator := apiVersion.Group("/coordinator")
	{
		coordinator.POST("/signup", CoordinatorHandler.CoordinatorSignup)
		coordinator.POST("/signup/verify", CoordinatorHandler.CoordinatorSignupVerify)
		coordinator.POST("/login", CoordinatorHandler.CoordinatorLogin)

		coordinator.GET("/dashbord/view", CoordinatorHandler.CoordinatorAuthenticate, CoordinatorHandler.ViewDashBord)

		coordinator.GET("/packages/view", CoordinatorHandler.CoordinatorAuthenticate, CoordinatorHandler.ViewPackages)

		coordinator.POST("/password/forget", CoordinatorHandler.ForgetPassword)
		coordinator.POST("/password/forget/verify", CoordinatorHandler.ForgetPasswordVerify)
		coordinator.POST("/password/forget/newpassword", CoordinatorHandler.NewPassword)

		coordinator.GET("/catagory/view", CoordinatorHandler.CoordinatorViewCatagory)
		coordinator.GET("/location/suggest", CoordinatorHandler.CoordinatorLocations)

		coordinator.POST("/package/add", CoordinatorHandler.CoordinatorAuthenticate, CoordinatorHandler.CoordinatorAddPackage)
		coordinator.GET("/package/view", CoordinatorHandler.CoordinatorAuthenticate, CoordinatorHandler.ViewPackage)
		
		coordinator.POST("/foodmenu/add", CoordinatorHandler.CoordinatorAuthenticate, CoordinatorHandler.CoordinatorAddFoodMenu)
		coordinator.GET("/foodmenu/view", CoordinatorHandler.CoordinatorAuthenticate, CoordinatorHandler.CoordinatorViewFoodMenus)

		coordinator.POST("/destination/add", CoordinatorHandler.CoordinatorAuthenticate, CoordinatorHandler.CoordinatorAddDestination)
		coordinator.GET("/destination/view", CoordinatorHandler.CoordinatorAuthenticate, CoordinatorHandler.ViewDestination)

		coordinator.POST("/activity/add", CoordinatorHandler.CoordinatorAuthenticate, CoordinatorHandler.CoordinatorAddActivity)
		coordinator.GET("/activity/view", CoordinatorHandler.CoordinatorAuthenticate, CoordinatorHandler.ViewActivity)

		coordinator.GET("/bookings/view", CoordinatorHandler.CoordinatorAuthenticate, CoordinatorHandler.ViewBookings)
		coordinator.GET("/booking/view", CoordinatorHandler.CoordinatorAuthenticate, CoordinatorHandler.ViewBooking)
		coordinator.GET("/traveller/view", CoordinatorHandler.CoordinatorAuthenticate, CoordinatorHandler.ViewTraveller)
	}
}

func (a *Coordinator) CoordinatorAuthenticate(ctx *gin.Context) {
	email, _, err := middleware.ValidateToken(ctx, "coordinator")
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

func (c *Coordinator) CoordinatorLogin(ctx *gin.Context) {
	handler.CoordinatorLoginHandler(ctx, c.client, "coordinator")
}

func (c *Coordinator) CoordinatorSignup(ctx *gin.Context) {
	handler.CoordinatorSignupHandler(ctx, c.client)
}

func (c *Coordinator) CoordinatorSignupVerify(ctx *gin.Context) {
	handler.VerifySignup(ctx, c.client)
}

func (c *Coordinator) ViewPackages(ctx *gin.Context) {
	handler.ViewPackages(ctx, c.client)
}

func (c *Coordinator) CoordinatorAddPackage(ctx *gin.Context) {
	handler.AddPackage(ctx, c.client)
}

func (c *Coordinator) CoordinatorAddDestination(ctx *gin.Context) {
	handler.AddDestination(ctx, c.client)
}

func (c *Coordinator) CoordinatorAddActivity(ctx *gin.Context) {
	handler.AddActivity(ctx, c.client)
}

func (c *Coordinator) ViewPackage(ctx *gin.Context) {
	handler.ViewPackage(ctx, c.client)
}

func (c *Coordinator) ViewDestination(ctx *gin.Context) {
	handler.ViewDestination(ctx, c.client)
}

func (c *Coordinator) ViewActivity(ctx *gin.Context) {
	handler.ViewActivity(ctx, c.client)
}

func (c *Coordinator) ForgetPassword(ctx *gin.Context) {
	handler.ForgetPassword(ctx, c.client)
}

func (c *Coordinator) ForgetPasswordVerify(ctx *gin.Context) {
	handler.ForgetPasswordVerify(ctx, c.client)
}

func (c *Coordinator) NewPassword(ctx *gin.Context) {
	handler.NewPassword(ctx, c.client)
}

func (c *Coordinator) CoordinatorLocations(ctx *gin.Context) {
	handler.SuggestLocation(ctx)
}

func (c *Coordinator) CoordinatorViewCatagory(ctx *gin.Context) {
	handler.CoordinatorViewCatagory(ctx, c.client)
}

func (c *Coordinator) CoordinatorAddFoodMenu(ctx *gin.Context) {
	handler.CoordinatorAddFoodMenu(ctx, c.client)
}

func (c *Coordinator) CoordinatorViewFoodMenus(ctx *gin.Context) {
	handler.CoordinatorViewFoodMenus(ctx, c.client)
}

func (c *Coordinator) ViewBookings(ctx *gin.Context) {
	handler.ViewBookings(ctx, c.client)
}

func (c *Coordinator) ViewBooking(ctx *gin.Context) {
	handler.ViewBooking(ctx, c.client)
}

func (c *Coordinator) ViewTraveller(ctx *gin.Context) {
	handler.ViewTraveller(ctx, c.client)
}

func (c *Coordinator) ViewDashBord(ctx *gin.Context) {
	handler.ViewDashBord(ctx, c.client)
}
