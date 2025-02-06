package config

type Config struct {
	App     AppConfig      `yaml:"app"`
	Youtube PlatformConfig `yaml:"youtube"`
	Tiktok  PlatformConfig `yaml:"tiktok"`
	//Database   DatabaseConfig   `yaml:"databases"`
	Monitoring MonitoringConfig `yaml:"monitoring"`
	DBConfig   DB               `yaml:"databases"`
	Logging    LoggingConfig    `yaml:"logging"`
}

type LoggingConfig struct {
	LogLevel string `yaml:"log_level"`
}

type AppConfig struct {
	VideoCategory          string `yaml:"videoCategory"`
	DownloadPath           string `yaml:"downloadPath"`
	OutputPath             string `yaml:"outputPath"`
	AdText                 string `yaml:"adText"`
	AdImage                string `yaml:"adImage"`
	MaxConcurrentDownloads int    `yaml:"maxConcurrentDownloads"`
	MaxConcurrentUploads   int    `yaml:"maxConcurrentUploads"`
	VeedAPIKey             string `yaml:"VEEDKey"`
}

type PlatformConfig struct {
	Accounts []AccountConfig `yaml:"accounts"`
}

type AccountConfig struct {
	ApiKey   string `yaml:"apiKey"`
	Username string `yaml:"username"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbName"`
}

type MonitoringConfig struct {
	Enabled bool
}
type DB struct {
	Type     string
	Postgres DatabaseConfig
	Mongo    DatabaseConfig
}
