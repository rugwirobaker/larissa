package app

import (
	"fmt"

	"github.com/rugwirobaker/larissa/pkg/config"
	"github.com/rugwirobaker/larissa/pkg/errors"
	"github.com/rugwirobaker/larissa/pkg/storage"
	"github.com/rugwirobaker/larissa/pkg/storage/fs"
	"github.com/spf13/afero"
)

//GetStorage returns as storage backend bases on configuration
func GetStorage(storageType string, storageConfig *config.StorageConfig) (storage.Backend, error) {
	const op errors.Op = "main.GetStorage"

	switch storageType {
	case "disk":
		if storageConfig.Disk == nil {
			return nil, errors.E(op, "Invalid Disk Storage Configuration")
		}
		rootLocation := storageConfig.Disk.RootPath
		s, err := fs.NewBackend(rootLocation, afero.NewOsFs())
		if err != nil {
			errStr := fmt.Sprintf("could not create new storage from os fs (%s)", err)
			return nil, errors.E(op, errStr)
		}
		return s, nil
	default:
		return nil, fmt.Errorf("storage type %s is unknown", storageType)
	}
}
