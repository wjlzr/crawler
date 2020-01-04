package GaodeMap

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
)

const (
	GAT_WAY  = "https://restapi.amap.com/v3/geocode/regeo"
	GaodeMapKey = "881b1b4f317ba7714b3956475625199f"
)

type GeoCode struct {
	Status    string           `json:"status"`
	Regeocode GeoCodeRegeocode `json:"regeocode"`
	Info      string           `json:"info"`
	Infocode  string           `json:"infocode"`
}

type GeoCodeRoads struct {
	ID        string `json:"id"`
	Location  string `json:"location"`
	Direction string `json:"direction"`
	Name      string `json:"name"`
	Distance  string `json:"distance"`
}

type GeoCodeRoadinters struct {
	SecondName string `json:"second_name"`
	FirstID    string `json:"first_id"`
	SecondID   string `json:"second_id"`
	Location   string `json:"location"`
	Distance   string `json:"distance"`
	FirstName  string `json:"first_name"`
	Direction  string `json:"direction"`
}

type GeoCodeStreetNumber struct {
	Number    string `json:"number"`
	Location  string `json:"location"`
	Direction string `json:"direction"`
	Distance  string `json:"distance"`
	Street    string `json:"street"`
}

type GeoCodeBusinessAreas struct {
	Location string `json:"location"`
	Name     string `json:"name"`
	ID       string `json:"id"`
}

type GeoCodeBuilding struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type GeoCodeNeighborhood struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type GeoCodeAddressComponent struct {
	City          []interface{}          `json:"city"`
	Province      string                 `json:"province"`
	Adcode        string                 `json:"adcode"`
	District      string                 `json:"district"`
	Towncode      string                 `json:"towncode"`
	StreetNumber  GeoCodeStreetNumber    `json:"streetNumber"`
	Country       string                 `json:"country"`
	Township      string                 `json:"township"`
	BusinessAreas []GeoCodeBusinessAreas `json:"businessAreas"`
	Building      GeoCodeBuilding        `json:"building"`
	Neighborhood  GeoCodeNeighborhood    `json:"neighborhood"`
	Citycode      string                 `json:"citycode"`
}

type GeoCodeAois struct {
	Area     string `json:"area"`
	Type     string `json:"type"`
	ID       string `json:"id"`
	Location string `json:"location"`
	Adcode   string `json:"adcode"`
	Name     string `json:"name"`
	Distance string `json:"distance"`
}

type GeoCodePois struct {
	ID           string `json:"id"`
	Direction    string `json:"direction"`
	Businessarea string `json:"businessarea"`
	Address      string `json:"address"`
	Poiweight    string `json:"poiweight"`
	Name         string `json:"name"`
	Location     string `json:"location"`
	Distance     string `json:"distance"`
	Tel          string `json:"tel"`
	Type         string `json:"type"`
}

type GeoCodeRegeocode struct {
	Roads            []GeoCodeRoads          `json:"roads"`
	Roadinters       []GeoCodeRoadinters     `json:"roadinters"`
	FormattedAddress string                  `json:"formatted_address"`
	AddressComponent GeoCodeAddressComponent `json:"addressComponent"`
	Aois             []GeoCodeAois           `json:"aois"`
	Pois             []GeoCodePois           `json:"pois"`
}

func FindCityByCoordinate(x_y string)(res GeoCodeAddressComponent){
	var geoCode GeoCode
	url := GAT_WAY + fmt.Sprintf("?location=%s&output=json&key=%s&extensions=all&batch=false&roadlevel=0", x_y, GaodeMapKey)
	fmt.Println("\n 逆地理编码url------>" + url+"\n")
	//返回json字符串
	response, _ := httplib.Get(url).String()

	_ = json.Unmarshal([]byte(string(response)), &geoCode)

	if geoCode.Status != "1" {
		panic("高德逆地理接口请求失败")
	}

	return geoCode.Regeocode.AddressComponent
}
