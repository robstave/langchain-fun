# Langchain Go Example

Example that uses LangChainGo and Chromadb to query a vector store.
In this case, we are using Gemini instead of OpenAi.

This is pretty much the same example as https://github.com/tmc/langchaingo/tree/main/examples/chroma-vectorstore-example with just a few differences:

- Using godotenv for environment variables
- Different documents
- Dazzling colors

## Setup

### Chromadb

You will need to have a version of Chromadb. In this case, I'm using a docker image.  From what I can tell, there are dependancies here.  Check with documentation for the latest compatibilies.

Im using port 8989 because my homelab is a bit crowded.  Use what makes sense.

```bash
docker pull ghcr.io/chroma-core/chroma:0.5.0
docker run -d --name chroma -p 8989:8000 ghcr.io/chroma-core/chroma:0.5.0
```

### Gemini API

You will need an Google Gemini key. You can pretty much sign into this with your Google account and
use that.  So far its been free!  But this is not a terribly taxing example.
Set this in the `.env` file along with the URL of Chroma.  

Create a `.env` file with the following content:

```bash
CHROMA_URL=http://localhost:8989
API_KEY=xyz
```

Again, Im using 8989...but you do you.




