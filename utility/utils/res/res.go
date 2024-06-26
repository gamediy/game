package res

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"math"
)

type PageRes struct {
	Data       interface{} `json:"data,omitempty"`
	Other      interface{} `json:"other,omitempty"`
	TotalCount int64       `json:"totalCount"`
	PageSize   int64       `json:"pageSize,omitempty"`
	TotalPage  int64       `json:"totalPage,omitempty"`
	CurrPage   int64       `json:"currPage,omitempty"`
	List       interface{} `json:"list"`
}

func PageList(r *ghttp.Request, page string, total int, list interface{}, info interface{}) {
	size := r.GetQuery("size").Int()
	content := r.GetPage(total, size).GetContent(3)
	if err := r.Response.WriteTpl(page, g.Map{
		"list":  list,
		"total": total,
		"page":  content,
		"c":     info,
	}); err != nil {
		g.Log().Error(r.Context(), err)
	}
	r.Exit()
}
func Page(r *ghttp.Request, page string, data ...interface{}) {
	d := g.Map{}
	if len(data) > 0 {
		d["data"] = data[0]
	}
	if err := r.Response.WriteTpl(page, d); err != nil {
		g.Log().Error(r.Context(), err)
	}
	r.Exit()
}
func GetPage(r *ghttp.Request, defaultSize ...int) (page, size int) {
	page = r.GetQuery("page").Int()
	size = r.GetQuery("size").Int()
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		if len(defaultSize) > 0 {
			size = defaultSize[0]
		} else {
			size = 10
		}
	}
	return page, size
}
func Err(err error, r *ghttp.Request) {
	r.Response.WriteJsonExit(g.Map{
		"code":    -1,
		"message": err.Error(),
	})
}
func Ok(r *ghttp.Request) {
	r.Response.WriteJsonExit(g.Map{
		"code":    0,
		"message": "ok",
	})
}

func OkMsg(msg string, r *ghttp.Request) {
	r.Response.WriteJsonExit(g.Map{
		"code":    0,
		"message": msg,
	})
}
func OkData(data interface{}, r *ghttp.Request) {
	r.Response.WriteJsonExit(g.Map{
		"code":    0,
		"message": "ok",
		"data":    data,
	})
}
func OkPage(page, size, total int, data interface{}, r *ghttp.Request) {
	if size == 0 {
		size = 10
	}
	totalPage := math.Ceil(float64(total) / float64(size)) //这里计算总页数时，要向上取整
	if totalPage <= 0 {
		totalPage = 1
	}
	if total == 0 || data == nil {
		data = make([]interface{}, 0)
	}
	OkData(PageRes{TotalCount: int64(total), PageSize: int64(size), CurrPage: int64(page), List: data, TotalPage: int64(totalPage)}, r)
}

// Tpl template
func Tpl(file string, data g.Map, r *ghttp.Request) {
	_ = r.Response.WriteTpl(file, data)

}
func RedirectTo(path string, r *ghttp.Request) {
	r.Response.RedirectTo(path)
}
