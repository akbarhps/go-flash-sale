package model

type FlashSaleItemIdResponse struct {
	Data FlashSaleItemIdData `json:"data"`
}

type FlashSaleItemIdData struct {
	ItemList []FlashSaleItemId `json:"item_brief_list"`
}

type FlashSaleItemId struct {
	Sold       bool  `json:"is_soldout"`
	CategoryId int   `json:"catid"`
	ItemId     int64 `json:"itemid"`
}
