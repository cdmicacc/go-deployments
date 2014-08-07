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

	notification.NotifyHipChat(*hipchatRoom, *hipchatToken)

}
