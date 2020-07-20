package handler

import (
	"app/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//GetAllNationalities ...
func GetAllNationalities(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	nationalities := []model.Nationality{}
	db.Find(&nationalities)
	respondJSON(w, http.StatusOK, nationalities)
}

//CreateNationality ...
func CreateNationality(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	nationality := model.Nationality{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&nationality); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&nationality).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, nationality)
}

//GetNationality ...
func GetNationality(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["NationalityID"])
	if err != nil {
		return
	}
	nationality := getGenderOr404(db, id, w, r)
	if nationality == nil {
		return
	}
	respondJSON(w, http.StatusOK, nationality)
}

//UpdateNationality ...
func UpdateNationality(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["NationalityID"])
	if err != nil {
		return
	}
	nationality := getNationalityOr404(db, id, w, r)
	if nationality == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&nationality); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&nationality).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, nationality)
}

//DeleteNationality ...
func DeleteNationality(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["NationalityID"])
	if err != nil {
		return
	}
	nationality := getNationalityOr404(db, id, w, r)
	if nationality == nil {
		return
	}
	if err := db.Delete(&nationality).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getNationalityOr404(db *gorm.DB, nationalityID int, w http.ResponseWriter, r *http.Request) *model.Nationality {
	nationality := model.Nationality{}
	if err := db.First(&nationality, model.Nationality{Model: gorm.Model{ID: uint(nationalityID)}}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &nationality
}
