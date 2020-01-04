package main

import (
	"crawler/common"
	"crawler/models"
	"crawler/services/GaodeMap"
	"encoding/json"
	"fmt"
	"math"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/astaxie/beego/httplib"
	"os"
	"strings"
)

const (
	GAT_WAY  = "https://restapi.amap.com/v3/place/text"
	geocodeRegeoEndPoint = "https://restapi.amap.com/v3/geocode/regeo"
	GaodeMapKey = "881b1b4f317ba7714b3956475625199f"
	PATH = "/Users/zhaoran/resource/高德品牌文件/"
)

type PoiRes struct {
	Status     string     `json:"status"`
	Count      string     `json:"count"`
	Info       string     `json:"info"`
	InfoCode   string     `json:"infocode"`
	Suggestion Suggestion `json:"suggestion"`
	Data       []Poi      `json:"pois"`
}

type Suggestion struct {
	Keywords json.RawMessage `json:"keywords"`
	Cities   []City          `json:"cities"`
}

type City struct {
	Name     string     `json:"name"`
	Num      string      `json:"num"`
	Citycode string     `json:"citycode"`
	Adcode   string     `json:"adcode"`
}

type Poi struct {
	ID           string          `json:"id"`
	Parent       json.RawMessage `json:"parent"`
	Childtype    json.RawMessage `json:"childtype"`
	Name         string          `json:"name"`
	Tag          json.RawMessage `json:"tag"`
	Type         string          `json:"type"`
	Typecode     string          `json:"typecode"`
	BizType      json.RawMessage `json:"biz_type"`
	Address      string          `json:"address"`
	Location     string          `json:"location"`
	Tel          string          `json:"tel"`
	Postcode     json.RawMessage `json:"postcode"`
	Website      json.RawMessage `json:"website"`
	Email        json.RawMessage `json:"email"`
	Pcode        string          `json:"pcode"`
	Pname        string          `json:"pname"`
	Citycode     string          `json:"citycode"`
	Cityname     string          `json:"cityname"`
	Adcode       string          `json:"adcode"`
	Adname       string          `json:"adname"`
	Importance   json.RawMessage `json:"importance"`
	Shopid       json.RawMessage `json:"shopid"`
	Shopinfo     string          `json:"shopinfo"`
	Poiweight    json.RawMessage `json:"poiweight"`
	Gridcode     string          `json:"gridcode"`
	Distance     json.RawMessage `json:"distance"`
	NaviPoiid    string          `json:"navi_poiid"`
	EntrLocation string          `json:"entr_location"`
	BusinessArea string          `json:"business_area"`
	ExitLocation json.RawMessage `json:"exit_location"`
	Match        string          `json:"match"`
	Recommend    string          `json:"recommend"`
	Timestamp    json.RawMessage `json:"timestamp"`
	Alias        string          `json:"alias"`
	IndoorMap    string          `json:"indoor_map"`
	IndoorData   IndoorData      `json:"indoor_data"`
	GroupbuyNum  string          `json:"groupbuy_num"`
	DiscountNum  string          `json:"discount_num"`
	BizExt       BizExt          `json:"biz_ext"`
	Event        json.RawMessage `json:"event"`
	Children     []Children      `json:"children"`
	Photos       []Photos        `json:"photos"`
}

type IndoorData struct {
	Cpid      json.RawMessage `json:"cpid"`
	Floor     json.RawMessage `json:"floor"`
	Truefloor json.RawMessage `json:"truefloor"`
	Cmsid     json.RawMessage `json:"cmsid"`
}

type BizExt struct {
	Rating json.RawMessage `json:"rating"`
	Cost   string          `json:"cost"`
}

type Children struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Sname    string `json:"sname"`
	Location string `json:"location"`
	Address  string `json:"address"`
	Distance string `json:"distance"`
	Subtype  string `json:"subtype"`
	Typecode string `json:"typecode"`
}

