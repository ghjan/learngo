package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFindNumOrLowerLetter(t *testing.T) {
	tests := []struct {
		args string
		want []string
	}{
		{"AnameNN", []string{"name"}},
		{"HAHA00azBAPabc09FGabHY99", []string{
			"00az", "abc09", "ab", "99",
		}},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			result := FindNumOrLowerLetter(tt.args)
			for i, r := range result {
				assert.Equal(t, tt.want[i], r)
			}
		})
	}
}

func TestFindNum(t *testing.T) {
	tests := []struct {
		args string
		want []string
	}{
		{"7顶视图31QB001LX805P.png", []string{"7", "31", "001", "805"}},
	}
	//a := assert.New(t)
	for index, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			result := FindNum(tt.args)
			for i, r := range result {
				assert.Equal(t, tt.want[i], r)
			}
		})
	}
}

func TestRemoveNum(t *testing.T) {
	tests := []struct {
		args string
		want string
	}{
		{"卧室1", "卧室"},
		{"卧室2", "卧室"},
		{"客厅1", "客厅"},
		{"客厅2", "客厅"},
	}
	//a := assert.New(t)
	for index, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			result := RemoveNum(tt.args)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestCase2Camel(t *testing.T) {
	tests := []struct {
		args string
		want string
	}{
		{"name", "Name"},
		{"name_value", "NameValue"},
		{"pcode", "Pcode"},
		{"quota_item_id", "QuotaItemId"},
	}
	//a := assert.New(t)
	for index, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			result := Case2Camel(tt.args)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestCamel2Case(t *testing.T) {
	tests := []struct {
		args string
		want string
	}{
		{"ShortUUID", "short_uuid"},
		{"Name", "name"},
		{"NameValue", "name_value"},
		{"Pcode", "pcode"},
		{"QuotaItemId", "quota_item_id"},
	}
	//a := assert.New(t)
	for index, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			result := Camel2Case(tt.args)
			assert.Equal(t, tt.want, result)
		})
	}
}

var testAppendPrefixIfNotExistArgs = []struct {
	arg         string
	prefixIndex int
}{
	{`
	1.混凝土墙需机器切割，费用另计
2.墙厚不足12厘米按12厘米计，超过12厘米按实际厚度计算
3.不足1平米按1平米计
4.承重墙严禁拆除
`, 2},
	{`
	 1.混凝土墙需机器切割，费用另计
2.墙厚不足12厘米按12厘米计，超过12厘米按实际厚度计算
3.不足1平米按1平米计
4.承重墙严禁拆除
`, 3},
	{`1.混凝土墙需机器切割，费用另计
2.墙厚不足12厘米按12厘米计，超过12厘米按实际厚度计算
3.不足1平米按1平米计
4.承重墙严禁拆除
`, 0},
}

func TestAppendPrefixIfNotExist(t *testing.T) {
	for index, tt := range testAppendPrefixIfNotExistArgs {
		t.Run(fmt.Sprint(tt.arg), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			prefix := "说明："
			result := AppendPrefixIfNotExist(tt.arg, "说明：")
			expected := prefix + tt.arg[tt.prefixIndex:]
			assert.Equal(t, expected, result)
		})
	}

}

func TestAppendNewComerToContent(t *testing.T) {
	var tests = []struct {
		Content         string
		Newcomer        string
		ExpectedContent string
	}{
		{Content: "veUgWzCuuQUv245U89erfN", Newcomer: "VGkVWNZP8eEKyc2MDgxQHh",
			ExpectedContent: "veUgWzCuuQUv245U89erfN,VGkVWNZP8eEKyc2MDgxQHh"},
		{Content: "", Newcomer: "VGkVWNZP8eEKyc2MDgxQHh",
			ExpectedContent: "VGkVWNZP8eEKyc2MDgxQHh"},
		{Content: "veUgWzCuuQUv245U89erfN,UVJ8DFDBQwrwbJQiAL2AnG", Newcomer: "VGkVWNZP8eEKyc2MDgxQHh",
			ExpectedContent: "veUgWzCuuQUv245U89erfN,UVJ8DFDBQwrwbJQiAL2AnG,VGkVWNZP8eEKyc2MDgxQHh"},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			result := AppendNewComerToContent(tt.Content, tt.Newcomer)
			assert.Equal(t, tt.ExpectedContent, result)
		})
	}

}

