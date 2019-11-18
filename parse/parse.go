/*
* @Author: scottxiong
* @Date:   2019-11-18 14:19:01
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-11-18 16:08:16
 */
package parse

import (
	"github.com/scott-x/gutils/cl"
	"github.com/scott-x/gutils/fs"
	"github.com/scott-x/table/defs"
	"log"
	"os"
	"strings"
)

func Parse(txt_path string) *defs.Lines {
	lines := &defs.Lines{}
	exist := fs.IsExist(txt_path)
	if !exist {
		cl.BoldCyan.Printf("You can define a.txt with syntax of Markdown, the cell content is separated by \"|\":\n")
		cl.BoldGreen.Printf("Suplier|Name|Email|Phone\n")
		cl.BoldGreen.Printf("QTOP|Lillian/Sue|xxx@qtop.com.cn / yyy@qtop.com.cn|Tel:(0)574 8921 xxxx\n")
		cl.BoldGreen.Printf("Danson|Lynn|zzz@dansondecor.com|Mobile : (86)15019468xxx Tel : (86) 0755-8272 xxxx\n")
		cl.BoldGreen.Printf("Merry Art|Tina|xxx@merryart.cn|TEL：+86 574 8790xxxx EXT：xxx\n")
		f, err := os.OpenFile("./a.txt", os.O_CREATE, 0755)
		if err != nil {
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
		return lines
	}
	content := fs.ReadFile(txt_path)

	line_data := strings.Split(content, "\n")
	for _, data := range line_data {
		line := &defs.Line{}
		dt := strings.Split(data, "|")
		line.Data = append(line.Data, dt...)
		*lines = append(*lines, *line)
	}
	return lines
}

func GenTable(lines *defs.Lines) {

	data := ""
	data += "<!DOCTYPE html>\n" +
		"<html lang=\"en\">\n" +
		"<head>\n" +
		tab(4) + "<meta charset=\"UTF-8\">\n" +
		tab(4) + "<title>table</title>\n" +
		tab(4) + "<style>\n" +
		tab(8) + "table.gridtable {\n" +
		tab(12) + "font-family: verdana,arial,sans-serif;\n" +
		tab(12) + "font-size:11px;\n" +
		tab(12) + "color:#333333;\n" +
		tab(12) + "border-width: 1px;\n" +
		tab(12) + "border-color: #666666;\n" +
		tab(12) + "border-collapse: collapse;\n" +
		tab(8) + " }\n" +
		tab(8) + "table.gridtable th {\n" +
		tab(12) + "border-width: 1px;\n" +
		tab(12) + "padding: 8px;\n" +
		tab(12) + "border-style: solid;\n" +
		tab(12) + "border-color: #666666;\n" +
		tab(12) + "background-color: #dedede;\n" +
		tab(8) + "}\n" +
		tab(8) + "table.gridtable td {\n" +
		tab(12) + "border-width: 1px;\n" +
		tab(12) + "padding: 8px;\n" +
		tab(12) + "border-style: solid;\n" +
		tab(12) + "border-color: #666666;\n" +
		tab(12) + "background-color: #ffffff;\n" +
		tab(8) + "}\n" +
		tab(4) + "</style>\n" +
		"</head>\n" +
		"<body>\n" +
		tab(4) + "<table class=\"gridtable\">\n"
	for _, line := range *lines {
		data += tab(8) + "<tr>\n"
		for _, cell := range line.Data {
			// fmt.Println(cell)
			data += tab(12) + "<td>" + strings.Trim(cell, " ") + "</td>\n"
		}
		data += tab(8) + "<tr>\n"
	}
	data += tab(4) + "</table>\n" +
		"<body>\n" +
		"</html>"
	fs.WriteString("table.html", data)
}

func tab(n int) string {
	return strings.Repeat(" ", n)
}
