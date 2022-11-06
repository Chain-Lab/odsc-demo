package main

import (
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"github.com/decision2016/go-crypto/ecdcs"
	"log"
	"os"
)

var priv = "c3aa13d2a7bef392af504cc2e5027515ff7601f0a2825d44a8124e814b3011d4"
var file_path = "G:\\repos\\Experiment\\files"

func main() {
	sk := ecdcs.PrivateKey{}
	_ = sk.FromHexString(priv)
	fmt.Println("{")

	for size := 100; size <= 1000; size += 100 {

		file_name := fmt.Sprintf("%s\\%d-0.test", file_path, size)

		data , err := os.ReadFile(file_name)

		if err != nil {
			log.Println(file_name)
			log.Fatalln("Open file failed.")
		}

		random, signature, err := sk.Signature(elliptic.P256(), data)
		hexSignature := hex.EncodeToString(signature)
		hexRandom := hex.EncodeToString(random)
		fmt.Println("\""+hexSignature+"\":[")
		fmt.Println("\""+hexRandom+"\",")


		for i := 1; i < 15; i++ {
			file_name := fmt.Sprintf("%s\\%d-%d.test", file_path, size, i)
			predata := data
			data , err = os.ReadFile(file_name)

			if err != nil {
				log.Fatalln("Open file failed.")
			}

			random := sk.ReSignature(elliptic.P256(), predata, data, signature)
			hexRandom = hex.EncodeToString(random)
			fmt.Println("\""+hexRandom+"\",")
		}
		fmt.Println("],")
	}
	fmt.Println("}")
}
