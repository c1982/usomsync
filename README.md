# USOM Zararlı Bağlantı Listesi Exim Blacklist Senkronizasyonu
Ulusal Siber Olaylara Müdahale Merkezi (USOM) Türkiye özelinde kendi yaptığı çalışmalar neticesinde veya ihbar yolu ile tespit ettiği zararlı (malware) yazılım yayan dağıtan veya oltalama saldırısında (Phising) bulunan IP adreslerini veya Domain isimlerini www.usom.gov.tr üzerinden yayınlamaktadır.

USOM Zararlı Bağlantı Listesi [https://www.usom.gov.tr/rss/zararli-baglanti.rss]() adresinde RSS standartlarında XML formatında ayrıca yayınlanmaktadır. 

Bu uygulama USOM'un yayınladığı bu zararlı listesini Exim Mail Sunucusu'nun Domain ve IP Black List dosyası ile eşleştirerek yeni tespit edilmiş olan zararlı yazılımlardan ve Phising saldırılarını korunmanızı sağlar.

## Gereksinimler
1. Exim Mail Server

## Parametreler
-url: USOM Zararlı Bağlantı Listesi'nin web adresini belirleyen parametre. Varsayılan olarak https://www.usom.gov.tr/rss/zararli-baglanti.rss değerini alır.

-d: Uygulamanın zararlı bulunan host isimlerinin hangi dosyaya kayıtedileceğini belirleyen parametre. Varsayılan olarak /etc/blockeddomains değerini alır.

-ip: Uygulamanın zararlı bulunan IP Adreslerinin hangi dosyaya kayıtedileceğini belirleyen parametre. Varsayılan olarak /etc/spammeripblocks değerini alır.

## Kullanım
`usomsync -url=https://www.usom.gov.tr/rss/zararli-baglanti.rss -d=/etc/blockeddomains -ip=/etc/spammeripblocks`

## Kurulum
1. usomsync binary dosyasını [https://github.com/c1982/usomsync/releases]() adresinden temin edin.
2. Sunucunuzun _/bin_ klasörü altına kopyalayın (/bin/usomsync).
3. `chmod 754 usomsync` komutu ile uygulamayı çalışır hale getirin.
4. Cronjob'a ekleyin. Bunun için;
	1. `crontab -e` ile zamanlanmış görevler dosyasını açın.
	2. dosyanın son satırına `@daily /bin/usomsync` ibaresini ekleyin.
	3. /etc/blockeddomains dosyası yok ise `touch /etc/blockeddomains` ile oluşturun.

> Not: Bu uygulama cPanel üzerinde test edilmiştir. blockeddomains ve spammeripblocks dosyalarında daha önce yaptığınız değişiklikler aynen korunur.