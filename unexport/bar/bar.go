package bar

import (
	"fmt"
	_ "unsafe"
)

//go:linkname pb bar.pb
func pb(){
	fmt.Println("welcome pb")
}
