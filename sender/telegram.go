package sender

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

    "golang.org/x/net/proxy"

	"github.com/mghgm/camelnews/collector"
	"github.com/mghgm/camelnews/config"
)

const (
    TelegramSendMessageAPI = "https://api.telegram.org/bot%s/sendMessage"
    ProxyAddr = "localhost:12345"

    SenderBufferSize = 5
)

var (
    _ Sender = &TelegramSender{}
)

type TelegramSender struct {
    botToken string
    proxy *http.Transport 
}

func NewTelegramSenderFromConfig(cfg config.SendersConfig) *TelegramSender {
    var socks5Proxy *http.Transport
    if (cfg.Proxy) {
        dialer, err := proxy.SOCKS5("tcp", ProxyAddr, nil, proxy.Direct)
	    if err != nil {
	    	log.Fatalf("Failed to create SOCKS5 dialer: %v", err)
            return nil
	    }

	    socks5Proxy = &http.Transport{
	    	DialContext: dialer.(proxy.ContextDialer).DialContext,
	    }
    }
    
    return NewTelegramSender(cfg.Token, socks5Proxy)
}

func NewTelegramSender(botToken string, socks5Proxy *http.Transport) *TelegramSender {
    return &TelegramSender{
        botToken: botToken,
        proxy: socks5Proxy,
    }
}

type telegramMessage struct {
    ChatID int64  `json:"chat_id"`
    Text   string `json:"text"`
    ParseMode string `json:"parse_mode"`
}

func createPostFromNews(news collector.News) Post {
    return Post{
        Title: news.Title,
        Desc: news.Desc,
        Tags: []string{"test1", "test2"},
        Ref: news.Link,
    }
}

func formatPost(post Post) string {
	message := fmt.Sprintf(
		"<b>%s</b> \n<a href='%s'>Reference</a>",
		post.Title,
		post.Ref,                  
	)
	return message
}


func (ts *TelegramSender) Start() chan collector.News {
    updates := make(chan collector.News, SenderBufferSize)

    go func() {
        for {
            select {
            case news := <- updates:
                post := createPostFromNews(news)
                
                go func(post Post) {
                    for i := 0; i < 5; i++ {
                        err := ts.send(1917075603, post)
                        if err == nil {
                            break
                        }
                        log.Println(err)
                    }
                } (post)
                
            }
        }
    }()

    return updates
}

func createTelegramMessage(chatID int64, post Post) *telegramMessage {
    if len(post.Title) == 0 {
        err := errors.New(fmt.Sprintf("Unable to create telegramMessage from %v", post))
        log.Println("Error: ", err)
        return nil
    }

    return &telegramMessage{
        Text: formatPost(post),
        ChatID: chatID,
        ParseMode: "HTML",
    }
}

func (ts *TelegramSender) send(chatID int64, post Post) error {
    url := fmt.Sprintf(TelegramSendMessageAPI, ts.botToken) 
    
    client := &http.Client{}
    // Set Socks5 Proxy
    client.Transport = ts.proxy
  
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
