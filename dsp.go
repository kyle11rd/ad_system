package main

import "math/rand"

const max_daily_budget_from_rand float32 = 10 //total allowance a campaign is allowed to spend
const max_bid_from_rand float32 = 5           //max cost of a single bid a campaign is allowed

//Demand Side Platform
//initialize a list of campaigns (daily budget & max bid)
//let's also assume each campaign (ad) has a fixed CTR for simplicity
func dsp(campaign_cnt int) (max_bids []float32, daily_budgets []float32, ctrs []float32) {
	max_bids = make([]float32, campaign_cnt)
	daily_budgets = make([]float32, campaign_cnt)
	ctrs = make([]float32, campaign_cnt)

	for i := 0; i < campaign_cnt; i++ {
		max_bids[i] = rand.Float32() * max_daily_budget_from_rand
		daily_budgets[i] = rand.Float32() * max_bid_from_rand
		ctrs[i] = rand.Float32() / 10 //Assume ad CTR is top at 10%
	}
	return
}
