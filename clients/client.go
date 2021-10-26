package clients

import (
	pb "api-grpc-articles-videogame/proto/users"
	"context"
	"fmt"
	"log"
	"os"

	"time"

	"google.golang.org/grpc"
)

func main() {

	HOST := os.Getenv("HOST_CLIENT")
	PORT := os.Getenv("PORT_CLIENT")

	url := fmt.Sprintf("%v:%v", HOST, PORT)
	conn, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Err connection grpc channel %v", err.Error())
	}

	defer conn.Close()

	cli := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	resp, err := cli.VerifyUser(ctx, &pb.VerifyUserRequest{
		UserId: 2,
	})

	if err != nil {
		log.Fatalf("Error creating a new task %v", err)
	}

	log.Printf("User: %v", resp.IsExist)
}
