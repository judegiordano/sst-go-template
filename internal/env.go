package env

type Stage = string

type Env struct {
	LogLevel string `json:"log_level"`
	Stage    Stage  `json:"stage"`
}

var env Env

func init() {
	env = Env{
		Stage: Stage("sdf"),
	}
}
