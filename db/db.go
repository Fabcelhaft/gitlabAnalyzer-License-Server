package db

import (
	"github.com/kikkirej/gitlab-analyzer/persistence"
	"gitlabAnalyzer-License-Server/db/rule_model"
	"gorm.io/gorm"
	"log"
	"sort"
	"strings"
)

var db = initDB()

func initDB() *gorm.DB {
	initDb := persistence.InitDb()
	errAutoMigrateMavenRule := initDb.AutoMigrate(&rule_model.MavenRule{})
	if errAutoMigrateMavenRule != nil {
		log.Fatalln("error while initializing to database:", errAutoMigrateMavenRule)
	}
	return initDb
}

func GetMavenRootElementNames() *[]string {
	return GetMavenChildElementNames(1, "")
}

func GetMavenChildElementNames(level uint, parent string) *[]string {
	var result []string
	db.Raw("select distinct split_part(group_id,'.', ?) as name from maven_dependencies where group_id like ? and split_part(group_id,'.', ?) != '' group by group_id, split_part(group_id,'.', ?)", level, parent+"%", level, level).Scan(&result)
	sort.Slice(result, func(i, j int) bool { return strings.ToLower(result[i]) < strings.ToLower(result[j]) })
	return &result
}

func GetArtifactsInGroup(group string) *[]string {
	var result []string
	db.Raw("select distinct artifact_id as name from maven_dependencies where group_id = ? order by artifact_id", group).Scan(&result)
	return &result
}
