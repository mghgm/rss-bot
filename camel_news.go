package main

import (
	"flag"
	"time"
	"log"

	"github.com/mghgm/camelnews/collector"
	"github.com/mghgm/camelnews/config"
	"github.com/mghgm/camelnews/sender"
)

var (
    configPath = flag.String("config", "~/.config/rss-bot/config.yaml", "Bot configuration")
)


func multiplexCollectorsToSenders(inputs []chan collector.News, outputs []chan collector.News) {
    for _, inCh := range inputs {
        go func(ch <- chan collector.News) {
            for msg := range ch {
                for _, outCh := range outputs {
                    outCh <- msg
                }
            }
        } (inCh)
    }
}

func main() {
    flag.Parse() 

    c, err := config.ReadConfig(*configPath)
    if err != nil {
        log.Fatal(err)
    }
    
    collectors := make([]collector.Collector, 0)
    senders := make([]sender.Sender, 0)
    
    for _, cfg := range c.Collectors {
        switch (cfg.Type) {
        case "rss":
            collectors = append(collectors, collector.NewRSSAgencyCollectorFromConfig(cfg))
        default:
            log.Println("Error: Invalid collector type")
        }
    }

    for _, cfg := range c.Senders {
        switch (cfg.Type) {
        case "telegrambot":
            senders = append(senders, sender.NewTelegramSenderFromConfig(cfg))
        default:
            log.Println("Error: Invalid sender type")
        }
    }

    inputs := make([]chan collector.News, 0)
    for _, c := range collectors {
        inputs = append(inputs, c.Start())
    }

    outputs := make([]chan collector.News, 0)
    for _, s := range senders {
        outputs = append(outputs, s.Start())
    }

    multiplexCollectorsToSenders(inputs, outputs) 
    time.Sleep(100 * time.Second)
}
