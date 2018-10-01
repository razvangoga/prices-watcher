package main


import (
	"github.com/antchfx/htmlquery"
	"strconv"
	"strings"
)

func GetPrice(page page) float64 {
	htmlDom, err := htmlquery.LoadURL(page.url)

	if err != nil {
		panic(err)
	}

	priceNode := htmlquery.FindOne(htmlDom, page.xpathQuery)

	rawPrice := htmlquery.InnerText(priceNode)
	price := strings.Replace(rawPrice, "€", "", -1)
	price = strings.Replace(price, ".", "", -1)
	price = strings.Replace(price, ",", ".", -1)
	price = strings.TrimSpace(price)

	priceN, err := strconv.ParseFloat(price, 64)

	if err != nil {
		panic(err)
	}

	return priceN
}
