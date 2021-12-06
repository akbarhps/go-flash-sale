package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go-flash-sale/model"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func sendRequest(method string, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("sendRequest:", err)
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("sendRequest:", err)
		return nil, err
	}

	return respBody, nil
}

func getItemIds() *model.FlashSaleItemIdResponse {
	url := "https://shopee.co.id/api/v2/flash_sale/get_all_itemids?sort_soldout=true"
	responseBody, err := sendRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal("getItemIds 1:", err)
	}

	itemResponse := &model.FlashSaleItemIdResponse{}
	err = json.Unmarshal(responseBody, itemResponse)
	if err != nil {
		log.Fatal("getItemIds 2:", err)
	}

	return itemResponse
}

func getItemByIds(itemIds []int64) *model.FlashSaleItemResponse {
	url := "https://shopee.co.id/api/v2/flash_sale/flash_sale_batch_get_items"
	requestBody, _ := json.Marshal(&model.FlashSaleItemRequest{
		ItemIds:       itemIds,
		Limit:         len(itemIds),
		SortBySoldOut: true,
		WithDpItems:   true,
	})
	responseBody, err := sendRequest(http.MethodPost, url, bytes.NewReader(requestBody))
	if err != nil {
		log.Fatal("getItemIds:", err)
	}

	fsResponse := &model.FlashSaleItemResponse{}
	err = json.Unmarshal(responseBody, fsResponse)
	if err != nil {
		log.Fatal("getItemByIds:", err)
	}

	return fsResponse
}

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
	itemResponse := getItemIds()
	catFlagLen := len(categories)
	var filteredIds []int64
	for _, item := range itemResponse.Data.ItemList {
		// item sold or not in categories flag
		if item.Sold || (catFlagLen > 0 && !categories.Contains(item.CategoryId)) {
			continue
		}
		filteredIds = append(filteredIds, item.ItemId)
	}

	counter := 0
	var paginatedIds []int64
	for i, item := range filteredIds {
		if *page > 1 && i < *page*(*limit) {
			continue
		}
		if counter == *limit {
			break
		}
		counter++
		paginatedIds = append(paginatedIds, item)
	}

	fsResponse := getItemByIds(paginatedIds)
	for _, item := range fsResponse.Data.Items {
		fmt.Println(item.ToString(*isShowVoucher), "\n\n")
	}
}
