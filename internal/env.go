package internal

import (
	"strings"

	"github.com/charmbracelet/log"
	"github.com/judegiordano/sst_template/pkg/dotenv"
)

type Stage string

const (
	LocalStage Stage = "LOCAL"
	DevStage   Stage = "DEV"
	ProdStage  Stage = "PROD"
)

func (c Stage) String() string {
	switch c {
	case LocalStage:
		return "LOCAL"
	case ProdStage:
		return "PROD"
	default:
		return "DEV"
	}
}

type Environment struct {
	LogLevel log.Level `json:"log_level"`
	Stage    Stage     `json:"stage"`
}

var Env Environment

func logLevel() log.Level {
	switch strings.ToUpper(dotenv.RequiredString("LOG_LEVEL")) {
	case "DEBUG":
		return log.DebugLevel
	case "WARN":
		return log.WarnLevel
	case "INFO":
		return log.InfoLevel
	default:
		return log.ErrorLevel
	}
}

func stage() Stage {
	val := strings.ToUpper(dotenv.RequiredString("STAGE"))
	switch val {
	case "LOCAL":
		return LocalStage
	case "PROD":
		return ProdStage
	case "DEV":
		return DevStage
	default:
		return Stage(val)
	}
}

func init() {
	dotenv.Load()
	Env = Environment{
		LogLevel: logLevel(),
		Stage:    stage(),
	}
}
