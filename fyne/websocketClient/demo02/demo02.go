package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode"

	"YHSoft/Demo/GUIDemo/FyneDemo/02_RenameFileName/models" //更换成自己的项目路径

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
)

func main() {
	//新建一个app
	a := app.New()
	//设置窗口栏，任务栏图标
	a.SetIcon(resourceIconPng)
	//新建一个窗口
	w := a.NewWindow("自动化更名程序V1.0")
	//主界面框架布局
	MainShow(w)
	//尺寸
	w.Resize(fyne.Size{Width: 500, Height: 100})
	//w居中显示
	w.CenterOnScreen()
	//循环运行
	w.ShowAndRun()

	err := os.Unsetenv("FYNE_FONT")
	if err != nil {
		return
	}
}

var tileInfo string
var done = make(chan bool)
var stop = make(chan int, 1)
var num int64

// MainShow 主界面函数
func MainShow(w fyne.Window) {
	//var ctrl *beep.Ctrl
	title := widget.NewLabel("自动化更名程序")
	hello := widget.NewLabel("目录文件:")
	entry1 := widget.NewEntry() //文本输入框
	//entry1.SetText("E:\\rename_temp2\\123.txt")

	dia1 := widget.NewButton("打开", func() { //回调函数：打开选择文件对话框
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if reader == nil {
				log.Println("Cancelled")
				return
			}

			entry1.SetText(reader.URI().Path()) //把读取到的路径显示到输入框中
		}, w)

		fd.SetFilter(storage.NewExtensionFileFilter([]string{".txt"})) //打开的文件格式类型
		fd.Show()                                                      //控制是否弹出选择文件目录对话框
	})

	text := widget.NewMultiLineEntry() //多行输入组件
	//text.Disable()                     //禁用输入框，不能更改数据

	labelLast := widget.NewLabel("发飙的蜗牛    ALL Right Reserved")
	//labelLast := widget.NewLabel(" ")
	label4 := widget.NewLabel("文件路径:")
	entry2 := widget.NewEntry()

	//entry2.SetText("E:\\rename_temp2\\宝岗路停车场项目幕墙施工图PDF")

	dia2 := widget.NewButton("打开", func() {
		dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if list == nil {
				log.Println("Cancelled")
				return
			}
			//设置输入框内容
			entry2.SetText(list.Path())
		}, w)
	})

	//开始更名按钮
	bt3 := widget.NewButton("开始 更名", func() {
		go func() {
			if (entry1.Text != "") && (entry2.Text != "") {
				text.SetText("")
				text.Refresh()
				if num != 0 {
					stop <- 1
					return
				} else {
					err := generateTxt(entry1.Text, entry2.Text, text)
					if err != nil {
						dialog.ShowError(err, w)
					}
					text.Refresh()
				}
			} else {
				dialog.ShowError(errors.New("读取TXT文件错误"), w)
			}
		}()

	})

	//停止更名按钮
	bt4 := widget.NewButton("停止 更名", func() {
		go func() {
			done <- false
		}()
	})

	head := container.NewCenter(title)

	v1 := container.NewBorder(layout.NewSpacer(), layout.NewSpacer(), hello, dia1, entry1)
	v4 := container.NewBorder(layout.NewSpacer(), layout.NewSpacer(), label4, dia2, entry2)

	v5 := container.NewHBox(bt3, bt4)
	v5Center := container.NewCenter(v5)

	ctnt := container.NewVBox(head, v1, v4, v5Center, text, labelLast) //控制显示位置顺序
	w.SetContent(ctnt)
}

// 设置字体
func init() {
	fontPaths := findfont.List()
	for _, fontPath := range fontPaths {
		//fmt.Println(fontPath)
		//楷体:simkai.ttf
		//黑体:simhei.ttf
		//微软雅黑：msyh.ttc
		if strings.Contains(fontPath, "simkai.ttf") {
			err := os.Setenv("FYNE_FONT", fontPath)
			if err != nil {
				return
			}
			break
		}
	}
}

