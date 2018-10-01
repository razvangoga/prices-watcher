package main

const Store_Bike24 = "bike24"
const Store_FahhradDE = "fahhrad.de"

const Product_Garmin_1030 = "Garmin Edge 1030"
const Product_Garmin_820 = "Garmin Edge 820"

func getPagesToWatch() pagesArray {
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