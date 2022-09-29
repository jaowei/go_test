package main

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"os/exec"

	"github.com/disintegration/imaging"
)

func Decode(r io.Reader) (image.Image, error) {
	fmt.Printf("%T\n", r)

	img, _, err := image.Decode(r)
	return img, err


	// pr, pw := io.Pipe()
	// r = io.TeeReader(r, pw)
	// done := make(chan struct{})
	// go func() {
	// 	defer close(done)
	// 	io.Copy(ioutil.Discard, pr)
	// }()

	// img, _, err := image.Decode(r)
	// pw.Close()
	// <-done
	// if err != nil {
	// 	return nil, err
	// }

	// return img, nil
}


func main() {
	cmd := exec.Command(
		"ffmpeg",
		"-f",
		"dshow",
		"-i",
		"video=@device_pnp_\\\\?\\usb#vid_04f2&pid_b65a&mi_00#6&fa823cd&0&0000#{65e8773d-8f56-11d0-a3b9-00a0c9223196}\\global",
		"-frames:v",
		"1",
		"-c:v",
		"png",
		"-f",
		"image2pipe",
		"-hide_banner",
		"-loglevel",
		"error",
		"pipe:1",
	)
	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Running ffmpeg failed: %v", err)
	}
	buf := bytes.NewBuffer(b)

	// err := ffmpeg.Input("video=@device_pnp_\\\\?\\usb#vid_04f2&pid_b65a&mi_00#6&fa823cd&0&0000#{65e8773d-8f56-11d0-a3b9-00a0c9223196}\\global", ffmpeg.KwArgs{"format": "dshow"}).
	// Output("pipe:", ffmpeg.KwArgs{"frames:v": 1, "format": "image2", "vcodec": "mjpeg"}).
	// WithOutput(buf, os.Stdout).
	// Run()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%T\n", buf)

	img , err := imaging.Decode(buf)

	if err != nil {
		fmt.Printf("%T\n",img)
		fmt.Println(err)
	}

	err = imaging.Save(img, "./out1.png")
	if err != nil {
		fmt.Println(err)
	}
	// var r io.Reader = buf
	// img, _, err := image.Decode(bytes.NewReader(buf.Bytes()))
	// if err != nil {
	// 	fmt.Printf("%T\n",img)
	// 	fmt.Println(err)
	// }
}