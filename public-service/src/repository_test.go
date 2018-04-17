package main

import (
	"testing"
	"time"

	"github.com/syedomair/api_micro/common"
	pb "github.com/syedomair/api_micro/public-service/proto"
	testdata "github.com/syedomair/api_micro/testdata"
)

func TestPublicDB(t *testing.T) {

	db, _ := common.CreateDBConnection()
	repo := &PublicRepository{db, common.GetLogger()}
	defer repo.db.Close()

	start := time.Now()
	repo.logger.Log("METHOD", "TestPublicDB", "SPOT", "method start", "time_start", start)

	user := &pb.User{
		FirstName: testdata.ValidFirstName,
		LastName:  testdata.ValidLastName,
		Email:     testdata.ValidEmail,
		Password:  testdata.ValidPassword,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339)}

	userId, err := repo.Create(user, testdata.NetworkId)
	var expected error = nil
	if expected != err {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, err)
	}
	repo.logger.Log("METHOD", "TestPublicDB", "userId", userId)

	loginReq := &pb.LoginRequest{Email: testdata.ValidEmail, Password: testdata.ValidPassword}
	userResponse, err := repo.Authenticate(loginReq, testdata.NetworkId)
	expected = nil
	if expected != err {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, err)
	}
	expectedString := userId
	if expectedString != userResponse.Id {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, userResponse.Id)
	}

	if err = repo.db.Delete(&user).Error; err != nil {
		repo.logger.Log("METHOD", "TestPublicDB", "Error in deleting", err)
	}
	repo.logger.Log("METHOD", "TestPublicDB", "SPOT", "method end", "time_spent", time.Since(start))
}
