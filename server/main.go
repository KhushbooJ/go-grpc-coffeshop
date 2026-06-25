package main

import (
	"context"
	pb "go-grpc-coffeshop/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCoffeeShopServer
}

func (s *server) GetMenu(request *pb.NoParam, srv pb.CoffeeShop_GetMenuServer) error {
	items := []*pb.Items{
		{
			Id:   "1",
			Name: "Black Coffee",
		},
		{
			Id:   "2",
			Name: "Chaii hLatte",
		},
		{
			Id:   "3",
			Name: "Americano",
		},
		{
			Id:   "3",
			Name: "Mocha",
		},
	}

	for i := range items {
		srv.Send(&pb.Menu{
			Item: items[0 : i+1],
		})
	}
	return nil
}

func (s *server) PlaceOrder(ctx context.Context, order *pb.Order) (*pb.Receipt, error) {
	return &pb.Receipt{
		Id: "abc123",
	}, nil
}
func (s *server) GetOrderStatus(ctx context.Context, reciept *pb.Receipt) (*pb.OrderStatus, error) {
	return &pb.OrderStatus{
		OrderId: reciept.Id,
		Status:  "IN PROGRESS",
	}, nil

}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Unable to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCoffeeShopServer(grpcServer, &server{})

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server %v", err)
	}
}
