package config

type ValidatorConfig struct {

}

type LoaderConfig struct {
	AppConfigFile  string
	Version        string
	PluginFilename string
	ObjectsDir     string
	AssetsDir      string
}

type PublisherConfig struct {
	AccessToken string
	URL         string
}