// 读取数据校验数据
func generateTxt(inPath, outPath string, text *widget.Entry) error {
	//标题
	tileInfo += "开始处理,正在读取文件...\n"
	text.SetText(tileInfo)

	nameList, err := models.GetStrList(inPath)
	if err != nil {
		tileInfo += "读取文件出错...\n"
		text.SetText(tileInfo)
	}

	//获取文件路径的文件
	files, _ := ioutil.ReadDir(outPath)

	var fileNameList []string

	for _, file := range files {
		// 带扩展名的文件名
		fullFilename := file.Name()

		//添加文件数据到切片中
		fileNameList = append(fileNameList, fullFilename)
	}

	if len(nameList) == 0 || len(fileNameList) == 0 {
		tileInfo += "已停止处理...\n"
		text.SetText(tileInfo) //设置多行显示控件中的内容
		return errors.New("找不到路径或路径下不存在此文件")
	}

	if len(nameList) != len(fileNameList) {
		tileInfo += "已停止处理...\n"
		text.SetText(tileInfo) //设置多行显示控件中的内容
		return errors.New("目录行数与文件行数不相等，请检查！")
	}

	renameFile(nameList, fileNameList, outPath, text)

	return nil
}

// 操作更名文件
func renameFile(nameList []string, fileNameList []string, outPath string, text *widget.Entry) {
	var newName string

	//遍历更改
	for index, fullFilename := range fileNameList {
		select {
		case <-done: //读管道中内容，没有内容前，阻塞
			//扩展名
			fileExt := filepath.Ext(fullFilename)

			//fmt.Println("nameList[index]=",nameList[index])
			defaultName := nameList[index] //文本文件初始名称
			var defaultNameCode []string
			var defaultNameNewCode string

			rune := []rune(defaultName)
			for i := len(rune); i > 0; i-- {
				if !unicode.Is(unicode.Han, rune[i-1]) {
					defaultNameCode = append(defaultNameCode, string(rune[i-1]))
				} else {
					break
				}
			}

			//重新编排名称编号
			for i := len(defaultNameCode); i > 0; i-- {
				defaultNameNewCode += defaultNameCode[i-1]
			}

			//fmt.Println("defaultName=", defaultName)

			//重组新名称
			newName = strconv.Itoa(index+1) + " " + strings.ReplaceAll(defaultName, defaultNameNewCode, "") + " " + defaultNameNewCode

			// 不带扩展名的文件名
			//filenameOnly := strings.TrimSuffix(fullFilename, fileExt)

			//将每个文件名后面加上1，扩展名不变
			err := os.Rename(outPath+`\`+fullFilename, outPath+`\`+newName+fileExt)
			if err != nil {
				fmt.Println("err=", err)
			}

			tileInfo += "处理数据文件:" + fullFilename + "\n"
			time.Sleep(time.Second * 1)
			text.SetText(tileInfo) //设置多行显示控件中的 内容

			tileInfo += "停止更名...\n"
			text.SetText(tileInfo) //设置多行显示控件中的内容
			num = num + 1
			<-stop //读管道中内容，没有内容前，阻塞
		default:
			//扩展名
			fileExt := filepath.Ext(fullFilename)

			//fmt.Println("nameList[index]=",nameList[index])
			defaultName := nameList[index] //文本文件初始名称
			var defaultNameCode []string
			var defaultNameNewCode string

			rune := []rune(defaultName)
			for i := len(rune); i > 0; i-- {
				if !unicode.Is(unicode.Han, rune[i-1]) {
					defaultNameCode = append(defaultNameCode, string(rune[i-1]))
				} else {
					break
				}
			}

			//重新编排名称编号
			for i := len(defaultNameCode); i > 0; i-- {
				defaultNameNewCode += defaultNameCode[i-1]
			}

			//fmt.Println("defaultName=", defaultName)

			//重组新名称
			newName = strconv.Itoa(index+1) + " " + strings.ReplaceAll(defaultName, defaultNameNewCode, "") + " " + defaultNameNewCode

			// 不带扩展名的文件名
			//filenameOnly := strings.TrimSuffix(fullFilename, fileExt)

			//将每个文件名后面加上1，扩展名不变
			err := os.Rename(outPath+`\`+fullFilename, outPath+`\`+newName+fileExt)
			if err != nil {
				fmt.Println("err=", err)
			}

			tileInfo += "处理数据文件:" + fullFilename + "\n"
			//time.Sleep(time.Second * 1)
			text.SetText(tileInfo) //设置多行显示控件中的 内容
		}
	}
	tileInfo += "处理完毕！"
	text.SetText(tileInfo) //设置多行显示控件中的内容
	num = 0
}
