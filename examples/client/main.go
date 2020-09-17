package main

import (
	"bufio"
	"fmt"
	"log"

	"gitlab.silkrode.com.tw/golang/fasthttp"
)

type customReader struct {
	*bufio.Reader
}

func (c *customReader) Read(p []byte) (n int, err error) {
	n, err = c.Reader.Read(p)
	fmt.Println(n, err)
	return
}

func main() {
	log.Println("start")
	var (
		strMethod     = []byte("GET")
		strRequestURI = []byte("http://ipv4.download.thinkbroadband.com/1GB.zip")
	)

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()

	req.Header.SetMethodBytes(strMethod)
	req.SetRequestURIBytes(strRequestURI)

	c := &customReader{}
	res.NewBodyReader = func(r *bufio.Reader, h fasthttp.ResponseHeader) fasthttp.Reader {
		log.Println(h.ContentLength())
		c.Reader = r
		return c
	}
	// res.SkipBody = true

	client := fasthttp.Client{}
	if err := client.Do(req, res); err != nil {
		log.Panic(err)
	}

	log.Println("end")
}
