package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color" // Added for colored output

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	llm, err := openai.New(

		openai.WithModel("gpt-4o-mini"),
	)
	if err != nil {
		log.Fatal(err)
	}

	joke := "Tell a chuck norris joke about ai."

	ctx := context.Background()
	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, joke),
	}
	color.Blue("Generating content...")
	output, err := llm.GenerateContent(ctx, content,
		llms.WithMaxTokens(4000),
		llms.WithTemperature(1),
		llms.WithN(3),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output.Choices[0].Content)
	fmt.Println(output.Choices[1].Content)
	fmt.Println(output.Choices[2].Content)
	// Print the number of tokens used using the CountTokens function
	// Note: The CountTokens function is not a part of the OpenAI API, but a utility function
	// provided by the langchaingo library to count tokens in a string.
	inputToken := llms.CountTokens("gpt-4", joke)
	outputToken := llms.CountTokens("gpt-4", output.Choices[0].Content+output.Choices[1].Content+output.Choices[2].Content)

	color.Green("\n\nToken usage:")
	color.Green("input: %v / output: %v\n", inputToken, outputToken)

	color.Blue("\ngeneration info:")
	json.NewEncoder(os.Stdout).Encode(output.Choices[0].GenerationInfo)

}
