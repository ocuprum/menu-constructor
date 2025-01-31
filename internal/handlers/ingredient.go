package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ocuprum/menu-constructor/internal/models"
	"github.com/ocuprum/menu-constructor/internal/services"
	pkgHTTP "github.com/ocuprum/menu-constructor/pkg/http"
)

type IngredientHandler struct {
	svc *services.IngredientService
}

func NewIngredientHandler(svc *services.IngredientService) *IngredientHandler {
	return &IngredientHandler{svc: svc}
}

func (h *IngredientHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/createingred", h.create)
}

func (h *IngredientHandler) create(resp http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest)
		return
	}
	
	ingred := models.Ingredient{}
	if err = json.Unmarshal(body, &ingred); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error unmarshalling the json")
		return
	}

	if err = h.svc.Create(req.Context(), ingred); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error creating ingredient")
		return
	}
}