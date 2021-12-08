package shopee

import (
	"fmt"
	"go-flash-sale/helper"
	"log"
	"net/http"
)

var flashSaleItemIdsEndpoint = "https://shopee.co.id/api/v2/flash_sale/get_all_itemids?sort_soldout=true"
var flashSaleItemEndpoint = "https://shopee.co.id/api/v2/flash_sale/flash_sale_batch_get_items"

type Shopee struct {
}

func (s *Shopee) GetFlashSaleItemIds() (*FlashSaleItemIdResponse, error) {
	responseBody, err := helper.Fetch(http.MethodGet, flashSaleItemIdsEndpoint, nil)
	if err != nil {
		log.Fatal("GetFlashSaleItemIds request:", err)
		return nil, err
	}

	fsItemIdResponse := &FlashSaleItemIdResponse{}
	err = helper.ResponseToInterface(responseBody, fsItemIdResponse)
	if err != nil {
		log.Fatal("GetFlashSaleItemIds marshall:", err)
		return nil, err
	}

	return fsItemIdResponse, nil
}

func (s *Shopee) GetFlashSaleItems(itemIds []int64) (*FlashSaleItemResponse, error) {
	requestBody, err := helper.InterfaceToRequest(&FlashSaleItemRequest{
		ItemIds:       itemIds,
		Limit:         len(itemIds),
		SortBySoldOut: true,
		WithDpItems:   true,
	})
	if err != nil {
		log.Fatal("GetFlashSaleItems marshall:", err)
		return nil, err
	}

	responseBody, err := helper.Fetch(http.MethodPost, flashSaleItemEndpoint, requestBody)
	if err != nil {
		log.Fatal("GetFlashSaleItems request:", err)
		return nil, err
	}

	fsItemResponse := &FlashSaleItemResponse{}
	err = helper.ResponseToInterface(responseBody, fsItemResponse)
	if err != nil {
		log.Fatal("GetFlashSaleItems unmarshall:", err)
		return nil, err
	}

	return fsItemResponse, nil
}
