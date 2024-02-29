package main

import "fmt"

type Decoder interface {
	Decode() string
}

type Chat struct {
	msg Decoder
}

func (c *Chat) Decode() string {
	return c.msg.Decode()
}

type TxtMessage struct {
	bts []byte
}

func (t *TxtMessage) Decode() string {
	return fmt.Sprintf("txt msg decoder %s\n", string(t.bts))
}

type ImgMessage struct {
	bts []byte
}

func (t *ImgMessage) Decode() string {
	return fmt.Sprintf("img msg decoder %s\n", string(t.bts))
}

func main() {
	t := &TxtMessage{bts: []byte("salasd")}
	i := &ImgMessage{bts: []byte("dssadas")}

	p := Chat{msg: t}
	fmt.Println(p.Decode())

	pp := Chat{msg: i}
	fmt.Println(pp.Decode())
}
