package main

import (
	"log"

	gorpc "github.com/laymer110/gorpc/proto"
)

func main() {
	s := gorpc.AuthRequest{}
	log.Println(s.AuthInfo)
}
