package bedrock

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

type Client struct {
	client *bedrockruntime.Client
	ctx    context.Context
}

type ModelParams struct {
	ModelID     string  `json:"model_id"`    // e.g., "anthropic.claude-v2:1"
	Temperature float64 `json:"temperature"` // creativity of the model's output (0.0 to 1.0)
	TopP        float64 `json:"top_p"`       // consider a broad range of possible words (0.0 to 1.0)
	TopK        int     `json:"top_k"`       // limits the number of probable words
	MaxTokens   int     `json:"max_tokens"`  // maximum number of tokens to generate
}

type ModelResponse struct {
	Type       string `json:"type"`
	Completion string `json:"completion"`
	StopReason string `json:"stop_reason"`
	Stop       string `json:"stop"`
}

func NewClient() (*Client, error) {
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	return &Client{
		client: bedrockruntime.NewFromConfig(cfg),
		ctx:    ctx,
	}, nil
}

func (c *Client) InvokeModel(prompt string, params ModelParams) (*ModelResponse, error) {
	requestBody := map[string]any{
		"prompt":               fmt.Sprintf("\n\nHuman: %s\n\nAssistant:", prompt),
		"temperature":          params.Temperature,
		"top_p":                params.TopP,
		"top_k":                params.TopK,
		"max_tokens_to_sample": params.MaxTokens,
	}

	bodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	input := &bedrockruntime.InvokeModelInput{
		ModelId:     &params.ModelID,
		Body:        bodyBytes,
		ContentType: aws.String("application/json"),
	}

	resp, err := c.client.InvokeModel(c.ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to invoke model: %w", err)
	}

	var modelResp ModelResponse
	if err := json.Unmarshal(resp.Body, &modelResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &modelResp, nil
}

// GetDefaultClaudeParams returns default parameters for Claude v2 model
func GetDefaultClaudeParams() ModelParams {
	return ModelParams{
		ModelID:     os.Getenv("MODEL_ID"),
		Temperature: 0.7,
		TopP:        1.0,
		TopK:        500,
		MaxTokens:   500,
	}
}
