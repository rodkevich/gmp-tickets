package configs

type Configs struct {
	Api      ApiConfig
	Database DatabaseConfig
	// Elk      Elasticsearch
}

func New() *Configs {
	return &Configs{
		Api:      API(),
		Database: DataStore(),
		// Elasticsearch: ElasticSearch(),
	}
}
