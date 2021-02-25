package main

import (
	"fmt"
	"os"

	"kmodules.xyz/auditor/pkg/controller/receiver"

	"github.com/nats-io/nats.go"
)

func main() {
	if len(os.Args) < 4 {
		panic("argument missing")
	}

	server := os.Args[1]
	subject := os.Args[2]
	creds := os.Args[3]

	fmt.Println(server, subject, creds)

	nc, err := getNatsClient(server, creds)
	if err != nil {
		panic(err)
	}

	if err = receiver.PublishEvent(nc, subject, "test", "Hello there"); err != nil {
		panic(err)
	}

	//if _, err = nc.Request(subject, []byte("Hello there"), time.Second*5); err != nil {
	//	panic(err)
	//}
}

func getNatsClient(server, creds string) (*nats.Conn, error) {
	var natsOpts = []nats.Option{nats.Name("Auditor")}

	natsOpts = append(natsOpts, nats.UserCredentials(creds))
	nc, err := nats.Connect(server, natsOpts...)
	if err != nil {
		return nil, err
	}

	return nc, nil
}
