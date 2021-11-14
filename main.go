package main

import (
	"fmt"
	"github.com/kikkirej/gitlab-analyzer/settings"
	"gitlabAnalyzer-License-Server/db"
	"time"
)

func main() {
	for settings.Struct.DBInitialized == false {
		time.Sleep(time.Second)
	}
	fmt.Println(db.GetMavenRootElementNames())
	fmt.Println(db.GetMavenChildElementNames(3, "de.comline"))
	fmt.Println(db.GetArtifactsInGroup("de.comline"))
}
