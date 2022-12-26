//package main

import "fmt"

type Product struct {
	ID          int32
	Name        string
	Price       float64
	Description string
	Category    string
}

var pro1 = Product{1, "Manzana", 2000, "Fruta roja de alta calidad", "Frutas"}
var pro2 = Product{2, "Banana", 3000, "Fruta amarilla de alta calidad", "Frutas"}
var pro3 = Product{3, "Pera", 2500, "Fruta verde de alta calidad", "Frutas"}
var Products = []Product{pro1, pro2, pro3}

func main11() {
	pro1.GetAll()
	var pro3 = Product{4, "Fresita", 2500, "Fruta dulce y roja de muy alta calidad", "Frutas"}
	pro3.Save()
	fmt.Println("El producto con el ID 4 es: ")
	fmt.Println((getById(4)))
	pro1.GetAll()

}

func (miProducto Product) Save() {
	Products = append(Products, miProducto)
	fmt.Println("Se ha guardado el producto: " + miProducto.Name)
}

func (miProducto Product) GetAll() {
	fmt.Println("Los productos del slide son: ")
	for _, productoActual := range Products {
		fmt.Println("El producto es: " + productoActual.Name)
	}
}

func getById(ID int32) Product {

	for _, productoActual := range Products {
		if productoActual.ID == ID {
			return productoActual
		}
	}

	prod := Product{}
	return prod

}
