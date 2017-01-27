// Copyright 2016 Marapongo, Inc. All rights reserved.

package testutil

import (
	"github.com/marapongo/mu/pkg/diag"
)

// TestDiagSink suppresses message output, but captures them, so that they can be compared to expected results.
type TestDiagSink struct {
	Pwd      string
	sink     diag.Sink
	errors   []string
	warnings []string
}

func NewTestDiagSink(pwd string) *TestDiagSink {
	return &TestDiagSink{Pwd: pwd, sink: diag.DefaultSink(pwd)}
}

func (d *TestDiagSink) Count() int {
	return d.Errors() + d.Warnings()
}

func (d *TestDiagSink) Errors() int {
	return len(d.errors)
}

func (d *TestDiagSink) ErrorMsgs() []string {
	return d.errors
}

func (d *TestDiagSink) Warnings() int {
	return len(d.warnings)
}

func (d *TestDiagSink) WarningMsgs() []string {
	return d.warnings
}

func (d *TestDiagSink) Success() bool {
	return d.Errors() == 0
}

func (d *TestDiagSink) Errorf(dia *diag.Diag, args ...interface{}) {
	d.errors = append(d.errors, d.Stringify(dia, diag.DefaultSinkErrorPrefix, args...))
}

func (d *TestDiagSink) Warningf(dia *diag.Diag, args ...interface{}) {
	d.warnings = append(d.warnings, d.Stringify(dia, diag.DefaultSinkWarningPrefix, args...))
}

func (d *TestDiagSink) Stringify(dia *diag.Diag, prefix string, args ...interface{}) string {
	return d.sink.Stringify(dia, prefix, args...)
}