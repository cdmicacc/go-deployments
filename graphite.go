package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type graphiteEvent struct {
	What string `json:"what"`
	Tags string `json:"tags"`
	Data string `json:"data"`
}

// curl -X POST "http://statsd.500px.net/events/" -d '{"what": "Deployment of indexer", "tags": "deploy,indexer", "data": "Deployment test"}'

func NotifyGraphite(n *Notification, graphiteUrl string) {
	evt := graphiteEvent{
		What: fmt.Sprintf("Deployment of %s", n.AppName),
		Tags: fmt.Sprintf("deployment %s", n.AppName),
		Data: graphiteMessage(n),
	}

	jsonBytes, _ := json.Marshal(evt)
	json := string(jsonBytes)

	if resp, err := http.Post(graphiteUrl, "application/json", strings.NewReader(json)); err != nil {
		log.Printf("Failed to notify graphite: %s", err.Error())
	} else {
		resp.Body.Close()
	}
}

func graphiteMessage(n *Notification) string {
	return fmt.Sprintf(
		"Deployment of <strong>%s</strong> by <strong>%s</strong><br>Revision <a href=\"%s\"><strong>%s</strong></a> (%s) deployed to <strong>%s</strong><br><ul><li>%s</li></ul>\n",
		n.AppName, n.UserName,
		n.CommitUrl, n.Commit, n.BranchName, n.TargetEnv,
		n.LastCommitMessage,
	)
}
