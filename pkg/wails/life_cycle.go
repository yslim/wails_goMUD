package wails

import (
    "context"
    "goMUD/pkg/app"
    "goMUD/pkg/log"
)

// startup is called at application startup
func startup(ctx context.Context) {
    app.App().SetCtx(ctx)
    log.Info("WAILS START UP")
}

// domReady is called after the front-end dom has been loaded
func domReady(ctx context.Context) {
    app.App().SetCtx(ctx)
    log.Info("WAILS DOM READY")
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue,
// false will continue shutdown as normal.
func beforeClose(ctx context.Context) (prevent bool) {
    app.App().SetCtx(ctx)
    log.Info("WAILS BEFORE CLOSE")
    return false
}

// shutdown is called at application termination
func shutdown(ctx context.Context) {
    app.App().SetCtx(ctx)
    log.Info("WAILS SHUTDOWN")
}

// suspend is called when Windows enters low power mode
func suspend() {
    log.Info("WAILS SUSPEND")
}

// resume is called when Windows resumes from low power mode
func resume() {
    log.Info("WAILS RESUME")
}
