package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/micmonay/keybd_event"
)

func TestWork(t *testing.T) {
	test := `1;TAB;11;TAB;13;TAB;21;TAB;ENTER
ываы1;TAB;2миттт ти;TAB;3выпапвп;TAB;4йцу;TAB
`
	in := bufio.NewReader(strings.NewReader(test))
	k, _ := keybd_event.NewKeyBonding()
	Work(in, k)

}
