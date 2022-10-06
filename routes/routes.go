package routes

import (
	"net/http"
	productController "webapp/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", productController.HandleListProducts)

	http.HandleFunc("/new", productController.HandleNewProduct)
	
	http.HandleFunc("/insert", productController.InsertNewProduct)
	
	http.HandleFunc("/delete", productController.HandleDeleteProduct)
	
	http.HandleFunc("/edit", productController.HandleEditProduct)

	http.HandleFunc("/update", productController.HandleUpdateProduct)
}
