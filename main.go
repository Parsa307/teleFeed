package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("Telegram Channel Exporter")
	fmt.Println("=========================")

	// Load configuration
	config, err := LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if len(config.Channels) == 0 {
		log.Fatal("No channels specified in config.json")
	}

	fmt.Printf("Found %d channels to process\n", len(config.Channels))

	// Process each channel
	for i, channel := range config.Channels {
		fmt.Printf("\n[%d/%d] Processing channel: %s\n", i+1, len(config.Channels), channel)
		
		// Fetch channel data
		channelData, err := fetchChannelDataWithColly(channel)
		if err != nil {
			log.Printf("Error fetching channel %s: %v", channel, err)
			continue
		}

		fmt.Printf("  - Channel: %s (@%s)\n", channelData.Info.Title, channelData.Info.Username)
		fmt.Printf("  - Posts found: %d\n", len(channelData.Posts))

		// Export to JSON
		if err := exportChannelData(channelData); err != nil {
			log.Printf("Error exporting channel %s: %v", channel, err)
			continue
		}

		// Add delay to avoid rate limiting
		if i < len(config.Channels)-1 {
			fmt.Println("  - Waiting 5 seconds before next channel...")
			time.Sleep(5 * time.Second)
		}
	}

	fmt.Println("\nExport completed! Check the 'export' folder for JSON files.")
}
