package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go-microservice/models"
	"go-microservice/services"
	"go-microservice/utils"
)

type UserHandler struct {
	Service *services.UserService
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.Service.GetAll())
}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, ok := h.Service.GetByID(id)
	if !ok {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	saved := h.Service.Create(user)
	go utils.LogAction("CREATE", saved.ID)
	json.NewEncoder(w).Encode(saved)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updated, ok := h.Service.Update(id, user)
	if !ok {
		http.NotFound(w, r)
		return
	}
	go utils.LogAction("UPDATE", updated.ID)
	json.NewEncoder(w).Encode(updated)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if !h.Service.Delete(id) {
		http.NotFound(w, r)
		return
	}
	go utils.LogAction("DELETE", id)
	w.WriteHeader(http.StatusNoContent)
}
