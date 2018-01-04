package confmgr


type WeChatSettings struct {
	// POST only
	Enable			bool		`yaml:"enable"`
	Url				string		`yaml:"url"`
	Header			string		`yaml:"header"`
}

type BasicSettings struct {
	EnableWebUi	bool		`yaml:"enableWeb"`
	ListenPort	string		`yaml:"listen"`
}

type SystemSettings struct {
	CpuLimit		float32		`yaml:"cpuLimit"`
	MemLimit		float32		`yaml:"memLimit"`
	DiskLimit		float32		`yaml:"diskLimit"`
	FileNumLimit	int 		`yaml:"fileNumLimit"`
	SystemLoad		float32		`yaml:"systemLoad"`
}

type ProcessSettings struct {
	// to identify the program's pid by its name
	Name			string		`yaml:"name"`
	// the following is the metric to alarm
	CpuLimit		int			`yaml:"cpuLimit"`
	MemLimit		int			`yaml:"memLimit"`
	ConnLimit		int			`yaml:"connLimit"`
	FileLimit		int			`yaml:"fileLimit"`
	ThreadLimit		int			`yaml:"threadLimit"`
	IsRun			bool		`yaml:"isRun"`
	Exists			bool		`yaml:"exists"`
}

type Config struct {
	Basic 		BasicSettings 		`yaml:"basic"`
	Process		[]ProcessSettings	`yaml:"processes"`
	SystemConf	SystemSettings		`yaml:"system"`
	WechatConf	WeChatSettings		`yaml:"wechat"`
	CleanerConf	FileCleanerSettings	`yaml:"cleaner"`
}

type FileCleanerSettings struct {
	Enabled			bool			`yaml:"enable"`
	FileDir			string			`yaml:"dir"`
	// only support the following format: 1h 1d
	Time			string			`yaml:"time"`
}