package coordinator

import (
	"log"

	pb "github.com/Shakezidin/pkg/coordinator/pb"
	"github.com/Shakezidin/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Configure) (pb.CoordinatorClient, error) {
	grpc, err := grpc.Dial(":"+cfg.GRPCCORDINATORPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc client: %s, ", cfg.GRPCCORDINATORPORT)
		return nil, err
	}
	log.Printf("succesfully Connected to Admin Client at port: %v", cfg.GRPCCORDINATORPORT)
	return pb.NewCoordinatorClient(grpc), nil
}
