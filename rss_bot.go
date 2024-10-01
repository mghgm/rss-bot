package main

import (
    "fmt"
    "log"

    "github.com/spf13/cobra"
)

var (
    configPath string 
)

var rootCmd = &cobra.Command{
    Use:   "rss-bot",
    Short: "A simple RSS bot",
}

var reconfigureCmd = &cobra.Command{
    Use:   "reconfigure",
    Short: "Reconfigure the RSS bot",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Reconfiguring the RSS bot...")
    },
}

var runCmd = &cobra.Command{
    Use:   "run",
    Short: "Run the RSS bot",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Bot is using config path %s\n", configPath)
    },
}

func init() {
    runCmd.Flags().StringVarP(&configPath, "config", "c", "~/.config/rss-bot/config.yaml", "Bot configuration")
    
    rootCmd.AddCommand(reconfigureCmd)
    rootCmd.AddCommand(runCmd)
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        log.Fatal(err)
    }
}
