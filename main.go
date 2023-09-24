package main

import "fmt"

func main() {
	//input = # of daily campaigns
	var campaigns int = 1000
	max_bids, daily_budgets, ctrs := dsp(campaigns)

	//daily_slots = how many impressions can be displayed per day
	//campaigns_per_selection = how many compaigns the simulator will randomly grab to choose a bid winner
	daily_slots, campaigns_per_selection := ssp(2000, 10)

	totalImp, totalClick, totalCost := aex(campaigns, max_bids, daily_budgets, ctrs, daily_slots, campaigns_per_selection)

	fmt.Printf("Total Impressions: %d\n", totalImp)
	fmt.Printf("Total Clicks: %d\n", totalClick)
	fmt.Printf("Total Cost: %.2f\n", totalCost)
	fmt.Printf("Ad CTRs: %.2f%%\n", float32(totalClick)/float32(totalImp)*100)
	fmt.Printf("CPM: $%.4f\n", totalCost/float32(totalImp)*1000)
	fmt.Printf("CPC: $%.4f\n", totalCost/float32(totalClick))

}
