package main

import (
	"flag"
	"os"
)

var (
	// HTTP server settings
	hipchatRoom  = flag.String("hipchat-room", "System Status", "HipChat room to notify")
	hipchatToken = flag.String("hipchat-token", "", "HipChat API v1 token")
	graphiteUrl  = flag.String("graphite-url", "", "Graphite URL to post events to")
	appName      = flag.String("appname", "", "Application name to display")
	commit       = flag.String("commit", "", "Git commit deployed")
	commitUrl    = flag.String("commit-url", "", "URL to git commit deployed")
	branchName   = flag.String("branch", "master", "Git branch deployed")
	userName     = flag.String("user", "", "User performing the deployment")
	targetEnv    = flag.String("target-env", "production", "Deployment target environment")
	urlTemplate  = flag.String("commit-url-template", "http://github.com/500px/%s/commit/%s", "URL template for commits")
)

/**
 * Set the value of the named flag with ENV(envName), if it exists
 */
func setFromEnv(flagName string, envName string) {
	if env := os.Getenv(envName); len(env) > 0 {
		flag := flag.Lookup(flagName)
		flag.Value.Set(env)
	}
}

func parseFlags() {
	// Set the flags from the env, then parse the command line (which will override the env)
	setFromEnv("hipchat-room", "DEPLOY_HIPCHAT_ROOM")
	setFromEnv("hipchat-token", "DEPLOY_HIPCHAT_TOKEN")
	setFromEnv("user", "USER")
	flag.Parse()
}
