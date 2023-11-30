package main


import(
	"fmt"
	"flag"
	"bufio"
	"os"

	"github.com/spawnzii/crtgo/crt"
)

func main(){

	domain := flag.String("domain","","Domain to enumerate.")
	flag.Parse()

	if *domain == "" {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			domain := scanner.Text()
			findings := crt.FetchCrt(domain)
			for sub := range findings{
				fmt.Println(findings[sub])
			}
		}
	}else {
		findings := crt.FetchCrt(*domain)
		for sub := range findings{
			fmt.Println(findings[sub])
		}
	}
}