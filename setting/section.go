package setting

type Config struct {
	Mysql  MySqlSetting `mapstructure:"mysql"`
	Server Server       `mapstructure:"server"`
	Logger Logger       `mapstructure:"log"`
}

type MySqlSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	Database        string `mapstructure:"database"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
}

type Server struct {
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type Logger struct {
	Level string `mapstructure:"level"`
	Env   string `mapstructure:"env"`
}
