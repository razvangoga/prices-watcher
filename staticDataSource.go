package main

const Store_Bike24 = "bike24"
const Store_FahhradDE = "fahhrad.de"

const Product_Garmin_1030 = "Garmin Edge 1030"
const Product_Garmin_820 = "Garmin Edge 820"

type StaticDataSource struct{}

func (StaticDataSource) get() pagesArray {
	return pagesArray{
		page{
			Site:       Store_Bike24,
			Product:    Product_Garmin_1030,
			Url:        "https://www.bike24.de/1.php?content=8;product=242780",
			XpathQuery: "//span[@itemprop='price']",
		},
		page{
			Site:       Store_Bike24,
			Product:    Product_Garmin_820,
			Url:        "https://www.bike24.de/1.php?content=8;product=182408",
			XpathQuery: "//span[@itemprop='price']",
		},
		page{
			Site:       Store_FahhradDE,
			Product:    Product_Garmin_1030,
			Url:        "https://www.fahrrad.de/garmin-edge-1030-gps-fahrradcomputer-bundle-schwarz-696317.html",
			XpathQuery: "//div[@class='cyc-margin_right-2 cyc-typo_display-3  cyc-color-text_sale']",
		},
		page{
			Site:       Store_FahhradDE,
			Product:    Product_Garmin_820,
			Url:        "https://www.fahrrad.de/garmin-edge-820-gps-fahrradcomputer-inkl-premium-hf-brustgurt-geschwindigtrittfrequenz-531556.html",
			XpathQuery: "//div[@class='cyc-margin_right-2 cyc-typo_display-3  cyc-color-text_sale']",
		},
	}
}