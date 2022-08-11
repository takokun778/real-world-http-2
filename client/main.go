package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
	"net/textproto"
	"net/url"
	"os"
	"strings"
)

const root = "http://localhost:18888"

func main() {
	fmt.Println("http.Get()")
	res1, err := http.Get(root)
	if err != nil {
		panic(err)
	}
	defer res1.Body.Close()
	body1, err := ioutil.ReadAll(res1.Body)
	if err != nil {
		panic(err)
	}
	log.Println("Status: ", res1.Status)
	log.Println("StatusCode: ", res1.StatusCode)
	log.Println("Headers: ", res1.Header)
	log.Println(string(body1))

	fmt.Println("http.Get()+query")
	values := url.Values{
		"query": {"hello world"},
	}
	res2, err := http.Get(root + "?" + values.Encode())
	if err != nil {
		panic(err)
	}
	defer res2.Body.Close()
	body2, _ := ioutil.ReadAll(res2.Body)
	log.Println(string(body2))

	fmt.Println("http.Head()")
	res3, err := http.Head(root)
	if err != nil {
		panic(err)
	}
	log.Println("Status: ", res3.Status)
	log.Println("StatusCode: ", res3.StatusCode)
	log.Println("Headers: ", res3.Header)

	fmt.Println("http.PostForm")
	values4 := url.Values{
		"test": {"value"},
	}
	res4, err := http.PostForm(root, values4)
	if err != nil {
		panic(err)
	}
	log.Println("Status: ", res4.Status)

	fmt.Println("http.Post()+body")
	file, err := os.Open("server/main.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	res5, err := http.Post(root, "text/plain", file)
	if err != nil {
		panic(err)
	}
	log.Println("Status: ", res5.Status)

	fmt.Println("http.Post()+io.Reader")
	reader := strings.NewReader("テキスト")
	res6, err := http.Post(root, "text/plain", reader)
	if err != nil {
		panic(err)
	}
	log.Println("Status: ", res6.Status)

	fmt.Println("http.Post()+multipart")
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")
	fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("photo.jpg")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	io.Copy(fileWriter, readFile)
	writer.Close()
	res7, err := http.Post(root, writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("Status: ", res7.Status)

	fmt.Println("http.Post()+mime")
	var buffer2 bytes.Buffer
	writer2 := multipart.NewWriter(&buffer)
	writer2.WriteField("name", "Michael Jackson")
	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "image/jpeg")
	part.Set("Content-Disposition", `form-data; name="thumbnail"; filename="photo.jpg"`)
	fileWriter2, err := writer.CreatePart(part)
	if err != nil {
		panic(err)
	}
	readFile2, err := os.Open("photo.jpg")
	if err != nil {
		panic(err)
	}
	defer readFile2.Close()
	io.Copy(fileWriter2, readFile2)
	writer2.Close()
	res8, err := http.Post(root, writer2.FormDataContentType(), &buffer2)
	if err != nil {
		panic(err)
	}
	log.Println("Status: ", res8.Status)

	fmt.Println("cookie")
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client := http.Client{
		Jar: jar,
	}
	for i := 0; i < 2; i++ {
		res, err := client.Get(root + "/cookie")
		if err != nil {
			panic(err)
		}
		dump, err := httputil.DumpResponse(res, true)
		if err != nil {
			panic(err)
		}
		log.Println(string(dump))
	}

	fmt.Println("DELETE")
	client2 := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, root, nil)
	if err != nil {
		panic(err)
	}
	req.AddCookie(&http.Cookie{Name: "test", Value: "test"})
	req.Header.Add("Content-Type", "application/json")
	res9, err := client2.Do(req)
	if err != nil {
		panic(err)
	}
	dump2, err := httputil.DumpResponse(res9, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump2))
}
