# Telegram Channel Exporter

A Go application that fetches posts from public Telegram channels and exports them as structured JSON files with automatic GitHub Actions integration.

## ✨ Features

- 🚀 **No API Required** - Uses Telegram's public RSS feed
- 📝 **Complete Post Data** - Message content, dates, views, media, hashtags, mentions
- 🖼️ **Media Extraction** - Photos, videos, and uploaded files with URLs
- 🎨 **Markdown Preserved** - HTML formatting, links, and emoji support
- 🕐 **Real Timestamps** - Accurate dates with Unix timestamps
- 🔄 **Auto-Export** - GitHub Actions runs every 30 minutes
- 🌿 **Clean Repository** - Separate export branch, main stays clean
- 🛡️ **Rate Limit Protection** - Random delays and user agents

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

## 📊 Output Structure

Each channel is exported to `export/{channel_name}.json` with comprehensive data:

```json
{
  "info": {
    "id": 0,
    "title": "Channel Title",
    "username": "channelname",
    "photo_url": "https://cdn.telesco.pe/file/..."
  },
  "posts": [
    {
      "id": 12345,
      "message": "Post content with <a href=\"links\">HTML formatting</a>",
      "date": "2026-05-03T09:30:00Z",
      "edited": false,
      "views": 1500,
      "forwards": 25,
      "replies": 10,
      "sender_name": "Sender Name",
      "media": [
        {
          "type": "photo",
          "url": "https://cdn.telesco.pe/file/...",
          "width": 800,
          "height": 600
        },
        {
          "type": "video", 
          "url": "https://cdn.telesco.pe/file/..."
        }
      ],
      "hashtags": ["#example", "#telegram"],
      "mentions": ["@user"],
      "links": ["https://example.com"]
    }
  ],
  "last_updated": 1777791234
}
```

## 🔄 Auto-Export with GitHub Actions

The project includes automated GitHub Actions that:

- **Run every 30 minutes** - Automatic channel updates
- **Separate branches** - Clean main branch, exports in `export` branch
- **Rate limiting** - Random delays and user agents
- **Error handling** - Continues even if some channels fail
- **Detailed logging** - Summary reports with file statistics

### Branch Strategy
- **main**: Source code and configuration only
- **export**: JSON export files (auto-updated)

## 🛠️ Local Development

### Manual Run
```bash
go run .
```

### Docker Support
```bash
# Build and run
docker-compose up telefeed

# Auto-run every 30 minutes
docker-compose --profile scheduled up telefeed-scheduled
```

## ⚙️ Configuration

Edit `config.json` to specify channels:

```json
{
  "channels": [
    "ircfspace",
    "vahidonline", 
    "your_channel"
  ]
}
```

## 🚀 Deployment

### GitHub Actions (Recommended)
1. Push code to GitHub
2. Enable Actions in repository settings
3. Automatic exports every 30 minutes

### Manual Deployment
```bash
# Build binary
go build -o telefeed

# Run with cron
*/30 * * * * cd /path/to/telefeed && ./telefeed
```

## 📋 Requirements

- Go 1.21 or higher
- Internet connection
- Public Telegram channels only

## ⚠️ Limitations

- **Public channels only** - No private channel support
- **Recent posts** - Limited by Telegram's public feed
- **Rate limits** - Built-in protection but may need adjustments
- **HTML dependency** - May break if Telegram changes web structure
