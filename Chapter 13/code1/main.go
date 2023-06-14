package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/xxx/ticketing"

	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTimeout(3*time.Second))
	conn, _ := grpc.Dial(fmt.Sprintf("%v:%v", "localhost", "8888"), opts...)
	for {
		getCustomerDetails(conn)
		time.Sleep(3 * time.Second)
	}
}

func getCustomerDetails(conn *grpc.ClientConn) {
	client := ticketing.NewCustomerControllerClient(conn)
	log.Println("Start GetCustomerDetails")
	defer log.Println("End GetCustomerDetails")
	zz, err := client.GetCustomer(context.Background(), &ticketing.GetCustomerRequest{})
	if err != nil {
		fmt.Println(err)
	}
	log.Println(zz)
}
