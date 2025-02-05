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

type MealHandler struct {
	svc *services.MealService
}

func NewMealHandler(svc *services.MealService) *MealHandler {
	return &MealHandler{svc: svc}
}

func (h *MealHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc(pkgHTTP.GetPath("/meal/{id}"), h.getByID)
	mux.HandleFunc(pkgHTTP.GetPath("/meals/paginate"), h.paginate)
	mux.HandleFunc(pkgHTTP.PostPath("/meal/create"), h.create)
	mux.HandleFunc(pkgHTTP.PutPath("/meal/change"), h.change)
	mux.HandleFunc(pkgHTTP.DeletePath("/meals/delete"), h.delete)
	mux.HandleFunc(pkgHTTP.PostPath("/meal/food/add"), h.addFood)
	mux.HandleFunc(pkgHTTP.DeletePath("/meal/food/delete"), h.deleteFood)
}

func (h *MealHandler) getByID(resp http.ResponseWriter, req *http.Request) {
	id, err := uuid.Parse(req.PathValue("id"))
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error parsing the id")
		return
	}

	meal, err := h.svc.GetByID(req.Context(), id)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error retrieving the meal from the db")
		return
	}

	mealJSON, err := json.MarshalIndent(meal, "", "  ")
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error preparing meal JSON")
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(mealJSON)	
}

func (h *MealHandler) paginate(resp http.ResponseWriter, req *http.Request) {
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

	meals, err := h.svc.Paginate(req.Context(), limit, offset)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error retrieving meals from the db")
		return
	}

	mealsJSON, err := json.MarshalIndent(meals, "", "  ")
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error preparing meals JSON")
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(mealsJSON)	
}

func (h *MealHandler) create(resp http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest)
		return
	}
	
	meal := models.Meal{}
	if err = json.Unmarshal(body, &meal); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error unmarshalling the json")
		return
	}

	if err = h.svc.Create(req.Context(), meal); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error creating meal")
		return
	}
}

func (h *MealHandler) change(resp http.ResponseWriter, req *http.Request) {	
	body, err := io.ReadAll(req.Body)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error reading the request body")
	}

	meal := models.Meal{}
	if err = json.Unmarshal(body, &meal); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error unmarshalling the json")
		return
	}

	if err = h.svc.Change(req.Context(), meal); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error changing the meal")
	}
}

func (h *MealHandler) delete(resp http.ResponseWriter, req *http.Request) {
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
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error deleting meals")
		return
	}
}

func (h *MealHandler) addFood(resp http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error reading the request body")
		return
	}
	
	fm := models.FoodMeal{}
	if err = json.Unmarshal(body, &fm); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error unmarshalling the json")
		return
	}

	if err = h.svc.AddFood(req.Context(), fm); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error adding food to meal")
		return
	}
}

func (h *MealHandler) deleteFood(resp http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error reading the request body")
		return
	}

	fm := models.FoodMeal{}
	if err = json.Unmarshal(body, &fm); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error unmarshalling the json")
		return
	}

	if err = h.svc.DeleteFood(req.Context(), fm); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error deleting food from meal")
		return
	}
}