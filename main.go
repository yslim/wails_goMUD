package main

import (
    "context"
    "embed"
    "github.com/wailsapp/wails/v2"
    "github.com/wailsapp/wails/v2/pkg/options"
    "github.com/wailsapp/wails/v2/pkg/options/assetserver"
    "github.com/wailsapp/wails/v2/pkg/options/linux"
    "github.com/wailsapp/wails/v2/pkg/options/mac"
    "github.com/wailsapp/wails/v2/pkg/options/windows"
    "github.com/wailsapp/wails/v2/pkg/runtime"
    "github.com/yslim/go-util"
    "goMUD/pkg/app"
    "goMUD/pkg/config"
    "goMUD/pkg/core/service"
    "goMUD/pkg/log"
)

//go:embed all:frontend/dist
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
    err := wails.Run(&options.App{
        Title:  config.Defaults.Application,
        Width:  1024,
        Height: 768,
        AssetServer: &assetserver.Options{
            Assets: assets,
        },
        BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
        OnStartup: func(ctx context.Context) {
            app.App().SetCtx(ctx)
            log.Info("WAILS START UP")
            runtime.EventsOn(ctx, "OnReady", func(optionalData ...interface{}) {
                err := service.GetMudService().Connect("KanghoMurim")
                if err != nil {
                    log.Error("MUD connect error = ", err)
                }
            })
        },
        OnDomReady: func(ctx context.Context) {
            app.App().SetCtx(ctx)
            log.Info("WAILS DOM READY")
        },
        OnBeforeClose: func(ctx context.Context) (prevent bool) {
            app.App().SetCtx(ctx)
            log.Info("WAILS BEFORE CLOSE")
            return false
        },
        OnShutdown: func(ctx context.Context) {
            app.App().SetCtx(ctx)
            log.Info("WAILS SHUTDOWN")
        },
        Bind: []interface{}{
            service.GetMudService(),
        },
        WindowStartState: options.Normal,
        Windows: &windows.Options{
            Theme: windows.SystemDefault,
            OnSuspend: func() {
                log.Info("WAILS SUSPEND")
            },
            OnResume: func() {
                log.Info("WAILS RESUME")
            },
        },
        Mac:          &mac.Options{},
        Linux:        &linux.Options{},
        Experimental: &options.Experimental{},
    })

    if err != nil {
        log.Error("wails.Run, error=", err.Error())
    }
}
