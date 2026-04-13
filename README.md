# Türkçe Domain Desteği

Bu fork, Gophish içinde `SMTP From` ve ilgili e-posta adresi doğrulama akışlarını güncelleyerek Türkçe karakter içeren alan adlarını destekler.

## Neden gerekliydi?

Orijinal akışta `SMTP From` alanı çok katı bir regex ile doğrulanıyordu. Bu nedenle aşağıdaki gibi uluslararasılaştırılmış domain adları kabul edilmiyordu:

- `enes@düzcecam.com`
- `info@örnek.com`

Sistem bu adresleri:

`Invalid SMTP From address because it is not an email address`

hatası ile reddediyordu.

## Bu fork neyi değiştiriyor?

Bu sürümde e-posta adreslerinin domain kısmı IDN uyumlu şekilde işlenir.

- Unicode domain kabul edilir
- Domain kısmı SMTP uyumlu ASCII/punycode biçimine çevrilir
- Gönderim ve doğrulama akışları ortak normalize mantığı kullanır
- Eski regex tabanlı kısıt kaldırılmıştır

Örnek:

- Giriş: `enesalbayrak@düzcecam.com`
- İç işleme / SMTP uyumlu form: punycode dönüşümü ile ASCII domain

## Hangi alanlarda düzeltildi?

Aşağıdaki akışlar güncellendi:

- Sending Profile doğrulaması
- Test email gönderimi
- Campaign mail üretimi
- Template context / From çözümleme akışı

## Orijinal sürümden farkı

Orijinal sürüm:

- Sadece klasik ASCII domainleri güvenle kabul ediyordu
- IDN domainlerde hataya düşebiliyordu
- `SMTP From` doğrulaması gereğinden katıydı

Bu fork:

- Türkçe domain kullanan adresleri destekler
- Domain kısmını SMTP için uygun hale getirir
- Modern IDN kullanımına daha uyumludur

## Sınırlar

Bu değişiklik özellikle domain kısmını hedefler.

Desteklenen:
- `kullanici@örnek.com`

Henüz sınırlı olabilecek durum:
- `çağrı@ornek.com`

Yani domain tarafında Türkçe karakter desteği vardır. Local-part tarafında (`@` öncesi bölüm) tam UTF-8 mail desteği, SMTP sunucusunun `SMTPUTF8` desteğine bağlıdır ve ayrı bir konudur.

## Kurulum

Kaynak kodu çektikten sonra binary yeniden derlenmelidir:

```bash
git pull
go build -o gophish
./gophish
