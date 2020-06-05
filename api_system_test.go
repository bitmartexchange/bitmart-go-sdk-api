package bitmart

import (
	"log"
	"testing"
)


// GET https://api-cloud.bitmart.com/system/time
func TestGetSystemTime(t *testing.T) {
	c := NewTestClient()
	ac, err := c.getSystemTime()
	if err != nil {
		log.Panic(err)
	} else {
		printResponse(ac)
	}
}

//GET https://api-cloud.bitmart.com/system/service
func TestGetSystemService(t *testing.T) {
	c := NewTestClient()
	ac, err := c.getSystemService()
	if err != nil {
		log.Panic(err)
	} else {
		printResponse(ac)
	}
}



