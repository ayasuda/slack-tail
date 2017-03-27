package main

import (
	"testing"
)

var rtm_start_json = `
{
    "ok": true,
    "self": {
        "id": "U00000001",
        "name": "bobby",
        "prefs": { },
        "created": 1000000000,
        "manual_presence": "active"
    },
    "team": {
        "id": "TEAMID",
        "name": "TeamSlackTail",
        "email_domain": "example.com",
        "domain": "example",
        "prefs": { },
        "icon": { },
        "over_storage_limit": false,
        "plan": "std",
        "avatar_base_url": "https:\/\/ca.slack-edge.com\/"
    },
    "channels": [ ],
    "groups": [ ],
    "ims": [ ],
    "users": [
        {
            "id": "U00000001",
            "team_id": "TEAMID",
            "name": "bobby",
            "color": "43761b",
            "real_name": "Bobby"
        },
        {
            "id": "U00000002",
            "team_id": "TEAMID",
            "name": "alice",
            "color": "bc3663",
            "real_name": "Alice"
        }
    ],
    "bots": [ ],
    "url": "wss:\/\/some.url"
}
`

const rtm_start_fail_json = `
{
    "ok": false,
    "error": "invalid_auth"
}
`

func TestSlack(t *testing.T) {
	var rtm = RTM{}

	err := rtm.applyJSON([]byte(rtm_start_json))
	if err != nil {
		t.Error(err)
	}

	actual := rtm.getWSSUrl()
	expected := "wss://some.url"
	if actual != expected {
		t.Errorf("got %s\nwant %s", actual, expected)
	}

	actual = rtm.getUserNameById("U00000001")
	expected = "bobby"
	if actual != expected {
		t.Errorf("got %s\nwant %s", actual, expected)
	}
}

func TestSlackRtmStartFailed(t *testing.T) {
	var rtm = RTM{}
	err := rtm.applyJSON([]byte(rtm_start_fail_json))
	if err == nil {
		t.Error("no Error occured")
	}
}
