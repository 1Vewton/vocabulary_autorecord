package llm

import (
	"context"
	"fmt"

	"github.com/1Vewton/vocabulary_autorecord/data_management/basic_config"
	"github.com/voocel/litellm"
)

// Request llm
func Request(prompt string, response_chan chan string, err_chan chan error) {
	response := ""
	fmt.Println("Start requesting...")
	ctx := context.Background()
	// Create client
	client, err := litellm.NewWithProvider(
		basic_config.BasicConfig.LLMProvider,
		litellm.ProviderConfig{
			APIKey:  basic_config.BasicConfig.LLMApiKey,
			BaseURL: basic_config.BasicConfig.LLMBaseURL,
		},
	)
	if err != nil {
		err_chan <- err
		return
	}
	// Start requesting
	resp, err := client.Chat(
		ctx,
		&litellm.Request{
			Model: basic_config.BasicConfig.LLMModelName,
			Messages: []litellm.Message{
				litellm.UserMessage(prompt),
			},
		},
	)
	if err != nil {
		err_chan <- err
		return
	}
	response = resp.Content
	response_chan <- response
}
