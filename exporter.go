package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func exportChannelData(channelData *ChannelData) error {
	// Generate filename (lowercase channel name)
	filename := strings.ToLower(channelData.Info.Username)
	if filename == "" {
		filename = strings.ToLower(channelData.Info.Title)
	}
	
	// Clean filename (remove special characters)
	filename = regexp.MustCompile(`[^a-zA-Z0-9_-]`).ReplaceAllString(filename, "_")
	filename = filename + ".json"

	// Export directly to root directory (no export folder)
	filePath := filename

	// Convert to export format (handles empty caption properly)
	exportData := toExportChannelData(*channelData)

	// Convert to JSON with indentation
	jsonData, err := json.MarshalIndent(exportData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	// Check if file already exists to prevent duplicates
	if _, err := os.Stat(filePath); err == nil {
		fmt.Printf("File %s already exists, skipping export for channel '%s'\n", filePath, channelData.Info.Title)
		return nil
	}

	// Write to file
	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	fmt.Printf("Exported channel '%s' to %s\n", channelData.Info.Title, filePath)
	return nil
}
