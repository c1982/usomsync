# USOM Zararlı Bağlantı Listesi Exim Mail Server Senkronizasyonu

Ulusal Siber Olaylara Müdahale Merkezi (USOM) Türkiye özelinde kendi yaptığı çalışmalar neticesinde veya ihbar yolu ile tespit ettiği zararlı (malware) yazılım yayan veya oltalama saldırısında (phising) bulunan IP adreslerini ve Domain isimlerini www.usom.gov.tr üzerinden yayınlamaktadır.

USOM Zararlı Bağlantı Listesi [https://www.usom.gov.tr/rss/zararli-baglanti.rss]() adresinde RSS standartlarında XML formatında ayrıca yayınlanmaktadır. 

Usomsync uygulaması USOM'un yayınladığı bu zararlı listesini [Exim Mail Sunucusu](http://www.exim.org/)  konfigürasyonuna ekleyerek yeni tespit edilmiş olan zararlı yazılımlardan ve phising saldırılarından korunmanızı sağlar.

## Gereksinimler

1. Exim Mail Server ([^](http://www.exim.org/))

## Parametreler

**-url:** USOM Zararlı Bağlantı Listesi'nin web adresini belirleyen parametre. Varsayılan olarak https://www.usom.gov.tr/rss/zararli-baglanti.rss değerini alır.

**-d:** Uygulamanın zararlı bulunan host isimlerinin hangi dosyaya kayıtedileceğini belirleyen parametre. Varsayılan olarak /etc/blockeddomains değerini alır.

**-ip:** Uygulamanın zararlı bulunan IP Adreslerinin hangi dosyaya kayıtedileceğini belirleyen parametre. Varsayılan olarak /etc/spammeripblocks değerini alır.

## Kullanım

`/bin/usomsync -url=https://www.usom.gov.tr/rss/zararli-baglanti.rss -d=/etc/blockeddomains -ip=/etc/spammeripblocks`

## Kurulum

1. Wget ile uygulamayı sunucunuza yükleyin  `wget https://github.com/c1982/usomsync/releases/download/v1.0/usomsync -O /bin/usomsync`

2. `chmod 754 /bin/usomsync` komutu ile uygulamayı çalışır hale getirin.

3. Uygulamayı crontab görevlerine günlük olarak ekleyin.

	1. `crontab -e` ile zamanlanmış görevler dosyasını açın.
	
	2. dosyanın son satırına `@daily /bin/usomsync` ibaresini ekleyin.
	
	3. /etc/blockeddomains dosyası yok ise `touch /etc/blockeddomains` ile oluşturun.

> Not: Bu uygulama cPanel üzerinde test edilmiştir. blockeddomains ve spammeripblocks dosyalarında daha önce yaptığınız değişiklikler aynen korunur.

##  Postfix Uyarlaması (by [@merttokgozoglu](https://twitter.com/merttokgozoglu))

usomsync çalıştıktan sonra hemen arkasından aşağıdaki scripti çalıştırarak Postfix'e uyumlu hale getirebilirsiniz.

```bash
#!/bin/bash
sed -i -e '/\#USOM KURALLARI/,$d' /etc/postfix/sndr
rm /tmp/blockeddomains.tmp /tmp/blockeddomains /tmp/spammeripblocks && touch /tmp/blockeddomains.tmp /tmp/blockeddomains /tmp/spammeripblocks
/bin/usomsync -url=https://www.usom.gov.tr/rss/zararli-baglanti.rss -d=/tmp/blockeddomains.tmp -ip=/tmp/spammeripblocks
cat /tmp/blockeddomains.tmp|grep -v "/">/tmp/blockeddomains
sed -i -e 's/^/\//' /tmp/blockeddomains && sed -i.bak 's/$/\/REJECT USOM/' /tmp/blockeddomains && sed -i '1s/^/\#USOM KURALLARI\n/' /tmp/blockeddomains
sed -i.bak 's/$/\ REJECT USOM/' /tmp/spammeripblocks
cat /tmp/spammeripblocks>/etc/postfix/client_ips_custom
postmap /etc/postfix/client_ips_custom
cat /tmp/blockeddomains>>/etc/postfix/sndr
/etc/init.d/postfix restart
```