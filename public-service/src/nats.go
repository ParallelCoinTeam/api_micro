package main

import (
	"log"
	"os"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	nats "github.com/nats-io/go-nats"
)

func CreateNATSConnection() (*nats.Conn, error) {

	serverList := os.Getenv("NATS_SERVER")
	rootCACertFile := os.Getenv("NATS_CACERT")
	clientCertFile := os.Getenv("NATS_CERT")
	clientKeyFile := os.Getenv("NATS_KEY")

	// Connect options
	rootCA := nats.RootCAs(rootCACertFile)
	clientCert := nats.ClientCert(clientCertFile, clientKeyFile)
	alwaysReconnect := nats.MaxReconnects(-1)

	var nc *nats.Conn
	var err error
	for {
		nc, err = nats.Connect(serverList, rootCA, clientCert, alwaysReconnect)
		if err != nil {
			log.Printf("Error while connecting to NATS, backing off for a sec... (error: %s)", err)
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
	return nc, err
}
