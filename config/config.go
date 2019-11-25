package config

const (
	// Parser names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ProfileParser = "ProfileParser"
	NilParser     = "NilParser"

	// Service ports
	ItemSaverPort = 12345
	WorkerPort0   = 9000

	// ElasticSearch
	ESIndex = "dating_profile"

	// RPC endpoints
	ItemSaverRpc     = "ItemSaverService.Save"
	SpiderServiceRpc = "SpiderService.Process"
)
