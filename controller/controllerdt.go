package controller

import (
	"encoding/json"
	"time"
)

type adreqbody struct {
	BarangID    int       `json:"barang_id"`
	Amount      uint      `json:"amount"`
	ExpiredDate time.Time `json:"expired_date"`
}

type redreqbody struct {
	BarangID int  `json:"barang_id"`
	Amount   uint `json:"amount"`
}

func (S *adreqbody) UnmarshalJSON(bts []byte) (err error) {
	type T adreqbody
	aux := &struct {
		*T
		ExpiredDate string `json:"expired_date"`
	}{
		T: (*T)(S),
	}

	if err = json.Unmarshal(bts, &aux); err != nil {
		return err
	}

	expDate, err := time.Parse("02-01-2006", aux.ExpiredDate)
	if err != nil {
		return err
	}

	S.ExpiredDate = expDate
	return nil
}
