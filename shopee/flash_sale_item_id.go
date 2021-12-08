package shopee

import "go-flash-sale/helper"

type FlashSaleItemIdResponse struct {
	Data FlashSaleItemIdData `json:"data"`
}

type FlashSaleItemIdData struct {
	Items []FlashSaleItemId `json:"item_brief_list"`
}

type FlashSaleItemId struct {
	Sold       bool  `json:"is_soldout"`
	CategoryId int   `json:"catid"`
	ItemId     int64 `json:"itemid"`
}

func (d *FlashSaleItemIdData) FilterItemsByCategories(categoryIds []int) *FlashSaleItemIdData {
	var filteredItems []FlashSaleItemId
	for _, item := range d.Items {
		if item.Sold || (len(categoryIds) > 0 && !helper.ContainsInt(categoryIds, item.CategoryId)) {
			continue
		}
		filteredItems = append(filteredItems,  item)
	}
	d.Items = filteredItems
	return d
}

func (d *FlashSaleItemIdData) PaginateItems(page, perPage int) *FlashSaleItemIdData {
	var paginatedItems []FlashSaleItemId
	counter := 0
	for i, item := range d.Items {
		if counter == perPage {
			break
		}
		if page > 1 && i < page * perPage {
			continue
		}
		counter++
		paginatedItems = append(paginatedItems, item)
	}
	d.Items = paginatedItems
	return d
}

func (d *FlashSaleItemIdData) GetItemIds() []int64 {
	var itemIds []int64
	for _, item := range d.Items {
		itemIds = append(itemIds, item.ItemId)
	}
	return itemIds
}