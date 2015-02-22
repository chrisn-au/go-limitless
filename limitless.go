package limitless

import (
  "encoding/binary"
  "net"
  "bytes"
  )

type LimitlessController struct {
  Host string `json:"host"`
  Name string `json:"name"`
  Groups []LimitlessGroup `json:"groups"`
}

type LimitlessGroup struct {
  Id int `json:"id"`
  Type string `json:"type"`
  Name string `json:"name"`
  Controller *LimitlessController `json:"-"`
}

type LimitlessMessage struct {
  Key uint8
  Value uint8
  Suffix uint8
}

const (
  LIMITLESS_ADMIN_PORT = "48899"
  LIMITLESS_PORT = "8899"
)

const MAX_BRIGHTNESS = 0x1b

func NewLimitlessMessage() *LimitlessMessage {
  msg := LimitlessMessage{}
  msg.Suffix = 0x55
  msg.Value = 0x00
  return &msg
}

func (g *LimitlessGroup) SetBri(b uint8) (error) {
  //if b > MAX_BRIGHTNESS {
    //return err
  //}
  msg := NewLimitlessMessage()
  msg.Key = 0x4e
  msg.Value = b
  return g.Controller.sendMsg(msg)
}
func (g *LimitlessGroup) White() (error) {
  msg := NewLimitlessMessage()
  msg.Key = 0x35 //all on
  return g.Controller.sendMsg(msg)
}

func (g *LimitlessGroup) On() (error) {
  msg := NewLimitlessMessage()
  msg.Key = 0x35
  return g.Controller.sendMsg(msg)
}

func (g *LimitlessGroup) Off() (error) {
  msg := NewLimitlessMessage()
  msg.Key = 0x39
  return g.Controller.sendMsg(msg)
}

func (g *LimitlessGroup) Activate() (error) {
  return g.On()
}

func (c *LimitlessController) sendMsg(msg *LimitlessMessage) (error) {
	conn, err := net.Dial("udp", c.Host + ":" + LIMITLESS_PORT)
  if err != nil {
    return err
  }
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, msg)
  untitled folder:
_, err = conn.Write(buf.Bytes())
  return err
}
