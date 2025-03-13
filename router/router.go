package router

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/gabrielalmir/arithmo/arithmo"
)

func HandleConnection(conn net.Conn, storage *arithmo.Storage) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}

		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		parts := strings.Fields(string(line))
		if len(parts) < 1 {
			continue
		}

		command := strings.ToUpper(parts[0])
		args := parts[1:]

		switch command {
		case "SET":
			if len(args) != 2 {
				fmt.Fprintln(conn, "-ERR wrong number of arguments for 'SET' command")
				continue
			}
			if val, err := strconv.Atoi(args[1]); err == nil {
				storage.Set(args[0], val)
			} else {
				storage.Set(args[0], args[1])
			}
			fmt.Fprintln(conn, "+OK")

		case "GET":
			if len(args) != 1 {
				fmt.Fprintln(conn, "-ERR wrong number of arguments for 'GET' command")
				continue
			}
			val, ok := storage.Get(args[0])
			if !ok {
				fmt.Fprintln(conn, "$-1")
				continue
			}
			fmt.Fprintf(conn, "$%d\r\n%s\r\n", len(fmt.Sprint(val)), val)

		case "INCR":
			if len(args) != 1 {
				fmt.Fprintln(conn, "-ERR wrong number of arguments for 'INCR' command")
				continue
			}
			newVal, err := storage.Incr(args[0])
			if err != nil {
				fmt.Fprintln(conn, "-"+err.Error())
				continue
			}
			fmt.Fprintln(conn, ":", newVal)

		case "DECR":
			if len(args) != 1 {
				fmt.Fprintln(conn, "-ERR wrong number of arguments for 'DECR' command")
				continue
			}
			newVal, err := storage.Decr(args[0])
			if err != nil {
				fmt.Fprintln(conn, "-"+err.Error())
				continue
			}
			fmt.Fprintln(conn, ":", newVal)

		case "TYPE":
			if len(args) != 1 {
				fmt.Fprintln(conn, "-ERR wrong number of arguments for 'TYPE' command")
				continue
			}
			fmt.Fprintln(conn, "+", storage.Type(args[0]))

		case "DEL", "EXISTS":
			if len(args) < 1 {
				fmt.Fprintln(conn, "-ERR wrong number of arguments for '"+command+"' command")
				continue
			}
			count := 0
			for _, key := range args {
				if command == "DEL" {
					if storage.Del(key) {
						count++
					}
				} else if command == "EXISTS" {
					if storage.Exists(key) {
						count++
					}
				}
			}
			fmt.Fprintln(conn, ":", count)

		case "QUIT":
			fmt.Fprintln(conn, "+OK")
			return

		default:
			fmt.Fprintf(conn, "-ERR unknown command '%s'\n", command)
		}
	}
}
