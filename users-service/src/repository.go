package main

import (
	"time"

	log "github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	pb "github.com/syedomair/api_micro/users-service/proto"
)

type Repository interface {
	Create(user *pb.User, networkId string) (string, error)
	GetAll(limit string, offset string, orderby string, sort string, networkId string) ([]*pb.User, string, error)
	Get(userId string, networkId string) (*pb.User, error)
	Update(user *pb.User, networkId string) error
	Delete(user *pb.User, networkId string) error
}

type UserRepository struct {
	db     *gorm.DB
	logger log.Logger
}

func (repo *UserRepository) Create(user *pb.User, networkId string) (string, error) {

	start := time.Now()
	repo.logger.Log("METHOD", "Create", "SPOT", "method start", "time_start", start)
	userId := uuid.NewV4().String()
	user = &pb.User{
		Id:        userId,
		NetworkId: networkId,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		IsAdmin:   user.IsAdmin,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339)}

	if err := repo.db.Create(user).Error; err != nil {
		return "", err
	}
	repo.logger.Log("METHOD", "Create", "SPOT", "method end", "time_spent", time.Since(start))
	return userId, nil
}
func (repo *UserRepository) GetAll(limit string, offset string, orderby string, sort string, networkId string) ([]*pb.User, string, error) {

	start := time.Now()
	repo.logger.Log("METHOD", "GetAll", "SPOT", "method start", "time_start", start)
	var users []*pb.User
	count := "0"
	if err := repo.db.Table("users").
		Select("*").
		Count(&count).
		Limit(limit).
		Offset(offset).
		Order(orderby+" "+sort).
		Where("network_id = ?", networkId).
		Scan(&users).Error; err != nil {
		return nil, "", err
	}
	repo.logger.Log("METHOD", "GetAll", "SPOT", "method end", "time_spent", time.Since(start))
	return users, count, nil
}

func (repo *UserRepository) Get(userId string, networkId string) (*pb.User, error) {
	start := time.Now()
	repo.logger.Log("METHOD", "Get", "SPOT", "method start", "time_start", start)
	user := pb.User{}
	if err := repo.db.Where("network_id = ?", networkId).Where("id = ?", userId).Find(&user).Error; err != nil {
		return nil, err
	}
	repo.logger.Log("METHOD", "Get", "SPOT", "method end", "time_spent", time.Since(start))
	return &user, nil
}

func (repo *UserRepository) Update(user *pb.User, networkId string) error {
	start := time.Now()
	repo.logger.Log("METHOD", "Update", "SPOT", "method start", "time_start", start)
	if err := repo.db.Model(user).Update(&user).Error; err != nil {
		return err
	}
	repo.logger.Log("METHOD", "Update", "SPOT", "method end", "time_spent", time.Since(start))
	return nil
}

func (repo *UserRepository) Delete(user *pb.User, networkId string) error {
	start := time.Now()
	repo.logger.Log("METHOD", "Delete", "SPOT", "method start", "time_start", start)
	userId := user.Id
	if err := repo.db.Where("network_id = ?", networkId).Where("id = ?", userId).Find(&user).Error; err != nil {
		return err
	}
	if err := repo.db.Delete(&user).Error; err != nil {
		return err
	}
	repo.logger.Log("METHOD", "Delete", "SPOT", "method end", "time_spent", time.Since(start))
	return nil
}
