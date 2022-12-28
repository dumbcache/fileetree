package filetree

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"math/rand"
	"os"

	"os/exec"
	"time"
)

var colors = []string{
	"\033[31m",
	"\033[32m",
	"\033[33m",
	"\033[34m",
	"\033[35m",
	"\033[36m",
	// "\033[37m",
}
var (
	vLine      = "\U00002502"
	hLine      = "\U00002500"
	colorBG    = "\033[40m"
	colorReset = "\033[0m"
	bgReset    = "\033[49m"
	space      = ""
)

func random() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(6)
}

func execCommand(c string, args ...string) {
	cmd := exec.Command(c, args...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func readFiles(path string, color string, padding string, args ...string) {

	padding = padding + color + vLine + colorReset
	if len(args) != 0 {
		padding = color + vLine + colorReset + space
		space = args[0]
	}
	files, _ := ioutil.ReadDir(path)
	dirList := []fs.FileInfo{}
	for _, file := range files {
		if !file.IsDir() {
			fmt.Print(padding)
			fmt.Println(hLine + " " + file.Name())
			continue
		}
		dirList = append(dirList, file)

	}
	dirColor := colors[random()]
	for {
		if dirColor != color {
			break
		}
		dirColor = colors[random()]
	}
	for _, dir := range dirList {
		fmt.Print(padding)
		fmt.Println(hLine + dirColor + "./" + dir.Name() + colorReset)
		readFiles(path+"/"+dir.Name(), dirColor, padding+space)
	}
}

func FileTree(path string) {
	initColor := colors[random()]
	fmt.Println(initColor + path + colorReset)
	readFiles(path, initColor, "", " ")
}
