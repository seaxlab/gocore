package mapx

import (
	"fmt"
	"testing"
)

// spy 2020/1/21

func TestDeleteMapItem(t *testing.T) {
	countryMap := make(map[string]string)

	countryMap["France"] = "巴黎"
	countryMap["Italy"] = "罗马"
	countryMap["Japan"] = "东京"
	countryMap["India "] = "新德里"

	/*使用键输出地图值 */
	for country := range countryMap {
		fmt.Println(country, "首都是", countryMap[country])
	}
	fmt.Println("----")

	delete(countryMap, "Japan")
	//DeleteMapItem(countryMap,"Japan","India")

	/*使用键输出地图值 */
	for country := range countryMap {
		fmt.Println(country, "首都是", countryMap[country])
	}
}
