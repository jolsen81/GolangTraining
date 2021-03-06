package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const input = `It so chanced, that after the Parsee's disappearance, I was he whom
the Fates ordained to take the place of Ahab's bowsman, when that
bowsman assumed the vacant post; the same, who, when on the last day
the three men were tossed from out of the rocking boat, was dropped
astern.  So, floating on the margin of the ensuing scene, and in full
sight of it, when the halfspent suction of the sunk ship reached me,
I was then, but slowly, drawn towards the closing vortex.  When I
reached it, it had subsided to a creamy pool.  Round and round, then,
and ever contracting towards the button-like black bubble at the axis
of that slowly wheeling circle, like another Ixion I did revolve.
Till, gaining that vital centre, the black bubble upward burst; and
now, liberated by reason of its cunning spring, and, owing to its
great buoyancy, rising with great force, the coffin life-buoy shot
lengthwise from the sea, fell over, and floated by my side.  Buoyed
up by that coffin, for almost one whole day and night, I floated on a
soft and dirgelike main.  The unharming sharks, they glided by as if
with padlocks on their mouths; the savage sea-hawks sailed with
sheathed beaks.  On the second day, a sail drew near, nearer, and
picked me up at last.  It was the devious-cruising Rachel, that in
her retracing search after her missing children, only found another
orphan.`

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
