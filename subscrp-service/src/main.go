package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/gogo/protobuf/proto"
	nats "github.com/nats-io/go-nats"
	pb "github.com/syedomair/api_micro/public-service/proto"
)

func main() {
	serverList := os.Getenv("NATS_SERVER")
	rootCACertFile := os.Getenv("NATS_CACERT")
	clientCertFile := os.Getenv("NATS_CERT")
	clientKeyFile := os.Getenv("NATS_KEY")
	// Connect options
	rootCA := nats.RootCAs(rootCACertFile)
	clientCert := nats.ClientCert(clientCertFile, clientKeyFile)
	alwaysReconnect := nats.MaxReconnects(-1)

	var nc *nats.Conn
	var err1 error
	for {
		fmt.Println("for loop")
		nc, err1 = nats.Connect(serverList, rootCA, clientCert, alwaysReconnect)
		if err1 != nil {
			log.Printf("Error while connecting to NATS, backing off for a sec... (error: %s)", err1)
			time.Sleep(1 * time.Second)
			continue
		}
		fmt.Println("Connected NATS")
		break
	}

	subject := "User.UserCreated"
	nc.Subscribe(subject, func(m *nats.Msg) {
		log.Printf("[Received on %q] %s", m.Subject, string(m.Data))
		userMsg := pb.UserMessage{}
		err := proto.Unmarshal(m.Data, &userMsg)
		if err != nil {
			fmt.Println("Khalid")
			fmt.Println(err)
			fmt.Println("Omair")
		}

		fmt.Println("Khalid1")
		fmt.Println(userMsg)
		fmt.Println("Omair1")

		data := url.Values{}
		data.Set("username", "omair511")
		request, err := http.NewRequest("POST", "http://kong-admin:8001/consumers", strings.NewReader(data.Encode()))
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		fmt.Println(response.Status)

		body, _ := ioutil.ReadAll(response.Body)

		var bodyInterface map[string]interface{}
		json.Unmarshal(body, &bodyInterface)
		//fmt.Println(string(bodyInterface))
	})
	runtime.Goexit()
}
