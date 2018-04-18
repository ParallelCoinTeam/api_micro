package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	testdata "github.com/syedomair/api_micro/testdata"
)

func main() {

	kongAdminURL, _ := minikubeServiceURL("kong-admin")
	kongProxyURL, _ := minikubeServiceURL("kong-proxy")
	publicURL, _ := minikubeServiceURL("public-srvc")
	userURL, _ := minikubeServiceURL("users-srvc")
	roleURL, _ := minikubeServiceURL("roles-srvc")

	fmt.Println("kong-admin:", kongAdminURL)
	fmt.Println("kong-proxy:", kongProxyURL)
	fmt.Println("public-srvc:", publicURL)
	fmt.Println("users-srvc:", userURL)
	fmt.Println("roles-srvc:", roleURL)
	testCases = []testCaseType{
		{"POST", publicURL + "/v1/register", `{"first_name":"` + testdata.ValidFirstName + `", "last_name":"` + testdata.ValidLastName + `", "email":"` + testdata.ValidEmail + `", "password":"` + testdata.ValidPassword + `"}`, `"success"`, ``},
		{"POST", publicURL + "/v1/authenticate", `{"email":"` + testdata.ValidEmail + `", "password":"` + testdata.ValidPassword + `"}`, `"success"`, ``},
		{"DELETE", userURL + "/v1/users/", `{"email":"` + testdata.ValidEmail + `", "password":"` + testdata.ValidPassword + `"}`, `"success"`, ``},
	}
	i := 0
	userId := ""
	for _, testCase := range testCases {
		req, err := http.NewRequest(testCase.method, testCase.url, strings.NewReader(testCase.requestBody))

		fmt.Println("method:", testCase.method)
		fmt.Println("url:", testCase.url)
		fmt.Println("body:", testCase.requestBody)
		if err != nil {
			print(err)
		}
		req.Header.Set("authorization", testdata.TestValidPublicToken)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			print(err)
		}

		//resp, _ := http.DefaultClient.Do(req)
		body, _ := ioutil.ReadAll(resp.Body)

		var bodyInterface map[string]interface{}
		json.Unmarshal(body, &bodyInterface)
		jsonData, _ := json.Marshal(bodyInterface["data"])
		jsonResult, _ := json.Marshal(bodyInterface["result"])
		fmt.Println(strconv.Itoa(i) + " " + testCase.method + " " + string(testCase.url))
		fmt.Println(string(jsonData))
		fmt.Println(string(jsonResult))
		if i == 0 {
			var userIdInterface map[string]interface{}
			json.Unmarshal(jsonData, &userIdInterface)
			jsonUserId, _ := json.Marshal(userIdInterface["user_id"])
			fmt.Println(string(jsonUserId))
			userId = string(jsonUserId)

			//jsonUserId, _ := json.Marshal(userIdInterface["user_id"])
			//fmt.Println(string(jsonUserId))
			//fmt.Println(bodyInterface)
			//var userIdInterface map[string]interface{}
			//userIdInterface = bodyInterface["data"]
		}

		/*
			if string(jsonData) != testCase.responseData {
				if testCase.responseData != "list" {
					t.Error("Expected:" + testCase.responseData + " got:" + string(jsonData))
				}
			}
			if string(jsonResult) != testCase.responseResult {
				t.Error("Expected:" + testCase.responseResult + " got:" + string(jsonResult))
			}
		*/
		i++
	}
}

func minikubeServiceURL(serviceName string) (string, error) {
	minikube := "minikube"
	service := "service"
	url := "--url"

	serviceURL, err := exec.Command(minikube, service, serviceName, url).Output()
	if err != nil {
		return "", err
	}
	return strings.Trim(string(serviceURL), "\n"), nil
}

type testCaseType struct {
	method         string
	url            string
	requestBody    string
	responseResult string
	responseData   string
}

var testCases []testCaseType

func TestServer() {

	//srv := httptest.NewServer(sh)
	//defer srv.Close()
	/*
	   	testCases = []testCaseType{
	   		{"GET", srv.URL + "/user/1", ``, `"success"`, `{"email":"john@gmail.com","first_name":"John","id":1,"last_name":"Smith"}`},
	   		{"POST", srv.URL + "/book", `{"name":"Test Book changed","description":"Desc","publish":true, "user_id":1}`, `"success"`, `null`},
	   		{"PATCH", srv.URL + "/book/1", `{"name":"Test Book changed","description":"Desc changed","publish":true, "user_id":1}`, `"success"`, `null`},
	   		{"GET", srv.URL + "/book/1", ``, `"success"`, `{"book_name":"Test Book changed","description":"Desc changed","first_name":"John","id":1,"last_name":"Smith","publish":true,"user_id":1}`},
	   		{"GET", srv.URL + "/my-books/1", ``, `"success"`, `list`},
	   		{"GET", srv.URL + "/books", ``, `"success"`, `list`},
	   		{"GET", srv.URL + "/public/books", ``, `"success"`, `list`},
	   		{"POST", srv.URL + "/book", `{"name":"","description":"Desc","publish":true, "user_id":1}`, `"error"`, `"name is a requird field"`},
	   		{"POST", srv.URL + "/book", `{"name":"Test Book","description":"Desc","publish":true}`, `"error"`, `"user_id is a requird field"`},
	   	}

	   	commonTest(testCases, t, "dHb%e@Bg0f8-API_KEY-&bE71jKoH=2", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG5AZ21haWwuY29tIiwicGFzc3dvcmQiOiJNVEl6TkRVPSIsImlzcyI6InRlc3QifQ.Fk0pm1AmEl7nNGH_7Xcs93r5V2nhPGnBCfWbIwkhTHk")
	   }
	   func commonTest(testCases []testCaseType, t *testing.T, xkey string, xjwt string) {

	   	i := 0
	   	for _, testCase := range testCases {
	   		req, _ := http.NewRequest(testCase.method, testCase.url, strings.NewReader(testCase.requestBody))
	   		if xkey != "" {
	   			req.Header.Set("x-key", xkey)
	   		}
	   		if xjwt != "" {
	   			req.Header.Set("x-jwt", xjwt)
	   		}
	   		resp, _ := http.DefaultClient.Do(req)
	   		body, _ := ioutil.ReadAll(resp.Body)

	   		var bodyInterface map[string]interface{}
	   		json.Unmarshal(body, &bodyInterface)
	   		jsonData, _ := json.Marshal(bodyInterface["data"])
	   		jsonResult, _ := json.Marshal(bodyInterface["result"])
	   		fmt.Println(strconv.Itoa(i) + " " + testCase.method + " " + string(testCase.url))
	   		//fmt.Println(string(jsonData))
	   		//fmt.Println(string(jsonResult))

	   		if string(jsonData) != testCase.responseData {
	   			if testCase.responseData != "list" {
	   				t.Error("Expected:" + testCase.responseData + " got:" + string(jsonData))
	   			}
	   		}
	   		if string(jsonResult) != testCase.responseResult {
	   			t.Error("Expected:" + testCase.responseResult + " got:" + string(jsonResult))
	   		}
	   		i++
	   	}
	*/
}
