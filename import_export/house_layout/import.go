package house_layout

import (
	"context"
	"encoding/json"
	"github.com/ghjan/learngo/pkg/cache/gredis"
	"github.com/ghjan/learngo/pkg/util"
	"github.com/ghjan/learngo/pkg/utils"
	"github.com/olivere/elastic/v7"
	"github.com/tidwall/gjson"
	"log"
	"strconv"
)

var client *elastic.Client

func initClient() {
	var err error
	client, err = elastic.NewClient(
		elastic.SetURL("http://elastic.davidzhang.xin:9200", "http://localhost:9200"),
		elastic.SetMaxRetries(10), elastic.SetSniff(false))
	if err != nil {
		log.Fatal(err)
	}
}
func SaveItem(client *elastic.Client, item HouseLayout, index string, typ string) (err error) {
	var (
		indexService *elastic.IndexService
	)
	if typ != "" {
		indexService = client.Index().Index(index).Type(typ).BodyJson(item)
	} else {
		indexService = client.Index().Index(index).BodyJson(item)
	}
	if item.Id == "" {
		_, err = indexService.Id(item.Id).Do(context.Background())
	} else {
		_, err = indexService.Do(context.Background())
	}
	return
}

func BulkIndex(index string, typ string, data []HouseLayout) (successItems []HouseLayout, err error) {
	successItems = make([]HouseLayout, 0)
	if client == nil {
		initClient()
	}
	for _, item := range data {
		if err = SaveItem(client, item, index, typ); err != nil {
			return
		} else {
			successItems = append(successItems, item)
		}
	}
	return
}


func GetHouseLayoutUploadSuccessObjMap(province string, city string, pageNo int) (successObjMap map[string]gjson.Result, err error) {
	var (
		v              interface{}
		result         string
		successObjList []gjson.Result
	)
	keyname4lo := GetKeyname4result(province, city, PREFIX_REDIS_NAME_KEY_LAYOUT_CONVERSION_AND_UPLOAD)
	if v, err = gredis.HGet(keyname4lo, strconv.Itoa(pageNo)); err != nil || v == nil {
		return
	} else {
		result = string(v.([]byte))
	}
	successObjList = gjson.Get(result, "success_obj_list").Array()
	if successObjList != nil && len(successObjList) > 0 {
		successObjMap = make(map[string]gjson.Result, 0)
		for _, so := range successObjList {
			id := gjson.Get(so.String(), "id").String()
			if id != "" {
				successObjMap[id] = so
			}
		}
	}
	return
}

func GetHouseLayout4CityPage(province, city string, pageNo int) (successLayouts, failLayouts []HouseLayout, err error) {
	var (
		successObjMap map[string]gjson.Result

		v            interface{}
		jsonFilePath string
		result       string
	)
	if successObjMap, err = GetHouseLayoutUploadSuccessObjMap("上海", "上海", 1); err != nil {
		utils.Checkerr2(err, "GetHouseLayoutUploadSuccessObjMap error")
		return
	}

	keyname4jsonfile := GetKeyname4result(province, city, PREFIX_REDIS_NAME_KEY_JSONFILES)
	if keyname4jsonfile == "" {
		return
	}
	if v, err = gredis.HGet(keyname4jsonfile, strconv.Itoa(pageNo)); err != nil || v == nil {
		return
	} else {
		jsonFilePath = string(v.([]byte))
	}
	if jsonFilePath != "" {
		if result, err = util.ReadJsonFileWithSearch(jsonFilePath, 3); err == nil && result != "" {
			successLayouts = make([]HouseLayout, 0)
			house := gjson.Get(result, "house").Array()
			for _, h := range house {
				var (
					layout HouseLayout
				)
				if err = json.Unmarshal([]byte(h.String()), &layout); err == nil && layout.IsValid() {
					if so, existed := successObjMap[layout.Id]; existed {
						timestamp := gjson.Get(so.String(), "timestamp").String()
						ossObjectPath := gjson.Get(so.String(), "oss_object_path").String()
						layout.Timestamp = timestamp
						layout.OssObjectPath = ossObjectPath
						layout.OssId = util.GetOssIdFromOssUrl([]byte(ossObjectPath))
						successLayouts = append(successLayouts, layout)
					} else {
						if failLayouts == nil {
							failLayouts = make([]HouseLayout, 0)
							failLayouts = append(failLayouts, layout)
						}
					}
				} else {
					utils.Checkerr2(err, "json.Unmarshal error")
				}
			}
		}

	}
	return
}

func GetMd5ProvinceCity(province string, city string) (md5 string) {
	if province != "" && city != "" {
		md5 = util.StringMD5(province + "_" + city)

	}
	return

}

func GetKeyname4result(province string, city string, keyNamePrefix string) (key string) {
	if keyNamePrefix == "" {
		keyNamePrefix = PREFIX_REDIS_NAME_KEY_JSONFILES
	}
	key = GetMd5ProvinceCity(province, city)
	if key != "" {
		key = keyNamePrefix + "_" + key
	}

	return
}
