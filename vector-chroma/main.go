package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color" // Added for colored output

	chroma_go "github.com/amikos-tech/chroma-go/types"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/chroma"
)

func main() {

	color.Blue("Loading")
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	fmt.Printf("url:" + os.Getenv("CHROMA_URL") + "\n")
	// Create a new Chroma vector store.
	store, errNs := chroma.New(
		chroma.WithChromaURL(os.Getenv("CHROMA_URL")),
		chroma.WithOpenAIAPIKey(os.Getenv("OPENAI_API_KEY")),
		chroma.WithDistanceFunction(chroma_go.COSINE),
		chroma.WithNameSpace(uuid.New().String()),
	)
	if errNs != nil {
		log.Fatalf("new: %v\n", errNs)
	}

	// Define documents
	type meta map[string]interface{}

	// Add documents to the vector store.
	_, errAd := store.AddDocuments(context.Background(), []schema.Document{
		{PageContent: "Chucks: American Fare. Tasty burgers and fries in an informal setting. Not Overpriced", Metadata: meta{"name": "Chucks", "cuisine": "American", "stars": 2}},
		{PageContent: "Bella Italia: Classic Italian. Authentic pasta and pizza in a cozy atmosphere.", Metadata: meta{"name": "Bella Italia", "cuisine": "Italian", "stars": 4}},
		{PageContent: "Sushi Zen: Japanese Cuisine. Fresh sushi and sashimi with a traditional vibe. $2 Sake bombs on Fridays.", Metadata: meta{"name": "Sushi Zen", "cuisine": "Japanese", "stars": 5}},
		{PageContent: "Taco Loco: Mexican Fiesta. Lively spot for tacos, burritos, and margaritas. Hottest salsa in town.", Metadata: meta{"name": "Taco Loco", "cuisine": "Mexican", "stars": 3}},
		{PageContent: "The Curry House: Indian Spices. Rich and flavorful curries and naan bread.", Metadata: meta{"name": "The Curry House", "cuisine": "Indian", "stars": 4}},
		{PageContent: "Le Bistro: French Delights. Elegant French dining with classic dishes.", Metadata: meta{"name": "Le Bistro", "cuisine": "French", "stars": 5}},
		{PageContent: "Peking Garden: Chinese Cuisine. Wide variety of Chinese dishes in a family-friendly setting. Authentic Bao", Metadata: meta{"name": "Peking Garden", "cuisine": "Chinese", "stars": 3}},
		{PageContent: "Greek Islands Taverna: Mediterranean Flavors. Fresh seafood and Greek specialties.", Metadata: meta{"name": "Greek Islands Taverna", "cuisine": "Greek", "stars": 4}},
		{PageContent: "Steakhouse 77: Premium Steaks. High-end steakhouse with aged cuts and fine wines. Dress code.", Metadata: meta{"name": "Steakhouse 77", "cuisine": "American", "stars": 5}},
		{PageContent: "Vegan Paradise: Plant-Based Meals. Creative vegan dishes in a modern setting.", Metadata: meta{"name": "Vegan Paradise", "cuisine": "Vegan", "stars": 4}},
		{PageContent: "Pasta Pronto: Quick Italian. Fast and casual Italian pasta and salads.", Metadata: meta{"name": "Pasta Pronto", "cuisine": "Italian", "stars": 3}},
		{PageContent: "Ramen Republic: Noodle Bar. Authentic ramen and Japanese appetizers. Famous chef. ", Metadata: meta{"name": "Ramen Republic", "cuisine": "Japanese", "stars": 4}},
		{PageContent: "Burro Burrito: Tex-Mex Grill. Large burritos and Tex-Mex favorites.", Metadata: meta{"name": "Burro Burrito", "cuisine": "Mexican", "stars": 2}},
		{PageContent: "Spice Merchant: Indian Cuisine. Authentic Indian spices and dishes, known for biryani.  Goat and Lamb too.", Metadata: meta{"name": "Spice Merchant", "cuisine": "Indian", "stars": 5}},
		{PageContent: "Cafe de Paris: French Cafe. Casual French cafe with pastries and light meals.", Metadata: meta{"name": "Cafe de Paris", "cuisine": "French", "stars": 3}},
		{PageContent: "Golden Dragon: Dim Sum House. Traditional dim sum and Cantonese dishes.", Metadata: meta{"name": "Golden Dragon", "cuisine": "Chinese", "stars": 4}},
		{PageContent: "Olive Branch Bistro: Mediterranean Cuisine. Healthy Mediterranean options with a focus on olive oil. Breadsticks, but limit 4", Metadata: meta{"name": "Olive Branch Bistro", "cuisine": "Greek", "stars": 4}},
		{PageContent: "Burger Barn: Classic Burgers. Casual and delicious burgers, fries and milkshakes. Nothing Fancy", Metadata: meta{"name": "Burger Barn", "cuisine": "American", "stars": 3}},
		{PageContent: "Noodle Nook: Asian Noodles. Variety of Asian noodle dishes from different regions.", Metadata: meta{"name": "Noodle Nook", "cuisine": "Asian Fusion", "stars": 3}},
		{PageContent: "Pizza Palace: Family Pizza. Large pizzas and family-friendly atmosphere. Retro games ", Metadata: meta{"name": "Pizza Palace", "cuisine": "Italian", "stars": 2}},
		{PageContent: "Dougs: Sammies and beans. Not recommended. Low prices.  Happy Hour ", Metadata: meta{"name": "Dougs", "cuisine": "American", "stars": 1}},
	})
	if errAd != nil {
		log.Fatalf("AddDocument: %v\n", errAd)
	}

	ctx := context.TODO()

	type exampleCase struct {
		name         string
		query        string
		numDocuments int
		options      []vectorstores.Option
	}

	type filter = map[string]any

	exampleCases := []exampleCase{
		{
			name:         "Informal American Fare",
			query:        "What is the name of the restaurant that serves informal American fare?",
			numDocuments: 5,
			options: []vectorstores.Option{
				vectorstores.WithScoreThreshold(0.8),
			},
		},
		{
			name:         "Cheap meals",
			query:        "Which of these are the cheapest meals?",
			numDocuments: 2,
			options: []vectorstores.Option{
				vectorstores.WithScoreThreshold(0.8),
			},
		},
		{
			name:         "Craving",
			query:        "Where can I get a Chimichanga?",
			numDocuments: 2,
			options: []vectorstores.Option{
				vectorstores.WithScoreThreshold(0.8),
			},
		},

		{
			name:         "Best to eat",
			query:        "Where should I eat thats rated well?",
			numDocuments: 3,
			options: []vectorstores.Option{
				vectorstores.WithFilters(filter{"stars": map[string]interface{}{"$gte": 4}}),
			},
		},
	}

	// run the example cases
	results := make([][]schema.Document, len(exampleCases))
	for ecI, ec := range exampleCases {
		docs, errSs := store.SimilaritySearch(ctx, ec.query, ec.numDocuments, ec.options...)
		if errSs != nil {
			log.Fatalf("query1: %v\n", errSs)
		}
		results[ecI] = docs
	}

	// Create SprintFuncs for each desired color.
	green := color.New(color.FgGreen).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	red := color.New(color.FgRed).SprintFunc()

	// Print out the resulgo ts with some color.
	color.Cyan("Results:")
	for ecI, ec := range exampleCases {
		texts := make([]string, len(results[ecI]))
		scores := make([]string, len(results[ecI]))
		stars := make([]string, len(results[ecI]))
		cuisine := make([]string, len(results[ecI]))
		for docI, doc := range results[ecI] {
			texts[docI] = doc.PageContent
			scores[docI] = fmt.Sprint(doc.Score)
			stars[docI] = fmt.Sprint(doc.Metadata["stars"])
			cuisine[docI] = fmt.Sprint(doc.Metadata["cuisine"])
		}
		color.Magenta("Question %d. %s", ecI+1, ec.name)
		color.Blue("    %s", ec.query)

		color.Yellow("    result: ")

		for i, text := range texts {
			// Print text and stars on the same line.
			// The stars appear at the end in blue.
			fmt.Printf("    %s %s %s \n", green(text), yellow(stars[i]), blue(cuisine[i]))
			// Print score on the next line in red.
			fmt.Printf("    Score: %s\n", red(scores[i]))
		}

	}
}
