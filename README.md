# AWS Bedrock Prompt Engineering

A comprehensive Go application demonstrating advanced prompt engineering techniques using AWS Bedrock and Claude AI models.

## 🚀 Overview

This project showcases three fundamental prompt engineering techniques through practical, reusable code examples. Each technique is implemented with real-world scenarios to help developers understand when and how to apply different prompting strategies.

## 🎯 Prompt Engineering Techniques

### 1. Zero-Shot Prompting
Direct task execution without providing examples, leveraging the model's pre-trained knowledge:
- **Text Classification** - Sentiment analysis without training data
- **Question Answering** - Factual responses from general knowledge
- **Language Translation** - Multi-language text conversion
- **Code Generation** - Programming solutions from natural language descriptions

### 2. Few-Shot Prompting
Pattern learning through curated examples for improved accuracy and consistency:
- **Sentiment Analysis** - Enhanced classification with contextual examples
- **Named Entity Recognition** - Structured data extraction
- **Code Completion** - Programming patterns and best practices
- **Email Classification** - Automated categorization systems
- **Creative Writing** - Style-consistent content generation

### 3. Chain-of-Thought Prompting
Step-by-step reasoning for complex problem-solving scenarios:
- **Mathematical Problem Solving** - Multi-step calculations with explanations
- **Logical Reasoning** - Structured argument analysis
- **Problem Decomposition** - Breaking complex tasks into manageable steps
- **Code Debugging** - Systematic error identification and resolution
- **Decision Analysis** - Structured decision-making frameworks

## 📁 Project Structure

```
aws-bedrock-prompt-engineering/
├── main.go                          # Interactive application with menu system
├── go.mod                           # Go module dependencies
├── go.sum                           # Dependency checksums
├── .env.example                     # Environment configuration template
├── .env                            # Your environment variables (gitignored)
├── Makefile                        # Build and development automation
├── README.md                       # Project documentation
└── internal/
    ├── bedrock/
    │   └── client.go               # AWS Bedrock client abstraction
    └── prompting/
        ├── zero_shot.go            # Zero-shot technique implementations
        ├── few_shot.go             # Few-shot technique implementations
        └── chain_of_thought.go     # Chain-of-thought technique implementations
```

## 🛠️ Prerequisites

- **Go 1.19+** - Latest stable version recommended
- **AWS CLI** - Configured with valid credentials
- **AWS Bedrock Access** - Claude model permissions enabled
- **Valid AWS Account** - With Bedrock service activated in your region

## ⚙️ Quick Start

### 1. Environment Setup
```bash
# Copy environment template
cp .env.example .env

# Edit with your AWS configuration
vim .env  # or your preferred editor
```

### 2. Install Dependencies
```bash
make install
# or manually: go mod tidy
```

### 3. Run the Application
```bash
make run
# or manually: go run main.go
```

## 🎮 Usage

The application provides an interactive menu system for exploring different prompting techniques:

```
📋 Main Menu:
1. 🎯 Zero-Shot Prompting Examples
2. 🎪 Few-Shot Prompting Examples  
3. 🧠 Chain-of-Thought Prompting Examples
4. 🌟 Run All Examples
5. 💬 Interactive Mode
6. 🚪 Exit
```

### Interactive Mode
Test custom prompts in real-time with immediate feedback and response analysis.

## 🔧 Configuration

### Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `AWS_REGION` | AWS region for Bedrock service | `us-east-1` | Yes |
| `AWS_ACCESS_KEY_ID` | AWS access credentials | - | Yes* |
| `AWS_SECRET_ACCESS_KEY` | AWS secret credentials | - | Yes* |
| `MODEL_ID` | Claude model identifier | `anthropic.claude-v2:1` | No |

*Required if not using IAM roles or AWS CLI profiles

### Model Parameters

Each technique uses optimized parameters for best results:

- **Zero-Shot**: Temperature 0.3 (focused responses)
- **Few-Shot**: Temperature 0.5 (balanced creativity)  
- **Chain-of-Thought**: Temperature 0.4, Max tokens 1000 (detailed reasoning)

## 🏗️ Architecture

### Modular Design
- **Separation of Concerns**: Each prompting technique isolated in dedicated modules
- **Reusable Components**: Bedrock client abstraction for easy extension
- **Clean Interfaces**: Well-defined APIs for consistent usage patterns

### Code Organization
```go
// Example usage
client, _ := bedrock.NewClient()
zeroShot := prompting.NewZeroShotPrompt(client)
zeroShot.ExecuteTextClassification()
```

## 📊 When to Use Each Technique

| Technique | Best For | Advantages | Considerations |
|-----------|----------|------------|----------------|
| **Zero-Shot** | Simple, well-defined tasks | Quick setup, no examples needed | May lack domain specificity |
| **Few-Shot** | Pattern recognition, consistency | Higher accuracy, controlled output | Requires good examples |
| **Chain-of-Thought** | Complex reasoning, multi-step problems | Explainable logic, detailed analysis | Higher token usage |

## 🚀 Development

### Available Make Commands
```bash
make help          # Show all available commands
make install       # Install dependencies
make build         # Build binary
make run           # Run application
make check         # Format code and run checks
make clean         # Clean build artifacts
```

### Extending the Project
1. **Add New Techniques**: Create new files in `internal/prompting/`
2. **Support New Models**: Extend the `bedrock.Client` interface
3. **Custom Parameters**: Modify model parameters for specific use cases

## 🔍 Troubleshooting

### Common Issues

**Authentication Errors**
```bash
# Verify AWS credentials
aws sts get-caller-identity
```

**Model Access Denied**
- Request Claude model access in AWS Bedrock console
- Ensure correct region configuration
- Verify IAM permissions for Bedrock service

**Rate Limiting**
- Implement exponential backoff for production use
- Monitor CloudWatch metrics for usage patterns
- Consider using different model variants

## 📚 Resources

- [AWS Bedrock Documentation](https://docs.aws.amazon.com/bedrock/)
- [Claude API Reference](https://docs.anthropic.com/claude/reference/)
- [Prompt Engineering Best Practices](https://www.promptingguide.ai/)
- [Go AWS SDK v2](https://aws.github.io/aws-sdk-go-v2/docs/)

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/new-technique`
3. Implement your changes with tests
4. Run code quality checks: `make check`
5. Submit a pull request with detailed description

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

---

**Built with ❤️ for the AI and prompt engineering community**
