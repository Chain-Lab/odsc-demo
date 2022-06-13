package ethereum

import (
	"log"
	"testing"
)

func TestGetBootstrap(t *testing.T) {
	bootstrap := ReadBootstrap()

	log.Println(bootstrap)
}
