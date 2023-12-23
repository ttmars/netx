package main

import (
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	//TestNetwork([]string{"https://www.baidu.com/", "http://blog.youthsweet.com/", "https://www.bilibili.com/", "https://www.google.com.hk/"}, "http://127.0.0.1:7890", 1)
	//TestNetwork([]string{"https://www.baidu.com/", "http://blog.youthsweet.com/", "https://www.bilibili.com/", "https://www.google.com.hk/"}, "", 1)
	run()
}

func run() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "netx",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		//BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
