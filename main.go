package main

import (
	"log"
)

func main() {
	parseFlags()

	if len(*commit) == 0 {
		log.Fatalf("-commit option is required")
	}
	if len(*userName) == 0 {
		log.Fatalf("-user option is required")
	}
	if len(*targetEnv) == 0 {
		log.Fatalf("-target-env option is required")
	}

	notification := &Notification{
		Commit:     *commit,
		CommitUrl:  *commitUrl,
		UserName:   *userName,
		AppName:    *appName,
		TargetEnv:  *targetEnv,
		BranchName: *branchName,
	}
	notification.PopulateCommitInfo(*urlTemplate)

	if len(*hipchatToken) > 0 {
		NotifyHipChat(notification, *hipchatRoom, *hipchatToken)
	}
	if len(*graphiteUrl) > 0 {
		NotifyGraphite(notification, *graphiteUrl)
	}

	log.Printf("Notified of deployment for %s #%s", notification.AppName, notification.Commit)

}
