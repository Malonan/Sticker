package libF

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

// Global koanf instance. Use "." as the key path delimiter. This can be "/" or any character.
var k = koanf.New(".")

func F() *koanf.Koanf {
	// Load YAML config and merge into the previously loaded config (because we can).
	k.Load(file.Provider("sticker.yml"), yaml.Parser())
	return k
}
