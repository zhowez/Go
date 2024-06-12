package maps

import "fmt"


func main() {

	websites := map[string]string{
		"Google": "https://google.com",
		"Amazon Web Service": "http://aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Google"])

	websites["Linkedin"] = "http://linkedin.com"
	fmt.Println(websites)

	delete(websites,"Google")
	fmt.Println(websites)
}