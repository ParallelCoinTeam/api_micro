package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"

	testdata "github.com/syedomair/api_micro/testdata"
)

type testCaseType struct {
	method         string
	url            string
	pathParam      string
	requestBody    string
	responseResult string
	responseData   string
}

var testCases []testCaseType

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
		{"POST", kongProxyURL + "/v1/register", "", `{"first_name":"` + testdata.ValidFirstName + `", "last_name":"` + testdata.ValidLastName + `", "email":"` + testdata.ValidEmail + `", "password":"` + testdata.ValidPassword + `"}`, `"success"`, ``},
		{"POST", kongProxyURL + "/v1/authenticate", "", `{"email":"` + testdata.ValidEmail + `", "password":"` + testdata.ValidPassword + `"}`, `"success"`, ``},
		{"POST", kongProxyURL + "/v1/roles", "", `{"title":"` + testdata.RoleTitle1 + `","role_type":"` + testdata.RoleType + `"}`, `"success"`, ``},
		{"GET", kongProxyURL + "/v1/roles/", "role_id", ``, `"success"`, ``},
		{"GET", kongProxyURL + "/v1/roles", "", ``, `"success"`, ``},
		{"PATCH", kongProxyURL + "/v1/roles/", "role_id", `{"title":"` + testdata.RoleTitle2 + `","role_type":"` + testdata.RoleType + `"}`, `"success"`, ``},
		{"GET", kongProxyURL + "/v1/roles/", "role_id", ``, `"success"`, ``},
		{"DELETE", kongProxyURL + "/v1/roles/", "role_id", ``, `"success"`, ``},
		{"GET", kongProxyURL + "/v1/users/", "user_id", ``, `"success"`, ``},
		{"GET", kongProxyURL + "/v1/users", "", ``, `"success"`, ``},
		{"PATCH", kongProxyURL + "/v1/users/", "user_id", `{"first_name":"` + testdata.ValidFirstName + `"}`, `"success"`, ``},
		{"GET", kongProxyURL + "/v1/users/", "user_id", ``, `"success"`, ``},
		{"DELETE", kongProxyURL + "/v1/users/", "user_id", ``, `"success"`, ``},
	}
	i := 0
	userId := ""
	roleId := ""
	token := ""
	for _, testCase := range testCases {
		fmt.Println("---------------------------------------------------------------------------")
		url := ""
		if testCase.pathParam == "user_id" {
			url = testCase.url + userId
		} else if testCase.pathParam == "role_id" {
			url = testCase.url + roleId
		} else {
			url = testCase.url
		}
		req, err := http.NewRequest(testCase.method, url, strings.NewReader(testCase.requestBody))

		fmt.Println("method:", testCase.method)
		fmt.Println("url:", url)
		fmt.Println("body:", testCase.requestBody)
		fmt.Println("token:", token)
		if err != nil {
			print(err)
		}
		if i > 1 {
			req.Header.Set("apikey", token)
			req.Header.Set("authorization", token)
		} else {
			req.Header.Set("authorization", testdata.TestValidPublicToken)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			print(err)
		}

		body, _ := ioutil.ReadAll(resp.Body)

		var bodyInterface map[string]interface{}

		fmt.Println(string(body))
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
			userId = string(bytes.Trim(jsonUserId, `"`))
			fmt.Println("userId=", userId)
		}
		if i == 1 {
			var tokenInterface map[string]interface{}
			json.Unmarshal(jsonData, &tokenInterface)
			jsonToken, _ := json.Marshal(tokenInterface["token"])
			token = string(bytes.Trim(jsonToken, `"`))
			fmt.Println("token:", token)
		}
		if i == 2 {
			var roleIdInterface map[string]interface{}
			json.Unmarshal(jsonData, &roleIdInterface)
			jsonUserId, _ := json.Marshal(roleIdInterface["role_id"])
			roleId = string(bytes.Trim(jsonUserId, `"`))
			fmt.Println("roleId:", roleId)
		}
		time.Sleep(2 * time.Second)
		fmt.Println("---------------------------------------------------------------------------")
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
