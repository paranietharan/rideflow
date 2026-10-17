package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

// Client wraps the MongoDB client and active database.
type Client struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// Config holds settings for MongoDB client initialization.
type Config struct {
	URI             string
	Database        string
	MinPoolSize     uint64
	MaxPoolSize     uint64
	MaxConnIdleTime time.Duration
	ConnectTimeout  time.Duration
}

// DefaultConfig returns sane production-friendly defaults.
func DefaultConfig(uri, database string) Config {
	return Config{
		URI:             uri,
		Database:        database,
		MinPoolSize:     5,
		MaxPoolSize:     100,
		MaxConnIdleTime: 5 * time.Minute,
		ConnectTimeout:  10 * time.Second,
	}
}

// NewClient establishes a MongoDB connection with connection pooling and pings the server.
func NewClient(ctx context.Context, cfg Config) (*Client, error) {
	opts := options.Client().
		ApplyURI(cfg.URI).
		SetMinPoolSize(cfg.MinPoolSize).
		SetMaxPoolSize(cfg.MaxPoolSize).
		SetMaxConnIdleTime(cfg.MaxConnIdleTime).
		SetConnectTimeout(cfg.ConnectTimeout)

	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to create mongodb client: %w", err)
	}

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := client.Ping(pingCtx, readpref.Primary()); err != nil {
		_ = client.Disconnect(ctx)
		return nil, fmt.Errorf("failed to ping mongodb: %w", err)
	}

	return &Client{
		Client:   client,
		Database: client.Database(cfg.Database),
	}, nil
}

// Ping checks if the MongoDB connection is alive.
func (c *Client) Ping(ctx context.Context) error {
	pingCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return c.Client.Ping(pingCtx, readpref.Primary())
}

// Close gracefully disconnects the MongoDB client.
func (c *Client) Close(ctx context.Context) error {
	return c.Client.Disconnect(ctx)
}
