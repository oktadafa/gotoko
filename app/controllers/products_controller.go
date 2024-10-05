package controllers

import (
	"fmt"
	"github.com/oktadafa/gotoko/app/models"
	"github.com/unrolled/render"
	"net/http"
)

func (server *Server) Products(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
		Layout: "layout",
	})

	productModel := models.Product{}
	products, err := productModel.GetProduct(server.DB)
	if err != nil {
		return
	}

	fmt.Print(err)

	_ = render.HTML(w, http.StatusOK, "products", map[string]interface{}{
		"products": products,
	})
}
