package main

import (
	"fmt"
	"github.com/andybons/hipchat"
	"log"
)

const (
	HipchatNotiferName = "Go Deployment"
)

func NotifyHipChat(n *Notification, room string, token string) {
	c := hipchat.Client{AuthToken: token}
	req := hipchat.MessageRequest{
		RoomId:        room,
		From:          HipchatNotiferName,
		Message:       hipchatMessage(n),
		Color:         hipchat.ColorPurple,
		MessageFormat: hipchat.FormatHTML,
		Notify:        true,
	}

	if err := c.PostMessage(req); err != nil {
		log.Printf("Failed to notify HipChat: %s", err.Error())
	}
}

func hipchatMessage(n *Notification) string {
	return fmt.Sprintf(
		"Deployment of <strong>%s</strong> by <strong>%s</strong><br>Revision <a href=\"%s\"><strong>%s</strong></a> (%s) deployed to <strong>%s</strong><br><ul><li>%s</li></ul>\n",
		n.AppName, n.UserName,
		n.CommitUrl, n.Commit, n.BranchName, n.TargetEnv,
		n.LastCommitMessage,
	)
}
