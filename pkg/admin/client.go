package admin

import (
	"log"

	pb "github.com/Shakezidin/pkg/admin/pb"
	"github.com/Shakezidin/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Configure) (pb.AdminClient, error) {
	grpc, err := grpc.Dial(cfg.ADMINPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error dialing to grpc client: %s, ", cfg.ADMINPORT)
		return nil, err
	}
	log.Printf("Succesfully Connected to Admin Client at port: %v", cfg.ADMINPORT)
	return pb.NewAdminClient(grpc), nil
}
