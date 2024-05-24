package model

import (
	"time"
)

type StokMasuk struct {
	ID          int       `json:"id" gorm:"primary_key;auto_increment"`
	StokMasuk   uint      `json:"stok_masuk" db:"stok_masuk"`
	ExpiredDate time.Time `json:"expired_date"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
	BarangID    int       `json:"barang-id"`
	Barang      Barang    `json:"barang,omitempty" gorm:"foreignKey:BarangID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type StokKeluar struct {
	ID         int       `json:"id" gorm:"primary_key;auto_increment"`
	StokKeluar uint      `json:"stok-keluar" db:"stok_keluar"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
	BarangID   int       `json:"barang-id"`
	Barang     Barang    `json:"barang,omitempty" gorm:"foreignKey:BarangID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
