# OpenAI Chat Simple Example

This example demonstrates how to use the [Langchain Go](https://tmc.github.io/langchaingo/docs/) library with OpenAI's GPT model to generate content. The program generates three variations of a Chuck Norris joke about AI and displays token usage statistics.

## Prerequisites

- Go installed on your system.
- An OpenAI API key.
- A `.env` file in the project directory containing your OpenAI API key:
  ```
  OPENAI_API_KEY=your-api-key-here
  ```

## How It Works

1. The program loads environment variables from a `.env` file.
2. It initializes the OpenAI LLM with the `gpt-4o-mini` model.
3. A prompt is sent to the model to generate three variations of a Chuck Norris joke about AI.
4. The program displays the generated content, token usage statistics, and generation information.

## Running the Example

1. Clone the repository and navigate to the `open-api-chat-simple` directory:
   ```bash
   git clone <repository-url>
   cd open-api-chat-simple
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the program:
   ```bash
   go run main.go
   ```

## Output

The program will display:

- Three variations of the generated joke.
- Token usage statistics (input and output tokens).
- Generation information for the first output.

## Notes

- The `CountTokens` function is used to calculate token usage. This is a utility provided by the Langchain Go library.
- The `fatih/color` package is used for colored terminal output.

## Example Output

```
Generating content...
Joke 1: <Generated Joke 1>
Joke 2: <Generated Joke 2>
Joke 3: <Generated Joke 3>

Token usage:
input: 10 / output: 50

generation info:
{
  "finish_reason": "stop",
  "index": 0
}
```
