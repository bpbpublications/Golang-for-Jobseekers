package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	testcontainers "github.com/testcontainers/testcontainers-go"
)

func Test_nats_ops(t *testing.T) {
	req, err := testcontainers.GenericContainer(context.TODO(), testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "nats:2.1.9",
			Name:         "some-nats",
			ExposedPorts: []string{"4222/tcp"},
		},
		Started: true,
	})
	time.Sleep(2 * time.Second)
	defer req.Terminate(context.TODO())
	if err != nil {
		t.Fatalf("Unable to set nats environment. Err: %v", err)
	}

	port, err := req.MappedPort(context.TODO(), "4222")
	connectionString := fmt.Sprintf("nats://localhost:%v", port)

	queueNats, err := NewNats(connectionString, "details")
	if err != nil {
		t.Fatalf("Unable to achieve connection to nats. ConnectionString: %v, Err: %v", connectionString, err)
	}

	testingString := "This is a test"

	err = queueNats.Add(context.TODO(), []byte(testingString))
	if err != nil {
		t.Errorf("Expected no errors from attempting to send message. Err: %v", err)
	}

	resp, err := queueNats.Pop(context.TODO())
	if err != nil {
		t.Errorf("Expected no errors from attempting to receive message. Err: %v", err)
	}
	if string(resp) != testingString {
		t.Errorf("Expected %v but received '%v'", testingString, string(resp))
	}
}
