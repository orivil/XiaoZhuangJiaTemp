package site

import (
	"gopkg.in/orivil/orivil.v1"
)

type Controller struct {
	*orivil.App
}

// @route {get}/
func (this *Controller) Index() {

	this.View("temp")
}