package config

type Config struct {
	App        AppConfig        `yaml:"app"`
	YtAccounts []YouTubeAccount `yaml:"yt_accounts"`
	Tiktok     PlatformConfig   `yaml:"tiktok"`
	Monitoring MonitoringConfig `yaml:"monitoring"`
	DBConfig   DBConfig         `yaml:"databases"`
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

type YouTubeAccount struct {
	ApiKey      string       `yaml:"apiKey"`
	Username    string       `yaml:"username"`
	Category    string       `yaml:"category"`
	Credentials *Credentials `yaml:"credentials,omitempty"`
}

type Credentials struct {
	Installed InstalledCredentials `yaml:"installed"`
}

type InstalledCredentials struct {
	ClientID            string   `yaml:"client_id"`
	ProjectID           string   `yaml:"project_id"`
	AuthURI             string   `yaml:"auth_uri"`
	TokenURI            string   `yaml:"token_uri"`
	AuthProviderCertURL string   `yaml:"auth_provider_x509_cert_url"`
	ClientSecret        string   `yaml:"client_secret"`
	RedirectURIs        []string `yaml:"redirect_uris"`
}

type AccountConfig struct {
	ApiKey   string `yaml:"apiKey"`
	Username string `yaml:"username"`
	Category string `yaml:"category"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbName"`
}

type MonitoringConfig struct {
	Enabled bool `yaml:"enabled"`
}

type DBConfig struct {
	Type     string         `yaml:"type"`
	Postgres DatabaseConfig `yaml:"postgres"`
	Mongo    DatabaseConfig `yaml:"mongo"`
}
