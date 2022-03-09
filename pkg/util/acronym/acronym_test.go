package acronym

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

type TestGetAcroArgs struct {
	ChineseWord string
	Want        string
}

func TestGetAcronym(t *testing.T) {
	targs := []TestGetAcroArgs{
		{ChineseWord: "洗衣机", Want: "XYJ"},
		{ChineseWord: "连体坐便器", Want: "LTZBQ"},
		{ChineseWord: "坐便器", Want: "ZBQ"},
		{ChineseWord: "电视柜", Want: "DSG"},
		{ChineseWord: "蹲便器", Want: "DBQ"},
		{ChineseWord: "浴霸", Want: "YB"},
		{ChineseWord: "排风", Want: "PF"},
		{ChineseWord: "台盆", Want: "TP"},
		{ChineseWord: "浴缸", Want: "YG"},
		{ChineseWord: "吸顶灯", Want: "XDD"},
		{ChineseWord: "吊顶灯", Want: "DDD"},
		{ChineseWord: "筒灯", Want: "TD"},
		{ChineseWord: "射灯", Want: "SD"},
		{ChineseWord: "集成吊顶灯", Want: "JCDDD"},
		{ChineseWord: "壁灯", Want: "BD"},
		{ChineseWord: "灯带", Want: "DD"},
		//hello,世界!
		{ChineseWord: "hello,世界!", Want: "SJ"},
		{ChineseWord: "800*800砖", Want: "Z"},
		{ChineseWord: "900*900", Want: ""},
		{ChineseWord: "600*1200", Want: ""},
		{ChineseWord: "750*1500", Want: ""},
	}
	for index, tt := range targs {
		t.Run(fmt.Sprint(tt.ChineseWord), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			result := GetAcronym(tt.ChineseWord)
			assert.Equal(t, result, tt.Want)

		})

	}
}
