package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	// sndr "github.com/mghgm/camelnews/sender"
	"github.com/mghgm/camelnews/collector"
)

var (
    runCmd = flag.NewFlagSet("run", flag.ExitOnError)
    configPath = runCmd.String("config", "~/.config/rss-bot/config.yaml", "Bot configuration")
    
    reconfigureCmd = flag.NewFlagSet("reconfigure", flag.ExitOnError)
)

func main() {
    flag.Parse() 

    if len(os.Args) < 2 {
        log.Fatal("Invalid cmd format!")
    }
    
    switch os.Args[1] {
    case "run":
        runCmd.Parse(os.Args[2:])
        log.Println("run")
        log.Printf("config: %v\n", *configPath)
        hckrNews := collector.NewRSSAgency("Rss Newest", "Programming", "https://hnrss.org/newest")
        news, err := hckrNews.Collect()
        if err != nil {
            log.Fatal(err)
        }
        
        for _, n := range news {
            fmt.Printf("title: %s link: %s\n", n.Title, n.Link)
        }

    case "reconfigure":
        reconfigureCmd.Parse(os.Args[2:])
        log.Println("reconfigure")
    
    default:
        log.Fatal("Invalid cmd format!")
    }
}
