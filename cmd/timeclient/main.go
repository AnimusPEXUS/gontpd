package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/AnimusPEXUS/gotimed"
)

func Log(txt string, out *os.File) {
	out.Write([]byte(time.Now().UTC().String() + " " + txt + "\n"))
}

func LogError(txt string) {
	Log(txt, os.Stderr)
}

//func LogInfo(txt string) {
//	Log(txt, os.Stdout)
//}

func main() {
	var (
		host string = gotimed.NTPD_TCP_LISTENING_HOST
		port int    = gotimed.NTPD_TCP_LISTENING_PORT
	)

	{
		fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

		err := fs.Parse(os.Args[1:])
		if err != nil {
			LogError("error: " + err.Error())
			os.Exit(10)
		}

		fs_args := fs.Args()
		len_args := len(fs_args)

		if len_args > 2 {
			LogError("error: too many arguments")
			os.Exit(12)
		}

		if len_args > 0 {
			host = fs_args[0]
		}

		if len_args > 1 {
			port, err = strconv.Atoi(fs_args[1])
			if err != nil {
				LogError("error: invalid value for port argument")
				os.Exit(13)
			}
		}

	}

	hostport := fmt.Sprintf("%s:%d", host, port)
	//	log.Printf("calling %s for time info", hostport)
	conn, err := net.Dial("tcp", hostport)
	if err != nil {
		LogError("failed calling Time server: " + err.Error())
		os.Exit(10)
	}

	var val uint32

	err = binary.Read(conn, binary.BigEndian, &val)
	if err != nil {
		LogError("failed reading Time server response: " + err.Error())
		os.Exit(10)
	}

	res := gotimed.UnixToRfc868(int64(val))

	fmt.Println(res)

	//	t := time.Unix(res, 0)
	//	t = t.UTC()
	//	fmt.Println(t.String())

}
