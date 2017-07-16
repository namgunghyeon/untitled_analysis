package config

type Config struct {
  Couchbase CouchbaseConfig `toml:"couchbase_test"`
}

type CouchbaseConfig struct {
	Host string `toml:"host"`
	Project string `toml:"project"`
	ProjectInfo string `toml:"projectInfo"`
}
