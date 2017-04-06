package main

import (
	//	"fmt"
	"testing"
)

func Test_genWordMap(t *testing.T) {
	text := `HTTP/2 is designed to address many of the failings 
	of HTTP/1.x. Modern web pages use many resources: HTML, stylesheets, scripts, 
	images, and so on. In HTTP/1.x, each of these resources must be requested explicitly. 
	This can be a slow process. The browser starts by fetching the HTML, then learns of 
	more resources incrementally as it parses and evaluates the page. Since the server 
	must wait for the browser to make each request, the network is often idle and underutilized.
	=========================own test========================================================
	products is decrese 12%; and $100.i don't konw anythings.`
	wordmap := genWordMap(text)
	t.Error(wordmap)
}
