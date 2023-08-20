package main

import "fmt"

func main() {
	bids := dsp()
	slots := ssp()
	deals := aex(bids, slots)

	fmt.Println(deals)

}

//Demand Side Platform
func dsp() (bids []float32) {
	bids = make([]float32, 2) //$ advertisers would like to pay for an impression or click
	bids[0] = 1
	bids[1] = 2
	return
}

//Supply Side Platform
func ssp() (slots []bool) {
	slots = append(slots, true) //if has a slot for ad
	return
}

//Ad Exchange
func aex(bids []float32, slots []bool) (deals []float32) {
	if bids[0] >= bids[1] {
		deals = append(deals, bids[0]) //confirmed cost for advertiser
	} else {
		deals = append(deals, bids[1])
	}
	return
}
