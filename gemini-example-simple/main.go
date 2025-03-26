package main

import (
	"context"
	"log"
	"os"

	"github.com/fatih/color" // Added for colored output

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
)

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	ctx := context.Background()

	apiKey := os.Getenv("API_KEY")
	llm, err := googleai.New(ctx, googleai.WithAPIKey(apiKey), googleai.WithDefaultModel("gemini-1.5-flash"))
	if err != nil {
		log.Fatal(err)
	}

	prompt := "Tell me a joke about a computer."
	answer, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		log.Fatal(err)
	}

	color.Blue(answer)
}
