package model

import (
	"github.com/KubeOperator/KubeOperator/pkg/model/common"
	uuid "github.com/satori/go.uuid"
)

type ClusterBackupStrategy struct {
	common.BaseModel
	ID              string        `json:"id"`
	Cron            int           `json:"cron"`
	SaveNum         int           `json:"saveNum"`
	BackupAccountID string        `json:"backupAccountId"`
	ClusterID       string        `json:"clusterId"`
	Status          string        `json:"status"`
	BackupAccount   BackupAccount `json:"_"`
}

func (c *ClusterBackupStrategy) BeforeCreate() error {
	c.ID = uuid.NewV4().String()
	return nil
}

func (c ClusterBackupStrategy) TableName() string {
	return "ko_cluster_backup_strategy"
}
