/*
 * Copyright (C) 2019 ~ 2019 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package xsettings

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	dmiSysVendor   = "/sys/class/dmi/id/sys_vendor"
	dmiBoardVendor = "/sys/class/dmi/id/board_vendor"
	dmiBiosVendor  = "/sys/class/dmi/id/bios_vendor"

	// scale mode available values: integer
	scaleModeSysFile  = "/etc/deepin/startdde/scale_mode"
	scaleModeUserFile = ".config/deepin/startdde/scale_mode"
)

func isOEMVersion() bool {
	return isHuaweiOEM()
}

func isHuaweiOEM() bool {
	files := []string{
		dmiSysVendor,
		dmiBoardVendor,
		dmiBiosVendor,
	}
	for _, filename := range files {
		if fileHasKey(filename, "huawei") {
			return true
		}
	}
	return false
}

func isIntegerScaleMode() bool {
	if fileHasKey(scaleModeSysFile, "integer") {
		return true
	}
	filename := filepath.Join(getHomeDir(), scaleModeUserFile)
	return fileHasKey(filename, "integer")
}

func setIntegerScaleMode() error {
	filename := filepath.Join(getHomeDir(), scaleModeUserFile)
	dir := filepath.Dir(filename)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, []byte("integer"), 0644)
}

func fileHasKey(filename, key string) bool {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return false
	}

	str := strings.ToLower(string(data))
	return strings.Contains(str, key)
}

func isFloatNum(num float64) bool {
	str := fmt.Sprint(num)
	items := strings.SplitN(str, ".", 2)
	if len(items) != 2 {
		return false
	}

	for _, ch := range items[1] {
		if ch > '0' {
			return true
		}
	}
	return false
}

func getHomeDir() string {
	d := os.Getenv("HOME")
	if len(d) != 0 {
		return d
	}
	return filepath.Join("/home", os.Getenv("USER"))
}
