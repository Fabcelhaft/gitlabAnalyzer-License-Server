package db

import (
	"github.com/fabcelhaft/gitlabAnalyzer-License-Server/db/rule_model"
	"github.com/kikkirej/gitlab-analyzer/persistence"
	"github.com/kikkirej/gitlab-analyzer/persistence/model"
	"github.com/mitchellh/go-spdx"
	"gorm.io/gorm"
	"log"
)

var db = initDB()

func initDB() *gorm.DB {
	initDb := persistence.InitDb()
	errAutoMigrateMavenRule := initDb.AutoMigrate(&rule_model.MavenRule{})
	if errAutoMigrateMavenRule != nil {
		log.Fatalln("error while initializing to database:", errAutoMigrateMavenRule)
	}
	initSPDXLicenses(initDb)
	return initDb
}

func initSPDXLicenses(db *gorm.DB) {
	spdxLicenseList, err := spdx.List()
	if err != nil {
		return
	}
	for _, spdxLicense := range spdxLicenseList.Licenses {
		var license *model.License
		db.Where("license_id=?", &spdxLicense.ID).Find(&license)
		if license == nil || license.LicenseID== ""{
			license = &model.License{LicenseID: spdxLicense.ID}
		}
		license.Name=spdxLicense.Name
		license.Url="https://spdx.org/licenses/"
		license.Deprecated=spdxLicense.Deprecated
		license.OsiApproved=spdxLicense.OSIApproved
		license.Spdx=true
		db.Save(license)
	}
}