type Photos struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type CityPoiRes struct {
	Status     string     `json:"status"`
	Count      string     `json:"count"`
	Info       string     `json:"info"`
	InfoCode   string     `json:"infocode"`
	Suggestion Suggestion `json:"suggestion"`
	Data       []CityPoi  `json:"pois"`
}

type CitySuggestion struct {
	Keywords json.RawMessage `json:"keywords"`
	Cities   []CityCity          `json:"cities"`
}

type CityCity struct {
	Name     string     `json:"name"`
	Num      string      `json:"num"`
	Citycode string     `json:"citycode"`
	Adcode   string     `json:"adcode"`
}

type CityPoi struct {
	ID           string          `json:"id"`
	Parent       json.RawMessage `json:"parent"`
	Childtype    json.RawMessage `json:"childtype"`
	Name         string          `json:"name"`
	Tag          json.RawMessage `json:"tag"`
	Type         string          `json:"type"`
	Typecode     string          `json:"typecode"`
	BizType      json.RawMessage `json:"biz_type"`
	Address      string          `json:"address"`
	Location     string          `json:"location"`
	Tel          string          `json:"tel"`
	Postcode     json.RawMessage `json:"postcode"`
	Website      json.RawMessage `json:"website"`
	Email        json.RawMessage `json:"email"`
	Pcode        string          `json:"pcode"`
	Pname        string          `json:"pname"`
	Citycode     string          `json:"citycode"`
	Cityname     string          `json:"cityname"`
	Adcode       string          `json:"adcode"`
	Adname       string          `json:"adname"`
	Importance   json.RawMessage `json:"importance"`
	Shopid       json.RawMessage `json:"shopid"`
	Shopinfo     string          `json:"shopinfo"`
	Poiweight    json.RawMessage `json:"poiweight"`
	Gridcode     string          `json:"gridcode"`
	Distance     json.RawMessage `json:"distance"`
	NaviPoiid    string          `json:"navi_poiid"`
	EntrLocation string          `json:"entr_location"`
	BusinessArea string          `json:"business_area"`
	ExitLocation json.RawMessage `json:"exit_location"`
	Match        string          `json:"match"`
	Recommend    string          `json:"recommend"`
	Timestamp    json.RawMessage `json:"timestamp"`
	Alias        string          `json:"alias"`
	IndoorMap    string          `json:"indoor_map"`
	IndoorData   IndoorData      `json:"indoor_data"`
	GroupbuyNum  string          `json:"groupbuy_num"`
	DiscountNum  string          `json:"discount_num"`
	BizExt       BizExt          `json:"biz_ext"`
	Event        json.RawMessage `json:"event"`
	Children     []CityChildren      `json:"children"`
	Photos       []CityPhotos        `json:"photos"`
}

type CityIndoorData struct {
	Cpid      json.RawMessage `json:"cpid"`
	Floor     json.RawMessage `json:"floor"`
	Truefloor json.RawMessage `json:"truefloor"`
	Cmsid     json.RawMessage `json:"cmsid"`
}

type CityBizExt struct {
	Rating json.RawMessage `json:"rating"`
	Cost   string          `json:"cost"`
}

type CityChildren struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Sname    string `json:"sname"`
	Location string `json:"location"`
	Address  string `json:"address"`
	Distance string `json:"distance"`
	Subtype  string `json:"subtype"`
	Typecode string `json:"typecode"`
}

type CityPhotos struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func main() {
	//beego.Run()
	args := os.Args
	if args[2] == "1"{
		query(args[1],args[1],"",args[2],args[4])
	} else {
		query(args[1],args[1],args[3],args[2],args[4])
	}
}

