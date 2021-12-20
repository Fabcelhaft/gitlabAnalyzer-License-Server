package db

import (
	"sort"
	"strings"
)

func GetMavenRootElementNames() *[]string {
	return GetMavenChildElementNames(1, "")
}

func GetMavenChildElementNames(level uint, parent string) *[]string {
	var result []string
	db.Raw("select distinct split_part(group_id,'.', ?) as name from maven_dependencies where group_id like ? and split_part(group_id,'.', ?) != '' group by group_id, split_part(group_id,'.', ?)", level, parent+"%", level, level).Scan(&result)
	sort.Slice(result, func(i, j int) bool { return strings.ToLower(result[i]) < strings.ToLower(result[j]) })
	return &result
}

func GetMavenArtifactsInGroup(group string) *[]string {
	var result []string
	db.Raw("select distinct artifact_id as name from maven_dependencies where group_id = ? order by artifact_id", group).Scan(&result)
	return &result
}
