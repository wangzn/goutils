// @Author: wangzn04@gmail.com
// @Date: 2018-08-13 15:35:51

package mymap

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// ParseMaps parses a string to a map[string]string
func ParseMaps(str, s1, s2 string) map[string]string {
	res := make(map[string]string)
	p1 := strings.Split(str, s1)
	for _, e := range p1 {
		p2 := strings.Split(e, s2)
		if len(p2) == 2 {
			res[p2[0]] = p2[1]
		} else {
			res[p2[0]] = ""
		}
	}
	return res
}

// StringMustString returns a must value of a map
func StringMustString(m map[string]string, k string) string {
	def := ""
	if v, ok := m[k]; ok {
		return v
	}
	return def
}

// StringMustInt returns a must value of int
func StringMustInt(m map[string]string, k string) int {
	def := 0
	str := StringMustString(m, k)
	if str == "" {
		return def
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		return def
	}
	return i
}

// StringMustFloat64 returns a must value of float64
func StringMustFloat64(m map[string]string, k string) float64 {
	def := 0.0
	str := StringMustString(m, k)
	if str == "" {
		return def
	}
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return def
	}
	return f
}

// FormatSlices returns the string of slices
func FormatSlices(header []string, body [][]string, format string) string {
	res := ""
	switch format {
	case "json":
		bs, _ := json.Marshal(body)
		res = string(bs)
	default:
		if len(body) == 0 {
			return res
		}
		b := new(bytes.Buffer)
		table := tablewriter.NewWriter(b)
		table.SetHeader(header)
		for _, r := range body {
			table.Append(r)
		}
		if len(header) > 2 {
			total := make([]string, len(header))
			total[len(total)-1] = strconv.Itoa(len(body))
			total[len(total)-2] = "TOTAL"
			table.SetFooter(total)
		}
		table.Render()
		res = b.String()
	}
	return res
}

// FormatMapslice returns the string of mapslice
func FormatMapslice(rs []map[string]string, format string) string {
	res := ""
	if rs == nil {
		return res
	}
	switch format {
	case "json":
		bs, _ := json.Marshal(rs)
		res = string(bs)
	default:
		if len(rs) == 0 {
			return res
		}
		b := new(bytes.Buffer)
		table := tablewriter.NewWriter(b)
		header := MapKeys(rs[0])
		table.SetHeader(header)
		for _, r := range rs {
			table.Append(MapValues(r, header))
		}
		if len(header) > 2 {
			total := make([]string, len(header))
			total[len(total)-1] = strconv.Itoa(len(rs))
			total[len(total)-2] = "TOTAL"
			table.SetFooter(total)
		}
		table.Render()
		res = b.String()
	}
	return res
}

// MapKeys returns the keys of a map
func MapKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// MapValues returns the values of a map
func MapValues(m map[string]string, ks []string) []string {
	vs := make([]string, 0, len(m))
	if ks == nil || len(ks) == 0 {
		for _, v := range m {
			vs = append(vs, v)
		}
	} else {
		for _, k := range ks {
			if v, ok := m[k]; ok {
				vs = append(vs, v)
			}
		}
	}
	return vs
}
