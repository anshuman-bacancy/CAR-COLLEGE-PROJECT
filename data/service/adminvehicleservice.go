package service

import (
	"encoding/base64"
	"encoding/json"
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

//GetAllBrand is...
func GetAllBrand(r *http.Request) []model.Company {
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	var companies []model.Company
	connection.Find(&companies)
	return companies
}

//DeleteOneBrand is..
func DeleteOneBrand(r *http.Request) {
	id := mux.Vars(r)["id"]
	var company model.Company
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	connection.Delete(&company, id)
}

//GetOneBrand is...
func GetOneBrand(r *http.Request) model.Company {
	id := mux.Vars(r)["id"]
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	var company model.Company
	connection.First(&company, id)
	return company
}

//GetOneBrandNameByID is...
func GetOneBrandNameByID(id uint) string {
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	var company model.Company
	connection.First(&company, id)
	return company.Name
}

//GetOneBrandImageByID is...
func GetOneBrandImageByID(id uint) string {
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	var company model.Company
	connection.First(&company, id)
	return company.Logo
}

//UpdateBrand is....
func UpdateBrand(r *http.Request) ([]byte, error) {
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	id := mux.Vars(r)["id"]
	var company model.Company
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

//GetParticlullarBrandVehicle is...
func GetParticlullarBrandVehicle(id uint) []model.Vehicle {
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	var vehicles []model.Vehicle
	connection.Where("company_id = ?", id).Find(&vehicles)
	return vehicles
}

//GetParticlullarBrandVehiclewithR is...
func GetParticlullarBrandVehiclewithR(r *http.Request) []model.Vehicle {
	id := mux.Vars(r)["id"]
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	var vehicles []model.Vehicle
	connection.Where("company_id = ?", id).Find(&vehicles)
	return vehicles
}

//GetOneVehicleNameByID is..
func GetOneVehicleNameByID(vehicleid uint) string {
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	var vehicle model.Vehicle
	connection.First(&vehicle, vehicleid)
	return vehicle.ModelName
}

//GetOneVehicleImageByID is..
func GetOneVehicleImageByID(vehicleid uint) string {
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	var vehicle model.Vehicle
	connection.First(&vehicle, vehicleid)
	return vehicle.Image
}

//GetVehicleBrandID is..
func GetVehicleBrandID(vehicleid uint) uint {
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	var vehicle model.Vehicle
	connection.First(&vehicle, vehicleid)
	return vehicle.CompanyID
}

//SaveAdmin is..
func SaveAdmin(r *http.Request) (bool, error) {
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
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
	salesPerson := model.SalesPerson{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
		Mobile:   r.FormValue("mobilenumber"),
		City:     r.FormValue("city"),
	}
	connection.Create(&salesPerson)
	return false, nil
}

//GetAllAdmin is..
func GetAllAdmin(r *http.Request) []model.SalesPerson {
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	var salesperson []model.SalesPerson
	connection.Find(&salesperson)
	return salesperson
}

//GetOneAdminBYemail is...
func GetOneAdminBYemail(email interface{}) model.SalesPerson {
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	var admin model.SalesPerson
	connection.Where("email = 	?", email).First(&admin)
	return admin
}

//AdminUpdate is...
func AdminUpdate(r *http.Request) ([]byte, error) {
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	id := mux.Vars(r)["id"]
	var admin model.SalesPerson
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
