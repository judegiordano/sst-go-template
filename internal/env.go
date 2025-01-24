package internal

import (
	"strings"

	"github.com/judegiordano/gogetem/pkg/dotenv"
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
	Stage Stage `json:"stage"`
}

var Env Environment

func stage() Stage {
	stage := dotenv.String("STAGE")
	if stage == nil {
		return LocalStage
	}
	val := strings.ToUpper(*stage)
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
	Env = Environment{
		Stage: stage(),
	}
}
