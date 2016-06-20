package tplengine

import (
	"encoding/json"
	"html/template"
)

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
}
