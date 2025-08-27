package prompting

import (
	"fmt"
	"log"
	"strings"

	"aws-bedrock-prompt-engineering/internal/bedrock"
)

type ChainOfThoughtPrompt struct {
	client *bedrock.Client
	params bedrock.ModelParams
}

// NewChainOfThoughtPrompt creates a new instance for chain-of-thought prompting.
// This approach divides tasks into clear reasoning steps for structured, coherent solutions.
func NewChainOfThoughtPrompt(client *bedrock.Client) *ChainOfThoughtPrompt {
	params := bedrock.GetDefaultClaudeParams()

	params.Temperature = 0.4 // Lower temperature for logical reasoning
	params.MaxTokens = 1000  // More tokens for step-by-step reasoning

	return &ChainOfThoughtPrompt{
		client: client,
		params: params,
	}
}

// ExecuteMathProblemSolving demonstrates chain-of-thought for math problems
func (c *ChainOfThoughtPrompt) ExecuteMathProblemSolving() error {
	prompt := `Solve the following math problem step by step. Show your reasoning process.

	Problem: A store is having a sale. Sarah buys 3 shirts that normally cost $25 each, but they're 20% off. She also buys 2 pairs of jeans that cost $40 each with no discount. If she pays with a $200 gift card, how much money will she have left on the card?
	
	Let me think through this step by step:
	
	Step 1: Calculate the original cost of the shirts
	3 shirts × $25 each = $75
	
	Step 2: Calculate the discount on the shirts
	20% of $75 = 0.20 × $75 = $15
	
	Step 3: Calculate the discounted price of the shirts
	$75 - $15 = $60
	
	Step 4: Calculate the cost of the jeans
	2 pairs × $40 each = $80
	
	Step 5: Calculate the total purchase amount
	Shirts: $60 + Jeans: $80 = $140
	
	Step 6: Calculate the remaining amount on the gift card
	$200 - $140 = $60
	
	Therefore, Sarah will have $60 left on her gift card.
	
	Now solve this problem using the same step-by-step approach:
	
	Problem: Tom is planning a party for 24 people. Each pizza serves 8 people and costs $12. He also wants to buy drinks that cost $3 per person. If he has a $150 budget, how much money will he have left after buying the food and drinks?
	
	Let me think through this step by step:`

	fmt.Println("🧠 Chain-of-Thought Prompting: Math Problem Solving")
	fmt.Println("Prompt:", prompt)
	fmt.Println(strings.Repeat("-", 80))

	response, err := c.client.InvokeModel(prompt, c.params)
	if err != nil {
		return fmt.Errorf("failed to execute math problem solving: %w", err)
	}

	fmt.Printf("Response: %s\n\n", response.Completion)
	return nil
}

// ExecuteLogicalReasoning demonstrates chain-of-thought for logical reasoning
func (c *ChainOfThoughtPrompt) ExecuteLogicalReasoning() error {
	prompt := `Solve the following logical reasoning problem by thinking through each step.

	Example:
	Problem: All cats are mammals. All mammals are animals. Fluffy is a cat. Is Fluffy an animal?
	
	Reasoning:
	1. Given: All cats are mammals
	2. Given: All mammals are animals  
	3. Given: Fluffy is a cat
	4. From 1 and 3: Since Fluffy is a cat, and all cats are mammals, Fluffy is a mammal
	5. From 2 and 4: Since Fluffy is a mammal, and all mammals are animals, Fluffy is an animal
	
	Conclusion: Yes, Fluffy is an animal.
	
	Now solve this problem using the same logical reasoning approach:
	
	Problem: All teachers at Riverside School speak at least two languages. Ms. Johnson teaches at Riverside School. Everyone who speaks at least two languages can tutor international students. Can Ms. Johnson tutor international students?
	
	Reasoning:`

	fmt.Println("🧠 Chain-of-Thought Prompting: Logical Reasoning")
	fmt.Println("Prompt:", prompt)
	fmt.Println(strings.Repeat("-", 80))

	response, err := c.client.InvokeModel(prompt, c.params)
	if err != nil {
		return fmt.Errorf("failed to execute logical reasoning: %w", err)
	}

	fmt.Printf("Response: %s\n\n", response.Completion)
	return nil
}

// ExecuteProblemDecomposition demonstrates breaking down complex problems
func (c *ChainOfThoughtPrompt) ExecuteProblemDecomposition() error {
	prompt := `Break down the following complex problem into smaller, manageable steps and solve it systematically.

	Example:
	Problem: Design a simple mobile app for a local restaurant
	
	Step-by-step breakdown:
	1. Identify core requirements
	   - Menu display
	   - Order placement
	   - Location and contact info
	   - User accounts
	
	2. Plan user interface
	   - Home screen with navigation
	   - Menu categories and items
	   - Shopping cart functionality
	   - Checkout process
	
	3. Consider technical requirements
	   - Database for menu items
	   - Payment processing integration
	   - Push notifications for order status
	   - Backend API for order management
	
	4. Implementation phases
	   - Phase 1: Basic menu display
	   - Phase 2: Order functionality
	   - Phase 3: User accounts and history
	   - Phase 4: Advanced features
	
	Now break down this complex problem using the same systematic approach:
	
	Problem: Plan a sustainable office renovation project for a 50-person company that wants to reduce their environmental impact while improving employee productivity.
	
	Step-by-step breakdown:`

	fmt.Println("🧠 Chain-of-Thought Prompting: Problem Decomposition")
	fmt.Println("Prompt:", prompt)
	fmt.Println(strings.Repeat("-", 80))

	response, err := c.client.InvokeModel(prompt, c.params)
	if err != nil {
		return fmt.Errorf("failed to execute problem decomposition: %w", err)
	}

	fmt.Printf("Response: %s\n\n", response.Completion)
	return nil
}

