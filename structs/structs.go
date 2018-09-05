// @Author: wangzn04@gmail.com
// @Date: 2017-08-10 21:22:52

package structs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/fatih/structs"
	"github.com/olekukonko/tablewriter"
)

// StructKeys returns all field names in struct
// b is a boolean for print embedded or not
func StructKeys(v interface{}, b bool) []string {
	s := structs.New(v)
	ks := make([]string, 0)
	for _, f := range s.Fields() {
		if f.IsEmbedded() {
			if b {
				tmp := StructKeys(f.Value(), b)
				ks = append(ks, tmp...)
			}
		} else {
			ks = append(ks, f.Name())
		}
	}
	return ks
}

// StructValues returns all values in struct
// b is a boolean for print embedded or not
func StructValues(v interface{}, b bool) []string {
	vs := make([]string, 0)
	s := structs.New(v)
	for _, f := range s.Fields() {
		if f.IsEmbedded() {
			if b {
				tmp := StructValues(f.Value(), b)
				vs = append(vs, tmp...)
			}
		} else {
			vs = append(vs, fmt.Sprintf("%v", f.Value()))
		}
	}
	return vs
}

// FormatResourceslice returns the string of resource slice with format
func FormatResourceslice(rs []interface{}, format string) string {
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
		header := StructKeys(rs[0], true)
		table.SetHeader(header)
		for _, r := range rs {
			table.Append(StructValues(r, true))
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
