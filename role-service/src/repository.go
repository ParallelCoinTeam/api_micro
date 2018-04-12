package main

import (
	"github.com/jinzhu/gorm"
	pb "github.com/syedomair/api_micro/role-service/proto"
)

type Repository interface {
	//Create(role *pb.Role, networkId string) (string, error)
	Get(roleId string, networkId string) (*pb.Role, error)
	/*
		GetAll(limit string, offset string, orderby string, sort string, networkId string) ([]*pb.Role, string, error)
		Update(role *pb.Role, networkId string) error
		Delete(role *pb.Role, networkId string) error
	*/
}

type RoleRepository struct {
	db *gorm.DB
}

/*
func (repo *RoleRepository) Create(role *pb.Role, networkId string) (string, error) {

	roleId := uuid.NewV4().String()
	role = &pb.Role{Id: roleId, NetworkId: networkId, Title: role.Title, RoleType: role.RoleType, CreatedAt: time.Now().Format(time.RFC3339), UpdatedAt: time.Now().Format(time.RFC3339)}

	if err := repo.db.Create(role).Error; err != nil {
		return "", err
	}
	return roleId, nil
}

func (repo *RoleRepository) GetAll(limit string, offset string, orderby string, sort string, networkId string) ([]*pb.Role, string, error) {

	var roles []*pb.Role
	count := "0"
	if err := repo.db.Table("roles").
		Select("*").
		Count(&count).
		Limit(limit).
		Offset(offset).
		Order(orderby+" "+sort).
		Where("network_id = ?", networkId).
		Scan(&roles).Error; err != nil {
		return nil, "", err
	}
	return roles, count, nil
}
*/
func (repo *RoleRepository) Get(roleId string, networkId string) (*pb.Role, error) {
	role := pb.Role{}
	if err := repo.db.Where("network_id = ?", networkId).Where("id = ?", roleId).Find(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

/*
func (repo *RoleRepository) Update(role *pb.Role, networkId string) error {
	if err := repo.db.Model(role).Update(&role).Error; err != nil {
		return err
	}
	return nil
}

func (repo *RoleRepository) Delete(role *pb.Role, networkId string) error {
	roleId := role.Id
	if err := repo.db.Where("network_id = ?", networkId).Where("id = ?", roleId).Find(&role).Error; err != nil {
		return err
	}
	if err := repo.db.Delete(&role).Error; err != nil {
		fmt.Println("Khalid")
		return err
	}
	return nil
}
*/
