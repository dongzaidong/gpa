package config

type Config struct {
}

func Load(path string) (*Config, error) {
	con := &Config{}
	err := con.LoadSettings(path)
	return con, err
}

//LoadSettings load configuration parameters from a file
func (c *Config) LoadSettings(path string) error {
	return nil
}

//CheckSettings check the validity of configuration parameters
func (c *Config) CheckSettings() error {
	return nil
}

/**
Examples：针对ymal文件

type Config struct {
	Log      LogConfig      `yaml:"*"`
	DataBase DataBaseConfig `yaml:"*"`
	Services ServicesConfig `yaml:"*"`
	Entrance EntranceConfig `yaml:"*"`
}

func (c *Config) LoadSettings(path string) error {
	cfgdata, err := ioutil.ReadFile(path)
	if err != nil {
	}
	if err = yaml.Unmarshal(cfgdata, c); err != nil {
	}
	return nil
}

func (c *Config) CheckSettings() error {
	if c.Log == nil {
		return errors.New("日志配置不能为空")
	}
	....
	return nil
}
*/
