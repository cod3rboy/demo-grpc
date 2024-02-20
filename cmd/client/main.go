package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	pb "github.com/cod3rboy/demo-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	action = flag.String("action", "", "specify invoice action (create,query)")
	server = flag.String("server", "localhost:8000", "server address")
)

func main() {
	flag.Parse()
	action := strings.ToLower(*action)

	switch action {
	case "create":
		handleCreate()
	case "query":
		handleQuery()
	default:
		fmt.Println("missing/invalid action! use '-action create' or '-action query'")
		os.Exit(1)
	}
}

func handleCreate() {
	currency := PromptCurrency()
	amount := PromptAmount(currency)
	service := PromptService()
	person := PromptPerson()

	conn, err := grpc.Dial(*server, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("failed server connection: %v\n", err)
		os.Exit(1)
	}
	client := pb.NewInvoicerServiceClient(conn)
	response, err := client.Create(context.Background(), &pb.CreateRequest{
		Amount: &pb.Amount{
			Currency: currency,
			Value:    amount,
		},
		ServiceAvailed: service,
		FromName:       person,
	})
	if err != nil {
		fmt.Printf("create invoice failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Invoice ID: %s\nUse above id to query your invoice\n", response.GetId())
}

func handleQuery() {
	invoiceId := PromptInvoiceId()

	conn, err := grpc.Dial(*server, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("failed server connection: %v\n", err)
		os.Exit(1)
	}
	client := pb.NewInvoicerServiceClient(conn)
	response, err := client.Get(context.Background(), &pb.InvoiceRequest{InvoiceId: invoiceId})
	if err != nil {
		fmt.Printf("query invoice failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Invoice Status:", response.Status)
	fmt.Println(string(response.Invoice))
}
