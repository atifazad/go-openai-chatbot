package handlers

import (
	"encoding/json"
	"go-chatbot/src/utils"
	"net/http"
)

type ChatRequest struct {
	UserInput string `json:"user_input"`
}

type ChatResponse struct {
	BotResponse string `json:"bot_response"`
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var chatRequest ChatRequest
	if err := json.NewDecoder(r.Body).Decode(&chatRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	botResponse, err := utils.GetGPT4Response(chatRequest.UserInput)
	if err != nil {
		http.Error(w, "Error getting response from GPT-3", http.StatusInternalServerError)
		return
	}

	response := ChatResponse{BotResponse: botResponse}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
