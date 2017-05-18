package main

import (
	"testing"
)

const base_json = `
{
    "type": "message",
    "ts": "1358878749.000002",
    "user": "U023BECGF",
    "text": "Hello"
}
`

var message_normal = `
{
	"type": "message",
	"channel": "foo",
	"user": "bar",
	"text": "Hello world",
	"ts": "1355517523.000005"
}
`

var other = `
{"type":"presence_change","presence":"active","user":"U02G7HVHA"}
`

func TestMessage(t *testing.T) {
	var rtm = RTM{}
	expected := "bar#foo: Hello world"
	evt, _ := trimEvent(message_normal, &rtm)
	actual := evt.toString()
	if actual != expected {
		t.Errorf("got %s\nwant %s", actual, expected)
	}
}
