package main

import (
	"fmt"
	"time"

	log "github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	pb "github.com/syedomair/api_micro/roles-service/proto"
)

type Repository interface {
	Create(role *pb.Role, networkId string) (string, error)
	Get(roleId string, networkId string) (*pb.Role, error)
	GetAll(limit string, offset string, orderby string, sort string, networkId string) ([]*pb.Role, string, error)
	Update(role *pb.Role, networkId string) error
	Delete(role *pb.Role, networkId string) error
}

type RoleRepository struct {
	db     *gorm.DB
	logger log.Logger
}

func (repo *RoleRepository) Create(role *pb.Role, networkId string) (string, error) {
	start := time.Now()
	repo.logger.Log("METHOD", "Create", "SPOT", "method start", "time_start", start)
	roleId := uuid.NewV4().String()
	role = &pb.Role{
		Id:        roleId,
		NetworkId: networkId,
		Title:     role.Title,
		RoleType:  role.RoleType,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339)}

	if err := repo.db.Create(role).Error; err != nil {
		return "", err
	}
	repo.logger.Log("METHOD", "Create", "SPOT", "method end", "time_spent", time.Since(start))
	return roleId, nil
}

func (repo *RoleRepository) GetAll(limit string, offset string, orderby string, sort string, networkId string) ([]*pb.Role, string, error) {
	start := time.Now()
	repo.logger.Log("METHOD", "GetAll", "SPOT", "method start", "time_start", start)
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
	repo.logger.Log("METHOD", "GetAll", "SPOT", "method end", "time_spent", time.Since(start))
	return roles, count, nil
}
func (repo *RoleRepository) Get(roleId string, networkId string) (*pb.Role, error) {
	start := time.Now()
	repo.logger.Log("METHOD", "Get", "SPOT", "method start", "time_start", start)
	role := pb.Role{}
	if err := repo.db.Where("network_id = ?", networkId).Where("id = ?", roleId).Find(&role).Error; err != nil {
		return nil, err
	}
	repo.logger.Log("METHOD", "Get", "SPOT", "method end", "time_spent", time.Since(start))
	return &role, nil
}

func (repo *RoleRepository) Update(role *pb.Role, networkId string) error {
	start := time.Now()
	repo.logger.Log("METHOD", "Update", "SPOT", "method start", "time_start", start)
	fmt.Println(role)
	if err := repo.db.Model(role).Update(&role).Error; err != nil {
		return err
	}
	repo.logger.Log("METHOD", "Update", "SPOT", "method end", "time_spent", time.Since(start))
	return nil
}

func (repo *RoleRepository) Delete(role *pb.Role, networkId string) error {
	start := time.Now()
	repo.logger.Log("METHOD", "Delete", "SPOT", "method start", "time_start", start)
	roleId := role.Id
	if err := repo.db.Where("network_id = ?", networkId).Where("id = ?", roleId).Find(&role).Error; err != nil {
		return err
	}
	if err := repo.db.Delete(&role).Error; err != nil {
		return err
	}
	repo.logger.Log("METHOD", "Delete", "SPOT", "method end", "time_spent", time.Since(start))
	return nil
}
