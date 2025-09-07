package model

import (
	"time"
)

type Client struct {
	ID        string `gorm:"primaryKey"`
	Ip        string
	OS        string
	Status    string
	FirstSeen time.Time
	LastSeen  time.Time
	Metadata  string
	CreatedAt time.Time
	DeletedAt time.Time
}

type Command struct {
	ID        string `gorm:"primaryKey"`
	ClientID  string
	Command   string
	Timestamp time.Time
	Status    string
}

type CommandLog struct {
	ID        string `gorm:"primaryKey"`
	CommandID string
	ClientID  string
	Output    string
	Timestamp time.Time
}

type ClientFile struct {
	ID        string `gorm:"primaryKey"`
	ClientID  string
	FileName  string
	FilePath  string
	FileType  string
	FileSize  int64
	Timestamp time.Time
}
