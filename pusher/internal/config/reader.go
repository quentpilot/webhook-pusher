package config

type EnvParser interface {
	Parse()
}

type EnvConfig interface {
	Get(key string)
	Has(key string)
}

type YamlEnvReader struct {
	Parameters map[string]any
}

func (c *YamlEnvReader) Read(filename string) []byte {
	return []byte{}
}

func (c *YamlEnvReader) Parse(filename string) {
	// read file and save key value to c.Parameters
}

func (c *YamlEnvReader) Get(key string) (any, error) {
	if val, ok := c.Parameters[key]; ok {
		return val, nil
	}
	return "", nil
}

func (c *YamlEnvReader) Has(key string) bool {
	_, ok := c.Parameters[key]
	return ok
}
