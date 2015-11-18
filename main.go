package main

import (
	"time"

	idl "go.iondynamics.net/iDlogger"
	"go.iondynamics.net/iDlogger/priority"
	"go.iondynamics.net/iDslackLog"

	"go.iondynamics.net/passPad/config"
	"go.iondynamics.net/passPad/filepath"
	"go.iondynamics.net/passPad/persistence"
	"go.iondynamics.net/passPad/server"
	"go.iondynamics.net/passPad/template"
)

func main() {
	defer func() {
		idl.Warn("shutdown")
		idl.Wait()
	}()
	if config.Std.Logger.SlackLogUrl != "" {

		prio := priority.Warning
		if config.Std.PassPad.DevelopmentEnv {
			prio = priority.Debugging
		}

		idl.AddHook(&iDslackLog.SlackLogHook{
			AcceptedPriorities: priority.Threshold(prio),
			HookURL:            config.Std.Logger.SlackLogUrl,
			IconURL:            "",
			Channel:            "",
			IconEmoji:          "",
			Username:           "PassPad Log",
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
			priority.Emergency,
			"Logger caught an internal error",
		})
		panic("Logger caught an internal error")
	})

	persistence.Init(filepath.GetPath("./bolt.db"))
	defer persistence.Close()

	err := template.Load()
	if err != nil {
		idl.Emerg("Couldn't load tempaltes: ", err)
	}

	server.Listen()
}
