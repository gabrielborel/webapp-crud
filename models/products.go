package models

import (
	"fmt"
	"webapp/database"
)

type Product struct {
	ID					int
	Name 				string
	Description string
	Price 			float64
	Quantity 		int
}

func QueryAllProducts() []Product {
	db := database.ConnectWithDatabase()
	defer db.Close()

	allProducts, err := db.Query("select * from products order by id asc")
	
	if err != nil {
		panic(err.Error())
	}

	product := Product{}
	products := []Product{}

	for allProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = allProducts.Scan(&id, &name, &description, &price, &quantity) 

		if err != nil {
			panic(err.Error())
		}

		product.ID = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity

		products = append(products, product)
	}

	return products
}

func CreateNewProduct(name, description string, price float64, quantity int) {
	db := database.ConnectWithDatabase()
	defer db.Close()

	newProductQuery, err := db.Prepare("insert into products(name, description, price, quantity) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	newProductQuery.Exec(name, description, price, quantity)

	fmt.Println("Produto criado com sucesso !")
}

func DeleteProduct(productId string) {
	db := database.ConnectWithDatabase()
	defer db.Close()

	deleteProductQuery, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteProductQuery.Exec(productId)

	fmt.Println("Produto deletado com sucesso !")
}

func FindById(productId string) Product {
	db := database.ConnectWithDatabase()
	defer db.Close()

	productDb, err := db.Query("select * from products where id=$1", productId)
	if err != nil {
		panic(err.Error())
	}

	product := Product{}

	for productDb.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productDb.Scan(&id, &name, &description, &price, &quantity) 
		if err != nil {
			panic(err.Error())
		}

		product.ID = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity
	}

	return product
}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	db := database.ConnectWithDatabase()
	defer db.Close()

	updateProductQuery, err := db.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5 ")
	if err != nil {
		panic(err.Error())
	}

	updateProductQuery.Exec(name, description, price, quantity, id)
}
