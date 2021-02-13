package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"project/common"
	"project/data/model"
	"strconv"

	"github.com/gorilla/mux"
)

//SaveVehicle is...
func SaveVehicle(r *http.Request) error {
	connection := common.GetDatabase()
	image, _, err := r.FormFile("image")
	defer common.Closedatabase(connection)
	if err != nil {
		return err
	}
	imagebyte, err := ioutil.ReadAll(image)
	if err != nil {
		return err
	}
	bookimagestring := base64.StdEncoding.EncodeToString(imagebyte)
	companyid, err := strconv.Atoi(r.FormValue("company"))
	if err != nil {
		return err
	}
	vehicle := model.Vehicle{
		Vin:         r.FormValue("vin"),
		Year:        r.FormValue("year"),
		ModelName:   r.FormValue("model"),
		TitleNumber: r.FormValue("title"),
		Price:       r.FormValue("price"),
		FuelType:    r.FormValue("fueltype"),
		Mileage:     r.FormValue("mileage"),
		Stock:       r.FormValue("stock"),
		Image:       bookimagestring,
		CompanyID:   uint(companyid),
	}
	connection.Create(&vehicle)
	return nil
}

//GetAllVehicle is...
func GetAllVehicle() []model.Vehicle {
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	var vehicles []model.Vehicle
	connection.Find(&vehicles)
	return vehicles
}

//GetOneVehicle is....
func GetOneVehicle(r *http.Request) model.Vehicle {
	id := mux.Vars(r)["id"]
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	var vehicle model.Vehicle
	connection.First(&vehicle, id)
	return vehicle
}

//DeleteOneVehicle is...
func DeleteOneVehicle(r *http.Request) {
	id := mux.Vars(r)["id"]
	var vehicle model.Vehicle
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	connection.Delete(&vehicle, id)
}

//UpdateVehicle is...
func UpdateVehicle(r *http.Request) ([]byte, error) {
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	id := mux.Vars(r)["id"]
	var findvehicle model.Vehicle
	connection.First(&findvehicle, id)
	bodydata, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(r.Body)
	err = json.Unmarshal(bodydata, &findvehicle)
	if err != nil {
		return nil, err
	}
	connection.Save(&findvehicle)
	bytedata, err := json.MarshalIndent(findvehicle, "", "  ")
	if err != nil {
		return nil, err
	}
	return bytedata, nil

}

//SaveCompanyLogo is...
func SaveCompanyLogo(r *http.Request) error {
	connection := common.GetDatabase()
	image, _, err := r.FormFile("logo")
	defer common.Closedatabase(connection)
	if err != nil {
		return err
	}
	imagebyte, err := ioutil.ReadAll(image)
	if err != nil {
		return err
	}
	logoimagestring := base64.StdEncoding.EncodeToString(imagebyte)
	company := model.Company{
		Name: r.FormValue("companyname"),
		Logo: logoimagestring,
	}
	connection.Create(&company)
	return nil
}

//GetAllVehicleComapnyLogo is...
func GetAllVehicleComapnyLogo(r *http.Request) []model.Company {
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	var companies []model.Company
	connection.Find(&companies)
	return companies
}
