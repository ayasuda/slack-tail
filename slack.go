package main

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const RTM_ENDPOINT string = "https://slack.com/api/rtm.start"

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Color    string `json:"color"`
	RealName string `json:"real_name"`
}

type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type RTM struct {
	Ok       bool      `json:"ok"`    // TODO: separate interface
	Error    string    `json:"error"` // TODO: separate interface
	Url      string    `json:"url"`
	Users    []User    `json:"users"`
	Channels []Channel `json:"channels"`
}

func newRTM(token string) (*RTM, error) {
	c := &http.Client{}
	resp, err := c.PostForm(RTM_ENDPOINT, url.Values{"token": {token}})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rtm = RTM{}
	err = rtm.applyJSON(b)
	if err != nil {
		return nil, errors.Wrap(err, "failed to start rtm")
	}

	return &rtm, nil
}

func (rtm *RTM) applyJSON(b []byte) error {
	err := json.Unmarshal(b, &rtm)
	if err != nil {
		return err
	}

	// if slack api returns fail response, return error.
	// cf, { "ok" false, "error": "invalid_auth" }
	if rtm.Ok == false {
		if rtm.Error != "" {
			return errors.New(rtm.Error)
		} else {
			return errors.New("Slack API returns fail")
		}
	}

	return nil
}

func (rtm *RTM) getWSSUrl() string {
	return rtm.Url
}

func (rtm *RTM) getUserNameById(id string) string {
	name := id
	for i := range rtm.Users {
		if rtm.Users[i].Id == id {
			name = rtm.Users[i].Name
		}
	}
	return name
}

func (rtm *RTM) getUserColorById(id string) string {
	return ""
}

// func getGroupNameById(repo *dataRepo, id string) {}
func (rtm *RTM) getChannelNameById(id string) string {
	name := id
	for i := range rtm.Channels {
		if rtm.Channels[i].Id == id {
			name = rtm.Channels[i].Name
		}
	}
	return name
}

// func getGroupsNameById(repo *dataRepo, id string) {}
// func getMPIMNameById(repo *dataRepo, id string) {}
// func getIMNameById(repo *dataRepo, id string) {}
// func getBotNameById(repo *dataRepo, id string) {}
