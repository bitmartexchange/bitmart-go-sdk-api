package bitmart

import (
	"log"
	"testing"
)


// GET https://api-cloud.bitmart.com/system/time
func TestGetSystemTime(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSystemTime()
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

//GET https://api-cloud.bitmart.com/system/service
func TestGetSystemService(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSystemService()
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}



