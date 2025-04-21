package main

import (
	"fmt"
	"net/url"
)

func main() {
	const urlStr string = "https://en.wikipedia.org/wiki/Dr._Jekyll_and_Mr._Hyde_(1887_play)"

	fmt.Println("Welcome to url handling")
	// fmt.Println(urlStr)

	// result, err := url.Parse(urlStr)
	// if err != nil {
	// 	fmt.Println("Error parsing URL:", err)
	// 	return
	// }

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)

	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	// qparam := result.Query()
	// fmt.Println("the type of qparam are :%T/n", qparam)
// Create another URL manually (like your LinkedIn profile URL)
newUrl := &url.URL{
	Scheme:   "https",
	Host:     "www.linkedin.com",
	Path:     "/feed/",
	RawQuery: "trk=sem-ga_campid.14650114788_asid.150089839322_crid.656569072777_kw.www%20linkedin_d.c_tid.kwd-2246447582_n.g_mt.e_geo.1007819",
}

// Convert URL struct to string
anotherUrl := newUrl.String()
fmt.Println("The constructed LinkedIn URL is:", anotherUrl)
	
	
}
