package tplengine

import (
	"encoding/json"
	"html/template"
	"math"
)
// Paging 分页参数
type Paging struct{
	// 要显示的页数(page numbers to display)
	PageNumbers []int
	PageNo      int
	RowsPerPage int
	TotalRows   int
	TotalPages  int
	StartPage, EndPage    int
	PrePageNo, NextPageNo int
}
var plugins = template.FuncMap{
	"json": func(d interface{}) template.HTML {
		b, err := json.Marshal(d)
		if err != nil {
			return ""
		}
		return template.HTML(string(b))
	},
	"list": func(vals ...interface{}) []interface{} {
		return append([]interface{}{}, vals...)
	},
	"map": func() map[string]interface{} {
		return map[string]interface{}{}
	},
	"addToMap": func(name string, val interface{}, m map[string]interface{}) map[string]interface{} {
		m[name] = val
		return m
	},
	"sum": func(args ...int) int {
		sum := 0
		for _, n := range args {
			sum += n
		}
		return sum
	},
	"minus": func(args ...int) int {
		val := 0
		for _, n := range args {
			val -= n
		}
		return val
	},
	"between": func(val, min, max int) bool {
		return val >= min && val <= max
	},
	// 三元运算 ?:
	"ternary": func(comp bool, trueOption, falseOption interface{}) interface{} {
		if comp {
			return trueOption
		}
		return falseOption
	},
	"divide": func(i, j int) int {
		return i / j
	},
	// 生成分页参数对象
	"paging": func(displayPageCount, pageNo, rowsPerPage, totalRows  int)*Paging{
		paging := &Paging{}
		totalPages := int(math.Ceil(float64(totalRows) / float64(rowsPerPage)))
		if totalPages < 1 {
			return paging
		}
		if pageNo > totalPages {
			pageNo = totalPages
		}
		paging.TotalPages = totalPages
		paging.TotalRows = totalRows
		paging.RowsPerPage = rowsPerPage
		paging.PageNo = pageNo
		sidePageCount := displayPageCount/2
		var upperCount, lowerCount int
		lowerCount, upperCount = sidePageCount, sidePageCount
		// 为偶数时总数会大于总长度,因此选择将起始页向后推一页.
		if sidePageCount * 2 == displayPageCount{
			lowerCount = sidePageCount - 1
		}
		paging.EndPage = pageNo + upperCount
		d := 0
		// 如果终止页超过总页数,则将多出的页数补偿到起始页
		if paging.EndPage > totalPages {
			d = paging.EndPage - totalPages
			paging.EndPage = totalPages
		}
		paging.StartPage = pageNo - lowerCount - d
		// 起始页补偿
		if paging.StartPage < 1 {
			d = 1 - paging.StartPage
			paging.StartPage = 1
			paging.EndPage += d
			if paging.EndPage > paging.TotalPages {
				paging.EndPage = paging.TotalPages
			}
		}
		if paging.StartPage > paging.PageNo {
			paging.StartPage = paging.PageNo
		}

		paging.PrePageNo = paging.StartPage - 1
		if paging.PrePageNo < 1 {
			paging.PrePageNo = 1
		}
		paging.NextPageNo = paging.EndPage + 1
		if paging.NextPageNo > paging.TotalPages {
			paging.NextPageNo = paging.TotalPages
		}
		for i := paging.StartPage; i <= paging.EndPage; i++ {
			paging.PageNumbers = append(paging.PageNumbers, i)
		}
		return paging
	},
}
