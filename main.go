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

	err := srv.Listen(":3011")

	if err != nil {
		panic(err)
	}
}
