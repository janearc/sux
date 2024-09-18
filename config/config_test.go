package config_test

import (
	"github.com/janearc/sux/config"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// TODO: this should be more useful but for now this works
	root := os.Getenv("SUX_ROOT")

	assert.NotNil(t, root)
	assert.DirExistsf(t, root, "root directory '%s' does not exist", root)

	configFileName := filepath.Join(root, "config/config.yml")
	versionFileName := filepath.Join(root, "config/version.yml")
	secretsFileName := filepath.Join(root, "config/secrets.yml")

	// fire up that constructor and read some yaml
	cfg, err := config.LoadConfig(configFileName, versionFileName, secretsFileName)

	assert.Nil(t, err)
	assert.NotNil(t, cfg)
}
