package services

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	pb "github.com/cod3rboy/demo-grpc/proto"
	"github.com/google/uuid"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type Invoice struct {
	ID     string
	Status pb.InvoiceStatusEnum
	Data   []byte
}

var (
	invoices map[string]*Invoice = make(map[string]*Invoice)
	mu       sync.Mutex
)

type Transaction struct {
	Amount        int64
	Currency      string
	From          string
	Service       string
	InvoiceIDChan chan string
}

type invoicerService struct {
	transactions chan Transaction
	pb.UnimplementedInvoicerServiceServer
}

func NewInvoicerService() pb.InvoicerServiceServer {
	service := &invoicerService{
		transactions: make(chan Transaction),
	}
	go service.invoiceGenerator()
	return service
}

func (s *invoicerService) Create(ctx context.Context, create *pb.CreateRequest) (*pb.CreateResponse, error) {
	log.Println("received request to generate invoice")
	// create transaction
	transaction := Transaction{
		Amount:        create.Amount.Value,
		Currency:      create.Amount.Currency,
		From:          create.FromName,
		Service:       create.ServiceAvailed,
		InvoiceIDChan: make(chan string),
	}
	// send to invoice generator
	s.transactions <- transaction
	// get invoice id
	invoiceId := <-transaction.InvoiceIDChan
	close(transaction.InvoiceIDChan) // not needed anymore, so closing
	// construct and send response
	response := &pb.CreateResponse{
		Id: invoiceId,
	}

	return response, nil
}

func (s *invoicerService) ListAll(_ *emptypb.Empty, request pb.InvoicerService_ListAllServer) error {
	mu.Lock()
	responses := make([]*pb.InvoiceResponse, 0, len(invoices))
	for _, invoice := range invoices {
		data := make([]byte, len(invoice.Data))
		copy(data, invoice.Data)
		responses = append(responses, &pb.InvoiceResponse{
			Id:      invoice.ID,
			Status:  invoice.Status,
			Invoice: data,
		})
	}
	mu.Unlock()

	// stream responses
	for _, response := range responses {
		time.Sleep(1 * time.Second)
		if err := request.Send(response); err != nil {
			return err
		}
	}

	return nil
}

func (s *invoicerService) Get(ctx context.Context, request *pb.InvoiceRequest) (*pb.InvoiceResponse, error) {
	log.Println("received query for invoice")
	// construct response
	response := &pb.InvoiceResponse{}
	// fill response from the invoice data
	id := request.InvoiceId
	mu.Lock()
	defer mu.Unlock()
	invoice, ok := invoices[id]
	if !ok {
		return response, nil
	}
	response.Id = id
	response.Status = invoice.Status
	response.Invoice = make([]byte, len(invoice.Data))
	copy(response.Invoice, invoice.Data)

	return response, nil
}

func (s *invoicerService) invoiceGenerator() {
	// process each incoming transaction
	for transaction := range s.transactions {
		// generate a unique invoice id
		invoiceId := uuid.NewString()
		transaction.InvoiceIDChan <- invoiceId
		mu.Lock()
		invoices[invoiceId] = &Invoice{
			ID:     invoiceId,
			Status: pb.InvoiceStatusEnum_Pending,
		}
		mu.Unlock()
		go s.makeInvoice(invoiceId, transaction)
	}
}

func (s *invoicerService) makeInvoice(id string, t Transaction) {
	// simulate time to generate invoice
	time.Sleep(10 * time.Second)
	// adding randomness for failure/success cases
	if rand.Float64() < 0.3 {
		log.Printf("failed to generate invoice %s", id)
		// invoice generation failed
		mu.Lock()
		invoices[id].Status = pb.InvoiceStatusEnum_Failed
		mu.Unlock()
		return
	}

	invoiceTemplate := `
	########### DIGITAL INVOICE ##########
	--------------------------------------
	INVOICE ID: %s
	TO: %s
	AMOUNT: %d %s
	FOR SERVICE: %s
	--------------------------------------
	########### DIGITAL INVOICE ##########
	`
	invoiceData := []byte(fmt.Sprintf(invoiceTemplate, id, t.From, t.Amount, t.Currency, t.Service))

	mu.Lock()
	invoices[id].Status = pb.InvoiceStatusEnum_Success
	invoices[id].Data = invoiceData
	mu.Unlock()
	log.Printf("generated invoice %s", id)
}
