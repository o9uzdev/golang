package environments

import (
	"fmt"
	"os"
)

func uNameOS() string {
	// for Windows
	// uHomePathWin := os.Getenv("USERNAME")
		
	return os.Getenv("USER")
}

func uHomePathOS() string {
	// for Windows
	// uHomePathWin := os.Getenv("HOMEPATH")
	
	return os.Getenv("HOME")
}

func main() {
	// for mac
	//uNameOS := os.Getenv("USER")
	//uHomePathOS := os.Getenv("HOME")
	
	fmt.Println(uNameOS())
	fmt.Println(uHomePathOS())
}
