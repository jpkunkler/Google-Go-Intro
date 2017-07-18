package main

import (
	"fmt"
	"log"
)

// NorgateMathError returns lat and long of Norgate and the error provided
type NorgateMathError struct {
	lat, long string
	err       error
}

// because we connect the Error() method to our NorgateMathError struct, it can be used as an error!
func (n *NorgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-17)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// create our error message
		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)

		// return 0 and our Error struct including our previous error method
		return 0, &NorgateMathError{"50.2289 N", "99.4656 W", nme}
	}

	return 42, nil
}
