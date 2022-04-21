# MiniGUI

一个 Golang 的 GUI 基础结构，依赖 fyne 和 systray 。

> 构建使用 png 文件最好

- [wails](https://github.com/wailsapp/wails)
- [fyne.io/fyne](https://github.com/fyne-io/fyne)
- [github.com/getlantern/systray](https://github.com/getlantern/systray)


### Linux

* Building apps requires gcc as well as the `gtk3` and `libayatana-appindicator3` development headers to be installed. For Debian or Ubuntu, you may install these using:

```sh
sudo apt-get install gcc libgtk-3-dev libayatana-appindicator3-dev
```

On Linux Mint, `libxapp-dev` is also required.

If you need to support the older `libappindicator3` library instead, you can pass the build flag `legacy_appindicator`
when building. For example:

```
go build -tags=legacy_appindicator
```

To build `webview_example`, you also need to install `libwebkit2gtk-4.0-dev` and remove `webview_example/rsrc.syso` which is required on Windows.

### Windows

* To avoid opening a console at application startup, use these compile flags:

```sh
go build -ldflags -H=windowsgui
```

### macOS

On macOS, you will need to create an application bundle to wrap the binary; simply folders with the following minimal structure and assets:

```
SystrayApp.app/
  Contents/
    Info.plist
    MacOS/
      go-executable
    Resources/
      SystrayApp.icns
```

When running as an app bundle, you may want to add one or both of the following to your Info.plist:

```xml
<!-- avoid having a blurry icon and text -->
	<key>NSHighResolutionCapable</key>
	<string>True</string>

	<!-- avoid showing the app on the Dock -->
	<key>LSUIElement</key>
	<string>1</string>
```

Consult the [Official Apple Documentation here](https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFBundles/BundleTypes/BundleTypes.html#//apple_ref/doc/uid/10000123i-CH101-SW1).

## Credits

- https://github.com/xilp/systray
- https://github.com/cratonica/trayhost