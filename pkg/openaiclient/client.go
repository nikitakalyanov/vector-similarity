package openaiclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"time"
)

type Client struct {
	baseURL string
	client  http.Client
}

func NewClient() *Client {
	return &Client{
		baseURL: "https://api.openai.com/",
		client: http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

type DataItem struct {
	Object    string `json:"object"`
	Index     int
	Embedding []float32
}

type EmbeddingsResponse struct {
	Object string     `json:"object"`
	Data   []DataItem `json:"data"`
}

type EmbeddingsRequest struct {
	Input string `json:"input"`
	Model string `json:"model"`
}

func (c *Client) GetEmbeddings(ctx context.Context, input string) ([]float32, error) {
	body := EmbeddingsRequest{
		Input: input,
		Model: "text-embedding-ada-002",
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.baseURL+"v1/embeddings", bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("OPENAI_APIKEY"))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("unexpected http code")
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res EmbeddingsResponse
	err = json.Unmarshal(respBytes, &res)
	if err != nil {
		return nil, err
	}
	return res.Data[0].Embedding, nil
}
