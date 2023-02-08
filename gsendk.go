package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/micmonay/keybd_event"
	"golang.design/x/clipboard"
)

var (
	sleeptime int
	pause     time.Duration
	hlp       bool
	textlines []string
	file      string
	helpStr   string = "\ngsendk designed by mrmioxin@gmail.com with golang v.1.19\r\nThis software comes with ABSOLUTELY NO WARRANTY.\r\ngsendk sending char-data to foreground window from input CSV-file.\r\nUsage: gsendk " +
		`<-h> <-n NNN> <-p NNN> file.csv
Flags: 
-n <seconds> (wait this many seconds before sending text)
-p <milliseconds> (wait this many milliseconds after each key sending)
Special words in input file meaning same keyboard button:
TAB as Tab;
ENTER as Enter;
BSP as Backspace;
ESC as Escape;
UP as Arrow up;
DOWN as Arrow down;
LEFT as Arrow left;
RIGHT as Arrow right;
PAGEDOWN as PageDown;
PAGEUP as PageUp;
HOME as Home;
END as End;
P-<milliseconds> meaning pause. Example P-100 (pause 100 ms).
`
)

func init() {
	flag.IntVar(&sleeptime, "t", 3, "wait this many seconds before sending text")
	flag.DurationVar(&pause, "p", 50, "wait this many milliseconds after each key sending")
	flag.BoolVar(&hlp, "h", false, "help text")
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

}

func main() {
	flag.Parse()
	if hlp || flag.Arg(0) == "" {
		fmt.Println(helpStr)
		return
	}

	file = flag.Arg(0)

	fi, err := os.Open(file)
	if err != nil {
		fmt.Printf("Cant opent file %v\r\n%v\r\n", fi, err)
		return
	}
	defer fi.Close()

	kb, err := keybd_event.NewKeyBonding()
	defer kb.Release()
	if err != nil {
		panic(err)
	}

	Work(fi, kb)

}

func send_CtrlV(kb keybd_event.KeyBonding) {
	// Select keys to be pressed
	kb.SetKeys(keybd_event.VK_V)

	// Set shift to be pressed
	kb.HasCTRL(true)
	err := kb.Launching()
	if err != nil {
		panic(err)
	}
	kb.HasCTRL(false)

	time.Sleep(pause * time.Millisecond)
}

func sendKey(key string, kb keybd_event.KeyBonding) error {
	switch key {
	case "TAB":
		kb.SetKeys(keybd_event.VK_TAB)
	case "ENTER":
		kb.SetKeys(keybd_event.VK_ENTER)
	case "ESC":
		kb.SetKeys(keybd_event.VK_ESC)
	case "BSPACE":
		kb.SetKeys(keybd_event.VK_BACKSPACE)
	case "UP":
		kb.SetKeys(keybd_event.VK_UP)
	case "DOWN":
		kb.SetKeys(keybd_event.VK_DOWN)
	case "LEFT":
		kb.SetKeys(keybd_event.VK_LEFT)
	case "RIGHT":
		kb.SetKeys(keybd_event.VK_RIGHT)
	case "PAGEDOWN":
		kb.SetKeys(keybd_event.VK_PAGEDOWN)
	case "PAGEUP":
		kb.SetKeys(keybd_event.VK_PAGEUP)
	case "HOME":
		kb.SetKeys(keybd_event.VK_HOME)
	case "END":
		kb.SetKeys(keybd_event.VK_END)
	default:
		return fmt.Errorf(`get %v, but expected TAB, ENTER, BSP, ESC, UP, DOWN,LEFT,RIGHT,PAGEDOWN,PAGEUP,HOME,END, P-xxx only`, key)
	}
	err := kb.Launching()
	if err != nil {
		panic(err)
	}
	time.Sleep(pause * time.Millisecond)
	return nil
}

func Work(fi io.Reader, k keybd_event.KeyBonding) {
	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		textlines = append(textlines, scanner.Text())
	}
	fmt.Printf("\ngot %d lines from standard input\n", len(textlines))

	fmt.Printf("\nSleeping for %d seconds. Let's fokus to target window.", sleeptime)
	for n := 0; n != sleeptime; n++ {
		time.Sleep(1 * time.Second)
		print(".")
	}
	println("\r\nSending keys!")
	for _, line := range textlines {
		for _, w := range strings.Split(line, ";") {
			fword := strings.Split(w, "-")
			switch fword[0] {
			case "TAB", "ENTER", "ESC", "BSPACE", "UP", "DOWN", "LEFT", "RIGHT", "PAGEDOWN", "PAGEUP", "HOME", "END":
				if err := sendKey(w, k); err != nil {
					println(err)
					return
				}
			case "P":
				if p, e := strconv.ParseInt(fword[1], 10, 0); e != nil {
					println(e)
					return
				} else {
					time.Sleep(time.Duration(p) * time.Millisecond)
				}
			default:
				clipboard.Write(clipboard.FmtText, []byte(w))
				send_CtrlV(k)
				time.Sleep(50 * time.Millisecond)
			}
		}
	}
	println("\ndone!")

}
