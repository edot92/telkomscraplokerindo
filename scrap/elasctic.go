package scrap

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var urlElastic = "http://localhost:9200"

type StructP1Cleaning struct {
	Tokens []struct {
		Token       string `json:"token"`
		StartOffset int    `json:"start_offset"`
		EndOffset   int    `json:"end_offset"`
		Type        string `json:"type"`
		Position    int    `json:"position"`
	} `json:"tokens"`
}
type StructStandardRequestPost struct {
	Analyzer string `json:"analyzer"`
	Text     string `json:"text"`
}

// InitElastic
// register stadar analizer
func InitElastic() {
	dat, err := ioutil.ReadFile("stopwords.txt")
	checkError(err)
	tokensizerArray := strings.Split(string(dat), "\r")
	// untuk data json indonesian_example
	bodyJSON := map[string]interface{}{
		"settings": map[string]interface{}{
			"analysis": map[string]interface{}{
				"filter": map[string]interface{}{
					"indonesian_stop": map[string]interface{}{
						"type":      "stop",
						"stopwords": "_indonesian_",
					},
					"indonesian_keywords": map[string]interface{}{
						"type":     "keyword_marker",
						"keywords": tokensizerArray,
					},
					"indonesian_stemmer": map[string]interface{}{
						"type":     "stemmer",
						"language": "indonesian",
					},
				},
				"analyzer": map[string]interface{}{
					"indonesian": map[string]interface{}{
						"tokenizer": "standard",
						"filter": []string{
							"lowercase",
							"indonesian_stop",
							"indonesian_keywords",
							"indonesian_stemmer",
						},
					},
				},
			},
		},
	}
	bodyJSON2 := map[string]interface{}{
		"settings": map[string]interface{}{
			"analysis": map[string]interface{}{
				"analysis": map[string]interface{}{
					"analyzer": map[string]interface{}{
						"my_stop_analyzer": map[string]interface{}{
							"type":      "stop",
							"stopwords": tokensizerArray,
						},
					},
				},
			},
		},
	}
	// delete
	// DELETE / indonesian_example
	mJson, err := json.Marshal(bodyJSON2)
	if err != nil {
		checkError(err)
	}
	contentReader := bytes.NewReader(mJson)
	req, err := http.NewRequest("DELETE", urlElastic+"/my_stopword_analyzer_telkom", nil)
	if err != nil {
		checkError(err)
	}
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Notes", "GoRequest is coming!")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		checkError(err)
	}
	_ = resp
	req2, err := http.NewRequest("PUT", urlElastic+"/my_stopword_analyzer_telkom", contentReader)
	if err != nil {
		checkError(err)
	}
	req2.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Notes", "GoRequest is coming!")
	client = &http.Client{}
	resp, err = client.Do(req2)
	if err != nil {
		checkError(err)
	}
	_ = bodyJSON
	// log.Println(resp)
}
func PStopWord(text string) {
	m := map[string]interface{}{
		"analyzer": "my_stopword_analyzer_telkom",
		"text":     text,
	}
	mJson, err := json.Marshal(m)
	if err != nil {
		checkError(err)
	}
	contentReader := bytes.NewReader(mJson)
	req2, err := http.NewRequest("PUT", urlElastic+"/my_stopword_analyzer_telkom", contentReader)
	if err != nil {
		checkError(err)
	}
	req2.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Notes", "GoRequest is coming!")
	client := &http.Client{}
	resp, err := client.Do(req2)
	if err != nil {
		checkError(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		checkError(err)
	}
	log.Println(string(body))
}

// P1Cleaning ...
//  see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-htmlstrip-charfilter.html
// https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-lowercase-tokenizer.html
func P1Cleaning(text string) (data StructP1Cleaning, err error) {
	m := map[string]interface{}{
		"tokenizer":   "lowercase",
		"char_filter": []string{"html_strip"},
		"filter":      []string{"lowercase"},
		"text":        text,
		"analyzer":    "stop",
		// "analyzer":    "indonesian_example",
	}
	mJson, err := json.Marshal(m)
	if err != nil {
		return data, err
	}
	contentReader := bytes.NewReader(mJson)
	req, _ := http.NewRequest("POST", urlElastic+"/_analyze", contentReader)
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Notes", "GoRequest is coming!")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	// log.Println(string(body))
	// log.Println(string(mJson))
	return data, err
}
