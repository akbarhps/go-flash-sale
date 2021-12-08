## Golang Flash Sale

Aplikasi ini dibuat untuk membantu mencari produk flash sale di website e-commerce indonesia

> Saat ini hanya support website [Shopee Indonesia](https://shopee.co.id)

## Roadmap:
- Add More Shopee Categories
- Support Tokopedia
- Support Blibli

## Usage:
```bash
go build .
```

then run the output file (for example .exe)

## Flags:
```bash
➜ .\go-flash-sale.exe -h

  -c value
        categoryId barang yang ingin ditampilkan, dipisahkan dengan koma (,)
        - 10 Serba Seribu
        - 35 Super Brand day
        - 39 HP & Acc
        - 41 Komputer & Acc
        - 50 Elektronik & Photography
        - Lainnya buka https://shopee.co.id/flashsale dan lihat query categoryId=? ketika memilih salah satu kategory
        Contoh:

        -c=10,35,39

        (default all)
  -l int
        limit item yang akan ditampilkan, maksimal 90
        Contoh:

        -l=10

         (default 10)
  -p int
        halaman pagination yang akan ditampilkan
        Contoh:

        -p=2

         (default 1)
  -v    menampilkan voucher yang (jika tersedia)
        Contoh:

        -v=true

        (default false)
```

> -c merupakan category item pada shopee indonesia

## Example:

```bash
~
➜ .\go-flash-sale.exe -l=1 -p=1 -v=true -c=41

https://shopee.co.id/IMSJBL-Tune-T500-On-Ear-Headphones-Garansi-Resmi-IMS-i.28615794.3865204840
Name: JBL TUNE500 Headphone On-Ear
OriginalPrice: Rp. 599000 | SalePrice: Rp. 259000 | Stock : 50
Voucher -> Code: Rp. DESOGO200 | Value: 200000 | MinSpend: 0
```