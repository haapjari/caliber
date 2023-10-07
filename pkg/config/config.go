package config

type Config struct{}

func New() *Config {
	return &Config{}
}

// TODO
func (c *Config) SetDefaults() {}

// TODO
func (c *Config) Load() error {
	c.SetDefaults()
	return nil
}

// TODO
func (c *Config) Initialize() error {
	return c.Load()
}
