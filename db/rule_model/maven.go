package rule_model

import (
	"github.com/kikkirej/gitlab-analyzer/persistence/model"
	"gorm.io/gorm"
)

type MavenRule struct {
	gorm.Model
	GroupIDPrefix string
	Inheritance   bool
	License       *model.License `gorm:"foreignKey:LicenseID"`
	LicenseID     *uint
}
