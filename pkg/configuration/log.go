package configuration

/*
Log 运行调试日志相关配置

示例：
filename: logs/app_api.log
level: debug
format: json
stack: true
maxsize: 100
maxage: 7
maxbackups: 10
localtime: 2006-01-02 15:04:05
compress: false
*/
type Log struct {
	Filename string `mapstructure:"filename" validate:"-"`
	Level    string `mapstructure:"level" validate:"required,eq=debug|eq=info|eq=warn|eq=error|eq=fatal|eq=panic"`
	Format   string `mapstructure:"format" validate:"required,eq=text|eq=json"`

	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int `mapstructure:"maxsize" validate:"-"`

	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int `mapstructure:"maxage" validate:"-"`

	// MaxBackups is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.)
	MaxBackups int `mapstructure:"maxbackups" validate:"-"`

	// LocalTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time.  The default is to use UTC
	// time.
	LocalTime bool `mapstructure:"localtime" validate:"-"`

	// Compress determines if the rotated log files should be compressed
	// using gzip. The default is not to perform compression.
	Compress bool `mapstructure:"compress" validate:"-"`

	// Analyzing the call stack information to get the code location of
	// the print log may not be accurate, but it can quickly locate the problem
	Stack bool `mapstructure:"stack" validate:"-"`
}

type RecordLog struct {
	Filename string `mapstructure:"filename" validate:"required"`
	// refer to time.ParseDuration
	RotateInterval string `mapstructure:"rotate_interval" validate:"required,duration"`
}

type SqlLog struct {
	Filename       string `mapstructure:"filename" validate:"-"`
	RotateInterval string `mapstructure:"rotate_interval" validate:"required,duration"`
	Level          string `mapstructure:"level" validate:"required,eq=debug|eq=info|eq=warn|eq=error|eq=fatal|eq=panic"`
	Format         string `mapstructure:"format" validate:"required,eq=text|eq=json"`

	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int `mapstructure:"maxsize" validate:"-"`

	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int `mapstructure:"maxage" validate:"-"`

	// refer to time.ParseDuration
	ExecTime string `mapstructure:"exec_time" validate:"required,duration"`
}
