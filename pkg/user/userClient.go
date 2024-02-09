package user

import (
	"log"

	pb "github.com/Shakezidin/pkg/user/userpb"
	"github.com/Shakezidin/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Configure) (pb.UserClient, error) {
	grpc, err := grpc.Dial(":"+cfg.GRPCUSERPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc client: %s, ", cfg.GRPCUSERPORT)
		return nil, err
	}
	log.Printf("succesfully Connected to Admin Client at port: %v", cfg.GRPCUSERPORT)
	return pb.NewUserClient(grpc), nil
}
