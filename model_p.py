from fastapi import FastAPI, File, UploadFile, Form
from fastapi.responses import JSONResponse
from dotenv import load_dotenv
import os
import google.generativeai as genai
from PIL import Image
import io

load_dotenv()  # take environment variables from .env.

app = FastAPI()

# os.getenv("GEMINI_API_KEY")
genai.configure(api_key="AIzaSyBhJ709JA3k4OIscgN1yWW5HmrKJrpPpes")

def get_gemini_response(input, image, prompt):
    model = genai.GenerativeModel('gemini-1.5-pro')
    response = model.generate_content([input, image[0], prompt])
    return response.text

def input_image_setup(file):
    image = Image.open(io.BytesIO(file))
    image_parts = [
        {
            "mime_type": "image/jpeg",  # Assuming JPEG format, adjust if needed
            "data": file
        }
    ]
    return image_parts

@app.post("/analyze_invoice")
async def analyze_invoice(file: UploadFile = File(...), input: str = Form(...)):
    input_prompt = """
               You are an expert in understanding invoices.
               You will receive input images as invoices &
               you will have to answer questions based on the input image
               """
    
    try:
        contents = await file.read()
        image_data = input_image_setup(contents)
        response = get_gemini_response(input_prompt, image_data, input)
        return JSONResponse(content={"response": response})
    except Exception as e:
        return JSONResponse(content={"error": str(e)}, status_code=400)

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000) 