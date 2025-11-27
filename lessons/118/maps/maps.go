package maps

import "fmt"

func Maps() {
	//websites := []string{"https://google.com", "https://facebook.com"}
	websites := map[string]string{
		"google":   "https://google.com",
		"facebook": "https://facebook.com",
	}
	websites["stackoverflow"] = "https://stackoverflow.com"
	delete(websites, "facebook")
	fmt.Println(websites["google"])
	fmt.Println(websites)
}
