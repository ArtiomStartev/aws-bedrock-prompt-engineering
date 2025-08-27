package prompting

import (
	"fmt"
	"log"
	"strings"

	"aws-bedrock-prompt-engineering/internal/bedrock"
)

type FewShotPrompt struct {
	client *bedrock.Client
	params bedrock.ModelParams
}

// NewFewShotPrompt creates a few-shot prompting instance.
// Few-shot prompting provides the model with several task examples to guide its output.
// Providing only one example is called one-shot prompting.
func NewFewShotPrompt(client *bedrock.Client) *FewShotPrompt {
	params := bedrock.GetDefaultClaudeParams()

	params.Temperature = 0.5 // Moderate temperature for balanced creativity
	params.MaxTokens = 800   // More tokens for detailed responses

	return &FewShotPrompt{
		client: client,
		params: params,
	}
}

// ExecuteSentimentAnalysis demonstrates few-shot sentiment analysis
func (f *FewShotPrompt) ExecuteSentimentAnalysis() error {
	prompt := `Analyze the sentiment of the following texts. Classify each as "positive", "negative", or "neutral".

	Examples:
	Text: "I love this product! It's amazing!"
	Sentiment: positive
	
	Text: "This is terrible. I hate it."
	Sentiment: negative
	
	Text: "The weather is okay today."
	Sentiment: neutral
	
	Text: "The customer service was outstanding and they resolved my issue quickly."
	Sentiment: positive
	
	Now classify this text:
	Text: "The movie was disappointing. The plot was confusing and the acting was mediocre."
	Sentiment:`

	fmt.Println("🎯 Few-Shot Prompting: Sentiment Analysis")
	fmt.Println("Prompt:", prompt)
	fmt.Println(strings.Repeat("-", 80))

	response, err := f.client.InvokeModel(prompt, f.params)
	if err != nil {
		return fmt.Errorf("failed to execute sentiment analysis: %w", err)
	}

	fmt.Printf("Response: %s\n\n", response.Completion)
	return nil
}

// ExecuteEntityExtraction demonstrates few-shot named entity recognition
func (f *FewShotPrompt) ExecuteEntityExtraction() error {
	prompt := `Extract named entities from the given text. Identify PERSON, ORGANIZATION, and LOCATION entities.

	Examples:
	Text: "John Smith works at Microsoft in Seattle."
	Entities:
	- PERSON: John Smith
	- ORGANIZATION: Microsoft
	- LOCATION: Seattle
	
	Text: "Apple Inc. was founded by Steve Jobs in Cupertino."
	Entities:
	- ORGANIZATION: Apple Inc.
	- PERSON: Steve Jobs
	- LOCATION: Cupertino
	
	Text: "The meeting with Google representatives will be held in San Francisco."
	Entities:
	- ORGANIZATION: Google
	- LOCATION: San Francisco
	
	Now extract entities from this text:
	Text: "Dr. Sarah Johnson from Harvard University will present her research at the conference in Boston next week."
	Entities:`

	fmt.Println("🎯 Few-Shot Prompting: Entity Extraction")
	fmt.Println("Prompt:", prompt)
	fmt.Println(strings.Repeat("-", 80))

	response, err := f.client.InvokeModel(prompt, f.params)
	if err != nil {
		return fmt.Errorf("failed to execute entity extraction: %w", err)
	}

	fmt.Printf("Response: %s\n\n", response.Completion)
	return nil
}

// ExecuteCodeCompletion demonstrates few-shot code completion
func (f *FewShotPrompt) ExecuteCodeCompletion() error {
	prompt := `Complete the following code snippets based on the pattern shown in the examples:

	Example 1:
	Input: Create a function to add two numbers
	Output:
	def add_numbers(a, b):
	    """Add two numbers and return the result."""
	    return a + b
	
	Example 2:
	Input: Create a function to multiply two numbers
	Output:
	def multiply_numbers(a, b):
	    """Multiply two numbers and return the result."""
	    return a * b
	
	Example 3:
	Input: Create a function to check if a number is even
	Output:
	def is_even(number):
	    """Check if a number is even."""
	    return number % 2 == 0
	
	Now complete this:
	Input: Create a function to find the maximum of three numbers
	Output:`

	fmt.Println("🎯 Few-Shot Prompting: Code Completion")
	fmt.Println("Prompt:", prompt)
	fmt.Println(strings.Repeat("-", 80))

	response, err := f.client.InvokeModel(prompt, f.params)
	if err != nil {
		return fmt.Errorf("failed to execute code completion: %w", err)
	}

	fmt.Printf("Response: %s\n\n", response.Completion)
	return nil
}

