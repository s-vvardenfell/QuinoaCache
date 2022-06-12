package config

type Config struct {
	Host           string `mapstructure:"host"`
	RedisPort      string `mapstructure:"redis_port"`
	ServerPort     string `mapstructure:"server_port"`
	Passw          string `mapstructure:"password"`
	DataBaseNum    int    `mapstructure:"db_num"`
	WithReflection bool   `mapstructure:"with_reflection"`
	Logrus         Logrus `mapstructure:"logrus"`
}

type Logrus struct {
	LogLvl int    `mapstructure:"log_level"`
	ToFile bool   `mapstructure:"to_file"`
	ToJson bool   `mapstructure:"to_json"`
	LogDir string `mapstructure:"log_dir"`
}
