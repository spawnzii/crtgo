package main


import(
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
			crt.FetchCrt(domain)
		}
	}else {
		crt.FetchCrt(*domain)
	}
}
