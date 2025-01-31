package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ocuprum/menu-constructor/internal/models"
	"github.com/ocuprum/menu-constructor/internal/services"
	pkgHTTP "github.com/ocuprum/menu-constructor/pkg/http"
)

type FoodHandler struct {
	svc *services.FoodService
}

func NewFoodHandler(svc *services.FoodService) *FoodHandler {
	return &FoodHandler{svc: svc}
}

func (h *FoodHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/addfi", h.addIngredient)
	mux.HandleFunc("/createfood", h.create)
}

func (h *FoodHandler) create(resp http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest)
		return
	}
	
	food := models.Food{}
	if err = json.Unmarshal(body, &food); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error unmarshalling the json")
		return
	}

	if err = h.svc.Create(req.Context(), food); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error creating ingredient")
		return
	}
}

func (h *FoodHandler) addIngredient(resp http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest)
		return
	}
	
	fi := models.IngredientFood{}
	if err = json.Unmarshal(body, &fi); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error unmarshalling the json")
		return
	}

	fmt.Println(fi)

	if err = h.svc.AddIngredient(req.Context(), fi); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error adding ingredient to food")
		return
	}
}
