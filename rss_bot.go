package main

import (
    "log"
    "flag"
    "os"

    sndr "github.com/mghgm/rss-bot/sender"
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
    case "reconfigure":
        reconfigureCmd.Parse(os.Args[2:])
        log.Println("reconfigure")
    
    default:
        log.Fatal("Invalid cmd format!")
    }
}
