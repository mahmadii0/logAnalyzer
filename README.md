# Log Analyzer

Log Analyzer is a Go service for ingesting application logs, storing them in MySQL, grouping repeated errors, and triggering notifications when error patterns cross configured thresholds. It also supports AI-assisted summaries and confidence scoring through a local OpenAI-compatible endpoint.

## Features

- REST API built with Gin.
- MySQL persistence with GORM auto-migrations.
- Request validation for allowed service and level values.
- Error grouping by `service|level|message`.
- Baseline-based alerting for repeated errors.
- Optional notifications through Telegram, Slack, and WhatsApp.
- Optional AI summaries and confidence analysis using a local LLM endpoint.

## Tech Stack

- Go 1.25
- Gin
- GORM
- MySQL
- validator/v10
- godotenv

## Project Structure

The application code lives under `src/`:

- `cmd/` - application entrypoint
- `api/` - HTTP server, handlers, routes, validation, and response helpers
- `data/db/` - database connection and migration setup
- `models/` - GORM models and database methods
- `services/` - log registration, grouping, alerting, and AI analysis
- `env/` - environment variable helpers
- `constants/` and `pkg/` - shared constants and utility helpers

## Requirements

- Go 1.25 or later
- MySQL database
- Optional: Telegram bot credentials, Slack bot token, or WhatsApp API credentials
- Optional: a local OpenAI-compatible server on `http://localhost:1234/v1` for AI analysis

## Setup

1. Change into the Go module directory:

   ```bash
   cd src
   ```

2. Create a `.env` file in `src/` with at least the database connection string:

   ```env
   DATABASE_SOURCE=user:password@tcp(127.0.0.1:3306)/log_analyzer?parseTime=true&charset=utf8mb4&loc=Local
   ```

3. Add optional notification and AI settings if you want those features enabled:

   ```env
   TELEGRAM_STATUS=0
   SLACK_STATUS=0
   WHATSAPP_STATUS=0

   TELEGRAM_BOT_TOKEN=
   CHAT_ID=

   SLACK_BOT_TOKEN=
   CHANNEL_ID=

   WHATSAPP_BOT_TOKEN=
   PHONE_ID=
   TO=

   SUMMARY_DEVELOPER_MESSAGE=
   CONFIDENCE_DEVELOPER_MESSAGE=
   ```

4. Download dependencies:

   ```bash
   go mod download
   ```

5. Run the service:

   ```bash
   go run ./cmd
   ```

The server starts on port `8000`.

## API

### POST `/api/v1/logs`

Registers a log entry and stores it in the database.

Request body:

```json
{
  "service": "payment",
  "level": "TIMEOUT",
  "message": "Payment gateway request timed out"
}
```

Validation rules currently enforced by the application:

- `service` must be one of: `payment`, `order`, `authorazation`
- `level` must be one of: `TIMEOUT`, `BAD REQUEST`
- `message` is required and must be at most 380 characters

## Behavior

When a log is stored, the service:

- creates or updates an error group keyed by `service|level|message`
- counts recent occurrences over the last 10 minutes
- computes a baseline from earlier 10-minute windows
- sends alerts when the configured threshold is reached
- optionally adds an AI-generated summary and confidence rating

## Notes

- The application writes logs to `src/logs/app.log`.
- Database tables are auto-migrated on startup.
- AI analysis expects a local LM Studio-style endpoint at `http://localhost:1234/v1` unless you change the implementation.

## License

No license file is included yet. Add one before publishing the repository publicly.