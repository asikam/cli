package query

import (
	"github.com/ignite/cli/v28/ignite/pkg/multiformatname"
	"github.com/ignite/cli/v28/ignite/templates/field"
)

// Options ...
type Options struct {
	AppName     string
	AppPath     string
	ModuleName  string
	ModulePath  string
	QueryName   multiformatname.Name
	Description string
	ResFields   field.Fields
	ReqFields   field.Fields
	Paginated   bool
}
