// Simple implementation of echo
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var s, sep string
	// The for loop is the only loop statement in Go.
	// It has a number of forms, one of which is illustrated here -
	start := time.Now()
	for j := 0; j < 1000; j++ {
		for i := 1; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
	}
	t := time.Now()
	elapsed := t.Sub(start)
	// any of the parts can be omited
	// for condition {
	// } // is also possible
	fmt.Printf("Time taken when implemented as a loop: %dns\n", elapsed.Nanoseconds()/1000)
	start = time.Now()
	for j := 0; j < 1000; j++ {
		s = strings.Join(os.Args[1:], " ")
	}
	t = time.Now()
	elapsed = t.Sub(start)
	fmt.Printf("Time taken when implemented using join: %dns\n", elapsed.Nanoseconds()/1000)
	fmt.Println(s)

}

// func main() {
//     // This is a range iterator - idx, val
//	   // Go does not permit unused local variables, hence _ is used
//     s, sep := "", ""
//     for _, arg := range os.Args[1:] {
//         s += sep + arg
//         sep = " "
//     }
//     fmt.Println(s)
// }
