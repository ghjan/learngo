package util

import (
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"strings"
	"testing"
)

//var dataPath = "../../test_data/test_data82.json"
var dataPath = "../../import_export/house_layout/上海_上海_1.json"

func TestReadJsonFile(t *testing.T) {
	result, err := ReadJsonFile(dataPath)
	if err != nil {
		t.Logf("error:%s\n", err.Error())
		t.Fail()
	} else {
		statusDescRuhumen := gjson.Get(result, "data.profile.keptItemList.0.statusDesc").String()
		statusDescRuhumen2 := gjson.Get(result, `data.profile.keptItemList.#[name="入户门"].statusDesc`).String()
		assert.Equal(t, statusDescRuhumen, statusDescRuhumen2)
		t.Logf("statusDescRuhumen:%s", statusDescRuhumen)
		t.Logf("statusDescRuhumen2:%s", statusDescRuhumen2)
	}
}

func TestReadJsonFileWithSearch(t *testing.T) {
	dataPath1 := strings.TrimPrefix(dataPath, "../../")
	result, err := ReadJsonFileWithSearch(dataPath1, 3)
	if err != nil {
		t.Logf("error:%s\n", err.Error())
		t.Fail()
	} else {
		house := gjson.Get(result, "house").Array()
		pageSize := gjson.Get(result, "pageSize").Int()
		assert.True(t, house != nil && int64(len(house)) == pageSize)
		t.Logf("----house---")
		for index, h := range house {
			t.Logf("index:%d,h:%s", index, h)
		}
		t.Logf("----house---")

		//house2 := gjson.Get(result, "house")
		//t.Logf("----house2---")
		//house2.ForEach(func(key, value gjson.Result) bool {
		//	println(value.String())
		//	return true // keep iterating
		//})
		//t.Logf("----house2---")

		firstCity := gjson.Get(result, "house.0.city").Int()
		assert.Equal(t, int64(321), firstCity)
		firstRegionName := gjson.Get(result, "house.0.regionName").String()
		assert.Equal(t, "静安区", firstRegionName)
		//厅 一厅
		firstHall := gjson.Get(result, "house.0.hall").String()
		assert.Equal(t, "一", firstHall)
		firstVillageName := gjson.Get(result, "house.0.villageName").String()
		assert.Equal(t, "静安鼎鑫佳园", firstVillageName)
		//房 二房
		firstRoom := gjson.Get(result, "house.0.room").String()
		assert.Equal(t, "二", firstRoom)
		//卫 二卫
		firstToilet := gjson.Get(result, "house.0.toilet").String()
		assert.Equal(t, "二", firstToilet)
		firstCityName := gjson.Get(result, "house.0.cityName").String()
		assert.Equal(t, "上海", firstCityName)
		firstProvinceName := gjson.Get(result, "house.0.provinceName").String()
		assert.Equal(t, "上海", firstProvinceName)

		t.Logf("firstCity:%d, firstRegionName:%s,firstHall:%s,firstVillageName:%s,firstRoom:%s,"+
			"firstToilet:%s,firstCityName:%s,firstProvinceName:%s", firstCity, firstRegionName,
			firstHall, firstVillageName, firstRoom, firstToilet, firstCityName, firstProvinceName)
		//t.Logf("house:%#v", house)
	}
}
