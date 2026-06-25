package main

import (
	"context"
	pb "go-grpc-coffeshop/proto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to connect %v\n", err)
	}

	defer conn.Close()

	c := pb.NewCoffeeShopClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	menuStream, err := c.GetMenu(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Error calling get menu %v", err)

	}

	done := make(chan bool)
	var items []*pb.Items

	go func() {
		for {
			resp, err := menuStream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				log.Fatalf("could not get menu %v", err)
			}
			items := resp.Item
			log.Printf("Items recieved %v", items)
		}

	}()
	<-done

	reciept, err2 := c.PlaceOrder(ctx, &pb.Order{
		Item: items,
	})
	if err2 != nil {
		log.Fatalf("could not place orders %v", err2)
	}
	log.Printf("Placed order, reciept id is %s", reciept.Id)

	status, err3 := c.GetOrderStatus(ctx, &pb.Receipt{
		Id: reciept.Id,
	})
	if err3 != nil {
		log.Fatalf("could not get status %v", err3)
	}
	log.Printf("Order status is %v", status)

}
