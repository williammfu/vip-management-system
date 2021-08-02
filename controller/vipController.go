package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/williammfu/vip-management-system/config"
	"github.com/williammfu/vip-management-system/model"
	"gorm.io/gorm"
)

var db *gorm.DB

type Response struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

type ResponseData struct {
	Ok      bool      `json:"ok"`
	Data    model.Vip `json:"data"`
	Message string    `json:"message"`
}

type ResponseDataList struct {
	Ok      bool        `json:"ok"`
	Data    []model.Vip `json:"data"`
	Message string      `json:"message"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func StoreVip(w http.ResponseWriter, r *http.Request) {
	var vip model.Vip

	w.Header().Set("Content-type", "application/json")
	json.NewDecoder(r.Body).Decode(&vip)
	guest := vip.GetGuestInfo()
	resGuest := db.Create(&guest)

	if resGuest.Error == nil {
		desc := vip.GetDescriptions(guest.ID)
		resDesc := db.Create(&desc)

		if resDesc.Error != nil {
			json.NewEncoder(w).Encode(Response{false, "Error"})
		} else {
			json.NewEncoder(w).Encode(Response{true, "Success"})
		}

	} else {
		json.NewEncoder(w).Encode(Response{false, "Error"})
	}
}

func RetrieveVip(w http.ResponseWriter, r *http.Request) {
	var guest model.Guest
	w.Header().Set("Content-Type", "application/json")

	idVip := mux.Vars(r)["id"]
	result := db.Where("id = ?", idVip).First(&guest)

	if result.Error != nil {
		json.NewEncoder(w).Encode(Response{false, "Failed"})
	} else {
		var descriptions []model.Description
		db.Where("id = ?", guest.ID).Find(&descriptions)
		json.NewEncoder(w).Encode(ResponseData{true, model.CreateVip(guest, descriptions), "Success"})
	}
}

func RetrieveAllVips(w http.ResponseWriter, r *http.Request) {
	var descriptions []model.Description
	var guests []model.Guest
	var vips []model.Vip

	w.Header().Set("Content-Type", "application/json")
	db.Find(&guests)

	for _, guest := range guests {
		db.Where("id = ?", guest.ID).Find(&descriptions)
		vips = append(vips, model.CreateVip(guest, descriptions))
	}

	json.NewEncoder(w).Encode(ResponseDataList{true, vips, "success"})
}

func UpdateVip(w http.ResponseWriter, r *http.Request) {
	var newVip model.Vip
	var oldGuest, newGuest model.Guest
	var description model.Description
	var descriptions []model.Description
	var err error

	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&newVip)
	newGuest = newVip.GetGuestInfo()
	newGuest.ID, err = strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		json.NewEncoder(w).Encode(Response{false, "Failed"})
	} else {
		db.Where("id = ?", mux.Vars(r)["id"]).First(&oldGuest)
		oldGuest.AssignGuest(newGuest)
		db.Save(&oldGuest)

		descriptions = newVip.GetDescriptions(newGuest.ID)
		db.Where("id = ?", mux.Vars(r)["id"]).Delete(&description)
		db.Create(&descriptions)

		json.NewEncoder(w).Encode(Response{true, "Success"})
	}
}

type Arrival struct {
	Arrived bool `json:"arrived"`
}

func ArrivedVip(w http.ResponseWriter, r *http.Request) {
	var guest model.Guest
	var arrival Arrival

	w.Header().Set("Content-Type", "application/json")
	db.Where("id = ?", mux.Vars(r)["id"]).First(&guest)

	json.NewDecoder(r.Body).Decode(&arrival)
	guest.Arrived = arrival.Arrived
	result := db.Save(&guest)

	if result.Error != nil {
		json.NewEncoder(w).Encode(Response{false, "Failed"})
	} else {
		json.NewEncoder(w).Encode(Response{true, "Success"})
	}
}
