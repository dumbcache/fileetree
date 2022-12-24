package filetree

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"

	"os/exec"
	"time"
)

var colors = map[string]string{
	"Reset":   "\033[0m",
	"Red":     "\033[31m",
	"Green":   "\033[32m",
	"Yellow":  "\033[33m",
	"Blue":    "\033[34m",
	"Purple":  "\033[35m",
	"Cyan":    "\033[36m",
	"Gray":    "\033[37m",
	"BG":      "\033[40m",
	"BgReset": "\033[49m",
	"White":   "\033[97m",
}
var (
	vLine = "\U00002502"
	hLine = "\U00002500"
)

func random() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func execCommand(c string, args ...string) {
	fmt.Print(colors["Blue"])
	cmd := exec.Command(c, args...)
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Print(colors["Reset"])
}

func readFiles(dir string, padding string, args ...string) {

	padding = vLine + padding
	space := ""
	if len(args) != 0 {
		space = args[0]
		padding = vLine + space
	}
	files, _ := ioutil.ReadDir(dir)
	// var tempFiles = make([]string, len(files))
	// for _, file := range files {
	// 	tempFiles = append(tempFiles, file.Name())
	// }
	// fmt.Println(tempFiles)
	// sort.Strings(tempFiles)
	for _, file := range files {
		fmt.Print(padding)
		if file.IsDir() {
			fmt.Println(hLine + " " + file.Name())
			readFiles(dir+"/"+file.Name(), " "+padding)
		} else {
			fmt.Println(hLine + " " + file.Name())
		}
	}
}

func FileTree(dir string) {
	execCommand("clear")
	fmt.Println(colors["Red"] + "files:" + colors["Reset"])
	readFiles(dir, "")
}
