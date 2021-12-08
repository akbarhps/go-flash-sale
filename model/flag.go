package model

import (
	"flag"
	"go-flash-sale/helper"
	"os"
)

type AppFlag struct {
	Limit                  int
	Page                   int
	IsShowAvailableVoucher bool
	FilterCategories       CategoriesFlag
}

var flagLimitName = "l"
var flagLimitUsage = `limit item yang akan ditampilkan, maksimal 90
Contoh: 

-l=10

`

var flagPageName = "p"
var flagPageUsage = `halaman pagination yang akan ditampilkan
Contoh: 

-p=2

`

var flagShowVoucherName = "v"
var flagShowVoucherUsage = `menampilkan voucher yang (jika tersedia)
Contoh: 

-v=true

(default false)`

var flagCategoriesName = "c"
var flagCategoriesUsage = `categoryId barang yang ingin ditampilkan, dipisahkan dengan koma (,)
- 10 Serba Seribu
- 35 Super Brand day
- 39 HP & Acc
- 41 Komputer & Acc
- 50 Elektronik & Photography
- Lainnya buka https://shopee.co.id/flashsale dan lihat query categoryId=? ketika memilih salah satu kategory
Contoh: 

-c=10,35,39

(default all)`

func (f *AppFlag) Init() {
	flag.Var(&f.FilterCategories, flagCategoriesName, flagCategoriesUsage)
	flag.IntVar(&f.Limit, flagLimitName, 10, flagLimitUsage)
	flag.IntVar(&f.Page, flagPageName, 1, flagPageUsage)
	flag.BoolVar(&f.IsShowAvailableVoucher, flagShowVoucherName, false, flagShowVoucherUsage)
	flag.Parse()

	helpSlice := &[]string{"h", "help", "-h", "--h", "-help", "--help"}
	// show usage when user type help, --h, or --help
	for _, i := range os.Args {
		if helper.ContainsString(helpSlice, i) {
			flag.Usage()
			os.Exit(1)
		}
	}
}
