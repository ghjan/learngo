package house_layout

import (
	"github.com/ghjan/learngo/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"strings"
	"testing"
)

func TestGetKeyname4result(t *testing.T) {
	key := GetKeyname4result("上海", "上海", "")
	assert.Equal(t, "js_d16a2d82ac903e1d03aa46abb4787a21", key)
	key2 := GetKeyname4result("上海", "上海", PREFIX_REDIS_NAME_KEY_LAYOUT_CONVERSION_AND_UPLOAD)
	assert.Equal(t, "lo_d16a2d82ac903e1d03aa46abb4787a21", key2)
}

func TestGetHouseLayout(t *testing.T) {
	var (
		successLayouts, failLayouts []HouseLayout
		err                         error
	)
	successLayouts, failLayouts, err = GetHouseLayout4CityPage("上海", "上海", 1)
	assert.Nil(t, err)
	assert.Equal(t, 19, len(successLayouts))
	firstLayout := successLayouts[0]
	assert.Equal(t, "静安区", firstLayout.RegionName)
	//厅 一厅
	assert.Equal(t, "一", firstLayout.Hall)
	assert.Equal(t, "静安鼎鑫佳园", firstLayout.VillageName)
	//房 二房
	assert.Equal(t, "二", firstLayout.Room)
	//卫 二卫
	assert.Equal(t, "二", firstLayout.Toilet)
	assert.Equal(t, "上海", firstLayout.CityName)
	assert.Equal(t, "上海", firstLayout.ProvinceName)
	assert.Equal(t, "上海", firstLayout.ProvinceName)
	for _, layout := range successLayouts {
		assert.True(t, layout.OssId != "")
		assert.True(t, layout.Timestamp != "")
		assert.True(t, strings.Contains(layout.OssObjectPath, "llib"))
	}
	if failLayouts != nil && len(failLayouts) > 0 {
		for _, fl := range failLayouts {
			assert.True(t, fl.OssId == "")
			assert.True(t, fl.Timestamp == "")
			assert.True(t, fl.OssObjectPath == "")
		}
	}

}

func TestGetHouseLayoutUploadSuccessObjMap(t *testing.T) {
	var (
		successObjMap map[string]gjson.Result
		err           error
	)
	if successObjMap, err = GetHouseLayoutUploadSuccessObjMap("上海", "上海", 1); err == nil {
		for id, so := range successObjMap {
			timestamp := gjson.Get(so.String(), "timestamp")
			ossObjectPath := gjson.Get(so.String(), "oss_object_path")
			t.Logf("--id:%s--", id)
			t.Logf("id:%s,timestamp:%s,ossObjectPath:%s", id, timestamp, ossObjectPath)
			t.Logf("id:%s,so:%s", id, so.String())
		}
	} else {
		utils.Checkerr2(err, "GetHouseLayoutUploadSuccessObjMap error")
	}
}

func TestBulkIndex(t *testing.T) {
	var (
		successLayouts, failLayouts []HouseLayout
		successItems                []HouseLayout
		err                         error
	)
	successLayouts, failLayouts, err = GetHouseLayout4CityPage("上海", "上海", 1)
	assert.Nil(t, err)
	assert.Equal(t, 19, len(successLayouts))
	firstLayout := successLayouts[0]
	assert.Equal(t, "静安区", firstLayout.RegionName)
	//厅 一厅
	assert.Equal(t, "一", firstLayout.Hall)
	assert.Equal(t, "静安鼎鑫佳园", firstLayout.VillageName)
	//房 二房
	assert.Equal(t, "二", firstLayout.Room)
	//卫 二卫
	assert.Equal(t, "二", firstLayout.Toilet)
	assert.Equal(t, "上海", firstLayout.CityName)
	assert.Equal(t, "上海", firstLayout.ProvinceName)
	assert.Equal(t, "上海", firstLayout.ProvinceName)
	for _, layout := range successLayouts {
		assert.True(t, layout.OssId != "")
		assert.True(t, layout.Timestamp != "")
		assert.True(t, strings.Contains(layout.OssObjectPath, "llib"))
	}
	if failLayouts != nil && len(failLayouts) > 0 {
		for _, fl := range failLayouts {
			assert.True(t, fl.OssId == "")
			assert.True(t, fl.Timestamp == "")
			assert.True(t, fl.OssObjectPath == "")
		}
	}

	if successItems, err = BulkIndex("house_test", "HouseLayout", successLayouts); err != nil {
		t.Errorf("TestBulkIndex, BulkIndex error:%s\n", err.Error())
		t.FailNow()
	} else {
		t.Logf("successItems:%#v\n", successItems)
	}

}