// ExecuteCodeDebugging demonstrates chain-of-thought for debugging
func (c *ChainOfThoughtPrompt) ExecuteCodeDebugging() error {
	prompt := `Debug the following code by thinking through the logic step by step.

	Example:
	Code with bug:
	def calculate_average(numbers):
	    total = 0
	    for num in numbers:
	        total += num
	    return total / len(numbers)
	
	# Test
	result = calculate_average([])
	print(result)
	
	Debugging process:
	1. Analyze what the function should do: Calculate the average of a list of numbers
	2. Trace through the code:
	   - Initialize total = 0
	   - Loop through numbers and add to total
	   - Return total divided by length of numbers
	3. Identify the problem: When numbers is empty list [], len(numbers) = 0
	4. Issue: Division by zero will raise ZeroDivisionError
	5. Solution: Add check for empty list
	
	Fixed code:
	def calculate_average(numbers):
	    if not numbers:  # Check if list is empty
	        return 0     # or raise ValueError("Cannot calculate average of empty list")
	    total = 0
	    for num in numbers:
	        total += num
	    return total / len(numbers)
	
	Now debug this code using the same systematic approach:
	
	Code with bug:
	def find_max_value(data):
	    max_val = 0
	    for item in data:
	        if item > max_val:
	            max_val = item
	    return max_val
	
	# Test cases
	print(find_max_value([1, 5, 3, 9, 2]))  # Should return 9
	print(find_max_value([-5, -2, -8, -1]))  # Should return -1
	print(find_max_value([]))  # Should handle empty list
	
	Debugging process:`

	fmt.Println("🧠 Chain-of-Thought Prompting: Code Debugging")
	fmt.Println("Prompt:", prompt)
	fmt.Println(strings.Repeat("-", 80))

	response, err := c.client.InvokeModel(prompt, c.params)
	if err != nil {
		return fmt.Errorf("failed to execute code debugging: %w", err)
	}

	fmt.Printf("Response: %s\n\n", response.Completion)
	return nil
}

// ExecuteDecisionMaking demonstrates chain-of-thought for decision analysis
func (c *ChainOfThoughtPrompt) ExecuteDecisionMaking() error {
	prompt := `Analyze the following decision scenario step by step, considering all factors.

	Example:
	Decision: Should I accept a job offer in another city?
	
	Analysis framework:
	1. Define the decision criteria
	   - Salary and benefits
	   - Career growth opportunities
	   - Cost of living
	   - Work-life balance
	   - Distance from family/friends
	
	2. Evaluate current situation
	   - Current salary: $70K
	   - Limited growth opportunities
	   - Low cost of living
	   - Close to family
	
	3. Evaluate new opportunity
	   - New salary: $95K (+$25K)
	   - Strong growth potential
	   - Higher cost of living (+$15K/year)
	   - 500 miles from family
	
	4. Calculate net benefits
	   - Financial: +$10K after cost of living
	   - Career: Significant improvement
	   - Personal: Some sacrifice in family time
	
	5. Consider long-term implications
	   - Career trajectory over 5 years
	   - Potential for remote work
	   - Family visit frequency
	
	Conclusion: Accept if career growth is priority and financial gain justifies personal costs.
	
	Now analyze this decision using the same systematic approach:
	
	Decision: Should a small business owner invest $50,000 in new equipment or hire two additional employees?
	
	Analysis framework:`

	fmt.Println("🧠 Chain-of-Thought Prompting: Decision Making")
	fmt.Println("Prompt:", prompt)
	fmt.Println(strings.Repeat("-", 80))

	response, err := c.client.InvokeModel(prompt, c.params)
	if err != nil {
		return fmt.Errorf("failed to execute decision making: %w", err)
	}

	fmt.Printf("Response: %s\n\n", response.Completion)
	return nil
}

// RunAllExamples executes all chain-of-thought prompting examples
func (c *ChainOfThoughtPrompt) RunAllExamples() {
	fmt.Println("=== CHAIN-OF-THOUGHT PROMPTING EXAMPLES ===\n")

	examples := []struct {
		name string
		fn   func() error
	}{
		{"Math Problem Solving", c.ExecuteMathProblemSolving},
		{"Logical Reasoning", c.ExecuteLogicalReasoning},
		{"Problem Decomposition", c.ExecuteProblemDecomposition},
		{"Code Debugging", c.ExecuteCodeDebugging},
		{"Decision Making", c.ExecuteDecisionMaking},
	}

	for _, example := range examples {
		if err := example.fn(); err != nil {
			log.Printf("Error in %s: %v", example.name, err)
		}
	}
}
