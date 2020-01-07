// SPDX-License-Identifier: Unlicense OR MIT

// +build android

package app

import "C"

import (
	"sync"

	"github.com/gop9/olt/gio/app/internal/window"
)

var (
	dataDirOnce sync.Once
	dataPath    string
)

func dataDir() (string, error) {
	dataDirOnce.Do(func() {
		dataPath = window.GetDataDir()
	})
	return dataPath, nil
}
