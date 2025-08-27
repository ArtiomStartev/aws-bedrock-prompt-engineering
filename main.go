package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"aws-bedrock-prompt-engineering/internal/bedrock"
	"aws-bedrock-prompt-engineering/internal/prompting"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	client, err := bedrock.NewClient()
	if err != nil {
		log.Fatal("Error creating Bedrock client: ", err)
	}

	// Display welcome message
	displayWelcomeMessage()

	// Main menu loop
	for {
		choice := displayMenu()

		switch choice {
		case "1":
			runZeroShotExamples(client)
		case "2":
			runFewShotExamples(client)
		case "3":
			runChainOfThoughtExamples(client)
		case "4":
			runAllExamples(client)
		case "5":
			runInteractiveMode(client)
		case "6":
			fmt.Println("👋 Thank you for using AWS Bedrock Prompt Engineering Demo!")
			return
		default:
			fmt.Println("❌ Invalid choice. Please try again.")
		}

		fmt.Println("\nPress Enter to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}

func displayWelcomeMessage() {
	fmt.Println("🚀 Welcome to AWS Bedrock Prompt Engineering Demo!")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("This application demonstrates three key prompting techniques:")
	fmt.Println("• Zero-Shot Prompting: Direct questions without examples")
	fmt.Println("• Few-Shot Prompting: Learning from provided examples")
	fmt.Println("• Chain-of-Thought: Step-by-step reasoning process")
	fmt.Println(strings.Repeat("=", 60))
}

func displayMenu() string {
	fmt.Println("\n📋 Main Menu:")
	fmt.Println("1. 🎯 Zero-Shot Prompting Examples")
	fmt.Println("2. 🎪 Few-Shot Prompting Examples")
	fmt.Println("3. 🧠 Chain-of-Thought Prompting Examples")
	fmt.Println("4. 🌟 Run All Examples")
	fmt.Println("5. 💬 Interactive Mode")
	fmt.Println("6. 🚪 Exit")
	fmt.Print("\nEnter your choice (1-6): ")

	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	return strings.TrimSpace(choice)
}

func runZeroShotExamples(client *bedrock.Client) {
	fmt.Println("\n" + strings.Repeat("=", 80))
	zeroShot := prompting.NewZeroShotPrompt(client)
	zeroShot.RunAllExamples()
	fmt.Println(strings.Repeat("=", 80))
}

func runFewShotExamples(client *bedrock.Client) {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fewShot := prompting.NewFewShotPrompt(client)
	fewShot.RunAllExamples()
	fmt.Println(strings.Repeat("=", 80))
}

func runChainOfThoughtExamples(client *bedrock.Client) {
	fmt.Println("\n" + strings.Repeat("=", 80))
	chainOfThought := prompting.NewChainOfThoughtPrompt(client)
	chainOfThought.RunAllExamples()
	fmt.Println(strings.Repeat("=", 80))
}

func runAllExamples(client *bedrock.Client) {
	fmt.Println("\n🌟 Running all prompting technique examples...")

	runZeroShotExamples(client)
	fmt.Println("\n⏳ Pausing between techniques...")

	runFewShotExamples(client)
	fmt.Println("\n⏳ Pausing between techniques...")

	runChainOfThoughtExamples(client)

	fmt.Println("\n✅ All examples completed!")
}

func runInteractiveMode(client *bedrock.Client) {
	fmt.Println("\n💬 Interactive Mode - Enter your own prompts!")
	fmt.Println("Type 'exit' to return to main menu")
	fmt.Println(strings.Repeat("-", 50))

	reader := bufio.NewReader(os.Stdin)
	params := bedrock.GetDefaultClaudeParams()

	for {
		fmt.Print("\n🤖 Enter your prompt: ")
		prompt, _ := reader.ReadString('\n')
		prompt = strings.TrimSpace(prompt)

		if strings.ToLower(prompt) == "exit" {
			break
		}

		if prompt == "" {
			fmt.Println("❌ Please enter a valid prompt.")
			continue
		}

		fmt.Println("\n🔄 Processing your request...")
		response, err := client.InvokeModel(prompt, params)
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			continue
		}

		fmt.Printf("\n🎯 Response:\n%s\n", response.Completion)
		fmt.Println(strings.Repeat("-", 50))
	}
}
