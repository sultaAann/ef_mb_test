package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	AGE    string = "https://api.agify.io/"
	GENDER string = "https://api.genderize.io/"
	NATION string = "https://api.nationalize.io/"
)

type Parsed_data struct {
	Age     int
	Gender  string
	Nations []string
}

type Parser struct {
	client *http.Client
}

func NewParser(NewClient *http.Client) *Parser {
	return &Parser{client: NewClient}
}

func (p *Parser) Parse(name string) (*Parsed_data, error) {
	age, err := p.parse_age(name)
	if err != nil {
		return nil, err
	}
	gender, err := p.parse_gender(name)
	if err != nil {
		return nil, err
	}
	nations, err := p.parse_nation(name)
	if err != nil {
		return nil, err
	}

	return &Parsed_data{age, gender, nations}, nil
}

func (p *Parser) do_request(todo, name string) ([]byte, error) {
	req, err := http.NewRequest("GET", todo, nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("name", name)
	req.URL.RawQuery = q.Encode()

	resp, err := p.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("Error doing request: response body: %s", body)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return body, nil
}

func (p *Parser) parse_age(name string) (int, error) {
	body, err := p.do_request(AGE, name)
	if err != nil {
		return 0, err
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)

	if err != nil {
		return -1, err
	}
	age, ok := data["age"].(float64)
	if !ok {
		return -1, fmt.Errorf("Person not found")
	}
	return int(age), nil
}

func (p *Parser) parse_gender(name string) (string, error) {
	body, err := p.do_request(GENDER, name)
	if err != nil {
		return "", err
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)

	if err != nil {
		return "", err
	}
	gender, ok := data["gender"].(string)
	if !ok {
		return "", fmt.Errorf("Person not found")
	}
	return gender, nil
}

func (p *Parser) parse_nation(name string) ([]string, error) {
	body, err := p.do_request(NATION, name)

	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, err
	}
	list, ok := data["country"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("Person not found")
	}
	res := []string{}
	for _, el := range list {
		res = append(res, el.(map[string]interface{})["country_id"].(string))
	}
	return res, nil
}
