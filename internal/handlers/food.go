package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/google/uuid"

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
	mux.HandleFunc(pkgHTTP.GetPath("/food/{id}"), h.getByID)
	mux.HandleFunc(pkgHTTP.GetPath("/foods/paginate"), h.paginate)
	mux.HandleFunc(pkgHTTP.PostPath("/food/create"), h.create)
	mux.HandleFunc(pkgHTTP.PutPath("/food/change"), h.change)
	mux.HandleFunc(pkgHTTP.DeletePath("/foods/delete"), h.delete)
	mux.HandleFunc(pkgHTTP.PostPath("/food/ingredient/add"), h.addIngredient)
	mux.HandleFunc(pkgHTTP.DeletePath("/food/ingredient/delete"), h.deleteIngredient)
}

func (h *FoodHandler) getByID(resp http.ResponseWriter, req *http.Request) {
	id, err := uuid.Parse(req.PathValue("id"))
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error parsing the id")
		return
	}

	food, err := h.svc.GetByID(req.Context(), id)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error retrieving the food from the db")
		return
	}

	foodJSON, err := json.MarshalIndent(food, "", "  ")
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error preparing food JSON")
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(foodJSON)	
}

func (h *FoodHandler) paginate(resp http.ResponseWriter, req *http.Request) {
	limitStr := req.URL.Query().Get("limit")
	offsetStr := req.URL.Query().Get("offset")

	if limitStr == "" || offsetStr == "" {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Missing limit or offset parameter")
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error parsing the limit")
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error parsing the offset")
		return
	}

	foods, err := h.svc.Paginate(req.Context(), limit, offset)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error retrieving foods from the db")
		return
	}

	foodsJSON, err := json.MarshalIndent(foods, "", "  ")
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error preparing foods JSON")
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(foodsJSON)	
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
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error creating food")
		return
	}
}

func (h *FoodHandler) change(resp http.ResponseWriter, req *http.Request) {	
	body, err := io.ReadAll(req.Body)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error reading the request body")
	}

	food := models.Food{}
	if err = json.Unmarshal(body, &food); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error unmarshalling the json")
		return
	}

	if err = h.svc.Change(req.Context(), food); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error changing the food")
	}
}

func (h *FoodHandler) delete(resp http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error reading the request body")
	}

	dr := models.DeleteRequest{}
	if err = json.Unmarshal(body, &dr); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error unmarshalling the json")
		return
	}

	if err = h.svc.Delete(req.Context(), dr.IDs); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error deleting foods")
		return
	}
}

func (h *FoodHandler) addIngredient(resp http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error reading the request body")
		return
	}
	
	fi := models.IngredientFood{}
	if err = json.Unmarshal(body, &fi); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error unmarshalling the json")
		return
	}

	if err = h.svc.AddIngredient(req.Context(), fi); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error adding ingredient to food")
		return
	}
}

func (h *FoodHandler) deleteIngredient(resp http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error reading the request body")
		return
	}

	fi := models.IngredientFood{}
	if err = json.Unmarshal(body, &fi); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error unmarshalling the json")
		return
	}

	if err = h.svc.DeleteIngredient(req.Context(), fi); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error deleting ingredient from food")
		return
	}
}