func TestRemoveSuffixPart(t *testing.T) {
	var tests = []struct {
		Value        string
		WantedResult string
		Sep          string
	}{
		{Value: "201-01", WantedResult: "201"},
		{Value: "201", WantedResult: "201"},
		{Value: "201-01", WantedResult: "201", Sep: "-"},
		{Value: "201", WantedResult: "201", Sep: "-"},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			result := RemoveSuffixPart(tt.Value, tt.Sep)
			assert.Equal(t, tt.WantedResult, result)
		})
	}
}

func TestReplaceModelName(t *testing.T) {
	var tests = []struct {
		sourceName         string
		sourceUUID         string
		targetUUID         string
		expectedTargetName string
	}{
		{sourceName: "杉木实木踢脚线A09\\18元/米_fVtYQoU5xxPN7ZZgtjLw8o", sourceUUID: "Gmuy2uGs2MvnzNE8u8zYDN",
			targetUUID: "ytd4W8t6F8KUFjbJVqkN2m", expectedTargetName: "杉木实木踢脚线A09\\18元/米ytd4W8t6F8KUFjbJVqkN2m"},
		{sourceName: "杉木实木踢脚线A09\\18元/米", sourceUUID: "MjgUcRhLoKfMRVirtdXAzC",
			targetUUID: "Gmuy2uGs2MvnzNE8u8zYDN", expectedTargetName: "杉木实木踢脚线A09\\18元/米_Gmuy2uGs2MvnzNE8u8zYDN"},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			result := ReplaceModelName(tt.sourceName, tt.sourceUUID, tt.targetUUID)
			assert.Equal(t, tt.expectedTargetName, result)
		})
	}

}

func TestGetItemNoFromSubItemNo(t *testing.T) {
	itemNo := GetItemNoFromSubItemNo([]byte("232-01"))
	assert.Equal(t, "232", itemNo)
}

var specification_4_test = `
型号：磁吸轨道灯（1.5m预埋轨道+3个泛光灯+100W内置电源）
规格：内径7cm*高5.2cm
说明：1.水电验收前确认，送货周期14天
2.根据米数不同适配电源
`

func TestGetTypeFromSpecification(t *testing.T) {
	typ := GetTypeFromSpecification([]byte(specification_4_test))
	assert.Equal(t, "磁吸轨道灯（1.5m预埋轨道+3个泛光灯+100W内置电源）", typ)
}

func TestGetSpecsFromSpecification(t *testing.T) {
	specs := GetSpecsFromSpecification([]byte(specification_4_test))
	assert.Equal(t, "内径7cm*高5.2cm", specs)
}

func TestGetDescriptionFromSpecification(t *testing.T) {
	dOld := `1.面积以实际测量为准，单扇不足1平米按1平米计
2.窗宽度小于85厘米无法做左右移窗
3.仿钢色费用增加30元/平米
4.开工后一周内确认，制作周期约15天
5.拆除费用另计`
	d := strings.ReplaceAll(dOld, `\r\n`, `\n`)
	assert.Equal(t, dOld, d)
	description := GetDescriptionFromSpecification([]byte(specification_4_test), false)
	assert.Equal(t, `1.水电验收前确认，送货周期14天
2.根据米数不同适配电源`, description)
}

func TestGetOssIdFromOssUrl(t *testing.T) {
	ossUrl := "llib/979da15566/data.db"
	ossId := GetOssIdFromOssUrl([]byte(ossUrl))
	assert.Equal(t, "979da15566", ossId)
}
