package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"webapp/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func HandleListProducts(w http.ResponseWriter, _ *http.Request) {
	products := models.QueryAllProducts()

	temp.ExecuteTemplate(w, "Index", products)
}

func HandleNewProduct(w http.ResponseWriter, _ *http.Request) {
	temp.ExecuteTemplate(w, "NewProduct", nil)
}

func InsertNewProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		quantity := r.FormValue("quantidade")
		price := r.FormValue("preco")
	
		formattedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Erro na conversão do preço: ", err)
		}

		formattedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			fmt.Println("Erro na conversão da quantidade: ", err)
		}

		models.CreateNewProduct(name, description, formattedPrice, formattedQuantity)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func HandleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id") 
	models.DeleteProduct(productId)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func HandleEditProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	product := models.FindById(productId)

	temp.ExecuteTemplate(w, "EditProduct", product)
}

func HandleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		quantity := r.FormValue("quantidade")
		price := r.FormValue("preco")

		formattedId, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Erro na conversão da quantidade: ", err)
		}

		formattedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Erro na conversão do preço: ", err)
		}

		formattedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			fmt.Println("Erro na conversão da quantidade: ", err)
		}

		models.UpdateProduct(formattedId, name, description, formattedPrice, formattedQuantity)
	}

	http.Redirect(w,r, "/", http.StatusMovedPermanently)
}
