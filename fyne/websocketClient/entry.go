package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// entry
var (
	urlEntry       = widget.NewEntry()
	messageEntry   = widget.NewMultiLineEntry()
	msgHeaderEntry = widget.NewMultiLineEntry()
	msgParamEntry  = widget.NewMultiLineEntry()

	recvMsgEntry = widget.NewMultiLineEntry()

	testEntry = widget.NewMultiLineEntry()
)

func initEntry() {
	// 地址 url
	urlEntry.SetPlaceHolder("ws://localhost:8002/ws")

	// message
	messageEntry.SetPlaceHolder("please enter your json message")

	// message header
	msgHeaderEntry.SetPlaceHolder("please enter message header")

	// message params
	msgParamEntry.SetPlaceHolder("please enter message params")

	// revc message
	recvMsgEntry.BaseWidget.MinSize()
	recvMsgEntry.Resize(fyne.NewSize(500, 500))
}
