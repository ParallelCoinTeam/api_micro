package main

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	pb "github.com/syedomair/api_micro/public-service/proto"
)

type Repository interface {
	Create(user *pb.User, networkId string) (string, error)
	Authenticate(req *pb.LoginRequest, networkId string) (*pb.User, error)
}

type PublicRepository struct {
	db *gorm.DB
}

func (repo *PublicRepository) Create(user *pb.User, networkId string) (string, error) {

	userId := uuid.NewV4().String()
	user = &pb.User{Id: userId, NetworkId: networkId, FirstName: user.FirstName, LastName: user.LastName, Email: user.Email, Password: user.Password, IsAdmin: user.IsAdmin, CreatedAt: time.Now().Format(time.RFC3339), UpdatedAt: time.Now().Format(time.RFC3339)}

	if err := repo.db.Create(user).Error; err != nil {
		return "", err
	}
	return userId, nil
}

func (repo *PublicRepository) Authenticate(req *pb.LoginRequest, networkId string) (*pb.User, error) {

	user := pb.User{}
	if err := repo.db.Where("network_id = ?", networkId).Where("email = ?", req.Email).Where("password = ?", req.Password).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}