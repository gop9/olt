// SPDX-License-Identifier: Unlicense OR MIT

// +build linux freebsd windows

package headless

import (
	"github.com/gop9/olt/gio/app/internal/egl"
)

func newContext() (context, error) {
	return egl.NewContext(egl.EGL_DEFAULT_DISPLAY)
}
