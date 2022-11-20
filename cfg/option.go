package cfg

type FileType string

const (
	FileTypeYaml = "yaml"
	FileTypeYml  = "yml"
)

var defaultOpt = Option{
	ConfigList:     make([]string, 0),
	NeedDatasource: true,
	Path:           "./conf",
	FileType:       FileTypeYml,
}

type Option struct {
	NeedDatasource bool
	Path           string
	FileType       FileType
	ConfigList     []string
}

func NewDefaultOption() Option {
	return defaultOpt
}

type OptionFn func(opt *Option)

func WithConfigList(configList []string) OptionFn {
	return func(opt *Option) {
		opt.ConfigList = configList
	}
}

func WithPath(path string) OptionFn {
	return func(opt *Option) {
		opt.Path = path
	}
}

func WithoutDatasource() OptionFn {
	return func(opt *Option) {
		opt.NeedDatasource = false
	}
}

func WithFileType(fileType FileType) OptionFn {
	return func(opt *Option) {
		opt.FileType = fileType
	}
}

func NewOption(optFn ...OptionFn) Option {
	opt := NewDefaultOption()
	for _, fn := range optFn {
		fn(&opt)
	}
	return opt
}
