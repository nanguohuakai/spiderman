package dto

type BuildList struct {
	CountryCode  string `json:"country_code"`
	CountryName  string `json:"country_name"`
	ProvinceCode string `json:"province_code"`
	ProvinceName string `json:"province_name"`
	CityCode     string `json:"city_code"`
	CityName     string `json:"city_name"`
	DistrictCode string `json:"district_code"`
	DistrictName string `json:"district_name"`
	BuildingCode string `json:"building_code"`
	BuildingName string `json:"building_name"`
	FloorCode    string `json:"floor_code"`
	FloorName    string `json:"floor_name"`
}
