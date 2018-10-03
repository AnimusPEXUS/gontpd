package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/AnimusPEXUS/gontpd"
)

func main() {

	var port int = gontpd.NTPD_TCP_LISTENING_PORT

	{
		fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

		fs.IntVar(
			&port,
			"p",
			gontpd.NTPD_TCP_LISTENING_PORT,
			fmt.Sprintf("override default TCP listening port"),
		)

		err := fs.Parse(os.Args[1:])
		if err != nil {
			fmt.Println("error", err)
			os.Exit(10)
		}

	}

	s, err := gontpd.NewServer(port)
	if err != nil {
		log.Println("error", "Error instantiating Time server:", err)
	}

	err = s.Run()
	if err != nil {
		log.Println("error", "Time server exiting with error:", err)
	}

}
