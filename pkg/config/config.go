package config

type Config struct {
	SlackToken     string
	DownloadToken  string
	RequestDelay   int64
	ChannelTypes   []string
	ExportBasePath string
	SplitMessages  bool
	ArchiveData    bool
	DownloadFiles  bool
	IncludeChannel []string
}
