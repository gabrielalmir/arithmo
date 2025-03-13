package router

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/gabrielalmir/arithmo/arithmo"
	"github.com/tidwall/resp"
)

func HandleConnection(conn net.Conn, storage *arithmo.Storage) {
	defer conn.Close()
	rd := resp.NewReader(conn)
	wr := resp.NewWriter(conn)

	for {
		val, _, err := rd.ReadValue()
		if err != nil {
			fmt.Println("Erro ao ler comando:", err)
			return
		}

		if val.Type() != resp.Array {
			wr.WriteError(fmt.Errorf("ERR comando deve ser um array"))
			continue
		}

		args := val.Array()

		if len(args) < 1 {
			wr.WriteError(fmt.Errorf("ERR comando vazio"))
			continue
		}

		cmdName := strings.ToUpper(args[0].String())
		args = args[1:]

		switch cmdName {
		case "SET":
			if len(args) != 2 {
				wr.WriteError(fmt.Errorf("ERR wrong number of arguments for 'SET' command"))
				continue
			}

			if val, err := strconv.Atoi(args[1].String()); err == nil {
				storage.Set(args[0].String(), val)
			} else {
				storage.Set(args[0].String(), args[1].String())
			}

			wr.WriteSimpleString("OK")

		case "GET":
			if len(args) != 1 {
				wr.WriteError(fmt.Errorf("ERR wrong number of arguments for 'GET' command"))
				continue
			}
			val, ok := storage.Get(args[0].String())
			if !ok {
				wr.WriteNull()
				continue
			}
			wr.WriteString(fmt.Sprintf("%v", val))

		case "INCR":
			if len(args) != 1 {
				wr.WriteError(fmt.Errorf("ERR wrong number of arguments for 'INCR' command"))
				continue
			}
			newVal, err := storage.Incr(args[0].String())
			if err != nil {
				wr.WriteError(err)
				continue
			}
			wr.WriteInteger(newVal)

		case "DECR":
			if len(args) != 1 {
				wr.WriteError(fmt.Errorf("ERR wrong number of arguments for 'DECR' command"))
				continue
			}
			newVal, err := storage.Decr(args[0].String())
			if err != nil {
				wr.WriteError(err)
				continue
			}
			wr.WriteInteger(newVal)

		case "LPUSH":
			if len(args) < 2 {
				wr.WriteError(fmt.Errorf("ERR wrong number of arguments for 'LPUSH' command"))
				continue
			}
			key := args[0].String()
			values := make([]interface{}, len(args)-1)
			for i, v := range args[1:] {
				values[i] = v.String()
			}
			newLen, err := storage.LPush(key, values...)
			if err != nil {
				wr.WriteError(err)
				continue
			}
			wr.WriteInteger(newLen)

		case "RPOP":
			if len(args) != 1 {
				wr.WriteError(fmt.Errorf("ERR wrong number of arguments for 'RPOP' command"))
				continue
			}
			val, err := storage.RPop(args[0].String())
			if err != nil {
				wr.WriteError(err)
				continue
			}
			wr.WriteString(fmt.Sprintf("%v", val))

		case "TYPE":
			if len(args) != 1 {
				wr.WriteError(fmt.Errorf("ERR wrong number of arguments for 'TYPE' command"))
				continue
			}
			wr.WriteSimpleString(storage.Type(args[0].String()))

		case "DEL":
			if len(args) < 1 {
				wr.WriteError(fmt.Errorf("ERR wrong number of arguments for 'DEL' command"))
				continue
			}
			count := 0
			for _, key := range args {
				if storage.Del(key.String()) {
					count++
				}
			}
			wr.WriteInteger(count)

		case "EXISTS":
			if len(args) < 1 {
				wr.WriteError(fmt.Errorf("ERR wrong number of arguments for 'EXISTS' command"))
				continue
			}
			count := 0
			for _, key := range args {
				if storage.Exists(key.String()) {
					count++
				}
			}
			wr.WriteInteger(count)

		case "QUIT":
			wr.WriteSimpleString("OK")
			return

		default:
			wr.WriteError(fmt.Errorf("ERR unknown command '%s'", cmdName))
		}
	}
}
