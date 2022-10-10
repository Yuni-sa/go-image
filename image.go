package main

import "net/textproto"

type Image struct {
	Name     string               `json:"imageName"`
	ImageUrl string               `json:"imageUrl"`
	Header   textproto.MIMEHeader `json:"header"`
	Size     int64                `json:"size"`
}
