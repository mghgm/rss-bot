package sender

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

const (
    telegramSendMessageAPI = "https://api.telegram.org/bot%s/sendMessage"
    proxyAddr = "localhost:12345"
)

type TelegramSender struct {
    botToken string
    proxy *http.Transport 
}

func NewTelegramSender(botToken string, proxy *http.Transport) *TelegramSender {
    return &TelegramSender{
        botToken: botToken,
        proxy: proxy,
    }
}

type telegramMessage struct {
    ChatID int64  `json:"chat_id"`
    Text   string `json:"text"`
}

func createTelegramMessage(chatID int64, post Post) *telegramMessage {
    if len(post.Title) == 0 {
        err := errors.New(fmt.Sprintf("Unable to create telegramMessage from %v", post))
        log.Println("Error: ", err)
        return nil
    }

    return &telegramMessage{
        Text: fmt.Sprintf("%s", post.Title),
        ChatID: chatID,
    }
}

func (ts TelegramSender) Send(chatID int64, post Post, useProxy bool) error {
    url := fmt.Sprintf(telegramSendMessageAPI, ts.botToken) 
    
    client := &http.Client{}
    if (useProxy) {
        if (ts.proxy == nil) {
            err := errors.New(fmt.Sprintf("No proxy is set"))
            log.Println("Error: ", err)
            return fmt.Errorf("No proxy set for %v", ts)
        }
        client.Transport = ts.proxy
    }
    
    telegramMsg := createTelegramMessage(chatID, post)
    jsonData, err := json.Marshal(telegramMsg)
    if err != nil {
        return err
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
    if err != nil {
        return err;
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("status code %v", resp.StatusCode)
    }

    return nil
}
