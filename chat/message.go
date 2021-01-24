package chat

import (
	"fmt"
	"github.com/kraxarn/website/user"
	"github.com/kraxarn/website/yt"
)

type Message struct {
	Type string `json:"type"`

	// message
	AvatarUrl string `json:"avatar_url,omitempty"`
	Message   string `json:"message,omitempty"`

	// video
	Audio string `json:"audio,omitempty"`
	Video string `json:"video,omitempty"`
}

func getAvatarUrl(sender *user.User) string {
	if sender == nil {
		// Default error icon
		return "/img/status/274c.svg"
	} else {
		return sender.AvatarPath()
	}
}

func getMessage(sender *user.User, message string) string {
	if sender == nil {
		return message
	} else {
		return fmt.Sprintf("%s: %s", sender.Name, message)
	}
}

func NewMessage(sender *user.User, message string) Message {
	return Message{
		Type:      "message",
		AvatarUrl: getAvatarUrl(sender),
		Message:   getMessage(sender, message),
	}
}

func NewVideoMessage(sender *user.User, videoId string) (Message, error) {
	videoInfo, err := yt.Info(videoId)
	if err != nil {
		return Message{}, err
	}

	return Message{
		Type:      "video",
		AvatarUrl: getAvatarUrl(sender),
		Audio:     videoInfo.Audio.Url,
		Video:     videoInfo.Video.Url,
	}, nil
}
