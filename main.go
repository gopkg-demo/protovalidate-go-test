package main

import (
	"fmt"
	"time"

	"github.com/bufbuild/protovalidate-go"
	pb "github.com/gopkg-demo/protovalidate-go-test/generated/protos/api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	msg := &pb.Transaction{
		Id:           1234,
		Price:        "$5.67",
		PurchaseDate: timestamppb.New(time.Now()),
		DeliveryDate: timestamppb.New(time.Now().Add(time.Hour)),
	}

	v, err := protovalidate.New()
	if err != nil {
		fmt.Println("failed to initialize validator:", err)
	}

	if err = v.Validate(msg); err != nil {
		fmt.Println("validation failed:", err)
	} else {
		fmt.Println("validation succeeded")
	}
}
