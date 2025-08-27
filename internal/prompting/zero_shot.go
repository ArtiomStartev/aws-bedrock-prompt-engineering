package prompting

import (
	"fmt"
	"log"
	"strings"

	"aws-bedrock-prompt-engineering/internal/bedrock"
)

type ZeroShotPrompt struct {
	client *bedrock.Client
	params bedrock.ModelParams
}

// NewZeroShotPrompt creates a zero-shot prompting instance.
// Zero-shot prompting presents a task to the model without examples or task-specific training,
// relying entirely on the model's general knowledge and capabilities.
func NewZeroShotPrompt(client *bedrock.Client) *ZeroShotPrompt {
	params := bedrock.GetDefaultClaudeParams()
	params.Temperature = 0.3 // Lower temperature for more focused responses

	return &ZeroShotPrompt{
		client: client,
		params: params,
	}
}

// ExecuteTextClassification demonstrates zero-shot text classification
func (z *ZeroShotPrompt) ExecuteTextClassification() error {
	prompt := `Classify the following text as either "positive", "negative", or "neutral":
	Text: "I absolutely love this new restaurant! The food was amazing and the service was excellent."
	Classification:`

	fmt.Println("🎯 Zero-Shot Prompting: Text Classification")
	fmt.Println("Prompt:", prompt)
	fmt.Println(strings.Repeat("-", 80))

	response, err := z.client.InvokeModel(prompt, z.params)
	if err != nil {
		return fmt.Errorf("failed to execute text classification: %w", err)
	}

	fmt.Printf("Response: %s\n\n", response.Completion)
	return nil
}

// ExecuteQuestionAnswering demonstrates zero-shot question answering
func (z *ZeroShotPrompt) ExecuteQuestionAnswering() error {
	prompt := `Answer the following question based on general knowledge:
	Question: What is the capital of Japan and what is it famous for?
	Answer:`

	fmt.Println("🎯 Zero-Shot Prompting: Question Answering")
	fmt.Println("Prompt:", prompt)
	fmt.Println(strings.Repeat("-", 80))

	response, err := z.client.InvokeModel(prompt, z.params)
	if err != nil {
		return fmt.Errorf("failed to execute question answering: %w", err)
	}

	fmt.Printf("Response: %s\n\n", response.Completion)
	return nil
}

// ExecuteLanguageTranslation demonstrates zero-shot translation
func (z *ZeroShotPrompt) ExecuteLanguageTranslation() error {
	prompt := `Translate the following English text to French:
	English: "Hello, how are you today? I hope you're having a wonderful day!"
	French:`

	fmt.Println("🎯 Zero-Shot Prompting: Language Translation")
	fmt.Println("Prompt:", prompt)
	fmt.Println(strings.Repeat("-", 80))

	response, err := z.client.InvokeModel(prompt, z.params)
	if err != nil {
		return fmt.Errorf("failed to execute language translation: %w", err)
	}

	fmt.Printf("Response: %s\n\n", response.Completion)
	return nil
}

// ExecuteCodeGeneration demonstrates zero-shot code generation
func (z *ZeroShotPrompt) ExecuteCodeGeneration() error {
	prompt := `Write a Python function that calculates the factorial of a number:
	Function name: calculate_factorial
	Input: integer n
	Output: factorial of n
	Include error handling for negative numbers.
	Code:`

	fmt.Println("🎯 Zero-Shot Prompting: Code Generation")
	fmt.Println("Prompt:", prompt)
	fmt.Println(strings.Repeat("-", 80))

	response, err := z.client.InvokeModel(prompt, z.params)
	if err != nil {
		return fmt.Errorf("failed to execute code generation: %w", err)
	}

	fmt.Printf("Response: %s\n\n", response.Completion)
	return nil
}

// RunAllExamples executes all zero-shot prompting examples
func (z *ZeroShotPrompt) RunAllExamples() {
	fmt.Println("=== ZERO-SHOT PROMPTING EXAMPLES ===\n")

	examples := []struct {
		name string
		fn   func() error
	}{
		{"Text Classification", z.ExecuteTextClassification},
		{"Question Answering", z.ExecuteQuestionAnswering},
		{"Language Translation", z.ExecuteLanguageTranslation},
		{"Code Generation", z.ExecuteCodeGeneration},
	}

	for _, example := range examples {
		if err := example.fn(); err != nil {
			log.Printf("Error in %s: %v", example.name, err)
		}
	}
}
