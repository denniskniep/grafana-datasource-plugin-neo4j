//+build mage

package main

import (
	"os"
	"path/filepath"

	"github.com/magefile/mage/sh"
	// mage:import
	build "github.com/grafana/grafana-plugin-sdk-go/build"
)

// Test runs backend tests.
func TestShort() error {
	return sh.RunV("go", "test", "./pkg/...", "-test.short")
}

// Coverage runs backend tests and makes a coverage report.
func CoverageShort() error {
	// Create a coverage file if it does not already exist
	if err := os.MkdirAll(filepath.Join(".", "coverage"), os.ModePerm); err != nil {
		return err
	}

	if err := sh.RunV("go", "test", "./pkg/...", "-v", "-test.short", "-cover", "-coverprofile=coverage/backend.out"); err != nil {
		return err
	}

	return sh.RunV("go", "tool", "cover", "-html=coverage/backend.out", "-o", "coverage/backend.html")
}

// Default configures the default target.
var Default = build.BuildAll
