package gemini

import (
	"bytes"
	"strings"

	"github.com/Bashsoft707/tweet-audit/src/config"
	"github.com/Bashsoft707/tweet-audit/src/models"

	"net/http"

	"encoding/json"
	"fmt"
	"io"
)

type GeminiRequest struct {
	Contents []GeminiBodyContent `json:"contents"`
}

type GeminiBodyContent struct {
	Parts []GeminiContentPart `json:"parts"`
}

type GeminiContentPart struct {
	Text string `json:"text"`
}

type GeminiResponse struct {
	Candidates []GeminiResponseCandidate `json:"candidates"`
}

type GeminiResponseCandidate struct {
	Content GeminiBodyContent `json:"content"`
}

func AnalyseTweet(tweet models.Tweet, apiKey string, baseURL string, cfg *config.Config) (string, error) {
	geminiUrl := baseURL + "?key=" + apiKey

	forbiddenWords := strings.Join(cfg.Criteria.ForbiddenWords, ", ")
	professionalCheck := "disabled"
	if cfg.Criteria.ProfessionalCheck {
		professionalCheck = "enabled"
	}

	prompt := fmt.Sprintf(`You are reviewing a tweet for content moderation.

Tweet: "%s"

Flag this tweet if it meets any of these criteria:
- Contains any of these words: %s
- Is unprofessional in tone: %s

Reply with exactly one of these two formats:
FLAGGED: <one sentence reason>
CLEAN`, tweet.FullText, forbiddenWords, professionalCheck)

	reqBody := GeminiRequest{
		Contents: []GeminiBodyContent{
			{
				Parts: []GeminiContentPart{
					{Text: prompt},
				},
			},
		},
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(geminiUrl, "application/json", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("Gemini API returned status %d: %s", resp.StatusCode, string(body))
	}

	var geminiResp GeminiResponse
	err = json.NewDecoder(resp.Body).Decode(&geminiResp)
	if err != nil {
		return "", err
	}

	if len(geminiResp.Candidates) > 0 && len(geminiResp.Candidates[0].Content.Parts) > 0 {
		return geminiResp.Candidates[0].Content.Parts[0].Text, nil
	}

	return "", nil
}
