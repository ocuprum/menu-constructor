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

type CategoryHandler struct {
	svc *services.CategoryService
}

func NewCategoryHandler(svc *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{svc: svc}
}

func (h *CategoryHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc(pkgHTTP.GetPath("/category/{id}"), h.getByID)
	mux.HandleFunc(pkgHTTP.GetPath("/categories/paginate"), h.paginate)
	mux.HandleFunc(pkgHTTP.PostPath("/category/create"), h.create)
	mux.HandleFunc(pkgHTTP.PutPath("/category/change"), h.change)
	mux.HandleFunc(pkgHTTP.DeletePath("/categories/delete"), h.delete)
	mux.HandleFunc(pkgHTTP.PostPath("/category/food/add"), h.addFood)
	mux.HandleFunc(pkgHTTP.DeletePath("/category/food/delete"), h.deleteFood)
}

func (h *CategoryHandler) getByID(resp http.ResponseWriter, req *http.Request) {
	id, err := uuid.Parse(req.PathValue("id"))
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error parsing the id")
		return
	}

	category, err := h.svc.GetByID(req.Context(), id)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error retrieving the category from the db")
		return
	}

	categoryJSON, err := json.MarshalIndent(category, "", "  ")
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error preparing category JSON")
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(categoryJSON)	
}

func (h *CategoryHandler) paginate(resp http.ResponseWriter, req *http.Request) {
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

	categories, err := h.svc.Paginate(req.Context(), limit, offset)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error retrieving categories from the db")
		return
	}

	categoriesJSON, err := json.MarshalIndent(categories, "", "  ")
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error preparing categories JSON")
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(categoriesJSON)	
}

func (h *CategoryHandler) create(resp http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest)
		return
	}
	
	category := models.Category{}
	if err = json.Unmarshal(body, &category); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error unmarshalling the json")
		return
	}

	if err = h.svc.Create(req.Context(), category); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error creating category")
		return
	}
}

func (h *CategoryHandler) change(resp http.ResponseWriter, req *http.Request) {	
	body, err := io.ReadAll(req.Body)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error reading the request body")
	}

	category := models.Category{}
	if err = json.Unmarshal(body, &category); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error unmarshalling the json")
		return
	}

	if err = h.svc.Change(req.Context(), category); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error changing the category")
	}
}

func (h *CategoryHandler) delete(resp http.ResponseWriter, req *http.Request) {
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
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error deleting categories")
		return
	}
}

func (h *CategoryHandler) addFood(resp http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error reading the request body")
		return
	}
	
	fc := models.FoodCategory{}
	if err = json.Unmarshal(body, &fc); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error unmarshalling the json")
		return
	}

	if err = h.svc.AddFood(req.Context(), fc); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error adding food to category")
		return
	}
}

func (h *CategoryHandler) deleteFood(resp http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error reading the request body")
		return
	}

	fc := models.FoodCategory{}
	if err = json.Unmarshal(body, &fc); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error unmarshalling the json")
		return
	}

	if err = h.svc.DeleteFood(req.Context(), fc); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error deleting food from category")
		return
	}
}