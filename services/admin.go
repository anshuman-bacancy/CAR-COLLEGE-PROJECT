package services

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"project/common"
	"project/models"
	"strconv"

	"github.com/gorilla/mux"
)

// saves vehicle to db
func SaveVehicle(r *http.Request) error {
	connection := common.GetDatabase()
	image, _, err := r.FormFile("image")
	defer common.CloseDatabase(connection)
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
	vehicle := models.Vehicle{
		Vin:         r.FormValue("vin"),
		Year:        r.FormValue("year"),
		ModelName:   r.FormValue("models"),
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

// returns all vehicles
func GetAllVehicle() []models.Vehicle {
	connection := common.GetDatabase()
	defer common.CloseDatabase(connection)
	var vehicles []models.Vehicle
	connection.Find(&vehicles)
	return vehicles
}

// returns one vehicles
func GetOneVehicle(r *http.Request) models.Vehicle {
	id := mux.Vars(r)["id"]
	connection := common.GetDatabase()
	defer common.CloseDatabase(connection)
	var vehicle models.Vehicle
	connection.First(&vehicle, id)
	return vehicle
}

// deletes one vehicle
func DeleteOneVehicle(r *http.Request) {
	id := mux.Vars(r)["id"]
	var vehicle models.Vehicle
	connection := common.GetDatabase()
	defer common.CloseDatabase(connection)
	connection.Delete(&vehicle, id)
}

// update one vehicle
func UpdateVehicle(r *http.Request) ([]byte, error) {
	connection := common.GetDatabase()
	defer common.CloseDatabase(connection)
	id := mux.Vars(r)["id"]
	var findvehicle models.Vehicle
	connection.First(&findvehicle, id)
	bodydata, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
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

// save company logo
func SaveCompanyLogo(r *http.Request) error {
	connection := common.GetDatabase()
	image, _, err := r.FormFile("logo")
	defer common.CloseDatabase(connection)
	if err != nil {
		return err
	}
	imagebyte, err := ioutil.ReadAll(image)
	if err != nil {
		return err
	}
	logoimagestring := base64.StdEncoding.EncodeToString(imagebyte)
	company := models.Company{
		Name: r.FormValue("companyname"),
		Logo: logoimagestring,
	}
	connection.Create(&company)
	return nil
}

// returns all brands
func GetAllBrands(r *http.Request) []models.Company {
	connection := common.GetDatabase()
	defer common.CloseDatabase(connection)
	var companies []models.Company
	connection.Find(&companies)
	return companies
}

// delete one brand
func DeleteOneBrand(r *http.Request) {
	id := mux.Vars(r)["id"]
	var company models.Company
	connection := common.GetDatabase()
	defer common.CloseDatabase(connection)
	connection.Delete(&company, id)
}

// returns one brand
func GetOneBrand(r *http.Request) models.Company {
	id := mux.Vars(r)["id"]
	connection := common.GetDatabase()
	defer common.CloseDatabase(connection)
	var company models.Company
	connection.First(&company, id)
	return company
}

// returns one brand by ID
func GetOneBrandNameByID(id uint) string {
	connection := common.GetDatabase()
	defer common.CloseDatabase(connection)
	var company models.Company
	connection.First(&company, id)
	return company.Name
}

// return brand image
func GetOneBrandImageByID(id uint) string {
	connection := common.GetDatabase()
	defer common.CloseDatabase(connection)
	var company models.Company
	connection.First(&company, id)
	return company.Logo
}

// update brand
func UpdateBrand(r *http.Request) ([]byte, error) {
	connection := common.GetDatabase()
	defer common.CloseDatabase(connection)
	id := mux.Vars(r)["id"]
	var company models.Company
	connection.First(&company, id)
	bodydata, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodydata, &company)
	if err != nil {
		return nil, err
	}
	connection.Save(&company)
	bytedata, err := json.MarshalIndent(company, "", "  ")
	if err != nil {
		return nil, err
	}
	return bytedata, nil
}

// return vehicles of brand
func GetParticularBrandVehicle(id uint) []models.Vehicle {
	connection := common.GetDatabase()
	defer common.CloseDatabase(connection)
	var vehicles []models.Vehicle
	connection.Where("company_id = ?", id).Find(&vehicles)
	return vehicles
}

// same as above. remove it ?
func GetParticularBrandVehiclewithR(r *http.Request) []models.Vehicle {
	id := mux.Vars(r)["id"]
	connection := common.GetDatabase()
	defer common.CloseDatabase(connection)
	var vehicles []models.Vehicle
	connection.Where("company_id = ?", id).Find(&vehicles)
	return vehicles
}

// returns one vehicle by ID
func GetOneVehicleNameByID(vehicleid uint) string {
	connection := common.GetDatabase()
	defer common.CloseDatabase(connection)
	var vehicle models.Vehicle
	connection.First(&vehicle, vehicleid)
	return vehicle.ModelName
}

// returns vehicle logo
func GetOneVehicleImageByID(vehicleid uint) string {
	connection := common.GetDatabase()
	defer common.CloseDatabase(connection)
	var vehicle models.Vehicle
	connection.First(&vehicle, vehicleid)
	return vehicle.Image
}

// return vehicle brand ID
func GetVehicleBrandID(vehicleid uint) uint {
	connection := common.GetDatabase()
	defer common.CloseDatabase(connection)
	var vehicle models.Vehicle
	connection.First(&vehicle, vehicleid)
	return vehicle.CompanyID
}

// save admin
func SaveAdmin(r *http.Request) (bool, error) {
	connection := common.GetDatabase()
	defer common.CloseDatabase(connection)
	sqldb := connection.DB()
	rows, err := sqldb.Query("SELECT email FROM sales_people")
	defer rows.Close()
	if err != nil {
		return false, err
	}
	emails := make(map[string]bool)
	for rows.Next() {
		var email string
		rows.Scan(&email)
		emails[email] = true
	}

	if emails[r.FormValue("email")] {
		return true, nil
	}
	salesPerson := models.SalesPerson{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
		Mobile:   r.FormValue("mobilenumber"),
		City:     r.FormValue("city"),
	}
	connection.Create(&salesPerson)
	return false, nil
}

// returns all admin
func GetAllAdmin(r *http.Request) []models.SalesPerson {
	connection := common.GetDatabase()
	defer common.CloseDatabase(connection)
	var salesperson []models.SalesPerson
	connection.Find(&salesperson)
	return salesperson
}

// returns one admin
func GetOneAdminByEmail(email interface{}) models.SalesPerson {
	connection := common.GetDatabase()
	defer common.CloseDatabase(connection)
	var admin models.SalesPerson
	connection.Where("email = 	?", email).First(&admin)
	return admin
}

// update admin
func AdminUpdate(r *http.Request) ([]byte, error) {
	connection := common.GetDatabase()
	defer common.CloseDatabase(connection)
	id := mux.Vars(r)["id"]
	var admin models.SalesPerson
	connection.First(&admin, id)
	bodydata, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodydata, &admin)
	if err != nil {
		return nil, err
	}
	connection.Save(&admin)
	bytedata, err := json.MarshalIndent(admin, "", "  ")
	if err != nil {
		return nil, err
	}
	return bytedata, nil
}
