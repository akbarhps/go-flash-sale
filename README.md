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

## Feature:
```bash
./go-flash-sale -h

Usage of sale:
  -c value
        optional!
        kode angka kategory dipisahkan dengan koma
        - 10 Serba Seribu
        - 35 Super Brand day
        - 39 HP & Acc
        - 41 Komputer & Acc
        - 50 Elektronik & Photography
  -l int
        required!
        limit item yang akan ditampilkan, maksimal 90 (default 10)
  -p int
        optional!
        halaman pagination yang akan ditampilkan, default 1 (default 1)
  -v    optional!
        menampilkan voucher yang tersedia, default false
```

> -c merupakan category item pada shopee indonesia