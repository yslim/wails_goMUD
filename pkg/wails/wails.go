package wails

import (
    "embed"
    "github.com/wailsapp/wails/v2"
    "github.com/wailsapp/wails/v2/pkg/options"
    "github.com/wailsapp/wails/v2/pkg/options/assetserver"
    "github.com/wailsapp/wails/v2/pkg/options/linux"
    "github.com/wailsapp/wails/v2/pkg/options/mac"
    "github.com/wailsapp/wails/v2/pkg/options/windows"
    "goMUD/pkg/config"
    "goMUD/pkg/core/service"
    "goMUD/pkg/log"
)

func run(assets embed.FS) {
    err := wails.Run(&options.App{
        Title:  config.Defaults.Application,
        Width:  1024,
        Height: 768,
        AssetServer: &assetserver.Options{
            Assets: assets,
        },
        BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
        OnStartup:        startup,
        OnDomReady:       domReady,
        OnBeforeClose:    beforeClose,
        OnShutdown:       shutdown,
        Bind: []interface{}{
            service.GetMudService(),
        },
        WindowStartState: options.Normal,
        Windows: &windows.Options{
            Theme:     windows.SystemDefault,
            OnSuspend: suspend,
            OnResume:  resume,
        },
        Mac:          &mac.Options{},
        Linux:        &linux.Options{},
        Experimental: &options.Experimental{},
    })

    if err != nil {
        log.Error("wails.Run, error=", err.Error())
    }
}
