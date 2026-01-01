package configloader

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

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
