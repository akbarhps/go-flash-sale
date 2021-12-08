package main

import (
	"flag"
	"fmt"
	"go-flash-sale/model"
	shopee "go-flash-sale/shopee"
	"log"
	"os"
)

var flagLimitName = "l"
var flagLimitUsage = "required!\nlimit item yang akan ditampilkan, maksimal 90"
var flagLimitDefault = 10

var flagPageName = "p"
var flagPageUsage = "optional!\nhalaman pagination yang akan ditampilkan"
var flagPageDefault = 1

var flagShowVoucherName = "v"
var flagShowVoucherUsage = "optional!\nmenampilkan voucher yang tersedia (default false)"

var flagCategoriesName = "c"
var flagCategoriesUsage = "optional!\nkode angka kategory dipisahkan dengan koma \n- 10 Serba Seribu\n- 35 Super Brand day\n- 39 HP & Acc\n- 41 Komputer & Acc\n- 50 Elektronik & Photography"

var categories model.CategoriesFlag
var limit *int
var page *int
var isShowVoucher *bool

func init() {
	flag.Var(&categories, flagCategoriesName, flagCategoriesUsage)
	limit = flag.Int(flagLimitName, flagLimitDefault, flagLimitUsage)
	page = flag.Int(flagPageName, flagPageDefault, flagPageUsage)
	isShowVoucher = flag.Bool(flagShowVoucherName, false, flagShowVoucherUsage)

	help := false
	for _, i := range os.Args {
		if i == "help" || i == "--h" || i == "--help" {
			help = true
		}
	}

	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	shopeeObj := shopee.Shopee{}
	fsItemIdsResponse, err := shopeeObj.GetFlashSaleItemIds()
	if err != nil {
		log.Fatal(err)
	}

	fsItemIds := fsItemIdsResponse.Data.
		FilterItemsByCategories(categories).
		PaginateItems(*page, *limit).
		GetItemIds()

	fsItems, err := shopeeObj.GetFlashSaleItems(fsItemIds)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range fsItems.Data.Items {
		fmt.Println(item.ToString(*isShowVoucher), "\n\n")
	}
}
