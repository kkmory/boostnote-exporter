package entity

import "time"

type AppConfig struct {
	FolderName          string
	BasePath            string
	PreviousProcessedAt time.Time
}
