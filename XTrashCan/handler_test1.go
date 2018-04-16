package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	pb "github.com/syedomair/api_micro/role-service/proto"
)

type mockDB struct {
	roles []*pb.Role
}

func (mdb *mockDB) initMockDb() ([]*pb.Role, error) {
	roles := make([]*pb.Role, 0)
	roles = append(roles, &pb.Role{Id: "04b58e6e-f910-4ff0-83f1-27fbfa85dc2f", NetworkId: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11", Title: "Role Title", RoleType: "1", CreatedAt: time.Now().Format(time.RFC3339), UpdatedAt: time.Now().Format(time.RFC3339)})
	roles = append(roles, &pb.Role{Id: "14b58e6e-f910-4ff0-83f1-27fbfa85dc2f", NetworkId: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11", Title: "Role Title", RoleType: "1", CreatedAt: time.Now().Format(time.RFC3339), UpdatedAt: time.Now().Format(time.RFC3339)})

	return roles, nil
}

func (mdb *mockDB) Get(roleId string, networkId string) (*pb.Role, error) {
	mdb.roles, _ = mdb.initMockDb()
	/*
		for _, v := range myconfig {
		    if v.Key == "key1" {
		            // Found!
			        }
				}
	*/
	role := &pb.Role{Id: "04b58e6e-f910-4ff0-83f1-27fbfa85dc2f", NetworkId: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11", Title: "Role Title", RoleType: "1", CreatedAt: time.Now().Format(time.RFC3339), UpdatedAt: time.Now().Format(time.RFC3339)}
	return role, nil
}
func TestGetRole(t *testing.T) {

	srv := Service{repo: &mockDB{}}
	role, _ := srv.GetRole(context.Background(), &pb.Role{})

	fmt.Println("Syed")
	fmt.Println(role)
	fmt.Println("Khalid")
	//expected := 1
	//t.Errorf("\n...expected = %v\n...obtained = %v", expected, "123")
}
