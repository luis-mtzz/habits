// internal/ports/api/discogs.go

// Package api defines the interfaces (ports) for interacting with external APIs.
// This file specifically defines the interface for the Discogs API and related data structures.

package api

import "context"

type DiscogsPort interface {
	GetUserCollection(ctx context.Context, username string) ([]Album, error)
}

type Album struct {
	Title  string   `json:"title"`
	Artist string   `json:"artist"`
	Genre  []string `json:"genre"`
}
