package main

import (
	"time"

	"golang.org/x/net/context"

	log "github.com/go-kit/kit/log"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/syedomair/api_micro/common"
	pb "github.com/syedomair/api_micro/users-service/proto"
)

type Env struct {
	repo   Repository
	nats   Nats
	logger log.Logger
}

func (env *Env) GetAll(ctx context.Context, req *pb.RequestQuery) (*pb.ResponseList, error) {

	start := time.Now()
	env.logger.Log("METHOD", "GetAll", "SPOT", "method start", "time_start", start)
	networkId, _ := ctx.Value("network_id").(string)

	limit, offset, orderby, sort, err := common.ValidateQueryString(req.Limit, "3", req.Offset, "0", req.Orderby, "created_at", req.Sort, "asc")
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
	networkId, _ := ctx.Value("network_id").(string)

	if err := validateUserId(req); err != nil {
		return &pb.ResponseUser{Result: common.FAILURE, Data: nil, Error: common.ErrorMessage("2004", err.Error())}, nil
	}
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
	networkId, _ := ctx.Value("network_id").(string)

	if err := validateUserId(req); err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.ErrorMessage("2004", err.Error())}, nil
	}

	if err := validateParameters(req); err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.ErrorMessage("2004", err.Error())}, nil
	}
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

	if err := validateUserId(req); err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.ErrorMessage("2004", err.Error())}, nil
	}
	err := env.repo.Delete(req, networkId)
	if err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.CommonError(err.Error())}, nil
	}
	responseUserId := map[string]string{"user_id": req.Id}
	env.logger.Log("METHOD", "DeleteUser", "SPOT", "method end", "time_spent", time.Since(start))
	return &pb.Response{Result: common.SUCCESS, Data: responseUserId, Error: nil}, nil
}

func validateParameters(role *pb.User) error {
	if err := validation.Validate(
		role.FirstName,
		validation.Required.Error("first_name is a required field"),
		validation.Length(1, 64).Error("first_name is a rqquired field with the max character of 32")); err != nil {
		return err
	}
	if err := validation.Validate(
		role.LastName,
		validation.Required.Error("last_name is a required field"),
		validation.Length(1, 64).Error("last_name is a rqquired field with the max character of 32")); err != nil {
		return err
	}
	return nil
}
func validateUserId(user *pb.User) error {
	if err := validation.Validate(
		user.Id,
		validation.Required.Error("user_id is a required field"),
		is.UUIDv4.Error("invalid user_id.")); err != nil {
		return err
	}
	return nil
}
