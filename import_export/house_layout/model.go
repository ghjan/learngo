package house_layout

type HouseLayout struct {
	Id            string  `json:"id"`
	OssId         string  `json:"oss_id"`
	ProvinceName  string  `json:"provinceName"`
	Province      int     `json:"province"`
	CityName      string  `json:"cityName"`
	City          int     `json:"city"`
	RegionName    string  `json:"regionName"`
	Region        int     `json:"region"`
	Area          float64 `json:"area"`
	Hall          string  `json:"hall"`
	VillageName   string  `json:"villageName"`
	Room          string  `json:"room"`
	Toilet        string  `json:"toilet"`
	Status        string  `json:"status"`
	OssObjectPath string  `json:"oss_object_path"`
	Timestamp     string  `json:"timestamp"`
}

func (h *HouseLayout) IsValid() (valid bool) {
	valid = h.Id != ""
	return
}


//func FromJsonObj(o interface{}) (HouseLayout, error) {
//	var house HouseLayout
//	s, err := json.Marshal(o)
//	if err != nil {
//		return house, err
//	}
//	err = json.Unmarshal(s, &house)
//	return house, err
//}

var layout_sample = `
{
"id":  "33001",
"provinceName":  "\u4e0a\u6d77",
"cityName":  "\u4e0a\u6d77",
"oss_object_path":  "llib/979da15566/data.db",
"timestamp":  "20220301T155638",
"output_layout_json_file_path":  "h:/gresources\\house/img\\apartmen/33001\\apartmen.json",
"output_layout_data_file_path":  "h:/gresources\\house/img\\apartmen/33001\\data.db",
"svg_path":  "h:/gresources\\house/img\\apartmen/33001/apartmen.svg",
"thumbnail_path":  "h:/gresources\\house/img\\apartmen/33001/thumbnail.jpg",
"svgUrl":  "https://3d-design-prod-1305072137.cos.ap-shanghai.myqcloud.com/apartmen/33001/apartmen.svg",
"thumbnailurl":  "https://3d-design-prod-1305072137.cos.ap-shanghai.myqcloud.com/apartmen/33001/thumbnail.jpg",
"override":  false
"ignored_svg":  true,
"ignored_thumbnail":  true,
}
`

var house_jsonfile = `
"area": 136.0,
"city": 321,
"svgEditionUrl": "",
"regionName": "\u9759\u5b89\u533a",
"svgUrl": "apartmen/33001/apartmen.svg",
"hall": "\u4e00",
"villageName": "\u9759\u5b89\u9f0e\u946b\u4f73\u56ed",
"room": "\u4e8c",
"toilet": "\u4e8c",
"province": 25,
"cityName": "\u4e0a\u6d77",
"provinceName": "\u4e0a\u6d77",
"id": "33001",
"region": 2710,
"thumbnailurl": "apartmen/33001/thumbnail.jpg",
"status": "1",
"rawApartmentId": 33001
`
