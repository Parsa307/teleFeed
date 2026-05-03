# Telegram Channel Exporter

A Go application that fetches posts from public Telegram channels and exports them as structured JSON files.

## Features

- Fetches posts from public Telegram channels without requiring API credentials
- Extracts comprehensive post information:
  - Sender information (ID and name)
  - Date and edit status
  - Message content and captions
  - Media attachments (photos, videos, documents)
  - View counts, forwards, and replies
  - Hashtags, mentions, and links
- Exports channel information (title, username, photo)
- Creates clean JSON structure with proper formatting
- Saves exports in `export/` folder with lowercase filenames

## Installation

1. Clone or download this project
2. Install Go (if not already installed)
3. Run the application:

```bash
go run .
```

Or build it:

```bash
go build -o telefeed
./telefeed
```

## 📁 Project Structure

```
telefeed/
├── config.json          # Channel configuration
├── export/               # Local export folder (gitignored)
├── .github/workflows/    # GitHub Actions
├── main.go              # Main application
├── fetcher_colly.go     # Telegram data fetcher
├── exporter.go          # JSON export functionality
└── .gitignore           # Git ignore rules
```

## 🌿 Git Branch Strategy

- **main**: Source code and configuration only
- **export**: Contains JSON export files (auto-updated)

The `export/` folder is gitignored to keep the main branch clean. Export files are automatically pushed to the `export` branch via GitHub Actions.

## Configuration

Edit `config.json` to specify which channels to export:

```json
{
  "channels": [
    "ircfspace",
    "vahidonline",
    "your_channel_name"
  ]
}
```

**Note:** Use only the username part without the `@` symbol.

## Output Structure

Each channel is exported to `export/{channel_name}.json` with the following structure:

```json
{
  "info": {
    "id": 123456789,
    "title": "Channel Title",
    "username": "channelname",
    "photo_url": "https://..."
  },
  "posts": [
    {
      "id": 12345,
      "message": "Post content",
      "caption": "Media caption",
      "date": "2026-05-03T09:30:00Z",
      "edited": false,
      "edit_date": null,
      "views": 1500,
      "forwards": 25,
      "replies": 10,
      "sender_id": 987654321,
      "sender_name": "Sender Name",
      "media": [
        {
          "type": "photo",
          "url": "https://...",
          "width": 800,
          "height": 600
        }
      ],
      "hashtags": ["#example", "#telegram"],
      "mentions": ["@user"],
      "links": ["https://example.com"]
    }
  ]
}
```

## How It Works

The application uses Telegram's public RSS feed feature:
- Accesses `https://t.me/s/{channel_name}` for each channel
- Parses the embedded JSON data from the webpage
- Extracts posts and channel information
- Converts to structured JSON format
- Saves to the export folder

## Rate Limiting

The application includes a 2-second delay between channel requests to avoid rate limiting.

## Requirements

- Go 1.19 or higher
- Internet connection
- Public Telegram channels (no private channels supported)

## Limitations

- Only works with public channels
- Limited to recent posts (Telegram's public feed restrictions)
- May not work if Telegram changes their web interface structure
