package main

import "fmt"

type Message interface {
	PrintMsg()
}

type AudioMessage struct {
	bytes string
}

func (a *AudioMessage) PrintMsg() {
	fmt.Printf("Audiomsg :%s\n", a.bytes)
}

type VideoMessage struct {
	bytes string
}

func (v *VideoMessage) PrintMsg() {
	fmt.Printf("Video :%s\n", v.bytes)
}

type Decoder interface {
	Decode([]byte) Message
}

type VideoDecoder struct{}

func (v *VideoDecoder) Decode(str []byte) Message {
	return &VideoMessage{bytes: string(str)}
}

type AudioDecoder struct{}

func (a *AudioDecoder) Decode(str []byte) Message {
	return &AudioMessage{bytes: string(str)}
}

func main() {
	a := AudioDecoder{}
	v := VideoDecoder{}

	aud := a.Decode([]byte("asdlasdl;asdl;asda"))
	aud.PrintMsg()

	vid := v.Decode([]byte("asokdpasd"))
	vid.PrintMsg()
}
