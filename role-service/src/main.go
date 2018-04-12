package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/syedomair/api_micro/common"
	pb "github.com/syedomair/api_micro/role-service/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

func main() {
	errors := make(chan error)
	httpPort := "8180"
	grpcPort := "50080"
	fmt.Println("HTTP PORT", httpPort)
	fmt.Println("GRPC PORT", grpcPort)

	go func() { errors <- startGRPC(grpcPort) }()
	go func() { errors <- startHTTP(httpPort, grpcPort) }()
	for err := range errors {
		log.Fatal(err)
		return
	}
}
func startGRPC(port string) error {
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	} else {
		fmt.Printf("Connected to DB")
	}
	repo := &RoleRepository{db}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	//s := grpc.NewServer()
	s := grpc.NewServer(grpc.UnaryInterceptor(AuthInterceptor))
	pb.RegisterRoleServiceServer(s, &Service{repo})
	return s.Serve(lis)
}

func startHTTP(httpPort, grpcPort string) error {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterRoleServiceHandlerFromEndpoint(ctx, gwmux, "127.0.0.1:"+grpcPort, opts); err != nil {
		return err
	}
	mux := http.NewServeMux()
	mux.Handle("/v1/", gwmux)
	http.ListenAndServe(":"+httpPort, mux)
	return nil
}

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "missing context metadata")
	}
	if len(meta["authorization"]) != 1 {
		return nil, grpc.Errorf(codes.Unauthenticated, "missing authorization token")
	}
	currentUserId, networkId, authErr := common.CheckAuth(meta["authorization"][0])
	if authErr != nil {
		return &pb.ResponseList{Result: common.FAILURE, Error: common.CommonError(authErr.Error()), Data: nil}, nil
	}
	ctx = context.WithValue(ctx, "current_user_id", currentUserId)
	ctx = context.WithValue(ctx, "network_id", networkId)
	return handler(ctx, req)
}
func CheckAuth(tokenString string) (string, string, error) {

	type Claims struct {
		CurrentUserId string `json:"current_user_id"`
		NetworkId     string `json:"network_id"`
		IsAdmin       string `json:"is_admin"`
		jwt.StandardClaims
	}
	tokenClaims := Claims{}

	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(common.SIGNING_KEY), nil
	})
	if err != nil {
		return "", "", errors.New(err.Error())
	}
	if token.Valid {
		return tokenClaims.CurrentUserId, tokenClaims.NetworkId, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return "", "", errors.New("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return "", "", errors.New("Timing is everything")
		} else {
			return "", "", errors.New("Couldn't handle this token")
		}
	} else {
		return "", "", errors.New("Couldn't handle this token")
	}
}
