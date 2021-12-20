package rule_handling

import (
	"github.com/fabcelhaft/gitlabAnalyzer-License-Server/db"
	"github.com/fabcelhaft/gitlabAnalyzer-License-Server/db/rule_model"
	"github.com/kikkirej/gitlab-analyzer/persistence/model"
)

var cache map[string]*rule_model.MavenRule

func UpdateMavenLicenses(){
	cache = make(map[string]*rule_model.MavenRule)
	dependencies := db.GetAllMavenDependencies()
	for _, dependency := range dependencies {
		rule := getRule(dependency.GroupID)
		if rule != nil{
			applyRule(&dependency, rule)
		}
	}
}

func getRule(groupID string) *rule_model.MavenRule {
	rule := cache[groupID]
	if rule != nil {
		return rule
	}
	rule = db.GetMavenRuleForGroupID(groupID)
	cache[groupID] = rule
	return rule
}

func applyRule(dependency *model.MavenDependency, rule *rule_model.MavenRule) {
	dependency.LicenseID = rule.LicenseID
	db.SaveMavenDependency(dependency)
}
