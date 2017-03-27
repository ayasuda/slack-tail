package main

import (
	"encoding/json"
	"fmt"
)

type Event struct {
	Type string `json:"type"`
	Message
}

type Message struct {
	Channel string `json:"channel"`
	User    string `json:"user"`
	Text    string `json:"text"`
}

func trimEvent(json_str string) (*Event, error) {
	var evt = Event{}
	err := json.Unmarshal([]byte(json_str), &evt)
	if err != nil {
		return nil, err
	}
	return &evt, nil
}

func (evt *Event) isPrint() bool {
	return (evt.Text != "")
}

func (evt *Event) toString() string {
	ch_id, name_id, body := evt.Channel, evt.User, evt.Text
	if body != "" {
		ch := ch_id     // rtm.getChannelNameById(ch_id)
		name := name_id // rtm.getUserNameById(name_id)
		return fmt.Sprintf("%s\t%s\t%s", ch, name, body)
	}
	return ""
}
