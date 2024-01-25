package main

import (
	"context"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// label
var (
	title = widget.NewLabel("Websocket Debug Util")
)

var window fyne.Window

func main() {
	a := app.New()
	window = a.NewWindow("WebSocket调试工具")

	initEntry()
	initButtons()

	// 第一行。标题
	head := container.NewCenter(title)

	// 第二行。url 、建立 WS 连接
	urlLine := container.NewBorder(
		layout.NewSpacer(),
		layout.NewSpacer(),
		widget.NewLabel("ws addr:"),
		connectOrDisconnectButton,
		urlEntry)

	// 第三行。追加指定项目 proto 目录
	appendProtoLine := container.NewGridWithColumns(
		6,
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
		appendProtoButton)

	// 第四行。message 参数
	genMsgBtnLine := container.NewGridWithColumns(
		3,
		msgJsonEditButton,
		msgHeaderEditButton,
		msgParamEditButton)

	// 第五行，message 参数内容
	genMsgContentLine := container.NewGridWithColumns(
		3,
		messageEntry,
		msgHeaderEntry,
		msgParamEntry)

	// 主体
	content := container.NewVBox(
		head,
		urlLine,
		appendProtoLine,
		genMsgBtnLine,
		genMsgContentLine,
		sendMsgButton,
		widget.NewLabel("receive message:"),
		recvMsgEntry,
	)

	window.SetContent(content)
	window.Resize(fyne.NewSize(1000, 500))
	window.ShowAndRun()
}

func receiveMessage(ctx context.Context) {
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				//dialog.ShowError(fmt.Errorf("无法读取消息：%window", err), window)
				log.Println("read message,err:", err)
				return
			}

			recvMsgEntry.SetText(recvMsgEntry.Text + "\n" + string(message))
		}
	}()
}
