// Package log
// Created by yslim on 2020. 09. 14
// Copyright (c) 2020. All rights reserved.
//
// Description
//
//	logging
package log

import (
	logger "github.com/yslim/go-logger"
)

type Config struct {
	Path          string `properties:"file-path"`
	FileLimitSize int    `properties:"file-limit-size"`
	NumFiles      int    `properties:"num-files"`
	Level         string `properties:"level"`
}

var (
	log *logger.Logger
)

func Init(config *Config) {
	lvl := logger.GetLevelByName(config.Level)
	fs := config.FileLimitSize * 1024 * 1024

	log = logger.InitLogger(lvl, int64(fs), config.NumFiles, config.Path, logger.RollSize, true, true)

	log.Info("logging.level = %v", logger.LogLevelName[lvl])
	log.Info("logging.path = %v", config.Path)
	log.Info("logging.file-limit-size = %d MB", config.FileLimitSize)
	log.Info("logging.num-files = %d", config.NumFiles)

	log.SetCallDepth(3)
}

func IsEnabled(lvl logger.LogLevel) bool {
	return log.IsEnabled(lvl)
}

func Trace(format string, v ...interface{}) {
	log.Tracef(format, v...)
}

func Debug(format string, v ...interface{}) {
	log.Debugf(format, v...)
}

func Info(format string, v ...interface{}) {
	log.Infof(format, v...)
}

func Warn(format string, v ...interface{}) {
	log.Warnf(format, v...)
}

func Error(format string, v ...interface{}) {
	log.Errorf(format, v...)
}

func Fatal(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}
