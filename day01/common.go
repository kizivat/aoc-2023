package main

import (
	"strings"
	"sync"
	"sync/atomic"
)

type DigitFinder func(string, chan uint8)

func calcCalibrationValues(in string, findFirst DigitFinder, findLast DigitFinder) uint32 {
	lines := strings.Split(in, "\n")
	n := len(lines)

	firstDigits := make(chan uint8, n)
	lastDigits := make(chan uint8, n)

	var sumFirstDigits, sumLastDigits uint32

	var wg sync.WaitGroup

	for _, line := range lines {
		// this below might be bad practice, not sure
		// it causes the findFirst and findLast functions to ought to
		// push a value to the channel, or else the program will hang
		// as the channel keeps waiting for a value, thus never
		// reaching the wg.Done() call
		wg.Add(2)
		go findFirst(line, firstDigits)
		go findLast(line, lastDigits)
		go func() {
			atomic.AddUint32(&sumFirstDigits, uint32(<-firstDigits))
			wg.Done()
		}()
		go func() {
			atomic.AddUint32(&sumLastDigits, uint32(<-lastDigits))
			wg.Done()
		}()
	}

	wg.Wait()

	return sumFirstDigits*10 + sumLastDigits
}
