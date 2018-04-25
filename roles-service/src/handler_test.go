package main

import (
	"context"
	"testing"
	"time"

	"github.com/syedomair/api_micro/common"
	pb "github.com/syedomair/api_micro/roles-service/proto"
	testdata "github.com/syedomair/api_micro/testdata"
)

func TestCreateRole(t *testing.T) {

	env := Env{repo: &mockDB{}, nats: &mockNATS{}, logger: common.GetLogger()}
	start := time.Now()
	env.logger.Log("METHOD", "TestCreateRole", "SPOT", "method start", "time_start", start)
	ctx := context.WithValue(context.Background(), "network_id", testdata.NetworkId)

	//ALL Good
	role := &pb.Role{Title: testdata.RoleTitle1, RoleType: testdata.RoleType}
	response, _ := env.CreateRole(ctx, role)

	expected := testdata.SUCCESS
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//ALL Invalid RoleTitle
	role = &pb.Role{Title: testdata.RoleTitleInvalid, RoleType: testdata.RoleType}
	response, _ = env.CreateRole(ctx, role)

	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//ALL Invalid Role Type
	role = &pb.Role{Title: testdata.RoleTitle1, RoleType: testdata.RoleTypeInvalid}
	response, _ = env.CreateRole(ctx, role)

	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
	env.logger.Log("METHOD", "TestCreateRole", "SPOT", "method end", "time_spent", time.Since(start))
}

func TestGetAll(t *testing.T) {

	env := Env{repo: &mockDB{}, nats: &mockNATS{}, logger: common.GetLogger()}
	start := time.Now()
	env.logger.Log("METHOD", "TestGetAll", "SPOT", "method start", "time_start", start)
	ctx := context.WithValue(context.Background(), "network_id", testdata.NetworkId)

	//ALL Good
	req := &pb.RequestQuery{Limit: "3", Offset: "0", Orderby: "title", Sort: "desc"}
	response, _ := env.GetAll(ctx, req)

	expected := testdata.SUCCESS
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid Offset
	req = &pb.RequestQuery{Limit: "3", Offset: "A", Orderby: "title", Sort: "desc"}
	response, _ = env.GetAll(ctx, req)

	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid Limit
	req = &pb.RequestQuery{Limit: "A", Offset: "0", Orderby: "title", Sort: "desc"}
	response, _ = env.GetAll(ctx, req)

	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid orderby
	req = &pb.RequestQuery{Limit: "3", Offset: "0", Orderby: "3", Sort: "desc"}
	response, _ = env.GetAll(ctx, req)

	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid sort
	req = &pb.RequestQuery{Limit: "3", Offset: "0", Orderby: "title", Sort: "123"}
	response, _ = env.GetAll(ctx, req)

	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
	env.logger.Log("METHOD", "TestGetAll", "SPOT", "method end", "time_spent", time.Since(start))
}

func TestGetRole(t *testing.T) {

	env := Env{repo: &mockDB{}, nats: &mockNATS{}, logger: common.GetLogger()}
	start := time.Now()
	env.logger.Log("METHOD", "TestGetRole", "SPOT", "method start", "time_start", start)
	ctx := context.WithValue(context.Background(), "network_id", testdata.NetworkId)

	//All Good
	req := &pb.Role{Id: testdata.RoleId1}
	response, _ := env.GetRole(ctx, req)

	expected := testdata.SUCCESS
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid RoleID
	req = &pb.Role{Id: testdata.InValidId}
	response, _ = env.GetRole(ctx, req)

	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
	env.logger.Log("METHOD", "TestGetRole", "SPOT", "method end", "time_spent", time.Since(start))
}

func TestUpdateRole(t *testing.T) {

	env := Env{repo: &mockDB{}, nats: &mockNATS{}, logger: common.GetLogger()}
	start := time.Now()
	env.logger.Log("METHOD", "TestUpdateRole", "SPOT", "method start", "time_start", start)
	ctx := context.WithValue(context.Background(), "network_id", testdata.NetworkId)

	//All Good
	role := &pb.Role{Id: testdata.RoleId1, Title: testdata.RoleTitle1, RoleType: testdata.RoleType}
	response, _ := env.UpdateRole(ctx, role)

	expected := testdata.SUCCESS
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
	//Invalid RoleId
	role = &pb.Role{Id: testdata.InValidId, Title: testdata.RoleTitle1, RoleType: testdata.RoleType}
	response, _ = env.UpdateRole(ctx, role)

	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//ALL Invalid RoleTitle
	role = &pb.Role{Id: testdata.RoleId1, Title: testdata.RoleTitleInvalid, RoleType: testdata.RoleType}
	response, _ = env.UpdateRole(ctx, role)

	expected = testdata.SUCCESS
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//ALL Invalid RoleTitle
	role = &pb.Role{Id: testdata.RoleId1, RoleType: testdata.RoleType}
	response, _ = env.UpdateRole(ctx, role)

	expected = testdata.SUCCESS
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//ALL Invalid Role Type
	role = &pb.Role{Id: testdata.RoleId1, Title: testdata.RoleTitle1, RoleType: testdata.RoleTypeInvalid}
	response, _ = env.UpdateRole(ctx, role)

	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
	env.logger.Log("METHOD", "TestUpdateRole", "SPOT", "method end", "time_spent", time.Since(start))

	//No Role Type
	role = &pb.Role{Id: testdata.RoleId1, Title: testdata.RoleTitle1}
	response, _ = env.UpdateRole(ctx, role)
	expected = testdata.SUCCESS
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
	env.logger.Log("METHOD", "TestUpdateRole", "SPOT", "method end", "time_spent", time.Since(start))
}

