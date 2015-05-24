package v1

import (
	"time"

	idl "go.iondynamics.net/iDlogger"
	"go.iondynamics.net/passPad/v1/config"
	"go.iondynamics.net/passPad/v1/filepath"
	"go.iondynamics.net/passPad/v1/passpad/persistence"
	"go.iondynamics.net/passPad/v1/server"
	"go.iondynamics.net/passPad/v1/template"

	"go.iondynamics.net/iDslackLog"
)

func Run() {
	defer func() {
		idl.Warn("shutdown")
		idl.Wait()
	}()
	if config.Std.Logger.SlackLogUrl != "" {

		level := idl.WarnLevel
		if config.Std.PassPad.DevelopmentEnv {
			level = idl.DebugLevel
		}

		idl.AddHook(&iDslackLog.SlackLogHook{
			AcceptedLevels: iDslackLog.LevelThreshold(level),
			HookURL:        config.Std.Logger.SlackLogUrl,
			IconURL:        "",
			Channel:        "",
			IconEmoji:      "",
			Username:       "PassPad Log",
		})
	}
	idl.StandardLogger().Async = true
	idl.SetPrefix("PassPad")
	idl.SetErrCallback(func(err error) {
		idl.StandardLogger().Async = true
		idl.Log(&idl.Event{
			idl.StandardLogger(),
			map[string]interface{}{
				"error": err,
			},
			time.Now(),
			idl.PanicLevel,
			"Logger caught an internal error",
		})
		panic("Logger caught an internal error")
	})

	template.Load(filepath.GetPath("./templates/*.tpl"))

	persistence.Init(filepath.GetPath("./bolt.db"))
	defer persistence.Close()

	server.Listen()
}
