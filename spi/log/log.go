/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package log

// Level is a log level for a logging message.
type Level int

// Log levels.
const (
	CRITICAL Level = iota
	ERROR
	WARNING
	INFO
	DEBUG
)

// Logger represents a general-purpose logger.
type Logger interface {
	Panicf(msg string, args ...any)
	Fatalf(msg string, args ...any)
	Errorf(msg string, args ...any)
	Warnf(msg string, args ...any)
	Infof(msg string, args ...any)
	Debugf(msg string, args ...any)
}

// LoggerProvider is a factory for moduled loggers.
type LoggerProvider interface {
	GetLogger(module string) Logger
}
