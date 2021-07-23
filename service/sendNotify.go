package service

import (
	"bytes"
	"errors"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
)

type Notify struct{}

func NewNotify() *Notify {
	return &Notify{}
}

func (n *Notify) Send(msg string, fileName string) *http.Response {
	////data := url.Values{}
	////data.Set("message", msg)
	////data.Set("imageFile",img)
	////
	//body := strings.NewReader(url.Values{
	//	"message":       []string{msg},
	//	"imageFile": []string{"https://engineering.linecorp.com/wp-content/uploads/2016/11/LINE_Notify_02_Cony.jpg"},
	//}.Encode())
	//
	//resp, _ := http.NewRequest("POST", "https://notify-api.line.me/api/notify", body)
	//resp.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//resp.Header.Set("Authorization", "Bearer gBA8EpM7DbfYftLZHILBzSBUklR69HSB76BMYMopJLP")
	//
	//client := &http.Client{}
	//response, _ := client.Do(resp)
	//fmt.Println(response)
	//defer response.Body.Close()
	//defer resp.Body.Close()

	URL := "https://notify-api.line.me/api/notify"
	accessToken := "gBA8EpM7DbfYftLZHILBzSBUklR69HSB76BMYMopJLP"
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	c := &http.Client{}

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	fw, err := w.CreateFormField("message")
	if err != nil {
		panic(err)
	}
	if _, err = fw.Write([]byte(msg)); err != nil {
		panic(err)
	}

	part := make(textproto.MIMEHeader)
	part.Set("Content-Disposition", `form-data; name="imageFile"; filename=`+fileName)

	imgBase, format, err := checkImageFormat(f, fileName)
	if err != nil {
		panic(err)
	}

	if format == "jpeg" {
		part.Set("Content-Type", "image/jpeg")
	} else if format == "png" {
		part.Set("Content-Type", "image/png")
	} else {
		panic("LINE Notify supports only jpeg/png image format")
	}

	fw, err = w.CreatePart(part)
	if err != nil {
		panic(err)
	}

	io.Copy(fw, imgBase)
	w.Close() // boundaryの書き込み
	req, err := http.NewRequest("POST", URL, &b)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		panic("failed to send image, get http status code: " + resp.Status)
	}

	return nil
}

func checkImageFormat(r io.Reader, filename string) (io.Reader, string, error) {
	ext := filepath.Ext(filename)

	var b bytes.Buffer
	if ext == ".jpeg" || ext == ".jpg" || ext == ".JPEG" || ext == ".JPG" {
		ext = "jpeg"
		img, err := jpeg.Decode(r)
		if err != nil {
			return nil, "", err
		}

		if err := jpeg.Encode(&b, img, &jpeg.Options{Quality: 100}); err != nil {
			return nil, "", err
		}
	} else if ext == ".png" || ext == ".PNG" {
		ext = "png"
		img, err := jpeg.Decode(r)
		if err != nil {
			return nil, "", err
		}

		if err = png.Encode(&b, img); err != nil {
			return nil, "", err
		}
	} else {
		return nil, "", errors.New("Image format must be jpeg or png")
	}

	return &b, ext, nil
}
