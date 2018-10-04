package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	commands := readLine("./commands.txt")
	errs := make([]error, 0)

	for i := 0; i < len(commands); i++ {
		cmd := exec.Command("C:/Windows/System32/wsl.exe", "sudo", commands[i], "start")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		f, _ := os.Create("./errors.txt")
		defer f.Close()

		w := bufio.NewWriter(f)

		for e := 0; e < len(errs); e++ {
			w.WriteString(fmt.Sprintf("%v\n", errs[e]))
		}

		w.Flush()
	}
}

func readLine(path string) []string {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	lines := make([]string, 0)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		lines = append(lines, scanner.Text())
	}

	return lines
}
