package main

import (
	"time"

	"golang.org/x/net/context"

	log "github.com/go-kit/kit/log"
	"github.com/syedomair/api_micro/common"
	pb "github.com/syedomair/api_micro/users-service/proto"
)

type Env struct {
	repo   Repository
	nats   Nats
	logger log.Logger
}

func (env *Env) Create(ctx context.Context, req *pb.User) (*pb.Response, error) {

	start := time.Now()
	env.logger.Log("METHOD", "Create", "SPOT", "method start", "time_start", start)
	networkId, _ := ctx.Value("network_id").(envtring)

	userId, err := env.repo.Create(req, networkId)
	if err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.DatabaseError()}, nil
	}
	//go publishOrderCreated(req)
	responseUserId := map[string]string{"user_id": userId}
	env.logger.Log("METHOD", "Create", "SPOT", "method end", "time_spent", time.Since(start))
	return &pb.Response{Result: common.SUCCESS, Data: responseUserId, Error: nil}, err
}

func (env *Env) GetAll(ctx context.Context, req *pb.RequestQuery) (*pb.ResponseList, error) {

	start := time.Now()
	env.logger.Log("METHOD", "GetAll", "SPOT", "method start", "time_start", start)
	networkId, _ := ctx.Value("network_id").(envtring)

	limit, offset, orderby, sort, err := common.ValidateQueryString(req.Limit, "3", req.Offset, "0", req.Orderby, "title", req.Sort, "asc")
	if err != nil {
		return &pb.ResponseList{Result: common.FAILURE, Error: common.CommonError(err.Error()), Data: nil}, nil
	}

	users, count, _ := env.repo.GetAll(limit, offset, orderby, sort, networkId)
	if err != nil {
		return &pb.ResponseList{Result: common.FAILURE, Error: common.CommonError(err.Error()), Data: nil}, nil
	}

	userList := &pb.UserList{Offset: offset, Limit: limit, Count: count, List: users}
	env.logger.Log("METHOD", "GetAll", "SPOT", "method end", "time_spent", time.Since(start))
	return &pb.ResponseList{Result: common.SUCCESS, Error: nil, Data: userList}, nil
}

func (env *Env) GetUser(ctx context.Context, req *pb.User) (*pb.ResponseUser, error) {

	start := time.Now()
	env.logger.Log("METHOD", "GetUser", "SPOT", "method start", "time_start", start)
	networkId, _ := ctx.Value("network_id").(envtring)

	user, err := env.repo.Get(req.Id, networkId)
	if err != nil {
		return &pb.ResponseUser{Result: common.FAILURE, Data: nil, Error: common.CommonError(err.Error())}, nil
	}
	env.logger.Log("METHOD", "GetUser", "SPOT", "method end", "time_spent", time.Since(start))
	return &pb.ResponseUser{Result: common.SUCCESS, Data: user, Error: nil}, nil
}

func (env *Env) UpdateUser(ctx context.Context, req *pb.User) (*pb.Response, error) {

	start := time.Now()
	env.logger.Log("METHOD", "UpdateUser", "SPOT", "method start", "time_start", start)
	networkId, _ := ctx.Value("network_id").(envtring)

	err := env.repo.Update(req, networkId)
	if err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.CommonError(err.Error())}, nil
	}
	responseUserId := map[string]string{"user_id": req.Id}
	env.logger.Log("METHOD", "UpdateUse", "SPOT", "method end", "time_spent", time.Since(start))
	return &pb.Response{Result: common.SUCCESS, Data: responseUserId, Error: nil}, nil
}

func (env *Env) DeleteUser(ctx context.Context, req *pb.User) (*pb.Response, error) {

	start := time.Now()
	env.logger.Log("METHOD", "DeleteUser", "SPOT", "method start", "time_start", start)
	networkId, _ := ctx.Value("network_id").(string)

	err := env.repo.Delete(req, networkId)
	if err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.CommonError(err.Error())}, nil
	}
	responseUserId := map[string]string{"user_id": req.Id}
	env.logger.Log("METHOD", "DeleteUser", "SPOT", "method end", "time_spent", time.Since(start))
	return &pb.Response{Result: common.SUCCESS, Data: responseUserId, Error: nil}, nil
}
