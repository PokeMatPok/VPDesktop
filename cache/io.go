package cache

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func getCacheFilePath() string {
	dir, err := os.UserCacheDir()
	if err != nil {
		// last-resort fallback
		return filepath.Join(os.TempDir(), "vpmobil")
	}
	return filepath.Join(dir, "vpmobil")
}

func EnsureCacheDirExists() error {
	cachePath := getCacheFilePath()
	_, err := os.Stat(cachePath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(cachePath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func ClearAllCache() error {
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

func ListSchoolCacheFiles(schoolID string) ([]os.DirEntry, error) {
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

func HasCacheFile(relativePath string) bool {
	cachePath := getCacheFilePath()
	fullPath := filepath.Join(cachePath, relativePath+".cache")
	_, err := os.Stat(fullPath)
	return !os.IsNotExist(err)
}

func WriteCacheFile(relativePath string, data []byte) error {
	cachePath := getCacheFilePath()
	fullPath := filepath.Join(cachePath, relativePath+".cache")

	tmp := fullPath + ".writing"
	if err := os.WriteFile(tmp, data, 0o600); err != nil {
		return err
	}

	return os.Rename(tmp, fullPath)
}

func DeleteCacheFile(relativePath string) error {
	cachePath := getCacheFilePath()
	fullPath := filepath.Join(cachePath, relativePath+".cache")
	return os.Remove(fullPath)
}

func ReadCacheFile(relativePath string) ([]byte, error) {
	cachePath := getCacheFilePath()
	fullPath := filepath.Join(cachePath, relativePath+".cache")
	return os.ReadFile(fullPath)
}

func ReadJSONCacheFile[T any](relativePath string) (T, error) {
	var result T
	data, err := ReadCacheFile(relativePath)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

func WriteJSONCacheFile[T any](relativePath string, data T) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return WriteCacheFile(relativePath, bytes)
}