func query(brand string, fileName string, cityName string, type1 string, g_types string) {

	var page, offset int8 = 1, 20

	var key_word, types string = brand, g_types

	//client := &http.Client{}

	if type1 == "1"{
		url1 := GAT_WAY + fmt.Sprintf("?key=%s&keywords=%s&types=%s&offset=%d&page=%d&extensions=all&children=1", GaodeMapKey, key_word, types, offset, page)
		fmt.Println("url1-------->"+url1)
		//返回json字符串
		response1, _ := httplib.Get(url1).String()

		var poiRes PoiRes
		var cityPoiRes CityPoiRes
		_ = json.Unmarshal([]byte(string(response1)), &poiRes)
		if poiRes.Status != "1" {
			panic("高德接口请求失败1")
		}
		//fmt.Println(poiRes.Suggestion)

		if len(poiRes.Suggestion.Cities) != 0{
			//fmt.Println(poiRes.Suggestion.Cities)
			//allPoi := []string{}
			row := 2
			f := excelize.NewFile()
			index := f.NewSheet("Sheet1")
			f.SetCellValue("Sheet1", "A1", "名称")
			f.SetCellValue("Sheet1", "B1", "地址")
			f.SetCellValue("Sheet1", "C1", "省")
			f.SetCellValue("Sheet1", "D1", "城市")
			f.SetCellValue("Sheet1", "E1", "区域")
			f.SetCellValue("Sheet1", "F1", "街道")
			f.SetCellValue("Sheet1", "G1", "经度")
			f.SetCellValue("Sheet1", "H1", "纬度")
			f.SetCellValue("Sheet1", "I1", "ID")
			f.SetCellValue("Sheet1", "J1", "照片")
			for _, city := range poiRes.Suggestion.Cities{
				//fmt.Println(city)
				page2 := 1
				for{
					url2 := GAT_WAY + fmt.Sprintf("?key=%s&keywords=%s&types=%s&city=%s&offset=%d&page=%d&extensions=all&children=1", GaodeMapKey, key_word, types, city.Name, offset, page2)
					fmt.Println("\n url2------>" + url2+"\n")
					//返回json字符串
					response2, _ := httplib.Get(url2).String()

					_ = json.Unmarshal([]byte(string(response2)), &cityPoiRes)
					if cityPoiRes.Status != "1" {
						panic("高德接口请求失败2")
					}
					//fmt.Println(cityPoiRes.Data)
					for _, poi := range cityPoiRes.Data{
						//fmt.Println("\n",poi)
						//fmt.Println(reflect.TypeOf(cityPoiV))
						//for _,poi := range cityPoiV {
						// fmt.Println(poi)
						//}
						fmt.Println(fmt.Sprintf("%s,%d","A",row))
						//fmt.Println(poi)
						//查adcode上级省code和citycode
						var china_division_join models.China_division_join
						code := models.FindParentsCodeByAdcode(poi.Adcode)
						if code == china_division_join{
							continue
						}
						f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","A",row), poi.Name)
						f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","B",row), poi.Address)
						//省
						pName, _ := models.GetChina_divisionByCode(code.ProvideCode)
						if pName.Name != ""{
							f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","C",row), pName.Name)
						} else {
							f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","C",row), poi.Pname)
						}
						//市
						cName, _ := models.GetChina_divisionByCode(code.CityCode)
						if cName.Name != ""{
							f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","D",row), cName.Name)
						} else {
							f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","D",row), poi.Cityname)
						}
						//区
						aName, _ := models.GetChina_divisionByCode(poi.Adcode)
						if aName.Name != ""{
							f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","E",row), aName.Name)
						} else {
							f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","E",row), poi.Adname)
						}
						//街道
						street := GaodeMap.FindCityByCoordinate(poi.Location)
						f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","F",row), street.Township)
						locations := strings.Split(poi.Location,",")
						if len(locations) > 0{
							f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","G",row), locations[0])
							f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","H",row), locations[1])
						}
						f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","I",row), poi.ID)
						if len(poi.Photos) >0 {
							f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","J",row), poi.Photos[0].URL)
						}
						//}
						row ++
					}
					page2 ++
					if len(cityPoiRes.Data) == 0 {
						break
					}
				}
			}
			f.SetActiveSheet(index)
			//根据指定路径保存文件
			rsu := common.Mkdir(PATH)
			if err := f.SaveAs(rsu+fileName+".csv"); err != nil {
				println(err.Error())
			}
			fmt.Println("高德城市数据爬取成功")
		}
	} else {
		//只查询城市
		url1 := GAT_WAY + fmt.Sprintf("?key=%s&keywords=%s&types=%s&city=%s&offset=%d&page=%d&extensions=all&children=1", GaodeMapKey, key_word, types, cityName ,offset, page)
		fmt.Println("url1-------->"+url1)
		//返回json字符串
		response1, _ := httplib.Get(url1).String()

		var poiRes PoiRes
		var cityPoiRes CityPoiRes
		_ = json.Unmarshal([]byte(string(response1)), &poiRes)
		if poiRes.Status != "1" {
			panic("1高德接口请求失败：根据城市")
		}
		//fmt.Println(poiRes.Suggestion)

		if len(poiRes.Data) != 0{
			row := 2
			f := excelize.NewFile()
			index := f.NewSheet("Sheet1")
			f.SetCellValue("Sheet1", "A1", "名称")
			f.SetCellValue("Sheet1", "B1", "地址")
			f.SetCellValue("Sheet1", "C1", "省")
			f.SetCellValue("Sheet1", "D1", "城市")
			f.SetCellValue("Sheet1", "E1", "区域")
			f.SetCellValue("Sheet1", "F1", "街道")
			f.SetCellValue("Sheet1", "G1", "经度")
			f.SetCellValue("Sheet1", "H1", "纬度")
			f.SetCellValue("Sheet1", "I1", "ID")
			f.SetCellValue("Sheet1", "J1", "照片")
			//page2 := 1
			count := math.Ceil((float64(899)/float64(20)))
			for i:=1; i < int(count); i++ {
				url2 := GAT_WAY + fmt.Sprintf("?key=%s&keywords=%s&types=%s&city=%s&offset=%d&page=%d&extensions=all&children=1", GaodeMapKey, key_word, types, cityName ,offset, i)
				fmt.Println("url2-------->"+url2)
				//返回json字符串
				response2, _ := httplib.Get(url2).String()

				_ = json.Unmarshal([]byte(string(response2)), &cityPoiRes)
				if cityPoiRes.Status != "1" {
					panic("2高德接口请求失败：根据城市")
				}
				fmt.Println(fmt.Sprintf("%s,%d","A",row))
				//fmt.Println(poi)
				//查adcode上级省code和citycode
				for _, poi := range cityPoiRes.Data{
					var china_division_join models.China_division_join
					code := models.FindParentsCodeByAdcode(poi.Adcode)
					if code == china_division_join{
						continue
					}
					f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","A",row), poi.Name)
					f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","B",row), poi.Address)
					//省
					pName, _ := models.GetChina_divisionByCode(code.ProvideCode)
					if pName.Name != ""{
						f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","C",row), pName.Name)
					} else {
						f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","C",row), poi.Pname)
					}
					//市
					cName, _ := models.GetChina_divisionByCode(code.CityCode)
					if cName.Name != ""{
						f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","D",row), cName.Name)
					} else {
						f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","D",row), poi.Cityname)
					}
					//区
					aName, _ := models.GetChina_divisionByCode(poi.Adcode)
					if aName.Name != ""{
						f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","E",row), aName.Name)
					} else {
						f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","E",row), poi.Adname)
					}
					//街道
					street := GaodeMap.FindCityByCoordinate(poi.Location)
					f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","F",row), street.Township)
					locations := strings.Split(poi.Location,",")
					if len(locations) > 0{
						f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","G",row), locations[0])
						f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","H",row), locations[1])
					}
					f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","I",row), poi.ID)
					if len(poi.Photos) >0 {
						f.SetCellValue("Sheet1", fmt.Sprintf("%s,%d","J",row), poi.Photos[0].URL)
					}
					//}
					row ++
				}
			}
			f.SetActiveSheet(index)
			//根据指定路径保存文件
			rsu := common.Mkdir(PATH)
			if err := f.SaveAs(rsu+fileName+".csv"); err != nil {
				println(err.Error())
			}
			fmt.Println("高德城市数据爬取成功")
		}
	}
}