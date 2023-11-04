package config

// Environment is used for listen settings
type Environment struct {
	Mode string
}

// API gin服务接口配置
type API struct {
	Host string
	Port string
}

// Log 日志配置
type Log struct {
	Level      string
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
}

// Mysql 数据库配置
type Mysql struct {
	Host         string
	Port         string
	Database     string
	User         string
	Password     string
	Timezone     string
	MaxOpenConns int
	MaxIdleConns int
	Debug        bool
}
