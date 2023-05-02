package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Axit88/UserGrpc/constants"
	"github.com/Axit88/UserGrpc/infrastructure"
	pb "github.com/Axit88/UserGrpc/storage-proto"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer) AddUser(ctx context.Context, input *pb.AddUserInput) (*pb.AddUserOutput, error) {
	db, err := infrastructure.CreateDbClient()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("INSERT INTO USER (UserId, UserName) VALUES (?, ?)", input.UserId, input.UserName)
	if err != nil {
		return nil, err
	}

	output := pb.AddUserOutput{
		Message: fmt.Sprintf("User %v Added Successfully", input.UserName),
	}

	return &output, nil
}

func (s *testApiServer) GetUser(ctx context.Context, input *pb.GetUserInput) (*pb.GetUserOutput, error) {
	db, err := infrastructure.CreateDbClient()
	if err != nil {
		return nil, err
	}

	queryResult, err := db.Query("SELECT * FROM USER WHERE UserId = ?", input.UserId)
	if err != nil {
		return nil, err
	}

	output := pb.GetUserOutput{}
	for queryResult.Next() {
		err = queryResult.Scan(&output.UserId, &output.UserName)
		if err != nil {
			return nil, err
		}
	}
	return &output, nil
}

func (s *testApiServer) DeleteUser(ctx context.Context, input *pb.DeleteUserInput) (*pb.DeleteUserOutput, error) {
	db, err := infrastructure.CreateDbClient()
	if err != nil {
		return nil, err
	}

	_, err = db.Query("DELETE FROM USER WHERE UserId = ?", input.UserId)
	if err != nil {
		return nil, err
	}

	output := pb.DeleteUserOutput{}
	output.Message = fmt.Sprintf("User with id %v Deleted", input.UserId)
	return &output, nil
}

func (s *testApiServer) UpdateUser(ctx context.Context, input *pb.UpdateUserInput) (*pb.UpdateUserOutput, error) {
	db, err := infrastructure.CreateDbClient()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("IUPDATE USER SET UserName = ? WHERE UserId = ?", input.UserName, input.UserId)
	if err != nil {
		return nil, err
	}

	output := pb.UpdateUserOutput{
		Message: fmt.Sprintf("User %v Updated Name %v", input.UserId, input.UserName),
	}
	
	return &output, nil
}



func RunServer(host string, port string) error {
	connection := fmt.Sprintf("%v:%v", constants.GRPC_HOST, constants.GRPC_PORT)
	listner, err := net.Listen("tcp", connection)
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterTestApiServer(grpcServer, &testApiServer{})

	err = grpcServer.Serve(listner)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := RunServer(constants.GRPC_HOST, constants.GRPC_PORT)
	if err != nil {
		fmt.Println(err)
	}
}

