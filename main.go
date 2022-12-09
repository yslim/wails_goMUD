package main

import (
    "embed"
    "github.com/yslim/go-util"
    "goMUD/pkg/config"
    "goMUD/pkg/log"
    "goMUD/pkg/wails"
)

//go:embed all:frontend/build
var assets embed.FS

func init() {
    log.Init(config.Log)
    log.Info("### active profile : %s", config.Profile.Active)
    log.Info("### logging level  : %s", config.Log.Level)
    log.Info("### IsDev = %v", config.IsDev)
    log.Info("### MUD Server = %v:%v", config.Mud.Server, config.Mud.Port)
    log.Info("### socket buffer-size=%s, %s Bytes",
        config.Socket.BufferSizeStr, util.RenderInteger(config.Socket.BufferSize))
}

func main() {
    wails.Run(assets)
}
