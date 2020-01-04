package services

const (
	TYPE = 1
	GAT_WAY  = "https://restapi.amap.com/v3/place/text"
	FILE_NAME = ""
	geocodeRegeoEndPoint = "https://restapi.amap.com/v3/geocode/regeo"
	GaodeMapKey = "881b1b4f317ba7714b3956475625199f"
)

func main() {
	//beego.Run()
	//brand := [] string {"爱依服"}
	//for _,v:= range brand{
	//	query(v,v)
	//}
}

//func query(brand string, fileName string) {
//
//	var page, offset int8 = 1, 20
//
//	var key_word, types string = brand, ""
//
//	//pois := [] string{}
//
//	if TYPE == 1{
//		url1 := GAT_WAY + fmt.Sprintf("?key=%s&keywords=%s&types=%s&offset=%d&page=%d&extensions=all&children=1", GaodeMapKey, key_word, types, offset, page)
//
//		response1 := httplib.Get(url1)
//
//		resp,_ := response1.String()
//
//		fmt.Println(resp)
//
//		//var dat map[string]interface{}
//		//if err := json.Unmarshal([]byte(resp), &dat); err == nil {
//		// fmt.Println("==============json str 转map=======================")
//		// fmt.Println(dat)
//		// mapTmp := dat["suggestion"].(map[string]interface{})
//		// mapTmp2 := mapTmp["cities"].([]interface {})
//		// //[0].(map[string]interface {})
//		// //fmt.Println(mapTmp2["name"])
//		//
//		// for key, city := range mapTmp2{
//		//  fmt.Println(key)
//		//  fmt.Println(city["adcode"])
//		// }
//		//}
//
//		//fmt.Println(reflect.TypeOf(resp))
//	}
//
//}