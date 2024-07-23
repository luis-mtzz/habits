// internal/adapters/api/discogs/client.go

// Package discogs provides an implementation of the Discogs API client.
// This file contains the Client struct and methods to interact with the Discogs API,
// serving as an adapter for the DiscogsPort interface.

package discogs

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/luis-mtzz/habits/internal/ports/api"
)

type Client struct {
	apiKey  string
	baseURL string
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:  apiKey,
		baseURL: "https://api.discogs.com/",
	}
}

func (c *Client) GetUserCollection(ctx context.Context, username string) ([]api.Album, error) {
	url := fmt.Sprintf("%s/users/%s/collection/folders/0/releases", c.baseURL, username)
	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", fmt.Sprintf("Discogs token=%s", c.apiKey))

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var result struct {
		Releases []struct {
			Basic_information struct {
				Title   string `json:"title"`
				Artists []struct {
					Name string `json:"name"`
				} `json:"artists"`
				Genres []string `json:"genres"`
			} `json:"basic_information"`
		} `json:"releases"`
	}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	var albums []api.Album
	for _, r := range result.Releases {
		albums = append(albums, api.Album{
			Title:  r.Basic_information.Title,
			Artist: r.Basic_information.Artists[0].Name,
			Genre:  r.Basic_information.Genres,
		})
	}

	return albums, nil
}
