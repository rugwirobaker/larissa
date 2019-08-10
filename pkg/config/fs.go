package config

// DiskConfig specifies the properties required to use Disk as the storage backend
type DiskConfig struct {
	RootPath string `validate:"required" envconfig:"LARISSA_DISK_STORAGE_ROOT"`
}
