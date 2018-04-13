package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/syedomair/api_micro/common"
	pb "github.com/syedomair/api_micro/public-service/proto"
)

func TestRepoCreate(t *testing.T) {

	/**/
	db, err := common.CreateDBConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	} else {
		fmt.Println("Connected to DB")
	}

	logger := common.GetLogger()

	repo := &PublicRepository{db, logger}
	user := &pb.User{Id: "04b58e6e-f910-4ff0-83f1-27fbfa85dc2f", NetworkId: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11", FirstName: "First Name 1", LastName: "Last Name 1", Email: "email1@gmail.com", Password: "123", IsAdmin: "1", CreatedAt: time.Now().Format(time.RFC3339), UpdatedAt: time.Now().Format(time.RFC3339)}

	userId, err := repo.Create(user, "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	logger.Log("METHOD", "TestRepoCreate", "userId", userId)
	/**/
}
