/*
* @Author: scottxiong
* @Date:   2019-08-23 05:25:01
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-08-23 05:35:03
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://raw.githubusercontent.com/scott-x/templates/master/css/table/tb.html")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error Status Code: %d", resp.StatusCode)
		return
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)
	f, err := os.Create("./tb.html")
	if err != nil {
		panic(err)
	}
	f.Write(all)
}
