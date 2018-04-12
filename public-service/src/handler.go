package main

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/go-kit/kit/log"
	"google.golang.org/grpc/metadata"

	"golang.org/x/net/context"

	common "github.com/syedomair/api_micro/common"
	pb "github.com/syedomair/api_micro/public-service/proto"
)

type Env struct {
	repo   Repository
	nats   Nats
	logger log.Logger
}

func (env *Env) Register(ctx context.Context, req *pb.User) (*pb.Response, error) {

	env.logger.Log("METHOD", "Register", "SPOT", "method start")
	start := time.Now()
	meta, _ := metadata.FromIncomingContext(ctx)
	apiKey, err := common.GetAPIKey(meta["authorization"][0])
	if err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.DatabaseError()}, nil
	}

	network, err := env.repo.GetNetworkFromApiKey(apiKey)
	if err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.DatabaseError()}, nil
	}
	_, err = common.ValidateJWTToken(meta["authorization"][0], network.Secret)
	if err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.DatabaseError()}, nil
	}
	networkId := network.Id

	userId, err := env.repo.Create(req, networkId)
	if err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.DatabaseError()}, nil
	}
	responseUserId := map[string]string{"user_id": userId}

	env.logger.Log("METHOD", "Register", "SPOT", "before NATS event", "time_spent", time.Since(start))
	/**/
	//NATS Event Publish
	go func() {
		err = env.nats.PublishRegisterEvent(userId, networkId)
		if err != nil {
			env.logger.Log("Error during publishing: ", err)
		}
	}()
	/**/
	env.logger.Log("METHOD", "Register", "SPOT", "after NATS event", "time_spent", time.Since(start))
	env.logger.Log("METHOD", "Register", "SPOT", "method end")
	return &pb.Response{Result: common.SUCCESS, Data: responseUserId, Error: nil}, err
}

func (env *Env) Authenticate(ctx context.Context, req *pb.LoginRequest) (*pb.Response, error) {

	env.logger.Log("METHOD", "Authenticate", "SPOT", "method start")
	meta, _ := metadata.FromIncomingContext(ctx)
	apiKey, err := common.GetAPIKey(meta["authorization"][0])
	if err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.DatabaseError()}, nil
	}

	network, err := env.repo.GetNetworkFromApiKey(apiKey)
	if err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.DatabaseError()}, nil
	}
	_, err = common.ValidateJWTToken(meta["authorization"][0], network.Secret)
	if err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.DatabaseError()}, nil
	}
	networkId := network.Id

	user, err := env.repo.Authenticate(req, networkId)
	if err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.DatabaseError()}, nil
	}

	signedJwtToken := createUserToken(user)

	//NATS Event Publish
	go func() {
		err = env.nats.PublishAuthEvent(user.Id, signedJwtToken)
		if err != nil {
			env.logger.Log("Error during publishing: ", err)
		}
	}()

	tokenStr := map[string]string{"token": signedJwtToken}
	env.logger.Log("METHOD", "Authenticate", "SPOT", "method end")
	return &pb.Response{Result: common.SUCCESS, Data: tokenStr, Error: nil}, nil
}

func createUserToken(user *pb.User) string {
	type Claims struct {
		CurrentUserId string `json:"current_user_id"`
		NetworkId     string `json:"network_id"`
		IsAdmin       string `json:"is_admin"`
		jwt.StandardClaims
	}

	claims := Claims{
		user.Id,
		user.NetworkId,
		user.IsAdmin,
		jwt.StandardClaims{
			Issuer: "MEEM",
		},
	}
	signingKey := []byte(common.SIGNING_KEY)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedJwtToken, _ := token.SignedString(signingKey)
	return signedJwtToken
}
