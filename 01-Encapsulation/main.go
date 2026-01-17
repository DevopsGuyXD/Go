//	Encapsuation is the process of bundling your data along with it's logic(methods) for manipulating or returning data.
//  It's primary goal is to hide data and prevent direct acccess to the underlying logic
// 	Mutable fields need to be modified using setters
//  Data should be accessed using getters 
 
package main

import (
	"time"

	"github.com/devopsguyXD/test/01-Encapsulation/music"
)

func main() {
	act1 := music.NewMember("John", "Vocals")

	act1.PlaySong()
	time.Sleep(3 * time.Second)
	act1.StopSong()
}