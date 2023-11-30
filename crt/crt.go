package crt


import(
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
)


func FetchCrt(domain string) {
	var findings []string

    url := fmt.Sprintf("https://crt.sh/?q=%s", domain)
    resp, _ := http.Get(url)
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)

    content := string(body)
    domains := strings.Split(content, ">")

    for _, domainItem := range domains {
        if strings.Contains(domainItem, "</TD") && strings.Contains(domainItem, domain) && !strings.Contains(domainItem, ":") {
            cleardomain := strings.Split(domainItem, "</TD")
            for _, clearItem := range cleardomain {
                alreadyExists := false
                for _, existingItem := range findings {
                    if clearItem == existingItem {
                        alreadyExists = true
                    }
                }
                if !alreadyExists && strings.Contains(clearItem,"@")!=true && strings.Contains(clearItem,"*")!=true{
			if len(clearItem)>=4{
				findings = append(findings, clearItem)
				fmt.Println(clearItem)
			}
                }
            }
        }
    }

}
