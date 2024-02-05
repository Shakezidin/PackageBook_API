package user

import (
	"log"
	"net/http"

	"github.com/Shakezidin/middleware"
	"github.com/Shakezidin/pkg/config"
	"github.com/Shakezidin/pkg/user/handler"
	pb "github.com/Shakezidin/pkg/user/pb"
	"github.com/gin-gonic/gin"
)

type User struct {
	cfg    *config.Configure
	client pb.UserClient
}

func NewUserRoute(c *gin.Engine, cfg config.Configure) {
	client, err := ClientDial(cfg)
	if err != nil {
		log.Fatalf("error Not connected with gRPC server, %v", err.Error())
	}

	UserHandler := &User{
		cfg:    &cfg,
		client: client,
	}

	apiVersion := c.Group("/api/v1")

	user := apiVersion.Group("/user")
	{
		user.POST("/signup", UserHandler.UserSignup)
		user.POST("/signup/verify", UserHandler.UserSignupVerify)
		user.POST("/login", UserHandler.UserLogin)

		user.POST("/password/forget", UserHandler.ForgetPassword)
		user.POST("/password/forget/verify", UserHandler.ForgetPasswordVerify)
		user.POST("/password/forget/newpassword", UserHandler.NewPassword)
		user.POST("/profile/update", UserHandler.UserAuthenticate, UserHandler.UpdateProfile)

		user.GET("/home/packages", UserHandler.ViewPackages)
		user.GET("/package/search", UserHandler.SearchPackage)
		user.GET("/package/search/filter", UserHandler.PackageFilter)
		user.GET("/catagories/view", UserHandler.ViewCatagories)
		user.GET("/package/view", UserHandler.ViewPackage)
		user.GET("/package/foodmenu/view", UserHandler.ViewFoodMenus)
		user.GET("/destination/view", UserHandler.ViewDestination)
		user.GET("/activity/view", UserHandler.ViewActivity)

		user.POST("/traveller/add", UserHandler.UserAuthenticate, UserHandler.AddTraveller)
		user.GET("/booking/payment/advance", UserHandler.AdvancePayment)

		user.GET("/booking/payment/full", UserHandler.OnlinePayment)
		user.GET("/payment/success", UserHandler.PaymentSuccess)
		user.GET("/success/render", UserHandler.PaymentSuccessPage)

		user.GET("/booking/history", UserHandler.UserAuthenticate, UserHandler.ViewHistory)
		user.GET("/booking/history/view", UserHandler.UserAuthenticate, UserHandler.ViewBooking)

		user.GET("/booking/history/cancel", UserHandler.UserAuthenticate, UserHandler.PackageCancel)
	}
}

func (a *User) UserAuthenticate(ctx *gin.Context) {
	email, _, err := middleware.ValidateToken(ctx, "user")
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

func (c *User) UserLogin(ctx *gin.Context) {
	handler.UserLoginHandler(ctx, c.client, "user")
}

func (c *User) UserSignup(ctx *gin.Context) {
	handler.UserSignupHandler(ctx, c.client, "user")
}

func (c *User) UserSignupVerify(ctx *gin.Context) {
	handler.VerifySignup(ctx, c.client)
}

func (c *User) ViewPackage(ctx *gin.Context) {
	handler.ViewPackage(ctx, c.client)
}

func (c *User) ForgetPassword(ctx *gin.Context) {
	handler.ForgetPassword(ctx, c.client)
}

func (c *User) ForgetPasswordVerify(ctx *gin.Context) {
	handler.ForgetPasswordVerify(ctx, c.client)
}

func (c *User) NewPassword(ctx *gin.Context) {
	handler.NewPassword(ctx, c.client)
}

func (c *User) UpdateProfile(ctx *gin.Context) {
	handler.UpdateProfile(ctx, c.client)
}

func (c *User) ViewDestination(ctx *gin.Context) {
	handler.ViewDestination(ctx, c.client)
}

func (c *User) ViewActivity(ctx *gin.Context) {
	handler.ViewActivity(ctx, c.client)
}

func (c *User) ViewCatagories(ctx *gin.Context) {
	handler.ViewCatagories(ctx, c.client)
}

func (c *User) SearchPackage(ctx *gin.Context) {
	handler.SearchPackage(ctx, c.client)
}

func (c *User) PackageFilter(ctx *gin.Context) {
	handler.PackageFilter(ctx, c.client)
}

func (c *User) AddTraveller(ctx *gin.Context) {
	handler.AddTraveller(ctx, c.client)
}

func (c *User) AdvancePayment(ctx *gin.Context) {
	handler.OnlinePayment(ctx, c.client, "advance")
}

func (c *User) ViewPackages(ctx *gin.Context) {
	handler.ViewPackages(ctx, c.client)
}

func (c *User) OnlinePayment(ctx *gin.Context) {
	handler.OnlinePayment(ctx, c.client, "full")
}

func (c *User) PaymentSuccess(ctx *gin.Context) {
	handler.PaymentSuccess(ctx, c.client)
}

func (c *User) PaymentSuccessPage(ctx *gin.Context) {
	handler.PaymentSuccessPage(ctx, c.client)
}

func (c *User) ViewFoodMenus(ctx *gin.Context) {
	handler.ViewFoodMenus(ctx, c.client)
}

func (c *User) ViewHistory(ctx *gin.Context) {
	handler.ViewHistory(ctx, c.client)
}

func (c *User) ViewBooking(ctx *gin.Context) {
	handler.ViewBooking(ctx, c.client)
}

func (c *User) PackageCancel(ctx *gin.Context) {
	handler.PackageCancel(ctx, c.client)
}
