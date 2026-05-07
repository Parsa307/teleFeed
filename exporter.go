package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func exportChannelData(channelData *ChannelData, suffix string) error {
	// Create export directory if it doesn't exist
	exportDir := "export"
	if err := os.MkdirAll(exportDir, 0755); err != nil {
		return fmt.Errorf("failed to create export directory: %v", err)
	}

	// Generate filename (lowercase channel name)
	filename := strings.ToLower(channelData.Info.Username)
	if filename == "" {
		filename = strings.ToLower(channelData.Info.Title)
	}
	
	// Clean filename (remove special characters)
	filename = regexp.MustCompile(`[^a-zA-Z0-9_-]`).ReplaceAllString(filename, "_")
	
	// Add suffix if provided
	if suffix != "" {
		filename = filename + "_" + suffix
	}
	filename = filename + ".json"

	filePath := filepath.Join(exportDir, filename)

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
