package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

func main() {
	serverAddr := os.Getenv("SERVER_ADDR")
	if serverAddr == "" {
		serverAddr = "localhost:8080"
	}

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	green := color.New(color.FgGreen).SprintFunc()   // pesan masuk
	cyan := color.New(color.FgCyan).SprintFunc()     // nama
	yellow := color.New(color.FgYellow).SprintFunc() // jam

	// minta nama user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan nama Anda: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// terima pesan dari server
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			msg := scanner.Text()
			timestamp := time.Now().Format("15:04")

			parts := strings.SplitN(msg, "]", 2)
			if len(parts) == 2 {
				user := strings.TrimPrefix(parts[0], "[")
				content := parts[1]
				fmt.Printf(
					"\r%s [%s ~ %s]%s\n",
					green("↓"),
					cyan(user),
					yellow(timestamp),
					content,
				)
			} else {
				fmt.Printf("\r↓ [%s]\n", msg)
			}
			fmt.Print("> ")
		}
	}()

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}

		// kirim ke server
		fmt.Fprintf(conn, "[%s]: %s\n", name, text)
	}
}
