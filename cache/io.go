package cache

import (
	"fmt"
	"os"
	"path/filepath"
)

func getCacheFilePath() string {
	dir, _ := os.UserConfigDir()
	return filepath.Join(dir, "vpmobil")
}

func EnsureCacheDirExists() error {
	cachePath := getCacheFilePath()
	_, err := os.Stat(cachePath)
	fmt.Print(cachePath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(cachePath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteCacheDir() error {
	cachePath := getCacheFilePath()
	return os.RemoveAll(cachePath)
}

func EnsureSchoolCacheDir(schoolID string) error {
	cachePath := getCacheFilePath()
	schoolCachePath := filepath.Join(cachePath, schoolID)
	_, err := os.Stat(schoolCachePath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(schoolCachePath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

func ReadSchoolCacheDir(schoolID string) ([]os.DirEntry, error) {
	cachePath := getCacheFilePath()
	schoolCachePath := filepath.Join(cachePath, schoolID)
	return os.ReadDir(schoolCachePath)
}

func SchoolCacheDirExists(schoolID string) bool {
	cachePath := getCacheFilePath()
	schoolCachePath := filepath.Join(cachePath, schoolID)
	_, err := os.Stat(schoolCachePath)
	return !os.IsNotExist(err)
}

func DeleteSchoolCacheDir(schoolID string) error {
	cachePath := getCacheFilePath()
	schoolCachePath := filepath.Join(cachePath, schoolID)
	return os.RemoveAll(schoolCachePath)
}

func CacheFileExists(relativePath string) bool {
	cachePath := getCacheFilePath()
	fullPath := filepath.Join(cachePath, relativePath+".tmp")
	_, err := os.Stat(fullPath)
	return !os.IsNotExist(err)
}

func WriteCacheFile(relativePath string, data []byte) error {
	cachePath := getCacheFilePath()
	fullPath := filepath.Join(cachePath, relativePath+".tmp")
	return os.WriteFile(fullPath, data, 0600)
}

func DeleteCacheFile(relativePath string) error {
	cachePath := getCacheFilePath()
	fullPath := filepath.Join(cachePath, relativePath+".tmp")
	return os.Remove(fullPath)
}

func ReadCacheFile(relativePath string) ([]byte, error) {
	cachePath := getCacheFilePath()
	fullPath := filepath.Join(cachePath, relativePath+".tmp")
	return os.ReadFile(fullPath)
}
