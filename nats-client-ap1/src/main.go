package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/go-nats"
	"github.com/nats-io/nuid"
)

func main() {
	var (
		serverList     string
		rootCACertFile string
		clientCertFile string
		clientKeyFile  string
	)
	flag.StringVar(&serverList, "s", "tls://nats-1.nats-cluster.default.svc:4222", "List of NATS of servers available")
	flag.StringVar(&rootCACertFile, "cacert", "./certs/ca.pem", "Root CA Certificate File")
	flag.StringVar(&clientCertFile, "cert", "./certs/client.pem", "Client Certificate File")
	flag.StringVar(&clientKeyFile, "key", "./certs/client-key.pem", "Client Private key")
	flag.Parse()

	log.Println("KHALID server starts ")
	log.Println("NATS endpoint:", serverList)
	log.Println("Root CA:", rootCACertFile)
	log.Println("Client Cert:", clientCertFile)
	log.Println("Client Key:", clientKeyFile)

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

	nc.Subscribe("discovery.*.status", func(m *nats.Msg) {
		log.Printf("[Received on %q] %s", m.Subject, string(m.Data))
	})

	discoverySubject := fmt.Sprintf("discovery.%s.status", nuid.Next())
	info := struct {
		InMsgs        uint64   `json:"in_msgs"`
		OutMsgs       uint64   `json:"out_msgs"`
		Reconnects    uint64   `json:"reconnects"`
		CurrentServer string   `json:"current_server"`
		Servers       []string `json:"servers"`
	}{}

	for range time.NewTicker(1 * time.Second).C {
		stats := nc.Stats()
		info.InMsgs = stats.InMsgs
		info.OutMsgs = stats.OutMsgs
		info.Reconnects = stats.Reconnects
		info.CurrentServer = nc.ConnectedUrl()
		info.Servers = nc.Servers()
		payload, err := json.Marshal(info)
		if err != nil {
			log.Printf("Error marshalling data: %s", err)
		}
		err = nc.Publish(discoverySubject, payload)
		if err != nil {
			log.Printf("Error during publishing: %s", err)
		}
		nc.Flush()
	}
}
