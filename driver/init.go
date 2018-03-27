package driver

import (
	"github.com/om732/kamimai/core"
)

const versionTableName = "schema_version"

func init() {
	core.RegisterDriver("mysql", &MySQL{})
}
