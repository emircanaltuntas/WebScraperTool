TR
# WebScraperTool
- Yıldız CTI takım içi görevde hazırladığım bir araçtır. 
- Görev Tanımı: Go Dilinde Web Sitesi Anlık Durum Takip Botu Geliştirme.
- Görev Talimatı Detaylı: Go dilinde bir web sitesi anlık durum takip botu geliştirmeniz gerekmektedir. Bot, her 5 dakikada bir hedef web sitesini anlık olarak kontrol edecek, sitenin erişilebilir olup olmadığını (up/down durumu) ve yanıt süresini (ms cinsinden) Telegram üzerinden bildirecek. Eğer yanıt süresi 100 ms’nin üzerinde olursa, sistem Telegram’a yavaşlık uyarısı gönderecek.

- TOOL HAKKINDA :
  * gopkg.in/tucnak/telebot.v2, Go dilinde Telegram botları oluşturmak için kullanılan bir kütüphanedir. Bu kütüphaneyi kullanarak çektiğim veriyi Telegram'a yönlendirdim.
  * Yazdığımız bot istekleri telegram üzerine göndereceğinden dolayı kod içerisinde; Bot token ve chat ID açık bir şekilde gözükmemesi için "os" paketini kullandım. 
 os.Getenv("TELEGRAMBOTTOKEN") ve os.Getenv("TELEGRAMCHATID") ile, sistemde tanımlı olan çevresel değişkenlerden Telegram botunun token'ını ve chat ID'sini alır. Bu, güvenli bir şekilde 
 hassas bilgileri kod içinde saklamadan kullanmanıza olanak tanır.
  * Kodların diğer açıklamaları main.go içerisinde mevcuttur.

- TOOL DETAYLI KULLANIMI :
  * Öncelikle Telegram bot token'ınızı ve chat ID'nizi ayarlamak için:
 Linux/Mac Kullanıyorsanız, terminale bu komutları girmeniz gerekir: export TELEGRAMBOTTOKEN="Tokeniniz" , export TELEGRAMCHATID= Chat ID'niz.
 Windows kullanıyorsanız CMD'den aynı komutları başına set koyarak girmelisiniz.
  * Telegram bot tokeniniz yoksa oluşturmak için: Telegram bot token'ını almak için BotFather ile iletişime geçip /newbot komutunu kullanarak bot oluşturun ve verilen token'ı kaydedin.
  * Chat ID'nizi öğrenmek için: Chat ID'yi öğrenmek için botunuza mesaj gönderip, ardından https://api.telegram.org/bot<TELEGRAMBOTTOKEN>/getUpdates URL'sini ziyaret ederek yanıt içindeki 
 "chat": {"id": ...} kısmını kontrol edin.
  * Bilgileri girdikten sonra terminale go run main.go yazıp ilgili verileri çekmek istediğiniz web sitesini girin.
  * İYİ ÇALIŞMALAR - EMİRCAN ALTUNTAŞ
---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
EN
# WebScraperTool
- This is a tool I prepared for the Yıldız CTI team task.
- Task Description: Developing a Real-Time Status Tracking Bot for Websites in Go.
- Detailed Task Instruction: You are required to develop a real-time status tracking bot for a website in Go. The bot will check the target website every 5 minutes, reporting whether the site is accessible (up/down status) and the response time (in ms) via Telegram. If the response time exceeds 100 ms, the system will send a slowness warning to Telegram.

- ABOUT THE TOOL:
  * gopkg.in/tucnak/telebot.v2 is a library used for creating Telegram bots in Go. I used this library to direct the data I collected to Telegram.
  * Since the bot we wrote will send requests via Telegram, I used the "os" package in the code to prevent the bot token and chat ID from being exposed. By using 
 os.Getenv("TELEGRAMBOTTOKEN") and os.Getenv("TELEGRAMCHATID"), it retrieves the Telegram bot's token and chat ID from the environment variables defined in the system. This allows you to 
 use sensitive information securely without storing it directly in the code.
  * Other explanations of the code are available in main.go.

- DETAILED USAGE OF THE TOOL:
  * First, to set your Telegram bot token and chat ID:
 If you are using Linux/Mac, you need to enter these commands into the terminal: export TELEGRAMBOTTOKEN="YourToken", export TELEGRAMCHATID="YourChatID".
 If you are using Windows, you should enter the same commands in CMD with set at the beginning.
  * If you do not have a Telegram bot token, to create one: Contact BotFather and use the /newbot command to create a bot and save the provided token.
  * To learn your Chat ID: To find out your Chat ID, send a message to your bot and then visit the URL https://api.telegram.org/bot<TELEGRAMBOTTOKEN>/getUpdates to check the part of the 
 response containing "chat": {"id": ...}.
  * After entering the information, type go run main.go in the terminal and enter the website from which you want to retrieve data.
  * GOOD LUCK - EMİRCAN ALTUNTAŞ

 
  
