/*
 * Copyright (C) 2016 ~ 2018 Deepin Technology Co., Ltd.
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

package watchdog

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDBusExists(t *testing.T) {
	Convey("Test dbus whether exists", t, func(c C) {
		err := initDBusObject()
		if err != nil {
			logger.Warning("failed to init dbusObject:", err)
		}
		if busObj == nil {
			t.Skip("busObj is nil")
		}
		exist, _ := isDBusServiceExist(orgFreedesktopDBus)
		c.So(exist, ShouldBeTrue)
		exist, _ = isDBusServiceExist(orgFreedesktopDBus + "111")
		c.So(exist, ShouldBeFalse)
	})
}

func TestStrInList(t *testing.T) {
	Convey("Test item whether in list", t, func(c C) {
		var list = []string{
			"abc",
			"xyz",
			"123",
		}
		c.So(isItemInList("abc", list), ShouldEqual, true)
		c.So(isItemInList("abcd", list), ShouldEqual, false)
	})
}

func TestTaskInfo(t *testing.T) {
	Convey("Test task create", t, func(c C) {
		c.So(newTaskInfo("test1", nil, nil), ShouldBeNil)
		c.So(newTaskInfo("test1",
			func() (bool, error) { return true, nil },
			func() error { return nil }), ShouldNotBeNil)
	})

	task1 := newTaskInfo("test1",
		func() (bool, error) { return false, nil },
		func() error { return nil })
	Convey("Test task state", t, func(c C) {
		task1.Enable(false)
		c.So(task1.CanLaunch(), ShouldEqual, false)
		task1.Enable(true)
		c.So(task1.CanLaunch(), ShouldEqual, true)
		task1.failed = true
		c.So(task1.CanLaunch(), ShouldEqual, false)
		task1.failed = false
		task1.isRunning = func() (bool, error) { return true, nil }
		c.So(task1.CanLaunch(), ShouldEqual, false)
	})

	task2 := newTaskInfo("test2",
		func() (bool, error) { return false, nil },
		func() error { return nil })
	Convey("Test manager", t, func(c C) {
		var m = &Manager{
			timedTasks: []*taskInfo{task1},
		}
		task1.failed = true
		c.So(m.hasAnyRunnableTimedTask(), ShouldEqual, false)
		m.timedTasks = append(m.timedTasks, task2)
		c.So(m.hasAnyRunnableTimedTask(), ShouldEqual, true)
	})
}
