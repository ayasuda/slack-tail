package main

import (
	"encoding/json"
	"fmt"
	"github.com/ttacon/chalk"
)

type Event struct {
	RTM  *RTM
	Type string `json:"type"`
	Message
}

// type AccountsChanged struct{}
// type BotAdded struct{}
// type BotChanged struct{}
// type ChannelArchive struct{}
// type ChannelCreated struct{}
// type ChannelDeleted struct{}
// type ChannelHistoryChanged struct{}
// type ChannelJoined struct{}
// type ChannelLeft struct{}
// type ChannelMarked struct{}
// type ChannelRename struct{}
// type ChannelUnarchive struct{}
// type CommandsChanged struct{}
// type DnDUpdated struct{}
// type DnDUpdatedUser struct{}
// type EmailDomainChanged struct{}
// type EmojiChanged struct{}
// type FileChange struct{}
// type FileCommentAdded struct{}
// type FileCommentDeleted struct{}
// type FileCommentEdited struct{}
// type FileCreated struct{}
// type FileDeleted struct{}
// type FilePulic struct{}
// type FileShared struct{}
// type FileUnshared struct{}
// type Goodbye struct{}
// type GroupArchive struct{}
// type GroupClose struct{}
// type GroupHistoryChanged struct{}
// type GroupJoined struct{}
// type GroupLeft struct{}
// type GroupMarked struct{}
// type GroupOpen struct{}
// type GroupRename struct{}
// type GroupUnarchive struct{}
// type Hello struct{}
// type IMClose struct{}
// type IMCreated struct{}
// type IMHistoryChanged struct{}
// type IMMarked struct{}
// type IMOpen struct{}
// type ManualPresenceChange struct{}

type Message struct {
	Channel string `json:"channel"`
	User    string `json:"user"`
	Text    string `json:"text"`
}

// type PinAdded struct{}
// type PinRemoved struct{}
// type PrefChange struct{}
// type PresenceChange struct{}
// type ReactionAdded struct{}
// type ReactionRemoved struct{}
// type ReconnectUrl struct{}
// type StarAdded struct{}
// type StarRemoved struct{}
// type SubteamCreated struct{}
// type SubteamSelfAdded struct{}
// type SubteamSelfRemoved struct{}
// type TeamDomainChange struct{}
// type TeamJoin struct{}
// type TeamMigrationStarted struct{}
// type TeamPlanChange struct{}
// type TeamPrefChange struct{}
// type TeamProfileChange struct{}
// type TeamProfileDelete struct{}
// type TeamProfileReorder struct{}
// type TeamRename struct{}
// type UserChange struct{}
// type UserTyping struct{}

func trimEvent(json_str string, rtm *RTM) (*Event, error) {
	var evt = Event{}
	evt.RTM = rtm
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
	rtm, ch_id, name_id, body := evt.RTM, evt.Channel, evt.User, evt.Text
	if body != "" {
		color := rtm.getUserColorById(name_id)
		ch := rtm.getChannelNameById(ch_id)
		name := rtm.getUserNameById(name_id)
		name = nameWithColor(name, color)
		return fmt.Sprintf("%s#%s: %s", name, ch, body)
	}
	return ""
}

func nameWithColor(name string, color string) string {
	if name != "" {
		cCode := int(name[0]) % 7
		switch cCode {
		case 0:
			return chalk.Red.Color(name)
		case 1:
			return chalk.Green.Color(name)
		case 2:
			return chalk.Yellow.Color(name)
		case 3:
			return chalk.Blue.Color(name)
		case 4:
			return chalk.Magenta.Color(name)
		case 5:
			return chalk.Cyan.Color(name)
		case 6:
			return chalk.White.Color(name)
		default:
			return chalk.Red.Color(name)
		}
	}
	return name
}
