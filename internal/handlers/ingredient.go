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

type IngredientHandler struct {
	svc *services.IngredientService
}

func NewIngredientHandler(svc *services.IngredientService) *IngredientHandler {
	return &IngredientHandler{svc: svc}
}

func (h *IngredientHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc(pkgHTTP.GetPath("/ingredients/{id}"), h.getByID)
	mux.HandleFunc(pkgHTTP.GetPath("/ingredients/paginate"), h.paginate)
	mux.HandleFunc(pkgHTTP.PostPath("/ingredients/create"), h.create)
	mux.HandleFunc(pkgHTTP.PutPath("/ingredients/change"), h.change)
	mux.HandleFunc(pkgHTTP.DeletePath("/ingredients/delete"), h.delete)
}

func (h *IngredientHandler) getByID(resp http.ResponseWriter, req *http.Request) {
	id, err := uuid.Parse(req.PathValue("id"))
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error parsing the id")
		return
	}

	ingredient, err := h.svc.GetByID(req.Context(), id)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error retrieving the ingredient from the db")
		return
	}

	ingredientJSON, err := json.MarshalIndent(ingredient, "", "  ")
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error preparing ingredient JSON")
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(ingredientJSON)
}

func (h *IngredientHandler) paginate(resp http.ResponseWriter, req *http.Request) {
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

	ingredients, err := h.svc.Paginate(req.Context(), limit, offset)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error retrieving ingredients from the db")
		return
	}

	ingredientsJSON, err := json.MarshalIndent(ingredients, "", "  ")
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error preparing ingredients JSON")
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(ingredientsJSON)
}

func (h *IngredientHandler) create(resp http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error reading the request body")
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

func (h *IngredientHandler) change(resp http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error reading the request body")
	}

	ingred := models.Ingredient{}
	if err = json.Unmarshal(body, &ingred); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusBadRequest, "Error unmarshalling the json")
		return
	}

	if err = h.svc.Change(req.Context(), ingred); err != nil {
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error changing the ingredient")
	}
}

func (h *IngredientHandler) delete(resp http.ResponseWriter, req *http.Request) {
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
		pkgHTTP.WriteResponse(resp, http.StatusInternalServerError, "Error deleting ingredients")
		return
	}
}