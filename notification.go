package main

import (
	"fmt"
	"github.com/andybons/hipchat"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

type Notification struct {
	AppName           string
	Commit            string
	CommitUrl         string
	BranchName        string
	UserName          string
	TargetEnv         string
	LastCommitMessage string
}

const (
	NotiferName = "Go Deployment"
)

var (
	originRegex = regexp.MustCompile("/([^/]*).git$")
)

func (n *Notification) NotifyHipChat(room string, token string) {
	c := hipchat.Client{AuthToken: token}
	req := hipchat.MessageRequest{
		RoomId:        room,
		From:          NotiferName,
		Message:       n.hipchatMessage(),
		Color:         hipchat.ColorPurple,
		MessageFormat: hipchat.FormatHTML,
		Notify:        true,
	}

	if err := c.PostMessage(req); err != nil {
		log.Printf("Failed to notify HipChat %q", err)
	}
}

func (n *Notification) hipchatMessage() string {
	return fmt.Sprintf(
		"Deployment of <strong>%s</strong> by <strong>%s</strong><br>Revision <a href=\"%s\"><strong>%s</strong></a> (%s) deployed to <strong>%s</strong><br><ul><li>%s</li></ul>\n",
		n.AppName, n.UserName,
		n.CommitUrl, n.Commit, n.BranchName, n.TargetEnv,
		n.LastCommitMessage,
	)
}

func (n *Notification) PopulateCommitInfo(commitUrlTemplate string) {
	cmd := exec.Command("git", "show", "-s", "--format=%an: %s", n.Commit)
	if output, err := cmd.Output(); err == nil {
		n.LastCommitMessage = strings.TrimSpace(string(output))
	} else {
		log.Fatalf("Error getting git commit message for commit %s: %s", n.Commit, err.Error())
	}

	cmd = exec.Command("git", "config", "--get", "remote.origin.url")
	if output, err := cmd.Output(); err == nil {
		origin := strings.TrimSpace(string(output))
		matches := originRegex.FindStringSubmatch(origin)
		if len(matches) >= 2 {
			if len(n.CommitUrl) == 0 {
				n.CommitUrl = fmt.Sprintf(commitUrlTemplate, matches[1], n.Commit)
			}
			if len(n.AppName) == 0 {
				n.AppName = matches[1]
			}
		} else {
			log.Fatalf("Error getting git repo name from origin '%s': %s", origin, err.Error())
		}
	} else {
		log.Fatalf("Error getting git origin for repository: %s", err.Error())
	}
}