// ExecuteEmailClassification demonstrates few-shot email classification
func (f *FewShotPrompt) ExecuteEmailClassification() error {
	prompt := `Classify emails into categories: "urgent", "marketing", "support", or "general".

	Examples:
	Email: "URGENT: Server is down! Please fix immediately!"
	Category: urgent
	
	Email: "Check out our amazing 50% off sale this weekend!"
	Category: marketing
	
	Email: "I'm having trouble logging into my account. Can you help?"
	Category: support
	
	Email: "Thank you for your purchase. Your order has been shipped."
	Category: general
	
	Email: "SPECIAL OFFER: Buy 2 get 1 free on all products!"
	Category: marketing
	
	Now classify this email:
	Email: "Hi, I need assistance with setting up my new account. The verification email never arrived."
	Category:`

	fmt.Println("🎯 Few-Shot Prompting: Email Classification")
	fmt.Println("Prompt:", prompt)
	fmt.Println(strings.Repeat("-", 80))

	response, err := f.client.InvokeModel(prompt, f.params)
	if err != nil {
		return fmt.Errorf("failed to execute email classification: %w", err)
	}

	fmt.Printf("Response: %s\n\n", response.Completion)
	return nil
}

// ExecuteCreativeWriting demonstrates few-shot creative writing
func (f *FewShotPrompt) ExecuteCreativeWriting() error {
	prompt := `Write a short story opening based on the given prompt. Follow the style shown in the examples:

	Example 1:
	Prompt: A mysterious package arrives
	Opening: The package sat on her doorstep like a riddle wrapped in brown paper. No return address, no delivery notice—just her name written in elegant script that seemed to shimmer in the morning light.
	
	Example 2:
	Prompt: First day at a new job
	Opening: The elevator climbed twenty-three floors, and with each passing number, Marcus felt his confidence slip another notch. By the time the doors opened, he was pretty sure he'd made a terrible mistake.
	
	Example 3:
	Prompt: Finding an old diary
	Opening: The diary's leather cover was worn smooth by decades of handling, its pages yellowed and brittle. As Emma opened it, the scent of lavender and old secrets escaped into the dusty attic air.
	
	Now write an opening for this prompt:
	Prompt: Waking up in a world where colors have disappeared
	Opening:`

	fmt.Println("🎯 Few-Shot Prompting: Creative Writing")
	fmt.Println("Prompt:", prompt)
	fmt.Println(strings.Repeat("-", 80))

	params := f.params
	params.Temperature = 0.8 // Increase temperature for more creativity

	response, err := f.client.InvokeModel(prompt, params)
	if err != nil {
		return fmt.Errorf("failed to execute creative writing: %w", err)
	}

	fmt.Printf("Response: %s\n\n", response.Completion)
	return nil
}

// RunAllExamples executes all few-shot prompting examples
func (f *FewShotPrompt) RunAllExamples() {
	fmt.Println("=== FEW-SHOT PROMPTING EXAMPLES ===\n")

	examples := []struct {
		name string
		fn   func() error
	}{
		{"Sentiment Analysis", f.ExecuteSentimentAnalysis},
		{"Entity Extraction", f.ExecuteEntityExtraction},
		{"Code Completion", f.ExecuteCodeCompletion},
		{"Email Classification", f.ExecuteEmailClassification},
		{"Creative Writing", f.ExecuteCreativeWriting},
	}

	for _, example := range examples {
		if err := example.fn(); err != nil {
			log.Printf("Error in %s: %v", example.name, err)
		}
	}
}
