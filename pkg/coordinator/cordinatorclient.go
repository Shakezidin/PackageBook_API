package coordinator

import (
	"log"

	"github.com/Shakezidin/pkg/config"
	pb "github.com/Shakezidin/pkg/coordinator/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Configure) (pb.CoordinatorClient, error) {
	grpc, err := grpc.Dial(cfg.GRPCCORDINATORPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error Dialing to grpc client: %s, ", cfg.GRPCCORDINATORPORT)
		return nil, err
	}
	log.Printf("Succesfully Connected to Admin Client at port: %v", cfg.GRPCCORDINATORPORT)
	return pb.NewCoordinatorClient(grpc), nil
}
