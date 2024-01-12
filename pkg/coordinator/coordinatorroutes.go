package coordinator

import (
	"log"

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

	admin := apiVersion.Group("/coordinator")
	{
		admin.POST("/signup", CoordinatorHandler.CoordinatorSignup)
		admin.POST("/signup/verify", CoordinatorHandler.CoordinatorSignupVerify)
		admin.POST("/login", CoordinatorHandler.CoordinatorLogin)
		admin.POST("/package/add", CoordinatorHandler.CoordinatorAddPackage)
	}
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

func (c *Coordinator) CoordinatorAddPackage(ctx *gin.Context) {
	handler.AddPackage(ctx, c.client)
}
