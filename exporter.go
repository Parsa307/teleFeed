package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func exportChannelData(channelData *ChannelData) error {
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
	filename = filename + ".json"

	filePath := filepath.Join(exportDir, filename)

	// Convert to export format (handles empty caption properly)
	exportData := toExportChannelData(*channelData)

	// Convert to JSON with indentation
	jsonData, err := json.MarshalIndent(exportData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	// Write to file
	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	fmt.Printf("Exported channel '%s' to %s\n", channelData.Info.Title, filePath)
	return nil
}
