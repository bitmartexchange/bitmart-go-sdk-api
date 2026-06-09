package bitmart

import (
	"bufio"
	"os"
	"strings"
)

// loadDotEnv reads the project's .env file (if present) and returns its
// key/value pairs. The .env file is NOT committed to git and is only used by
// the unit tests to inject real credentials. Lines that are blank or start with
// '#' are ignored; values may optionally be wrapped in matching quotes.
func loadDotEnv() map[string]string {
	env := make(map[string]string)

	file, err := os.Open(".env")
	if err != nil {
		// No .env file: tests fall back to placeholder credentials.
		return env
	}
	defer func() { _ = file.Close() }()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		idx := strings.Index(line, "=")
		if idx < 0 {
			continue
		}

		key := strings.TrimSpace(line[:idx])
		value := strings.TrimSpace(line[idx+1:])

		// Strip a single pair of surrounding quotes, if present.
		if len(value) >= 2 {
			if (value[0] == '"' && value[len(value)-1] == '"') ||
				(value[0] == '\'' && value[len(value)-1] == '\'') {
				value = value[1 : len(value)-1]
			}
		}

		if key != "" {
			env[key] = value
		}
	}

	return env
}

// envOr returns the value for key from env if it exists and is non-empty,
// otherwise it returns fallback.
func envOr(env map[string]string, key string, fallback string) string {
	if v, ok := env[key]; ok && v != "" {
		return v
	}
	return fallback
}

func GetDefaultConfig(url string) *Config {
	// Prefer credentials from the project's .env file (not committed to git).
	// If .env is absent or a key is missing, fall back to the placeholders so
	// that public-data tests keep working out of the box.
	env := loadDotEnv()

	var config Config
	config.Url = url
	config.ApiKey = envOr(env, "BITMART_API_KEY", "Your API KEY")
	config.SecretKey = envOr(env, "BITMART_SECRET_KEY", "Your Secret KEY")
	config.Memo = envOr(env, "BITMART_MEMO", "Your Memo")
	config.TimeoutSecond = 30
	config.Headers = map[string]string{
		"X-Custom-Header1": "HeaderValue1",
		"X-Custom-Header2": "HeaderValue2",
	}

	// Initialize another logger to output to the console with a log level of DEBUG
	config.CustomLogger = NewCustomLogger(DEBUG, os.Stdout)

	return &config
}

func NewTestClient() *CloudClient {
	// Spot REST base URL: prefer .env BITMART_API_URL, fall back to API_URL_PRO.
	env := loadDotEnv()
	return NewClient(*GetDefaultConfig(envOr(env, "BITMART_API_URL", API_URL_PRO)))
}

func NewTestFuturesClient() *CloudClient {
	// Futures REST base URL: prefer .env BITMART_API_URL, fall back to API_URL_V2_PRO.
	env := loadDotEnv()
	return NewClient(*GetDefaultConfig(envOr(env, "BITMART_API_URL", API_URL_V2_PRO)))
}
