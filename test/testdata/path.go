package testdata

import (
	"path/filepath"
	"runtime"
)

var (
	Postgres = filepath.Join("postgres", "ddl.sql")
	Redis    = filepath.Join("redis", "redis.conf")
)

func init() {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return
	}
	dir := filepath.Dir(file)
	Postgres = filepath.Join(dir, "postgres", "ddl.sql")
	Redis = filepath.Join(dir, "redis", "redis.conf")
}
