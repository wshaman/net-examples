package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	slackGetMessagesURL = "https://slack.com/api/conversations.history"
)

type slackResp struct {
	Ok       bool `json:"ok"`
	Messages []struct {
		Text string `json:"text"`
	} `json:"messages"`
	Error string `json:"error"`
}

func (s slackResp) GetMessages() []string {
	messages := make([]string, 0, len(s.Messages))
	for _, v := range s.Messages {
		messages = append(messages, v.Text)
	}
	return messages
}

func readMessages() ([]string, error) {
	// Create a new request to slack API endpoint
	r, err := http.NewRequest(http.MethodGet, slackGetMessagesURL, nil)
	if err != nil {
		return nil, err
	}
	p := url.Values{}
	// Add mandatory params
	var token, channel string
	if token = os.Getenv("token"); token == "" {
		return nil, errors.New("no token set")
	}
	if channel = os.Getenv("channel"); channel == "" {
		return nil, errors.New("no channel set")
	}
	p.Add("token", token)
	p.Add("channel", channel)
	// Encode params to "?token=aaaaa&channel=BBB"
	r.URL.RawQuery = p.Encode()
	c := &http.Client{}
	// This is an actual request run
	resp, err := c.Do(r)
	if err != nil {
		return nil, err
	}
	// reading RESPONSE
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	m := &slackResp{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	if !m.Ok {
		return nil, errors.New(m.Error)
	}

	return m.GetMessages(), nil
}

func main() {
	msgs, err := readMessages()
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range msgs {
		fmt.Println(v)
	}
}
