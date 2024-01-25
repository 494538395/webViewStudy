package main

import view "github.com/webview/webview"

func main() {
	// 创建一个新的 WebView 窗口
	webView := view.New(view.Settings{
		Title:     "WebView Demo",
		Width:     800,
		Height:    600,
		Resizable: true,
		URL:       "https://www.example.com",
	})

	// 运行事件循环
	webView.Run()
}
