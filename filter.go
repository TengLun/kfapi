package kfapi

const (
	AccountId    = "accountId"
	AppId        = "appId"
	NetworkId    = "networkId"
	SiteId       = "siteId"
	TrackerId    = "trackerId"
	In           = "IN"
	NotIn        = "NOT IN"
	GreaterThan  = "greaterThan"
	LessThan     = "lessThan"
	InstallCount = "installct"
	ClickCount   = "clickct"
	Auto         = "auto"
)

func CreateFilter(dimension string, modifier string, values ...string) filter {

	filter := filter{
		Dimension: dimension,
		Modifier:  modifier,
		Values:    values,
	}

	return filter
}
