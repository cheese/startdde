package display

import "github.com/BurntSushi/xgb/randr"

const (
	DisplayModeUnknow  = -100
	DisplayModeMirrors = -1
	DisplayModeCustom  = 0
	DisplayModeOnlyOne = 1
)

func (dpy *Display) SwitchMode(mode int16) {
	if dpy.DisplayMode == mode {
		return
	}
	dpy.setPropDisplayMode(mode)

	if mode == DisplayModeMirrors {
		n := len(dpy.Monitors)
		if n <= 1 {
			return
		}
		for ; n > 1; n = len(dpy.Monitors) {
			dpy.JoinMonitor(dpy.Monitors[n-1].Name, dpy.Monitors[n-2].Name)
		}
		dpy.apply(false)
	} else if mode == DisplayModeCustom {
		dpy.ResetChanges()
	} else {
		func() {
			dpy.lockMonitors()
			if mode >= DisplayModeOnlyOne && int(mode) <= len(dpy.Monitors) {
				dpy.unlockMonitors()

				for _, m := range dpy.Monitors {
					dpy.SplitMonitor(m.Name)
				}

				for i, m := range dpy.Monitors {
					if i+1 == int(mode) {
						m.changePos(0, 0)
						m.changeMode(randr.Mode(m.BestMode.ID))
						m.changeSwitchOn(true)
						dpy.SetPrimary(m.Name)
					}
				}
				for i, m := range dpy.Monitors {
					if i+1 != int(mode) {
						m.changeSwitchOn(false)
					}
				}
				dpy.apply(false)
			}
		}()
	}
}
