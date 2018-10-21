package translate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const apiUrl = "https://translate.yandex.net/api/v1.5/tr.json"

type Translate struct {
	apiKey     string
	httpClient *http.Client
}

func New(apiKey string) *Translate {
	return &Translate{
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}
}

// https://tech.yandex.com/translate/doc/dg/reference/getLangs-docpage/
func (t *Translate) GetLanguages(ui string) (map[string]string, error) {
	if ui == "" {
		ui = "en"
	}

	params := make(url.Values)
	params.Set("ui", ui)

	var jsonResp struct {
		Code    uint16            `json:"code"`
		Message string            `json:"message"`
		Langs   map[string]string `json:"langs"`
	}

	if err := t.call("/getLangs", params, &jsonResp); err != nil {
		return nil, err
	}

	if jsonResp.Code != 0 {
		return nil, fmt.Errorf("%d: %s", jsonResp.Code, jsonResp.Message)
	}

	return jsonResp.Langs, nil
}

// https://tech.yandex.com/translate/doc/dg/reference/detect-docpage/
func (t *Translate) Detect(text string, hint []string) (string, error) {
	params := make(url.Values)
	params.Set("text", text)
	if len(hint) > 0 {
		params.Set("hint", strings.Join(hint, ","))
	}

	var jsonResp struct {
		Code    uint16 `json:"code"`
		Message string `json:"message"`
		Lang    string `json:"lang"`
	}

	if err := t.call("/detect", params, &jsonResp); err != nil {
		return "", err
	}

	if jsonResp.Code != 200 {
		return "", fmt.Errorf("%d: %s", jsonResp.Code, jsonResp.Message)
	}

	return jsonResp.Lang, nil
}

// https://tech.yandex.com/translate/doc/dg/reference/translate-docpage/
func (t *Translate) Translate(texts []string, lang string, format string) ([]string, error) {
	params := make(url.Values)
	for _, s := range texts {
		params.Add("text", s)
	}
	params.Set("lang", lang)
	params.Set("format", format)

	var jsonResp struct {
		Code    uint16   `json:"code"`
		Message string   `json:"message"`
		Text    []string `json:"text"`
	}

	if err := t.call("/translate", params, &jsonResp); err != nil {
		return nil, err
	}

	if jsonResp.Code != 200 {
		return nil, fmt.Errorf("%d: %s", jsonResp.Code, jsonResp.Message)
	}

	return jsonResp.Text, nil
}

func (t *Translate) call(method string, params url.Values, dst interface{}) error {
	params.Set("key", t.apiKey)

	resp, err := t.httpClient.PostForm(apiUrl+method, params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(dst); err != nil {
		return err
	}

	return nil
}
