# WebScraperTool
- Yıldız CTI takım içi görevde hazırladığım bir araçtır. 
- Görev Tanımı: Go Dilinde Web Sitesi Anlık Durum Takip Botu Geliştirme.
- Görev Talimatı Detaylı: Go dilinde bir web sitesi anlık durum takip botu geliştirmeniz gerekmektedir. Bot, her 5 dakikada bir hedef web sitesini anlık olarak kontrol edecek, sitenin erişilebilir olup olmadığını (up/down durumu) ve yanıt süresini (ms cinsinden) Telegram üzerinden bildirecek. Eğer yanıt süresi 100 ms’nin üzerinde olursa, sistem Telegram’a yavaşlık uyarısı gönderecek.

- TOOL HAKKINDA : * gopkg.in/tucnak/telebot.v2, Go dilinde Telegram botları oluşturmak için kullanılan bir kütüphanedir. Bu kütüphaneyi kullanarak çektiğim veriyi Telegram'a yönlendirdim.
  * Yazdığımız bot istekleri telegram üzerine göndereceğinden dolayı kod içerisinde; Bot token ve chat ID açık bir şekilde gözükmemesi için "os" paketini kullandım. os.Getenv("TELEGRAMBOTTOKEN") ve os.Getenv("TELEGRAMCHATID") ile, sistemde tanımlı olan çevresel değişkenlerden Telegram botunun token'ını ve chat ID'sini alır. Bu, güvenli bir şekilde hassas bilgileri kod içinde saklamadan kullanmanıza olanak tanır.
  * Kodların diğer açıklamaları main.go içerisinde mevcuttur.

- TOOL DETAYLI KULLANIMI : * Öncelikle Telegram bot token'ınızı ve chat ID'nizi ayarlamak için:
  Linux/Mac Kullanıyorsanız, terminale bu komutları girmeniz gerekir: export TELEGRAMBOTTOKEN="Tokeniniz" , export TELEGRAMCHATID= Chat ID'niz.
  Windows kullanıyorsanız CMD'den aynı komutları başına set koyarak girmelisiniz.
  * Telegram bot tokeniniz yoksa oluşturmak için: Telegram bot token'ını almak için BotFather ile iletişime geçip /newbot komutunu kullanarak bot oluşturun ve verilen token'ı kaydedin.
  * Chat ID'nizi öğrenmek için: Chat ID'yi öğrenmek için botunuza mesaj gönderip, ardından https://api.telegram.org/bot<TELEGRAMBOTTOKEN>/getUpdates URL'sini ziyaret ederek yanıt içindeki "chat": {"id": ...} kısmını kontrol edin.
  * Bilgileri girdikten sonra terminale go run main.go yazıp ilgili verileri çekmek istediğiniz web sitesini girin.
  * İYİ ÇALIŞMALAR - EMİRCAN ALTUNTAŞ
    
  
