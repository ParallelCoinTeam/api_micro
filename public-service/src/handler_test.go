package main

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	gokitlog "github.com/go-kit/kit/log"
	pb "github.com/syedomair/api_micro/public-service/proto"
	"google.golang.org/grpc/metadata"
)

type mockNATS struct {
}

func (mnats *mockNATS) PublishRegisterEvent(userId string, networkId string) error {
	fmt.Println("Publish Event")
	return nil
}
func (mnats *mockNATS) PublishAuthEvent(userId string, networkId string) error {
	fmt.Println("Publish Event")
	return nil
}

type mockDB struct {
	users []*pb.User
}

func (mdb *mockDB) initMockDb() ([]*pb.User, error) {

	users := make([]*pb.User, 0)
	users = append(users, &pb.User{Id: "04b58e6e-f910-4ff0-83f1-27fbfa85dc2f", NetworkId: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11", FirstName: "First Name 1", LastName: "Last Name 1", Email: "email1@gmail.com", Password: "123", IsAdmin: "1", CreatedAt: time.Now().Format(time.RFC3339), UpdatedAt: time.Now().Format(time.RFC3339)})
	users = append(users, &pb.User{Id: "04b58e6e-f910-4ff0-83f1-27fbfa85dc2f", NetworkId: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11", FirstName: "First Name 2", LastName: "Last Name 2", Email: "email2@gmail.com", Password: "123", IsAdmin: "1", CreatedAt: time.Now().Format(time.RFC3339), UpdatedAt: time.Now().Format(time.RFC3339)})
	users = append(users, &pb.User{Id: "04b58e6e-f910-4ff0-83f1-27fbfa85dc2f", NetworkId: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11", FirstName: "First Name 3", LastName: "Last Name 3", Email: "email3@gmail.com", Password: "123", IsAdmin: "1", CreatedAt: time.Now().Format(time.RFC3339), UpdatedAt: time.Now().Format(time.RFC3339)})

	return users, nil
}

func (mdb *mockDB) Create(user *pb.User, networkId string) (string, error) {
	//mdb.users, _ = mdb.initMockDb()
	return user.Id, nil
}
func (mdb *mockDB) Authenticate(user *pb.LoginRequest, networkId string) (*pb.User, error) {
	//mdb.users, _ = mdb.initMockDb()
	userObj := &pb.User{Id: "04b58e6e-f910-4ff0-83f1-27fbfa85dc2f", NetworkId: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11", FirstName: "First Name 1", LastName: "Last Name 1", Email: "email1@gmail.com", Password: "123", IsAdmin: "1", CreatedAt: time.Now().Format(time.RFC3339), UpdatedAt: time.Now().Format(time.RFC3339)}
	return userObj, nil
}
func (mdb *mockDB) GetNetworkFromApiKey(apiKey string) (*pb.Network, error) {
	network := &pb.Network{Id: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11", Name: "Network Name", ApiKey: "the$network#api*key", Secret: "the$network#api*secret", Status: "1", CreatedAt: time.Now().Format(time.RFC3339), UpdatedAt: time.Now().Format(time.RFC3339)}
	return network, nil
}

func TestCreate(t *testing.T) {
	var logger gokitlog.Logger
	{
		logger = gokitlog.NewLogfmtLogger(os.Stdout)
		logger = gokitlog.With(logger, "TIME", gokitlog.DefaultTimestamp)
		logger = gokitlog.With(logger, "CALLER", gokitlog.DefaultCaller)
	}
	md := metadata.New(map[string]string{"authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcGlfa2V5IjoidGhlJG5ldHdvcmsjYXBpKmtleSIsImlzcyI6Ik1FRU0ifQ.TAFZabSWpnmmXThkRZ1FIQZvRKzESL4jER2dj_h30oc"})

	ctx := metadata.NewIncomingContext(context.Background(), md)

	srv := Env{repo: &mockDB{}, nats: &mockNATS{}, logger: logger}

	user := &pb.User{Id: "04b58e6e-f910-4ff0-83f1-27fbfa85dc2f", NetworkId: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11", FirstName: "First Name 3", LastName: "Last Name 3", Email: "email3@gmail.com", Password: "123", IsAdmin: "1", CreatedAt: time.Now().Format(time.RFC3339), UpdatedAt: time.Now().Format(time.RFC3339)}
	response, _ := srv.Register(ctx, user)

	fmt.Println(response)
	fmt.Println(response.Result)
	expected := "success"
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
}
