package main

import (
	"golang.org/x/net/context"

	"github.com/syedomair/api_micro/common"
	pb "github.com/syedomair/api_micro/user-service/proto"
)

const (
	port      = ":50051"
	aggregate = "Order"
	event     = "OrderCreated"
)

type service struct {
	repo Repository
}

func (s *service) Create(ctx context.Context, req *pb.User) (*pb.Response, error) {

	networkId, _ := ctx.Value("network_id").(string)

	userId, err := s.repo.Create(req, networkId)
	if err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.DatabaseError()}, nil
	}
	//go publishOrderCreated(req)
	responseUserId := map[string]string{"user_id": userId}
	return &pb.Response{Result: common.SUCCESS, Data: responseUserId, Error: nil}, err
}

func (s *service) GetAll(ctx context.Context, req *pb.RequestQuery) (*pb.ResponseList, error) {

	networkId, _ := ctx.Value("network_id").(string)

	limit, offset, orderby, sort, err := common.ValidateQueryString(req.Limit, "3", req.Offset, "0", req.Orderby, "title", req.Sort, "asc")
	if err != nil {
		return &pb.ResponseList{Result: common.FAILURE, Error: common.CommonError(err.Error()), Data: nil}, nil
	}

	users, count, _ := s.repo.GetAll(limit, offset, orderby, sort, networkId)
	if err != nil {
		return &pb.ResponseList{Result: common.FAILURE, Error: common.CommonError(err.Error()), Data: nil}, nil
	}

	userList := &pb.UserList{Offset: offset, Limit: limit, Count: count, List: users}
	return &pb.ResponseList{Result: common.SUCCESS, Error: nil, Data: userList}, nil
}

/*
// publishOrderCreated publish an event via NATS server
func publishOrderCreated(order *pb.User) {
	// Connect to NATS server
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)
	defer natsConnection.Close()
	eventData, _ := json.Marshal(order)
	event := pbO.EventStore{
		AggregateId:   order.Id,
		AggregateType: aggregate,
		EventId:       uuid.NewV4().String(),
		EventType:     event,
		EventData:     string(eventData),
	}
	subject := "Order.OrderCreated"
	data, _ := proto.Marshal(&event)
	// Publish message on subject
	natsConnection.Publish(subject, data)
	log.Println("Published message on subject " + subject)
}
*/
func (s *service) GetUser(ctx context.Context, req *pb.User) (*pb.ResponseUser, error) {

	networkId, _ := ctx.Value("network_id").(string)

	user, err := s.repo.Get(req.Id, networkId)
	if err != nil {
		return &pb.ResponseUser{Result: common.FAILURE, Data: nil, Error: common.CommonError(err.Error())}, nil
	}
	return &pb.ResponseUser{Result: common.SUCCESS, Data: user, Error: nil}, nil
}

func (s *service) UpdateUser(ctx context.Context, req *pb.User) (*pb.Response, error) {

	networkId, _ := ctx.Value("network_id").(string)

	err := s.repo.Update(req, networkId)
	if err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.CommonError(err.Error())}, nil
	}
	responseUserId := map[string]string{"user_id": req.Id}
	return &pb.Response{Result: common.SUCCESS, Data: responseUserId, Error: nil}, nil
}

func (s *service) DeleteUser(ctx context.Context, req *pb.User) (*pb.Response, error) {

	networkId, _ := ctx.Value("network_id").(string)

	err := s.repo.Delete(req, networkId)
	if err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.CommonError(err.Error())}, nil
	}
	responseUserId := map[string]string{"user_id": req.Id}
	return &pb.Response{Result: common.SUCCESS, Data: responseUserId, Error: nil}, nil
}
