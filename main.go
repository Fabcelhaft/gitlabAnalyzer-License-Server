package main

import (
	"github.com/fabcelhaft/gitlabAnalyzer-License-Server/db"
	"github.com/fabcelhaft/gitlabAnalyzer-License-Server/rule_handling"
	"time"
)

func main() {
	//REST öffnen
	//Job zum refreshen starten
	time.Sleep(2* time.Second)
	db.GetMavenRuleForGroupID("org.apache.logging.log4j")
	rule_handling.UpdateMavenLicenses()
}
