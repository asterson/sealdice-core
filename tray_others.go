//go:build !windows
// +build !windows

package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net"
	"os"
	"regexp"
	"sealdice-core/dice"
)

func trayInit() {

}

func hideWindow() {
}

func showWindow() {
}

func TestRunning() bool {
	return false
}

func tempDirWarn() {
	fmt.Println("当前工作路径为临时目录，因此拒绝继续执行。")
}

func httpServe(e *echo.Echo, dm *dice.DiceManager) {
	portStr := "3211"
	rePort := regexp.MustCompile(`:(\d+)$`)
	m := rePort.FindStringSubmatch(dm.ServeAddress)
	if len(m) > 0 {
		portStr = m[1]
	}

	ln, err := net.Listen("tcp", ":"+portStr)
	if err != nil {
		logger.Errorf("端口已被占用，即将自动退出: %s", dm.ServeAddress)
		os.Exit(1)
	}
	_ = ln.Close()

	//exec.Command(`cmd`, `/c`, `start`, fmt.Sprintf(`http://localhost:%s`, portStr)).Start()
	fmt.Println("如果浏览器没有自动打开，请手动访问:")
	fmt.Println(fmt.Sprintf(`http://localhost:%s`, portStr)) // 默认:3211
	err = e.Start(dm.ServeAddress)
	if err != nil {
		logger.Errorf("端口已被占用，即将自动退出: %s", dm.ServeAddress)
		return
	}
}
