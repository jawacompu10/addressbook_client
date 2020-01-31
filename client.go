package main

import (
	"context"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"

	"bitbucket.org/jawacompu10/addressbook_client/addressbook"
)

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial("localhost:9999", opts...)
	if err != nil {
		log.Fatal("Failed to dial gRPC")
	}
	defer conn.Close()
	client := addressbook.NewAddressBookClient(conn)
	getUserAddresses(client, &addressbook.UserID{ID: "jawahar"})
}

func getUserAddresses(client addressbook.AddressBookClient, userID *addressbook.UserID) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.GetUserAddresses(ctx, userID)
	if err != nil {
		log.Fatal("Failed to get user addresses")
	}
	for {
		address, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Failed in stream recieve")
		}
		log.Printf("%+v", address)
	}
}
