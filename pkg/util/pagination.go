package util

import (
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"strconv"
)

func GetOnlyActivated(c *gin.Context, defaultValue bool) bool {
	result := defaultValue // config.PageSize
	if c.Query("only_activated") != "" {
		need3d := com.StrTo(c.Query("need3d")).MustInt()
		result = need3d >= 0
	}

	return result
}

func GetBoolQuery(c *gin.Context, key string, defaultValue bool) (result bool) {
	result = defaultValue // config.PageSize
	if c.Query(key) != "" {
		result = utils.ParseBool(c.Query(key))
	}
	return

}

func GetPageSize(c *gin.Context, defaultValue int) int {
	result := defaultValue // config.PageSize
	pageSize := com.StrTo(c.Query("page_size")).MustInt()
	if pageSize > 0 {
		result = pageSize
	}

	return result
}

// GetPageNum get page parameters
// return:
// 		page:  页数
// 		pageNum: (page - 1) * pageSize
func GetPageNum(c *gin.Context, defaultValue int) (pageNum int, page int) {
	pageNum = defaultValue //0
	page = com.StrTo(c.Query("page")).MustInt()
	if page > 0 {
		pageSize := GetPageSize(c, 0)
		if pageSize > 0 {
			pageNum = (page - 1) * pageSize
		}
	}

	return
}

// GetUnscoped get unscoped parameters
func GetUnscoped(c *gin.Context, defaultValue bool) (unscoped bool, err error) {
	unscoped = defaultValue //0
	unscoped, err = strconv.ParseBool(c.Query("include_deleted"))
	return
}

func GetNeedFix(c *gin.Context, defaultValue bool) (value bool, err error) {
	value = defaultValue //0
	value, err = strconv.ParseBool(c.Query("need_fix"))
	return
}

func GetOnlyPublic(c *gin.Context, defaultValue bool) (value bool, err error) {
	value = defaultValue //0
	value, err = strconv.ParseBool(c.Query("only_public"))
	return
}

func GetOverride(c *gin.Context, defaultValue bool) (value bool, err error) {
	value = defaultValue
	if c.PostForm("override") != "" && c.PostForm("override") != "undefined" {
		value, err = strconv.ParseBool(c.PostForm("override"))
	}
	return
}
func GetWithdraw(c *gin.Context, defaultValue bool) (value bool, err error) {
	value = defaultValue
	if c.PostForm("withdraw") != "" && c.PostForm("withdraw") != "undefined" {
		value, err = strconv.ParseBool(c.PostForm("withdraw"))
	}
	return
}

func GetWithdrawModelsOfCategory3(c *gin.Context, defaultValue bool) (value bool, err error) {
	value = defaultValue
	if c.PostForm("withdraw_models_of_category3") != "" && c.PostForm("withdraw_models_of_category3") != "undefined" {
		value, err = strconv.ParseBool(c.PostForm("withdraw_models_of_category3"))
	}
	return
}

func GetNeedProcessCategories(c *gin.Context, defaultValue bool) (value bool, err error) {
	value = defaultValue
	if c.PostForm("need_process_categories") != "" && c.PostForm("need_process_categories") != "undefined" {
		value, err = strconv.ParseBool(c.PostForm("need_process_categories"))
	}
	return
}

func GetNeedProcessBrands(c *gin.Context, defaultValue bool) (value bool, err error) {
	value = defaultValue
	if c.PostForm("need_process_brands") != "" && c.PostForm("need_process_brands") != "undefined" {
		value, err = strconv.ParseBool(c.PostForm("need_process_brands"))
	}
	return
}

func GetNeedProcessNoChange(c *gin.Context, defaultValue bool) (value bool, err error) {
	value = defaultValue
	if c.PostForm("need_process_no_change") != "" && c.PostForm("need_process_no_change") != "undefined" {
		value, err = strconv.ParseBool(c.PostForm("need_process_no_change"))
	}
	return
}
