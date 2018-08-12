/**
 * @Author codeAC
 * @Time: 2018/8/11 21:57
 * @Description
 */

package main

import (
	"fmt"
	"net/http"
)

func main() {
	request, _ := http.NewRequest(http.MethodGet,
		"http://www.imooc.com", nil)
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	//dumpResponse, e := httputil.DumpResponse(response, true)
	//if e != nil{
	//	panic(e)
	//}
	//fmt.Printf("%s\n", dumpResponse)
}
