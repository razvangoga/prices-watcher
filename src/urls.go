package main

import (
	"github.com/antchfx/htmlquery"
	"strconv"
	"strings"
)

const Store_Bike24 = "bike24"
const Store_FahhradDE = "fahhrad.de"

const Product_Garmin_1030 = "Garmin Edge 1030"
const Product_Garmin_820 = "Garmin Edge 820"

type page struct {
	site       string
	product    string
	url        string
	xpathQuery string
}

type pagesArray []page

func GetDataSource() pagesArray {
	return pagesArray{
		page{
			site:       Store_Bike24,
			product:    Product_Garmin_1030,
			url:        "https://www.bike24.de/1.php?content=8;product=242780",
			xpathQuery: "//span[@itemprop='price']",
		},
		page{
			site:       Store_Bike24,
			product:    Product_Garmin_820,
			url:        "https://www.bike24.de/1.php?content=8;product=182408",
			xpathQuery: "//span[@itemprop='price']",
		},
		page{
			site:       Store_FahhradDE,
			product:    Product_Garmin_1030,
			url:        "https://www.fahrrad.de/garmin-edge-1030-gps-fahrradcomputer-bundle-schwarz-696317.html",
			xpathQuery: "//div[@class='cyc-margin_right-2 cyc-typo_display-3  cyc-color-text_sale']",
		},
		page{
			site:       Store_FahhradDE,
			product:    Product_Garmin_820,
			url:        "https://www.fahrrad.de/garmin-edge-820-gps-fahrradcomputer-inkl-premium-hf-brustgurt-geschwindigtrittfrequenz-531556.html",
			xpathQuery: "//div[@class='cyc-margin_right-2 cyc-typo_display-3  cyc-color-text_sale']",
		},
	}
}

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
