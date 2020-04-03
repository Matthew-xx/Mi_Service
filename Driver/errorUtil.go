package main

import "os"

func HandlerError(err error)  {
	if err != nil {
		os.Exit(1)
	}
}
