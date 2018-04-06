package main

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/metadata"

	"golang.org/x/net/context"

	jwt "github.com/dgrijalva/jwt-go"
	nats "github.com/nats-io/go-nats"
	common "github.com/syedomair/api_micro/common"
	pb "github.com/syedomair/api_micro/public-service/proto"
)

type service struct {
	repo Repository
	nats *nats.Conn
}

func (s *service) Register(ctx context.Context, req *pb.User) (*pb.Response, error) {

	meta, _ := metadata.FromIncomingContext(ctx)
	apiKey, err := common.GetAPIKey(meta["authorization"][0])
	if err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.DatabaseError()}, nil
	}

	network, err := s.repo.GetNetworkFromApiKey(apiKey)
	if err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.DatabaseError()}, nil
	}
	fmt.Println(network)
	fmt.Println(network.Id)
	token, err := ValidateJWTToken(meta["authorization"][0], network.Secret)
	if err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.DatabaseError()}, nil
	}
	networkId := network.Id
	fmt.Println(token)
	//networkId := "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"
	userId, err := s.repo.Create(req, networkId)
	if err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.DatabaseError()}, nil
	}
	responseUserId := map[string]string{"user_id": userId}

	/*
		subject := "User.UserCreated"
		err = s.nats.Publish(subject, []byte("Hello NATS"))
		if err != nil {
			log.Printf("Error during publishing: %s", err)
		}
		s.nats.Flush()
	*/
	return &pb.Response{Result: common.SUCCESS, Data: responseUserId, Error: nil}, err
}

func CheckAuthWithSecret(tokenString string, secretString string) error {

	type Claims struct {
		ApiKey string `json:"api_key"`
		jwt.StandardClaims
	}
	tokenClaims := Claims{}

	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretString), nil
	})

	if err != nil {
		return errors.New(err.Error())
	}
	if token.Valid {
		return nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return errors.New("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return errors.New("Timing is everything")
		} else {
			return errors.New("Couldn't handle this token")
		}
	} else {
		return errors.New("Couldn't handle this token")
	}
}

func CheckAuth(tokenString string) (string, error) {

	type Claims struct {
		ApiKey string `json:"api_key"`
		jwt.StandardClaims
	}
	tokenClaims := Claims{}

	token, _ := jwt.ParseWithClaims(tokenString, &tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(common.SIGNING_KEY), nil
	})
	fmt.Println(token)
	return tokenClaims.ApiKey, nil
}

func (s *service) Authenticate(ctx context.Context, req *pb.LoginRequest) (*pb.Response, error) {

	networkId := "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"

	user, err := s.repo.Authenticate(req, networkId)
	if err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.DatabaseError()}, nil
	}
	fmt.Println(user)

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

	tokenStr := map[string]string{"token": signedJwtToken}
	return &pb.Response{Result: common.SUCCESS, Data: tokenStr, Error: nil}, nil
}
