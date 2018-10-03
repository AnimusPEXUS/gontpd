package main

import (
	"encoding/binary"
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

		if len_args > 2 {
			fmt.Println("error", "too many arguments")
			os.Exit(12)
		}

		if len_args > 0 {
			host = fs_args[0]
		}

		if len_args > 1 {
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

	var val uint32

	err = binary.Read(conn, binary.BigEndian, &val)
	if err != nil {
		log.Println("failed reading Time server response: " + err.Error())
		os.Exit(10)
	}

	res := gontpd.UnixToRfc(int64(val))

	fmt.Println(res)

	t := time.Unix(res, 0)
	t = t.UTC()
	fmt.Println(t.String())

}
