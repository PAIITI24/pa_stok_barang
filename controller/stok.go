package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/hakushigo/stok_barang/model"
	"gorm.io/gorm"
)

/**
TODO: Create Adding-Reducing Stock Number, and recored them in Stok Masuk & Keluar Table.
*/

func AddStok(ctx *fiber.Ctx) error {
	var request adreqbody
	var barang model.Barang
	var newInsertRecord model.StokMasuk

	// get the request body data
	err := json.Unmarshal(ctx.Body(), &request)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status_code": fiber.StatusInternalServerError,
			"error":       err,
		})
	}

	// then we decided to fetch the barang data.
	search := db.Where(&model.Barang{ID: request.BarangID}).First(&barang)
	if search.RowsAffected == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status_code": fiber.StatusNotFound,
			"error":       "The Barang can't be found",
		})
	}

	// because it's adding we don't need to amount checking beforehand
	// and then we added to record

	err = db.Transaction(func(tx *gorm.DB) error { // transaction as a condom of SQL
		barang.JumlahStok += request.Amount // first we update the amount of barang in the barang field
		if err := tx.Save(&barang).Error; err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status_code": fiber.StatusInternalServerError,
				"error":       err,
			})
		}

		// add item to the add stock table
		// get the barang that was meant

		newInsertRecord = model.StokMasuk{
			StokMasuk:   request.Amount,
			Barang:      barang,
			ExpiredDate: request.ExpiredDate,
		}
		if err := tx.Create(&newInsertRecord).Error; err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status_code": fiber.StatusInternalServerError,
				"error":       err,
			})
		}

		return nil
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status_code": fiber.StatusInternalServerError,
			"error":       err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status_code": fiber.StatusOK,
		"barang":      barang,
		"record_info": newInsertRecord,
	})
}

func ReduceStok(ctx *fiber.Ctx) error {
	var request redreqbody

	var obat model.Barang
	var newOutRecord model.StokKeluar

	err := json.Unmarshal(ctx.Body(), &request)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status_code": fiber.StatusInternalServerError,
			"error":       err,
		})
	}

	// then we decided to fetch the obat data.
	search := db.Where(&model.Barang{ID: request.BarangID}).First(&obat)
	if search.RowsAffected == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status_code": fiber.StatusNotFound,
			"error":       "The Barang can't be found",
		})
	}

	if obat.JumlahStok < request.Amount {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status_code": fiber.StatusBadRequest,
			"error":       "The amount taken larger than what available",
		})
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		obat.JumlahStok -= request.Amount
		if err := tx.Save(&obat).Error; err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status_code": fiber.StatusInternalServerError,
				"error":       err,
			})
		}

		newOutRecord = model.StokKeluar{
			StokKeluar: request.Amount,
			Barang:     obat,
		}

		if err := tx.Create(&newOutRecord).Error; err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status_code": fiber.StatusInternalServerError,
				"error":       err,
			})
		}

		return nil
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status_code": fiber.StatusInternalServerError,
			"error":       err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status_code": fiber.StatusOK,
		"obat":        obat,
		"record_info": newOutRecord,
	})
}
