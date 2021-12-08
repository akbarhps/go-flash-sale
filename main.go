package main

import (
	"fmt"
	"go-flash-sale/model"
	"go-flash-sale/shopee"
	"log"
)

func main() {
	appFlag := model.AppFlag{}
	appFlag.Init()

	shopeeObj := shopee.Shopee{}
	fsItemIdsResponse, err := shopeeObj.GetFlashSaleItemIds()
	if err != nil {
		log.Fatal(err)
	}

	fsItemIds := fsItemIdsResponse.Data.
		FilterItemsByCategories(appFlag.FilterCategories).
		PaginateItems(appFlag.Page, appFlag.Limit).
		GetItemIds()

	fsItems, err := shopeeObj.GetFlashSaleItems(fsItemIds)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range fsItems.Data.Items {
		fmt.Println(item.ToString(appFlag.IsShowAvailableVoucher), "\n\n")
	}
}
