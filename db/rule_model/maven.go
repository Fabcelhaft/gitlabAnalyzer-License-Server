package rule_model

import (
	"github.com/kikkirej/gitlab-analyzer/persistence/model"
	"gorm.io/gorm"
)

type MavenRule struct {
	gorm.Model
	groupIDPrefix string
	inheritance   bool
	License       *model.License `gorm:"foreignKey:LicenseID"`
	LicenseID     *uint
}
