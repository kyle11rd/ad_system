package main

import "math/rand"

//Ad Exchange
func aex(campaigns int, max_bids []float32, daily_budgets []float32, ctrs []float32, daily_slots int, campaigns_per_selection int) (totalImp int, totalClick int, totalCost float32) {
	totalImp = 0
	totalClick = 0
	totalCost = 0
	for i := 0; i < daily_slots; i++ {
		//at the beginning of each impression, initialize a list of campaigns to choose from
		chosenCampaignIndexes := make([]int, campaigns_per_selection)
		for j := 0; j < campaigns_per_selection; j++ {
			chosenCampaignIndexes[j] = rand.Intn(campaigns) //may repeat, but as long as we have campaigns>>campaigns_per_selection, we have no issue
		}

		//let's use ad_ranking = max_cpc x ctr to select chosen campaign
		ad_ranks := make([]float32, campaigns_per_selection)
		for j := 0; j < campaigns_per_selection; j++ {
			ad_ranks[j] = max_bids[chosenCampaignIndexes[j]] * ctrs[chosenCampaignIndexes[j]]
		}

		//let's rank the campaigns based on the calculated ad rank score
		//can't simply choose the top 1 since the chosen one might not have enough budget to get displayed
		chosenCampaignIndexesRanked := make([]int, campaigns_per_selection)
		for j := 0; j < campaigns_per_selection; j++ {
			maxIndx := findIndexOfMaxInFloat32Array(ad_ranks)
			chosenCampaignIndexesRanked[j] = chosenCampaignIndexes[maxIndx]
			ad_ranks[maxIndx] = -1
		}

		//and calculate how much the campaign needs to pay:
		//if the chosen campaign doesn't have the highest max_bid, use max_bid
		//else, use the next highest max_bid + 1 cent
		//in either of the above cases, if the remaining budget for the campaign is not available to pay the bill, we choose the next campaign
		//if no chosen campaign can pay the bid, we show no ad
		maxBidsInSelection := make([]float32, campaigns_per_selection)
		for j := 0; j < campaigns_per_selection; j++ {
			maxBidsInSelection[j] = max_bids[chosenCampaignIndexesRanked[j]]
		}
		chosenCampaignIndx := -1
		var chosenCost float32 = -1
		for j := 0; j < campaigns_per_selection; j++ {
			tempMaxBid := maxBidsInSelection[j]
			maxBidsInSelection[j] = -1 //ignore itself when looking for the max from the array
			cost := calcCost(tempMaxBid, maxBidsInSelection)
			if daily_budgets[chosenCampaignIndexesRanked[j]] >= cost {
				chosenCampaignIndx = chosenCampaignIndexesRanked[j]
				daily_budgets[chosenCampaignIndx] = daily_budgets[chosenCampaignIndx] - cost
				chosenCost = cost
				break
			}
		}
		if chosenCampaignIndx != -1 { //if have an ad impression available
			totalImp = totalImp + 1
			clickChance := rand.Float32()
			if clickChance < ctrs[chosenCampaignIndx] { // randomly generate a number and if this number is within CTR of the campaign, we say it has a click
				totalClick = totalClick + 1
				totalCost = totalCost + chosenCost
			}
		}
	}
	return
}

func findIndexOfMaxInFloat32Array(a []float32) (index int) {
	var maxVal float32 = -1
	for i := 0; i < len(a); i++ {
		if a[i] > maxVal {
			maxVal = a[i]
			index = i
		}
	}
	return
}

func calcCost(maxBid float32, listOfBidsToCompare []float32) (cost float32) {
	cost = maxBid
	var maxInList float32 = -1
	for i := 0; i < len(listOfBidsToCompare); i++ {
		if maxInList < listOfBidsToCompare[i] {
			maxInList = listOfBidsToCompare[i]
		}
	}
	if cost > maxInList && maxInList > 0 {
		cost = maxInList + 0.01
	}
	return
}
