package util

import (
	"fmt"
	"testing"
)

// spy 2020/1/21

func TestDeleteMapItem(t *testing.T) {
	countryCapitalMap := make(map[string]string)

	countryCapitalMap [ "France" ] = "巴黎"
	countryCapitalMap [ "Italy" ] = "罗马"
	countryCapitalMap [ "Japan" ] = "东京"
	countryCapitalMap [ "India " ] = "新德里"

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [country])
	}
	logger.Info("----")

	delete(countryCapitalMap, "Japan")
	//DeleteMapItem(countryCapitalMap,"Japan","India")

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [country])
	}
}
