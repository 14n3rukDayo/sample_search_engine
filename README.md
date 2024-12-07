# Sample Search Engine

This is a sample project where I created a full-text search engine tailored for Japanese text.

## Project Highlights

- **Index-Based Search Implementation with Inverted Index**  
  Built an inverted index for efficient retrieval of search results.
  
- **Score Calculation Using Okapi BM25**  
  Implemented a scoring mechanism to rank search results based on relevance.
  
- **Layered Architecture**  
  The project adopts a clean, layered architecture for better separation of concerns and maintainability.
  
- **Domain-Driven Design (DDD)**  
  Applied DDD principles to structure the project effectively.
  
- **Japanese Tokenizer Integration**  
  Unlike English, Japanese does not use spaces to separate words. To address this, a morphological analyzer is used to tokenize Japanese text properly.

- **Simple Data Persistence with Redis**  
  For simplicity, Redis is used for data storage.
  
- **Synonym Support**  
  Implemented synonym handling to improve search relevance. Synonyms are managed using a JSON-based configuration for flexibility and easy updates.

## Current Features

- **Basic Data Addition and Search Functionality**  
  You can add data to the search engine.

- **AND Search**  
  Supports AND-based query logic for precise matching.

## Why Japanese Text Requires Special Handling

Japanese differs significantly from English in terms of text processing.  
While English uses spaces to separate words, Japanese text is written continuously without spaces. This requires the use of a morphological analyzer to break sentences into meaningful tokens for effective searching.

## How to Get Started

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/sample-search-engine.git
   ```
2. Navigate to the project directory:
    ```bash
    cd sample-search-engine
    ```
3. Start the application using Docker Compose:
    ```bash
    docker compose up
    ```
4. Check `./docs/openapi.yml`
5. Add a dcoument
    ```bash
    curl -X POST http://localhost:8000/document \
    -H "Content-Type: application/json" \
    -d '@request_body_sample/add.json'
    ```

6. Search documents
    ```bash
    curl -X POST http://localhost:8000/search \
    -H "Content-Type: application/json" \
    -d '@request_body_sample/search.json'
    ```

7. To check Redis data, visit

    check `http://http://localhost:8081`


