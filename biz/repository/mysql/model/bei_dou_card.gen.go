// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameBeiDouCard = "bei_dou_card"

// BeiDouCard mapped from table <bei_dou_card>
type BeiDouCard struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Number    string    `gorm:"column:number;not null" json:"number"`
	Type      bool      `gorm:"column:type;not null" json:"type"`
	Station   int32     `gorm:"column:station" json:"station"`
	ShipID    int32     `gorm:"column:ship_id" json:"ship_id"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	Extra     string    `gorm:"column:extra" json:"extra"`
}

// TableName BeiDouCard's table name
func (*BeiDouCard) TableName() string {
	return TableNameBeiDouCard
}
