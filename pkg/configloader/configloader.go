package configloader

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// Load loads configuration into cfg from the given filePath or, if the file
// does not exist, from environment variables.
//
// It first checks the existence of filePath. If the file exists, it reads
// configuration using cleanenv.ReadConfig(filePath, cfg). If the file does not
// exist, it falls back to cleanenv.ReadEnv(cfg). cfg is expected to be a
// pointer to a struct compatible with cleanenv's tags. Any error encountered
// is returned wrapped to provide context ("failed to load config: ...").
func Load(cfg any, filePath string) error {
	_, err := os.Stat(filePath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if os.IsNotExist(err) {
		err := cleanenv.ReadEnv(cfg)
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}
	} else {
		err := cleanenv.ReadConfig(filePath, cfg)
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}
	}

	return nil
}
