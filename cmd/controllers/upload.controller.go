package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)


type ReceiptData struct {
	Date    string `json:"Date"`
	Name    string `json:"Name"`
	Address string `json:"Address"`
	City    string `json:"City"`
	Zip     string `json:"Zip"`
	Phone   string `json:"Phone"`
	Author  string `json:"Author"`
	Total   string `json:"Total"`
}


func UploadImage(c *fiber.Ctx) error {
	// Parse the multipart form data
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid form data"})
	}

	// Get the image file and input
	files := form.File["image"]
	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "No image file provided"})
	}

	inputText := c.FormValue("input")
	if inputText == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "No input text provided"})
	}


	file := files[0]
	fileData, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to open image file"})
	}
	defer fileData.Close()

	imgData := make([]byte, file.Size)
	_, err = fileData.Read(imgData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to read image file"})
	}


	var requestBody bytes.Buffer
	multipartWriter := multipart.NewWriter(&requestBody)

	
	fileWriter, err := multipartWriter.CreateFormFile("file", file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to create form file"})
	}
	_, err = fileWriter.Write(imgData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to write file data"})
	}

	
	err = multipartWriter.WriteField("input", inputText)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to write input field"})
	}


	multipartWriter.Close()


	resp, err := http.Post("http://0.0.0.0:8000/analyze_invoice", multipartWriter.FormDataContentType(), &requestBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to call FastAPI endpoint"})
	}
	defer resp.Body.Close()


	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to read FastAPI response"})
	}

	
	var response map[string]interface{}
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to parse FastAPI response"})
	}

	
	responseContent, ok := response["response"].(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Invalid response format"})
	}

	
	jsonStr := strings.TrimPrefix(responseContent, "```json\n")
	jsonStr = strings.TrimSuffix(jsonStr, "\n```")


	var receiptData ReceiptData
	err = json.Unmarshal([]byte(jsonStr), &receiptData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to parse JSON data"})
	}


	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   receiptData,
	})
}
