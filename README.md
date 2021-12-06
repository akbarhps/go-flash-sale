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

## Flags:
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

## Example:

```bash
~
➜ ./go-flash-sale.exe -l=4 -c=41
https://shopee.co.id/MicPhilips-SHE-8105-Earphone-with-Mic-i.25707188.310545596
Name: Philips SHE 8105 Earphone with Mic
OriginalPrice: Rp. 375000 | SalePrice: Rp. 199000 | Stock : 59



https://shopee.co.id/Keyboard-and-Mouse-Combo-Set-Ultra-Slim-Wireless-Multimedia-Silent-CLIPtec-RZK360-Mini-AirXilent-i.71042912.13445845191
Name: Keyboard & Mouse Combo Set Ultra-Slim Wireless Multimedia Silent Cliptec Rzk360 (Mini-Airxilent)
OriginalPrice: Rp. 220000 | SalePrice: Rp. 176000 | Stock : 20



https://shopee.co.id/FlexibleMini-Proyektor-Unitech-LED-YS500-Micro-Projector-Support-1080P-Mini-Theater-Gratis-Tripod-Flexible-i.1596789.8852351920
Name: Mini Proyektor Unitech LED YS500 Micro Projector Support 1080P Mini Theater
OriginalPrice: Rp. 700000 | SalePrice: Rp. 419900 | Stock : 15



https://shopee.co.id/VIVOHeadset-Gaming-hedset-BASIKE-Earphone-iPhone-Bass-In-ear-3-5mm-dengan-Mikrofon-tablet-Android-VIVO-i.227057943.4842942289
Name: Headset Gaming Basike iPhone Bass In-Ear 3.5mm dengan Mikrofon
OriginalPrice: Rp. 60000 | SalePrice: Rp. 8900 | Stock : 1192
```