package weimiao

import "time"

type InfoflowMedia struct {
	ID          int       `json:"id" gorm:"column:id"`
	Mid         string    `json:"mid" gorm:"column:mid"`
	MediaName   string    `json:"media_name" gorm:"column:media_name"`
	Status      string    `json:"status" gorm:"column:status"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
	Pid         int       `json:"pid" gorm:"column:pid"`
	AccountName string    `json:"account_name" gorm:"column:account_name"`
	Mark        string    `json:"mark" gorm:"column:mark"`
	IsReturn    string    `json:"is_return" gorm:"column:is_return"`
	ReleaseID   string    `json:"release_id" gorm:"column:release_id"`
	SpaceID     string    `json:"space_id" gorm:"column:space_id"`
}

func (InfoflowMedia) TableName() string {
	return "mh_infoflow_media"
}
