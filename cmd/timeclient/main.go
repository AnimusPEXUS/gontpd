package main

import (
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/AnimusPEXUS/gontpd"
)

func main() {
	var (
		host string = gontpd.NTPD_TCP_LISTENING_HOST
		port int    = gontpd.NTPD_TCP_LISTENING_PORT
	)

	{
		fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

		err := fs.Parse(os.Args[1:])
		if err != nil {
			fmt.Println("error", err)
			os.Exit(10)
		}

		fs_args := fs.Args()
		len_args := len(fs_args)

		if len_args < 1 {
			fmt.Println("error", "too few arguments")
			os.Exit(11)
		}

		if len_args > 2 {
			fmt.Println("error", "too many arguments")
			os.Exit(12)
		}

		if len_args >= 1 {
			host = fs_args[0]
		}

		if len_args == 2 {
			port, err = strconv.Atoi(fs_args[1])
			if err != nil {
				fmt.Println("error", "invalid value for port argument")
				os.Exit(13)
			}
		}

	}

	hostport := fmt.Sprintf("%s:%d", host, port)
	log.Printf("calling %s for time info", hostport)
	conn, err := net.Dial("tcp", hostport)
	if err != nil {
		log.Println("failed calling Time server: " + err.Error())
		os.Exit(10)
	}

	buff := make([]byte, 4)

	count, err := conn.Read(buff)
	if count != 4 {
		err = errors.New("response too short: " + strconv.Itoa(count))
	}

	if err != nil {
		log.Println("failed reading Time server response: " + err.Error())
		os.Exit(10)
	}

	nint := binary.LittleEndian.Uint32(buff)

	fmt.Println(nint)

	t := time.Unix(int64(nint), 0)
	t = t.UTC()
	fmt.Println(t.String())

}
