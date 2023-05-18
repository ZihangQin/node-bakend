package config

type Config struct {
	Mysql *MysqlConfig `yaml:"mysql"`
}

type MysqlConfig struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Addr     string `yaml:"addr"`
	Db       string `yaml:"db"`
}
