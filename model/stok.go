package model

import (
	"encoding/json"
	"time"
)

type StokMasuk struct {
	ID          int       `json:"id" gorm:"primary_key;auto_increment"`
	StokMasuk   uint      `json:"stok_masuk"`
	ExpiredDate time.Time `json:"expired_date"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
	BarangID    int       `json:"barang_id"`
	Barang      Barang    `json:"barang,omitempty" gorm:"foreignKey:BarangID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type StokKeluar struct {
	ID          int       `json:"id" gorm:"primary_key;auto_increment"`
	StokKeluar  uint      `json:"stok_keluar"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
	StokMasukID int       `json:"stok_masuk_id"`
	StokMasuk   StokMasuk `json:"stok_masuk" gorm:"foreignKey:StokMasukID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BarangID    int       `json:"barang_id"`
	Barang      Barang    `json:"barang,omitempty" gorm:"foreignKey:BarangID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (S *StokMasuk) MarshalJSON() ([]byte, error) {
	type Alias StokMasuk

	return json.Marshal(&struct {
		ExpiredDate string `json:"expired_date"`
		CreatedAt   string `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
		UpdatedAt   string `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
		*Alias
	}{
		ExpiredDate: S.ExpiredDate.Format("02/01/2006"),
		CreatedAt:   S.CreatedAt.Format("02/01/2006"),
		UpdatedAt:   S.UpdatedAt.Format("02/01/2006"),
		Alias:       (*Alias)(S),
	})
}

func (S *StokKeluar) MarshalJSON() ([]byte, error) {
	type Alias StokKeluar

	return json.Marshal(&struct {
		CreatedAt string `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
		UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
		*Alias
	}{
		CreatedAt: S.CreatedAt.Format("02/01/2006"),
		UpdatedAt: S.UpdatedAt.Format("02/01/2006"),
		Alias:     (*Alias)(S),
	})
}
