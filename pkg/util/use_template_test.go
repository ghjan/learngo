package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"testing"
	"text/template"
)

type MyRenderInventory struct {
	Material string
	Count    uint
}

type MyRenderContextObject struct {
	HouseQuantDoubleControl int
}

type RenderStringArgs struct {
	TemplateString string
	Context        interface{}
}

var jsonObj = gjson.Result{}.Map()

var tests []struct {
	args RenderStringArgs
	want string
}

func TestRenderStringWithTemplateString(t *testing.T) {

	//a := assert.New(t)
	for index, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			tmpl, result, err := RenderStringWithTemplateString(tt.args.TemplateString, tt.args.Context)
			t.Log(result)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.want, result)
			assert.Equal(t, "test", tmpl.Name())
		})
	}

}

func TestRenderStringWithTemplate(t *testing.T) {
	//a := assert.New(t)
	for index, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			tmpl, err := template.New("test").Parse(tt.args.TemplateString)
			result, err := RenderStringWithTemplate(tmpl, tt.args.Context)
			t.Log(result)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.want, result)
			assert.Equal(t, "test", tmpl.Name())
		})
	}

}

func TestMain(m *testing.M) {
	fmt.Println("begin util")
	jsonObj["house_quant_doubleControl"] = gjson.Result{Type: gjson.Number, Num: 6.0}
	tests = []struct {
		args RenderStringArgs
		want string
	}{
		{RenderStringArgs{"{{.Count}} items are made of {{.Material}}",
			MyRenderInventory{"wool", 17},},
			"17 items are made of wool"},
		{RenderStringArgs{`1.已包含{{.HouseQuantDoubleControl}}组双联回路，每增加一组双联增加25米
				2.潮湿环境应顶面架设布线
				3.先安装pvc线管，后穿线，严禁逆操作
				4.家具内嵌辅助照明时按实增加米数`,
			MyRenderContextObject{6}},
			`1.已包含6组双联回路，每增加一组双联增加25米
				2.潮湿环境应顶面架设布线
				3.先安装pvc线管，后穿线，严禁逆操作
				4.家具内嵌辅助照明时按实增加米数`},
		{RenderStringArgs{`1.已包含{{.house_quant_doubleControl}}组双联回路，每增加一组双联增加25米
				2.潮湿环境应顶面架设布线
				3.先安装pvc线管，后穿线，严禁逆操作
				4.家具内嵌辅助照明时按实增加米数`,
			jsonObj},
			`1.已包含6组双联回路，每增加一组双联增加25米
				2.潮湿环境应顶面架设布线
				3.先安装pvc线管，后穿线，严禁逆操作
				4.家具内嵌辅助照明时按实增加米数`},
	}
	m.Run()
	fmt.Println("end  util")
}
