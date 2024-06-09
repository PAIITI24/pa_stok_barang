package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hakushigo/stok_barang/controller"
	"github.com/hakushigo/stok_barang/helper"
)

func main() {
	helper.Migrator()

	srv := fiber.New(
		fiber.Config{
			Immutable: false,
			AppName:   "stok_Management_obat",
		})

	srv.Put("/barang/stok/add", controller.AddStok)
	srv.Put("/barang/stok/reduce", controller.ReduceStok)
	srv.Get("/barang/stok/add/history", controller.ListStokMasuk)
	srv.Get("/barang/stok/reduce/history", controller.ListStokKeluar)
	srv.Get("/barang/stok/add/history/:id", controller.ListStokMasukOfBarang)
	srv.Get("/barang/stok/reduce/history/:id", controller.ListStokKeluarOfBarang)

	err := srv.Listen(":3011")

	if err != nil {
		panic(err)
	}
}
