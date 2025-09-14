# Zeni - RSS Feed Aggregator

A modern RSS feed aggregator built with Go that allows users to subscribe to RSS feeds, manage their subscriptions, and view posts from their followed feeds.

## Project Overview

Zeni is a backend service that provides RSS feed aggregation functionality through a RESTful API. Users can create accounts, add RSS feeds, follow feeds, and receive posts from their subscribed feeds. The application automatically scrapes RSS feeds in the background to keep posts up to date.

## Features

- **User Management**: Create user accounts with API key authentication
- **Feed Management**: Add and manage RSS feeds
- **Feed Following**: Subscribe to feeds and manage subscriptions
- **Post Aggregation**: Automatically fetch and store posts from RSS feeds
- **RESTful API**: Clean REST API for all operations
- **Background Scraping**: Automated RSS feed scraping with configurable intervals
- **PostgreSQL Database**: Robust data persistence with SQLC for type-safe queries

## Tech Stack

- **Backend**: Go 1.21.1
- **Database**: PostgreSQL
- **ORM**: SQLC for type-safe database queries
- **HTTP Router**: Chi router
- **Authentication**: API Key-based authentication
- **RSS Parsing**: Custom XML parsing for RSS feeds

## Project Structure

```
zeni/
├── main.go                 # Application entry point
├── models.go              # Data models and conversion functions
├── middleware_auth.go     # Authentication middleware
├── handler_*.go          # HTTP request handlers
├── scraper.go            # RSS feed scraping logic
├── rss.go                # RSS feed parsing
├── internal/
│   ├── auth/             # Authentication utilities
│   └── database/         # Generated database queries (SQLC)
├── sql/
│   ├── schema/           # Database migration files
│   └── queries/          # SQL query definitions
└── vendor/               # Go dependencies
```

## API Endpoints

### Authentication
All protected endpoints require an `Authorization` header with format: `APIKey <your-api-key>`

### User Management
- `POST /v1/users` - Create a new user
- `GET /v1/users` - Get current user (requires auth)

### Feed Management
- `POST /v1/feeds` - Create a new RSS feed (requires auth)
- `GET /v1/feeds` - Get all available feeds

### Feed Following
- `POST /v1/feed_follows` - Follow a feed (requires auth)
- `GET /v1/feed_follows` - Get user's followed feeds (requires auth)
- `DELETE /v1/feed_follows/{feedFollowID}` - Unfollow a feed (requires auth)

### Posts
- `GET /v1/posts` - Get posts from user's followed feeds (requires auth)

### Health Check
- `GET /v1/ready` - Health check endpoint
- `GET /v1/err` - Error testing endpoint

## Database Schema

The application uses PostgreSQL with the following main tables:

- **users**: User accounts with API keys
- **feeds**: RSS feed definitions
- **feed_follows**: User subscriptions to feeds
- **posts**: Individual posts from RSS feeds

## Setup and Installation

### Prerequisites
- Go 1.21.1 or later
- PostgreSQL database
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd zeni
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   Create a `.env` file in the project root:
   ```env
   PORT=8080
   DB_URL=postgres://username:password@localhost/dbname?sslmode=disable
   ```

4. **Set up the database**
   - Create a PostgreSQL database
   - Run the database migrations in `sql/schema/` directory
   - Generate database queries using SQLC:
     ```bash
     sqlc generate
     ```

5. **Run the application**
   ```bash
   go run main.go
   ```

The server will start on the port specified in your `PORT` environment variable (default: 8080).

## Usage Examples

### Create a User
```bash
curl -X POST http://localhost:8080/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe"}'
```

### Create a Feed
```bash
curl -X POST http://localhost:8080/v1/feeds \
  -H "Content-Type: application/json" \
  -H "Authorization: APIKey <your-api-key>" \
  -d '{"name": "Tech News", "url": "https://example.com/rss"}'
```

### Follow a Feed
```bash
curl -X POST http://localhost:8080/v1/feed_follows \
  -H "Content-Type: application/json" \
  -H "Authorization: APIKey <your-api-key>" \
  -d '{"feed_id": "<feed-uuid>"}'
```

### Get Posts
```bash
curl -X GET http://localhost:8080/v1/posts \
  -H "Authorization: APIKey <your-api-key>"
```

## Background Scraping

The application includes a background scraper that automatically fetches new posts from RSS feeds. The scraper:

- Runs every minute by default
- Uses 10 concurrent goroutines for efficient scraping
- Marks feeds as fetched to prevent duplicate processing
- Handles RSS parsing errors gracefully
- Stores new posts in the database

## Development

### Database Migrations
Database schema changes should be added as migration files in the `sql/schema/` directory following the goose migration format.

### SQLC Queries
Add new database queries in the `sql/queries/` directory and regenerate the Go code:
```bash
sqlc generate
```

### Adding New Endpoints
1. Create handler functions in appropriate `handler_*.go` files
2. Add routes in `main.go`