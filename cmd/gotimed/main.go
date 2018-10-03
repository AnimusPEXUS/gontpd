package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/AnimusPEXUS/gotimed"
)

func main() {

	var port int = gotimed.TIMED_TCP_LISTENING_PORT

	{
		fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

		fs.IntVar(
			&port,
			"p",
			gotimed.TIMED_TCP_LISTENING_PORT,
			fmt.Sprintf("override default TCP listening port"),
		)

		err := fs.Parse(os.Args[1:])
		if err != nil {
			fmt.Println("error", err)
			os.Exit(10)
		}

	}

	s, err := gotimed.NewServer(port)
	if err != nil {
		log.Println("error", "Error instantiating Time server:", err)
		os.Exit(11)
	}

	err = s.Run()
	if err != nil {
		log.Println("error", "Time server exiting with error:", err)
		os.Exit(12)
	}

}
