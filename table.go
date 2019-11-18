/*
* @Author: scottxiong
* @Date:   2019-08-23 05:25:01
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-11-18 19:12:53
 */
package main

import (
	"flag"
	"fmt"
	"github.com/scott-x/gutils/fs"
	"github.com/scott-x/table/parse"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	customed := flag.Bool("c", false, "costomed:default false")
	flag.Parse()
	if *customed {
		// 用户自定义table
		flag, lines := parse.Parse("./a.txt")
		if flag {
			parse.GenTable(lines)
		}

	} else {
		genTable()
	}

}

func genTable() {
	resp, err := http.Get("https://raw.githubusercontent.com/scott-x/templates/master/css/table/tb.html")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		tb_local := ""
		fmt.Printf("Error Status Code: %d on github\n", resp.StatusCode)
		fmt.Printf("trying to fetch local data...\n")

		gopath, flag := os.LookupEnv("GOPATH")
		if flag {
			tb_local = gopath + "/src/github.com/scott-x/templates/css/table/tb.html"
		}
		content, err := fs.ReadFile1(tb_local)
		if err != nil {
			panic(err)
		}
		f, err := os.Create("./tb.html")
		if err != nil {
			panic(err)
		}
		_, err = f.Write([]byte(content))
		if err != nil {
			panic(err)
		}
		fmt.Printf("done...\n")
		return
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%s\n", all)
	f, err := os.Create("./tb.html")
	if err != nil {
		panic(err)
	}
	_, err = f.Write(all)
	if err != nil {
		panic(err)
	}
}
