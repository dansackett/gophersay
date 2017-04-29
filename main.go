//go:generate go-bindata -o gopherart/gopherart.go -pkg gopherart gopherart/gopher.ascii
package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/dansackett/gophersay/gopherart"
)

var sayings []string
var gopherArt string
var stdIn bool

func init() {

	// Load in gopher ascii art from go-bindata
	gopherArtBytes, _ := gopherart.Asset("gopherart/gopher.ascii")
	gopherArt = string(gopherArtBytes)
	sayings = []string{
		"There is no friend as loyal as a book.",
		"There is nothing to writing. All you do is sit down at a typewriter and bleed.",
		"Happiness in intelligent people is the rarest thing I know.",
		"I love sleep. My life has the tendency to fall apart when I'm awake, you know?",
		"The best way to find out if you can trust somebody is to trust them.",
		"Always do sober what you said you'd do drunk. That will teach you to keep your mouth shut.",
		"There is nothing noble in being superior to your fellow man; true nobility is being superior to your former self.",
		"The world breaks everyone, and afterward, many are strong at the broken places.",
		"The first draft of anything is shit.",
		"The most painful thing is losing yourself in the process of loving someone too much, and forgetting that you are special too.",
		"All you have to do is write one true sentence. Write the truest sentence that you know.",
		"I drink to make other people more interesting.",
		"When people talk, listen completely. Most people never listen.",
		"All good books are alike in that they are truer than if they had really happened and after you are finished reading one you will feel that all that happened to you and afterwards it all belongs to you: the good and the bad, the ecstasy, the remorse and sorrow, the people and the places and how the weather was. If you can get so that you can give that to people, then you are a writer.",
		"Every day is a new day. It is better to be lucky. But I would rather be exact. Then when luck comes you are ready.",
		"All thinking men are atheists.",
		"Courage is grace under pressure.",
		"Never think that war, no matter how necessary, nor how justified, is not a crime.",
		"But man is not made for defeat, A man can be destroyed but not defeated.",
		"Every man's life ends the same way. It is only the details of how he lived and how he died that distinguish one man from another.",
		"The world breaks every one and afterward many are strong at the broken places.",
		"Never confuse movement with action.",
		"A cat has absolute emotional honesty: human beings, for one reason or another, may hide their feelings, but a cat does not.",
		"You can't get away from yourself by moving from one place to another.",
		"Forget your personal tragedy. We are all bitched from the start and you especially have to be hurt like hell before you can write seriously. But when you get the damned hurt, use it-don't cheat with it.",
		"I’m not brave any more darling. I’m all broken. They’ve broken me.",
	}
	flag.Parse()
	if terminal.IsTerminal(int(os.Stdin.Fd())) {
		stdIn = false
	} else {
		stdIn = true
	}
}

func main() {
	var saying string

	// If there are any command line arguments, join them together with spaces
	// and output them as the saying
	// otherwise output a saying from the list
	if stdIn {
		data := os.Stdin
		scan := bufio.NewScanner(data)
		var stdInSlice []string
		for scan.Scan() {
			stdInSlice = append(stdInSlice, scan.Text()+"\n")
		}
		// Take the newline off the last line
		lastLine := stdInSlice[len(stdInSlice)-1]
		lastLine = strings.TrimSuffix(lastLine, "\n")
		stdInSlice[len(stdInSlice)-1] = lastLine

		// Join the lines
		saying = strings.Join(stdInSlice, "")

	} else if len(flag.Args()) > 0 {
		saying = strings.Join(flag.Args(), " ")
	} else {
		// Seed rand for psudo random numbers
		rand.Seed(time.Now().UnixNano())

		saying = sayings[rand.Intn(len(sayings))]
	}

	fmt.Println(" ------------------------")
	fmt.Println(saying)
	fmt.Print(gopherArt)
}
