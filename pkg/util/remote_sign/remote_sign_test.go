package remote_sign

import (
	"fmt"
	"gcherry-server/config"
	"gcherry-server/pkg/util"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

var TestCalcSignMd5Args = []map[string]string{
	{"access_key": "7b76ba543b054ac083f620669fdaa428", "nonce": "60a955c0c8ec11eb8bbe9cda3e0e0006",
		"timestamp": "1623220312447", "id": "DpQfzu5VynSpfrzbA7wadU", "sign": "0358bed0f43641fe4d11677fef3b44c2"},
	{"access_key": "7b76ba543b054ac083f620669fdaa428", "nonce": "60a955c1c8ec11ebbd6b9cda3e0e0006",
		"timestamp": "1623220312447", "id": "X6RLaWMmU4Yci8aHw2FViU", "sign": "480756ed29d7f28aeb9c9d7b1c207a13"},
	{"access_key": "7b76ba543b054ac083f620669fdaa428", "nonce": "60a955c2c8ec11eb92ee9cda3e0e0006",
		"timestamp": "1623220312447", "id": "xpYduccL6AzKVazshNGGhN", "sign": "82b86b696b4976733e58c58131548c8c"},
	{"access_key": "7b76ba543b054ac083f620669fdaa428", "nonce": "60a97cdcc8ec11ebbf4a9cda3e0e0006",
		"timestamp": "1623220312448", "sign": "6a7f40d0e05e3e3181d88a7e5e687131"},
}

func TestCalcSignMd5(t *testing.T) {
	for index, bd := range TestCalcSignMd5Args {
		t.Run(fmt.Sprint(bd), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			if bd["access_key"] == config.GCHERRY1_SK_ACCESS_KEY {
				sign := CalcSignMd5(bd, config.GCHERRY1_SK_SECRET_KEY)
				assert.Equal(t, sign, bd["sign"], "TestCalcSignMd5 error")
			}
		})
	}
}

func TestGetSignedQueryString(t *testing.T) {
	_, urlQuerystring := GetSignedQueryString(nil, false)
	nonce, _ := util.GetQueryParameterValue(urlQuerystring, "nonce")
	timestamp, _ := util.GetQueryParameterValue(urlQuerystring, "timestamp")
	qp := map[string]string{
		"access_key": config.GCHERRY1_SK_ACCESS_KEY,
		"nonce":      nonce,
		"timestamp":  timestamp,
	}
	_, urlQuerystring2 := GetSignedQueryString(qp, false)
	assert.Equal(t, urlQuerystring, urlQuerystring2, "urlQuerystring=urlQuerystring2")

}

func TestVerifySign(t *testing.T) {
	_, urlQuerystring := GetSignedQueryString(nil, false)
	success := VerifySignedQueryString(urlQuerystring, true, false)
	assert.Equal(t, true, success, "success=true")
	success = VerifySignedQueryString(urlQuerystring, true, false)
	assert.Equal(t, false, success, "success=false")

}

func TestVerifySign2(t *testing.T) {
	company_paras := map[string]string{
		"user_id":      "qdgRGZivuoS97KGgSEbZed",
		"company_name": url.QueryEscape("上海集岳建筑工程技术有限公司2"),
	}
	qps := []map[string]string{
		nil,
		company_paras,
	}
	for index, qp := range qps {
		t.Run(fmt.Sprint(qp), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			_, urlQuerystring := GetSignedQueryString(qp, false)
			success := VerifySignedQueryString(urlQuerystring, true, false)
			assert.Equal(t, true, success, "success=true")
			success = VerifySignedQueryString(urlQuerystring, true, false)
			assert.Equal(t, false, success, "success=false")
		})
	}
}

func TestVerifySign3(t *testing.T) {
	company_paras := map[string]string{
		"user_id":      "qdgRGZivuoS97KGgSEbZed",
		"company_name": url.QueryEscape("上海集岳建筑工程技术有限公司2"),
	}
	qps := []map[string]string{
		nil,
		company_paras,
	}
	for index, qp := range qps {
		t.Run(fmt.Sprint(qp), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			_, urlQuerystring := GetSignedQueryString(qp, true)
			success := VerifySignedQueryString(urlQuerystring, true, true)
			assert.Equal(t, true, success, "success=true")
			success = VerifySignedQueryString(urlQuerystring, true, true)
			assert.Equal(t, false, success, "success=false")
		})
	}
}

func TestMain(m *testing.M) {
	fmt.Println("remote_sign begin")
	config.InitConfig(false)

	m.Run()
	fmt.Println("remote_sign end")
}
