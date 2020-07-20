package handler

import (
	"encoding/json"
	"goproject/app/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//GetAllPersons ...
func GetAllPersons(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	persons := []model.Person{}
	db.Find(&persons)
	respondJSON(w, http.StatusOK, persons)
}

//CreatePerson ...
func CreatePerson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	person := model.Person{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&person); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&person).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, person)
}

//GetPerson ...
func GetPerson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	person := getPersonOr404(db, id, w, r)
	if person == nil {
		return
	}
	respondJSON(w, http.StatusOK, person)
}

//UpdatePerson ...
func UpdatePerson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	person := getPersonOr404(db, id, w, r)
	if person == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&person); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&person).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, person)
}

//DeletePerson ...
func DeletePerson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	person := getPersonOr404(db, id, w, r)
	if person == nil {
		return
	}
	if err := db.Delete(&person).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getPersonOr404(db *gorm.DB, personID int, w http.ResponseWriter, r *http.Request) *model.Person {
	person := model.Person{}
	if err := db.First(&person, model.Person{Model: gorm.Model{ID: uint(personID)}}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &person
}

func checkPerson(db *gorm.DB, personID int, w http.ResponseWriter, r *http.Request) *model.Person {
	person := model.Person{}
	if err := db.First(&person, model.Person{Model: gorm.Model{ID: uint(personID)}}).Error; err != nil {
		return nil
	}
	return &person
}
