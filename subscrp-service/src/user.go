package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gogo/protobuf/proto"
	nats "github.com/nats-io/go-nats"
	pb "github.com/syedomair/api_micro/public-service/proto"
)

func HandleUserRegister(m *nats.Msg) {
	log.Printf("[Received on %q] %s", m.Subject, string(m.Data))
	userMsg := pb.UserMessage{}
	err := proto.Unmarshal(m.Data, &userMsg)
	if err != nil {
		fmt.Println(err)
	}
	data := url.Values{}
	data.Set("username", userMsg.UserId)
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
}

func HandleUserLogin(m *nats.Msg) {
	log.Printf("[Received on %q] %s", m.Subject, string(m.Data))
	userMsg := pb.UserTokenMessage{}
	err := proto.Unmarshal(m.Data, &userMsg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Khalid")
	fmt.Println(usrMsg.UserId)
	fmt.Println("Khalid1")
	/*
		data := url.Values{}
		data.Set("username", userMsg.UserId)
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
	*/
}
