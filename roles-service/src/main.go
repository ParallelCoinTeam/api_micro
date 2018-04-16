package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	common "github.com/syedomair/api_micro/common"
	pb "github.com/syedomair/api_micro/roles-service/proto"
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
	db, err := common.CreateDBConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	} else {
		fmt.Println("Connected to DB")
	}

	logger := common.GetLogger()

	repo := &RoleRepository{db, logger}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	natsCon, _ := common.CreateNATSConnection()
	nats := &NatsWrapper{natsCon, logger}

	s := grpc.NewServer(grpc.UnaryInterceptor(AuthInterceptor))
	pb.RegisterRolesServiceServer(s, &Env{repo, nats, logger})

	return s.Serve(lis)
}
func startHTTP(httpPort, grpcPort string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterRolesServiceHandlerFromEndpoint(ctx, gwmux, "127.0.0.1:"+grpcPort, opts); err != nil {
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