func TestDeleteRole(t *testing.T) {

	env := Env{repo: &mockDB{}, nats: &mockNATS{}, logger: common.GetLogger()}
	start := time.Now()
	env.logger.Log("METHOD", "TestDeleteRole", "SPOT", "method start", "time_start", start)
	ctx := context.WithValue(context.Background(), "network_id", testdata.NetworkId)

	//All Good
	role := &pb.Role{Id: testdata.RoleId1}
	response, _ := env.DeleteRole(ctx, role)

	expected := testdata.SUCCESS
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid RoleId
	role = &pb.Role{Id: testdata.InValidId}
	response, _ = env.DeleteRole(ctx, role)

	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
	env.logger.Log("METHOD", "TestDeleteRole", "SPOT", "method end", "time_spent", time.Since(start))
}

type mockNATS struct {
}

func (mnats *mockNATS) PublishDeleteEvent(roleId string, networkId string) error {
	return nil
}

type mockDB struct {
	roles []*pb.Role
}

func (mdb *mockDB) initMockDb() []*pb.Role {
	mdb.roles = append(mdb.roles, &pb.Role{Id: testdata.RoleId1, NetworkId: testdata.NetworkId, Title: testdata.RoleTitle1, RoleType: testdata.RoleType, CreatedAt: time.Now().Format(time.RFC3339), UpdatedAt: time.Now().Format(time.RFC3339)})
	mdb.roles = append(mdb.roles, &pb.Role{Id: testdata.RoleId2, NetworkId: testdata.NetworkId, Title: testdata.RoleTitle2, RoleType: testdata.RoleType, CreatedAt: time.Now().Format(time.RFC3339), UpdatedAt: time.Now().Format(time.RFC3339)})
	mdb.roles = append(mdb.roles, &pb.Role{Id: testdata.RoleId3, NetworkId: testdata.NetworkId, Title: testdata.RoleTitle3, RoleType: testdata.RoleType, CreatedAt: time.Now().Format(time.RFC3339), UpdatedAt: time.Now().Format(time.RFC3339)})
	mdb.roles = append(mdb.roles, &pb.Role{Id: testdata.RoleId4, NetworkId: testdata.NetworkId, Title: testdata.RoleTitle4, RoleType: testdata.RoleType, CreatedAt: time.Now().Format(time.RFC3339), UpdatedAt: time.Now().Format(time.RFC3339)})
	mdb.roles = append(mdb.roles, &pb.Role{Id: testdata.RoleId5, NetworkId: testdata.NetworkId, Title: testdata.RoleTitle5, RoleType: testdata.RoleType, CreatedAt: time.Now().Format(time.RFC3339), UpdatedAt: time.Now().Format(time.RFC3339)})

	return mdb.roles
}
func (mdb *mockDB) Create(role *pb.Role, networkId string) (string, error) {
	roles := mdb.initMockDb()
	return roles[0].Id, nil
}
func (mdb *mockDB) GetAll(limit string, offset string, orderby string, sort string, networkId string) ([]*pb.Role, string, error) {
	return mdb.roles, "5", nil
}
func (mdb *mockDB) Get(roleId string, networkId string) (*pb.Role, error) {
	role := &pb.Role{Id: testdata.RoleId1, NetworkId: testdata.NetworkId, Title: testdata.RoleTitle1, RoleType: testdata.RoleType, CreatedAt: time.Now().Format(time.RFC3339), UpdatedAt: time.Now().Format(time.RFC3339)}
	return role, nil
}
func (mdb *mockDB) Update(role *pb.Role, networkId string) error {
	return nil
}
func (mdb *mockDB) Delete(role *pb.Role, networkId string) error {
	return nil
}
