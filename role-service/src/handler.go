package main

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/syedomair/api_micro/common"
	pb "github.com/syedomair/api_micro/role-service/proto"
	"google.golang.org/grpc/metadata"
)

const (
	port      = ":50051"
	aggregate = "Order"
	event     = "OrderCreated"
)

type service struct {
	repo Repository
}

func (s *service) Create(ctx context.Context, req *pb.Role) (*pb.Response, error) {
	meta, _ := metadata.FromIncomingContext(ctx)
	fmt.Println(meta["authorization"][0])

	roleId, err := s.repo.Create(req, "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	if err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.DatabaseError()}, nil
	}
	//go publishOrderCreated(req)
	responseRoleId := map[string]string{"role_id": roleId}
	return &pb.Response{Result: common.SUCCESS, Data: responseRoleId, Error: nil}, err
}

func (s *service) GetAll(ctx context.Context, req *pb.RequestQuery) (*pb.ResponseList, error) {

	limit, offset, orderby, sort, err := common.ValidateQueryString(req.Limit, "3", req.Offset, "0", req.Orderby, "title", req.Sort, "asc")
	if err != nil {
		return &pb.ResponseList{Result: common.FAILURE, Error: common.CommonError(err.Error()), Data: nil}, nil
	}

	roles, count, _ := s.repo.GetAll(limit, offset, orderby, sort, "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	if err != nil {
		return &pb.ResponseList{Result: common.FAILURE, Error: common.CommonError(err.Error()), Data: nil}, nil
	}

	roleList := &pb.RoleList{Offset: offset, Limit: limit, Count: count, List: roles}
	return &pb.ResponseList{Result: common.SUCCESS, Error: nil, Data: roleList}, nil
}

/*
// publishOrderCreated publish an event via NATS server
func publishOrderCreated(order *pb.Role) {
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
func (s *service) GetRole(ctx context.Context, req *pb.Role) (*pb.ResponseRole, error) {

	role, err := s.repo.Get(req.Id, "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	if err != nil {
		return &pb.ResponseRole{Result: common.FAILURE, Data: nil, Error: common.CommonError(err.Error())}, nil
	}
	return &pb.ResponseRole{Result: common.SUCCESS, Data: role, Error: nil}, nil
}

func (s *service) UpdateRole(ctx context.Context, req *pb.Role) (*pb.Response, error) {

	err := s.repo.Update(req, "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	if err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.CommonError(err.Error())}, nil
	}
	responseRoleId := map[string]string{"role_id": req.Id}
	return &pb.Response{Result: common.SUCCESS, Data: responseRoleId, Error: nil}, nil
}

func (s *service) DeleteRole(ctx context.Context, req *pb.Role) (*pb.Response, error) {

	err := s.repo.Delete(req, "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	if err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.CommonError(err.Error())}, nil
	}
	responseRoleId := map[string]string{"role_id": req.Id}
	return &pb.Response{Result: common.SUCCESS, Data: responseRoleId, Error: nil}, nil
}
