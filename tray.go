package main

import (
	"ctrlc/icon"

	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
)

func onReady(s *params) func() {
	return func() {
		systray.SetTemplateIcon(icon.Data, icon.Data)
		//systray.SetTitle("CtrlC")
		systray.SetTooltip("CtrlC")

		mPref := systray.AddMenuItem("Open Web", "Open Web")
		mBackup := systray.AddMenuItem("BackUp", "Run backup.sh")
		mQuit := systray.AddMenuItem("Quit", "Quit")
		systray.AddSeparator()
		mAboutLink := systray.AddMenuItem("About", "About Me!")

		for {
			select {
			case <-mAboutLink.ClickedCh:
				open.Run("https://space.bilibili.com/278413353")
			case <-mPref.ClickedCh:
				open.Run(s.webUrl)

			case <-mBackup.ClickedCh:
				go backup(s.homeDir, s.dbPath)

			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}
}

func onExit(ser *params) func() {
	return func() {
		ser.Stop()
	}
}
