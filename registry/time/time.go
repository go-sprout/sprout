package time

import "github.com/go-sprout/sprout"

type TimeRegistry struct {
	handler sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of conversion registry.
func NewRegistry() *TimeRegistry {
	return &TimeRegistry{}
}

// Uid returns the unique identifier of the registry.
func (tr *TimeRegistry) Uid() string {
	return "go-sprout/sprout.time"
}

// LinkHandler links the handler to the registry at runtime.
func (tr *TimeRegistry) LinkHandler(fh sprout.Handler) error {
	tr.handler = fh
	return nil
}

// RegisterFunctions registers all functions of the registry.
func (tr *TimeRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) error {
	sprout.AddFunction(funcsMap, "date", tr.Date)
	sprout.AddFunction(funcsMap, "dateInZone", tr.DateInZone)
	sprout.AddFunction(funcsMap, "duration", tr.Duration)
	sprout.AddFunction(funcsMap, "dateAgo", tr.DateAgo)
	sprout.AddFunction(funcsMap, "now", tr.Now)
	sprout.AddFunction(funcsMap, "unixEpoch", tr.UnixEpoch)
	sprout.AddFunction(funcsMap, "dateModify", tr.DateModify)
	sprout.AddFunction(funcsMap, "durationRound", tr.DurationRound)
	sprout.AddFunction(funcsMap, "htmlDate", tr.HtmlDate)
	sprout.AddFunction(funcsMap, "htmlDateInZone", tr.HtmlDateInZone)
	return nil
}

func (tr *TimeRegistry) RegisterAliases(aliasesMap sprout.FunctionAliasMap) error {
	sprout.AddAlias(aliasesMap, "dateModify", "mustDateModify")
	return nil
}

func (tr *TimeRegistry) RegisterNotices(notices *[]sprout.FunctionNotice) error {
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustDateModify", "please use `dateModify` instead"))
	return nil
}
