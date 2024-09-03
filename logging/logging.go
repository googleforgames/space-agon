// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logging

import (
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	// Default logging configuration: json that plays nicely with Google Cloud Run
	logger = &logrus.Logger{
		Out: os.Stdout,
	}
)

func GetSharedLogger() *logrus.Logger {
	// Get previously setup logger.
	return logger
}

func NewSharedLogger(cfg *viper.Viper) *logrus.Logger {
	// Set up structured logging

	// Default logging configuration is json that plays nicely with Google Cloud Run.
	// This text formatter is much nicer for local debugging.
	logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	})
	// Configured for text formatter instead
	if cfg.IsSet("LOGGING_FORMAT") &&
		strings.ToLower(cfg.GetString("LOGGING_FORMAT")) == "text" {
		logger.Warnf("enabling text formatter")
		logger.SetFormatter(&logrus.TextFormatter{})
	}

	// Configured to include source line in logs
	if cfg.IsSet("LOG_CALLER") {
		logger.SetReportCaller(cfg.GetBool("LOG_CALLER"))
	}

	// Set minimum log level to output
	var err error
	level := logrus.InfoLevel
	if cfg.IsSet("LOGGING_LEVEL") {
		level, err = logrus.ParseLevel(cfg.GetString("LOGGING_LEVEL"))
		if err != nil {
			logger.Warn("Unable to parse LOGGING_LEVEL; defaulting to Info")
			level = logrus.InfoLevel
		}
	}
	logger.SetLevel(level)
	logger.Infof("Logging level set to %v", logger.Level)

	return logger
}
