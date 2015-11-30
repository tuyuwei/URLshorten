package helpers

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	_ "log"
	"strings"
)

func URLshorten(url string) [4]string {
	//加密字符串
	var key string = "URLshorten"
	text := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	s := strings.Split(text, "")
	h := md5.New()
	h.Write([]byte(url + key))
	md5text := hex.EncodeToString(h.Sum(nil))
	md5textByte := []byte(md5text)
	var shortURLs [4]string
	var x, num int32
	data := make(chan string, 4)
	for i := 0; i < 4; i++ {
		binary.Read(bytes.NewBuffer(md5textByte[i*8:i*8+8]), binary.BigEndian, &x)
		num = x & 0x3FFFFFFF
		go handler(num, s, data)
		shortURLs[i] = <-data
	}
	close(data)
	return shortURLs
}

func handler(num int32, source []string, data chan<- string) {
	shortURL := ""
	for j := 0; j < 6; j++ {
		shortURL += source[num&0x0000003D]
		num >>= 5
	}
	data <- shortURL
}
