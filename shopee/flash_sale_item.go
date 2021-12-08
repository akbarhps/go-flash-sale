package shopee

import (
	"fmt"
	"strconv"
)

type FlashSaleItemRequest struct {
	ItemIds       []int64 `json:"itemids"`
	Limit         int     `json:"limit"`
	SortBySoldOut bool    `json:"sort_soldout"`
	WithDpItems   bool    `json:"with_dp_items"`
}

type FlashSaleItemResponse struct {
	Data FlashSaleItemData `json:"data"`
}

type FlashSaleItemData struct {
	Items []FlashSaleItem `json:"items"`
}

type FlashSaleItem struct {
	ItemId         int64   `json:"itemid"`
	ShopId         int64   `json:"shopid"`
	OriginalPrice  int64   `json:"price_before_discount"`
	SalePrice      int64   `json:"price"`
	Stock          int32   `json:"stock"`
	PromoName      string  `json:"promo_name"`
	Name           string  `json:"name"`
	IsMart         bool    `json:"is_mart"`
	IsShopOfficial bool    `json:"is_shop_official"`
	Voucher        Voucher `json:"voucher"`
}

type Voucher struct {
	MinSpend           int64  `json:"min_spend"`
	Code               string `json:"voucher_code"`
	DiscountPercentage int32  `json:"discount_percentage"`
	DiscountValue      int64  `json:"discount_value"`
}

func (i *FlashSaleItem) GenerateLink() string {
	var wordList []string
	tempStr := ""
	for _, char := range i.Name {
		if char >= 65 && char <= 90 || char >= 97 && char <= 122 || char >= 48 && char <= 57 {
			tempStr += string(char)
			continue
		}
		if tempStr != "" {
			wordList = append(wordList, tempStr)
		}
		tempStr = ""
	}
	if tempStr != "" {
		wordList = append(wordList, tempStr)
	}
	for _, word := range wordList {
		tempStr += word + "-"
	}
	return fmt.Sprintf("https://shopee.co.id/%si.%s.%s", tempStr, strconv.FormatInt(i.ShopId, 10), strconv.FormatInt(i.ItemId, 10))
}

func (i *FlashSaleItem) ToString(isShowVoucher bool) string {
	formatted := fmt.Sprintf("%s\nName: %s\nOriginalPrice: Rp. %d | SalePrice: Rp. %d | Stock : %d\n", i.GenerateLink(), i.PromoName, i.OriginalPrice/100000, i.SalePrice/100000, i.Stock)
	if isShowVoucher && i.Voucher != (Voucher{}) {
		v := i.Voucher
		formatted += fmt.Sprintf("Voucher -> Code: Rp. %s | Value: %d | MinSpend: %d\n", v.Code, v.DiscountValue/100000, v.MinSpend/100000)
	}
	return formatted
}
