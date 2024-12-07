from fastapi import FastAPI
from pydantic import BaseModel
import spacy

nlp = spacy.load("ja_ginza")

class TextRequest(BaseModel):
    text: str

app = FastAPI()

@app.post("/tokenize")
async def tokenize_text(request: TextRequest):
    doc = nlp(request.text)
    
    tokens = [token.text for token in doc]
    print(tokens)
    return {"tokens": tokens}
