package main

import "testing"

func TestGetAllArticles(t *testing.T) {
	a := getAllArticles()
	if len(a) == 0 {
		t.Fail()
	}

}
