package main

import (
	"fmt"
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

var (
	originRegex = regexp.MustCompile("/([^/]*?)(?:.git)?$")
)

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
			log.Fatalf("Error getting git repo name from origin '%s'", origin)
		}
	} else {
		log.Fatalf("Error getting git origin for repository: %s", err.Error())
	}
}
