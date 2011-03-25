// Copyright (c) 2011 Dmitry Chestnykh
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package porter2

import (
	"testing"
	"bufio"
	"os"
)

func TestStem(t *testing.T) {
	voc, err := os.Open("test_voc.txt", os.O_RDONLY, 0600)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	defer voc.Close()
	output, err := os.Open("test_output.txt", os.O_RDONLY, 0600)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	defer output.Close()
	bvoc := bufio.NewReader(voc)
	bout := bufio.NewReader(output)
	for {
		vocline, err := bvoc.ReadString('\n')
		if err != nil {
			switch err {
			case os.EOF:
				return
			default:
				t.Errorf("%s", err)
				return
			}
		}
		outline, err := bout.ReadString('\n')
		if err != nil {
			switch err {
			case os.EOF:
				return
			default:
				t.Errorf("%s", err)
				return
			}
		}
		vocline = vocline[:len(vocline)-1]
		outline = outline[:len(outline)-1]
		st := Stemmer.Stem(vocline)
		if st != outline {
			t.Errorf("\"%s\" expected %q got %q", vocline, outline, st)
		}
	}
}
