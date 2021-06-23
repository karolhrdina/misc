package env

import (
    "os"
)

// GetVar returns the value of the `key` environment variable if it's set,
// fallback value otherwise.
func GetVar(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}
