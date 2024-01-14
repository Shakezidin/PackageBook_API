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
	}
}

func (a *Admin) AdminAuthenticate(ctx *gin.Context) {
	email, _,err := middleware.ValidateToken(ctx,"admin")
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
