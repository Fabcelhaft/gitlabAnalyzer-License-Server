package db

import (
	"github.com/fabcelhaft/gitlabAnalyzer-License-Server/db/rule_model"
	"github.com/kikkirej/gitlab-analyzer/persistence/model"
	"strings"
)

func GetMavenRuleForGroupID(groupid string) *rule_model.MavenRule {
	var rule *rule_model.MavenRule
	db.Where("group_id_prefix = ?", groupid).Find(&rule)
	if rule==nil || rule.ID == 0{
		var upperGroupID = getUpperMavenGroupID(groupid)
		if upperGroupID != ""{
			return GetMavenRuleForGroupID(upperGroupID)
		}
	}
	if rule.ID==0{
		return nil
	}
	return rule
}

func getUpperMavenGroupID(groupid string) string {
	splittedGroupID := strings.Split(groupid, ".")
	var result = ""
	for i := 0; i < (len(splittedGroupID)-1); i++ {
		result += splittedGroupID[i]
		if i <(len(splittedGroupID)-2) {
			result += "."
		}
	}
	return result
}

func GetAllMavenDependencies() []model.MavenDependency{
	var dependencies []model.MavenDependency
	db.Find(&dependencies)
	return dependencies
}

func SaveMavenDependency(dependency *model.MavenDependency) {
	if dependency.LicenseID != nil {
		db.Where("id = ?", dependency.LicenseID).Find(&dependency.License)
	}
	db.Save(dependency)
}