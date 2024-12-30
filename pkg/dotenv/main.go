package dotenv

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/judegiordano/sst_template/pkg/logger"
)

func Load() {
	if err := godotenv.Load(); err != nil {
		logger.Fatal("error loading .env %v", err)
	}
}

func RequiredString(key string) string {
	var normalized string = strings.TrimSpace(strings.ToUpper(key))
	value, found := os.LookupEnv(normalized)
	if !found {
		logger.Fatal(fmt.Sprintf(".env %v not set", normalized))
	}
	return value
}

func RequiredInt(key string) int {
	var normalized string = strings.TrimSpace(strings.ToUpper(key))
	value, found := os.LookupEnv(normalized)
	if !found {
		logger.Fatal(fmt.Sprintf(".env %v not set", normalized))
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		logger.Fatal(fmt.Sprintf("cannot parse %v as int %v", value, err))
	}
	return i
}

func OptionalString(key string) *string {
	var normalized string = strings.TrimSpace(strings.ToUpper(key))
	value := os.Getenv(normalized)
	if len(strings.TrimSpace(value)) == 0 {
		return nil
	}
	return &value
}

func OptionalInt(key string) *int {
	var normalized string = strings.TrimSpace(strings.ToUpper(key))
	value := os.Getenv(normalized)
	if len(strings.TrimSpace(value)) == 0 {
		return nil
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot parse %v as int %v", value, err))
		return nil
	}
	return &i
}

func RequiredBool(key string) bool {
	var normalized string = strings.TrimSpace(strings.ToUpper(key))
	value, found := os.LookupEnv(normalized)
	if !found {
		logger.Fatal(fmt.Sprintf(".env %v not set", normalized))
	}
	b, err := strconv.ParseBool(value)
	if err != nil {
		logger.Fatal(fmt.Sprintf("cannot parse %v as bool %v", value, err))
	}
	return b
}

func OptionalBool(key string) *bool {
	var normalized string = strings.TrimSpace(strings.ToUpper(key))
	value := os.Getenv(normalized)
	if len(strings.TrimSpace(value)) == 0 {
		return nil
	}
	b, err := strconv.ParseBool(value)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot parse %v as bool %v", value, err))
		return nil
	}
	return &b
}
