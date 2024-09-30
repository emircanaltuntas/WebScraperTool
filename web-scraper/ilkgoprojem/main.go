package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"gopkg.in/tucnak/telebot.v2"
)

func main() {

	var websiteurl string //İstediğiniz bir web sitesinin URL'sini girmenizi sağlayan fonksiyon.
	fmt.Print("Bilgilerini çekmek istediğiniz web sitesinin URL'sini girin:")
	_, err := fmt.Scan(&websiteurl)
	if err != nil {
		log.Fatal("Geçersiz URL:", err)
	}

	botToken := os.Getenv("TELEGRAMBOTTOKEN") //Bilgilerin gelmesini istediğiniz Telegram hesap ID'si ve Telegram Bot Tokeni çeker.
	chatIDStr := os.Getenv("TELEGRAMCHATID")
	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		log.Fatal("Geçersiz Chat ID:", err)
	}
	if botToken == "" {
		log.Fatal("Bot Token veya Chat ID tanımlı değil.")
	}

	bot, err := telebot.NewBot(telebot.Settings{Token: botToken}) //Telegram bot API'mizin çalıştığı fonksiyon.
	if err != nil {
		log.Fatal("Bot çalıştırılamadı:", err)
	}

	for { //5 dakikada bir web sitesini kontrol etmeye yarayan fonksiyon.
		checkWebsite(bot, chatID, websiteurl)
		time.Sleep(5 * time.Minute)
	}
}

func checkWebsite(bot *telebot.Bot, chatID int64, websiteurl string) { //Web sitesine istek atıp durumunu kontrol eden fonksiyon.
	startTime := time.Now()
	resp, err := http.Get(websiteurl)
	if err != nil {
		sendMessage(bot, chatID, fmt.Sprintf("Web sitesine ulaşılamıyor: %s", err))
		return
	}

	defer resp.Body.Close() //Yanıtı bitirmeye yarayan fonksiyon.

	duration := time.Since(startTime).Milliseconds() //Yanıtın süresini hesaplayan fonksiyon.

	if resp.StatusCode == http.StatusOK { //Site açıksa ya da kapalıysa istediğimiz bilgileri veren fonksiyon.
		sendMessage(bot, chatID, fmt.Sprintf("Web Sitesi Erişilebilir! Yanıt süresi: %d ms", duration))
		if duration > 100 {
			sendMessage(bot, chatID, fmt.Sprintf("Web Sitesi Yavaş!: %d ms", duration))
		}
	} else {
		sendMessage(bot, chatID, fmt.Sprintf("Web Sitesine Erişilemedi!: %d", resp.StatusCode))
	}
}

func sendMessage(bot *telebot.Bot, chatID int64, message string) { //Son olarak çektiğimiz bilgileri seçtiğimiz Telegram hesabına yollar.
	_, err := bot.Send(&telebot.Chat{ID: chatID}, message)
	if err != nil {
		log.Printf("Mesaj Gönderilemedi: %s", err)
	}
}

// OLUŞTURAN = Emircan Altuntaş - YILDIZ CTI - github.com/emircanaltuntas
