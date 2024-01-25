package main

import (
	"context"
	"fmt"
	"log"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/gorilla/websocket"
)

// button
var (
	connectOrDisconnectButton *widget.Button
	sendMsgButton             *widget.Button
	appendProtoButton         *widget.Button

	msgJsonEditButton   *widget.Button
	msgHeaderEditButton *widget.Button
	msgParamEditButton  *widget.Button
)

func initButtons() {
	// 建立连接按钮
	connectOrDisconnectButton = widget.NewButton("connect", connectButtonClick)

	// 编辑 JSON 消息按钮
	msgJsonEditButton = widget.NewButton("Message", msgJsonEditButtonClick)
	// 编辑 Header 按钮
	msgHeaderEditButton = widget.NewButton("Headers", msgHeaderEditButtonClick)
	//
	msgParamEditButton = widget.NewButton("Params", msgParamEditButtonClick)

	// 发送消息按钮
	sendMsgButton = widget.NewButton("send message", sendMsgButtonClick)
	sendMsgButton.Disable()

	// 追加 proto 目录
	appendProtoButton = widget.NewButton("append proto", appendProtoButtonClick)
}

// region button click
func connectButtonClick() {
	// 连接存在，此时为断开连接按钮
	if conn != nil {
		conn.Close()
		conn = nil
		// 将按钮名改为 connect
		connectOrDisconnectButton.SetText("connect")
		// 禁用发送消息 按钮
		sendMsgButton.Disable()
		dialog.ShowInformation("connect break", "connect break", window)
		return
	}

	var formControl *dialog.CustomDialog

	requestProto, responseProto := widget.NewEntry(), widget.NewEntry()

	selectWidget := widget.NewSelect([]string{"WFL", "AOW"}, func(s string) {})

	confirmButton := widget.NewButton("Confirm", func() {
		// 连接不存在，此时为建立连接按钮
		var err error
		conn, err = connectTOWs("ws://localhost:8002/ws", nil, nil)
		if err != nil {
			dialog.ShowError(fmt.Errorf("connect to server failed,err:%w", err), window)
			return
		}

		// 将按钮改成 disconnect
		connectOrDisconnectButton.SetText("disconnect")
		// 启用 sendMsg 按钮
		sendMsgButton.Enable()
		// 清空历史接收消息
		recvMsgEntry.SetText("")
		// 持续接收消息
		receiveMessage(context.TODO())
		// 关闭表单
		formControl.Hide()
		// 打印提示
		dialog.ShowInformation("connection success", "conn success", window)
	})
	cancelButton := widget.NewButton("Cancel", func() {
		formControl.Hide()
	})

	selectContainer := container.NewGridWithColumns(
		2,
		layout.NewSpacer(),
		layout.NewSpacer(),
		widget.NewLabel("select application"),
		selectWidget)

	btnContainer := container.NewGridWithColumns(
		2,
		layout.NewSpacer(),
		layout.NewSpacer(),
		confirmButton,
		cancelButton,
	)

	formItems := []*widget.FormItem{
		{Text: "RequestProtoMessage", Widget: requestProto},
		{Text: "ResponseProtoMessage", Widget: responseProto},
	}

	container := container.NewVBox(
		widget.NewForm(formItems...),
		selectContainer,
		btnContainer,
	)

	formControl = dialog.NewCustom("set proto and app", "", container, window)

	formControl.Show()

	//var formControl *dialog.CustomDialog
	//
	//// 指定收发 proto 协议名、项目名
	//form := &widget.Form{
	//	Items: []*widget.FormItem{
	//		{Text: "RequestProtoMessage", Widget: requestProto},
	//		{Text: "ResponseProtoMessage", Widget: responseProto},
	//		{},
	//	},
	//	OnSubmit: func() {
	//		// 连接不存在，此时为建立连接按钮
	//		var err error
	//		conn, err = connectTOWs("ws://localhost:8002/ws", nil, nil)
	//		if err != nil {
	//			dialog.ShowError(fmt.Errorf("connect to server failed,err:%w", err), window)
	//			return
	//		}
	//
	//		// 将按钮改成 disconnect
	//		connectOrDisconnectButton.SetText("disconnect")
	//		// 启用 sendMsg 按钮
	//		sendMsgButton.Enable()
	//		// 清空历史接收消息
	//		recvMsgEntry.SetText("")
	//		// 持续接收消息
	//		receiveMessage(context.TODO())
	//
	//		dialog.ShowInformation("connection success", "conn success", window)
	//		log.Println("建立 WS 连接")
	//		formControl.Hide()
	//	},
	//	OnCancel: func() {
	//		log.Println("enen ")
	//		formControl.Hide()
	//	},
	//}
	//
	//formControl = dialog.NewCustom("set proto info", "", form, window)
	//formControl.Show()
}

func sendMsgButtonClick() {
	if conn == nil {
		dialog.ShowError(fmt.Errorf("WebSocket未连接"), window)
		return
	}

	err := conn.WriteMessage(websocket.TextMessage, []byte(messageEntry.Text))
	if err != nil {
		dialog.ShowError(fmt.Errorf("message send failed,err:%window", err), window)
		return
	}
}

func appendProtoButtonClick() {
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Field 1", Widget: widget.NewEntry()},
			{Text: "Field 2", Widget: widget.NewPasswordEntry()},
			// 添加更多表单字段
		},
		SubmitText: "Submit",
		CancelText: "Cancel",
		OnSubmit: func() {
			// 处理表单提交逻辑
		},
		OnCancel: func() {
			// 处理取消逻辑
		},
	}

	dialog.ShowCustomConfirm("Append Protocol", "Submit", "Cancel", form, func(b bool) {
		if b {
			form.OnSubmit()
		} else {
			form.OnCancel()
		}
	}, window)
}

func msgJsonEditButtonClick() {
}

func msgHeaderEditButtonClick() {
	log.Println("你好")
}

func msgParamEditButtonClick() {
	log.Println("msgParamEditButtonClick 点击")
}

// endregion button click
