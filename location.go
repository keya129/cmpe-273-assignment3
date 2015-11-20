package main

type Location struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	City      string    `json:"city"`
	State   	string    `json:"state"`
	Zip      	string    		`json:"zip"`
	Coordinate	Coordinate		`json:"cordinate"`
}
type Coordinate struct{
	Lat float64				`json:"lat"`
	Lng float64				`json:"lng"`
}
type BestRoutes struct{
Starting_from_location_id	string `json:"starting_from_location_id"`
Location_ids	[]string `json:"location_ids"`
}
type PriceDist struct{
Price	int `json:"price"`
Distance float64 `json:"distance"`
Duration float64 `json:"duration"`

}
type Trip struct{
	Id          int       `json:"id"`
	Status      string    `json:"status"`
	Starting_from_location_id   string    `json:"starting_from_location_id"`
	Best_route_location_ids  []string    `json:"best_route_location_ids"`
	Total_uber_costs  int    `json:"total_uber_costs"`
	Total_uber_duration	float64	`json:"total_uber_duration"`
	Total_distance	float64	`json:"total_distance"`
	Uber_wait_time_eta int `json:"eta"`
	Count int `json:"count"`

}
type TripFinal struct{
	Id          int       `json:"id"`
	Status      string    `json:"status"`
	Starting_from_location_id   string    `json:"starting_from_location_id"`
	Next_destination_location_id string `json:"next_destination_location_id"`
	Best_route_location_ids  []string    `json:"best_route_location_ids"`
	Total_uber_costs  int    `json:"total_uber_costs"`
	Total_uber_duration	float64	`json:"total_uber_duration"`
	Total_distance	float64	`json:"total_distance"`
	Uber_wait_time_eta int `json:"eta"`
	//Count int `json:"count"`


}
type TripInter struct{
	Id          int       `json:"id"`
	Status      string    `json:"status"`
	Starting_from_location_id   string    `json:"starting_from_location_id"`
	Best_route_location_ids  []string    `json:"best_route_location_ids"`
	Total_uber_costs  int    `json:"total_uber_costs"`
	Total_uber_duration	float64	`json:"total_uber_duration"`
	Total_distance	float64	`json:"total_distance"`
	
}

type TripResp struct{
	Index []int 	`json:"index"`
	SumDistance float64	`json:"sumdistance"`
	SumDuration float64	`json:"sumduration"`
	SumPrice int	`json:"sumprice"`
	Count int `json:"count"`


}
type Buffer struct{
	Product_id string `json:"product_id"`
	Start_latitude float64 `json:"start_latitude"`
	Start_longitude float64 `json:"start_longitude"`
	End_latitude float64 `json:"end_latitude"`
	End_longitude float64 `json:"end_longitude"`		
}
type EtaResp struct{
	Status string `json:"status"`
	Request_id string `json:"request_id"`
	Driver string `json:"driver"`
	Eta int `json:"eta"`
	Location string `json:"location"`
	Vehicle string `json:"vehicle"`
	Surge_multiplier float64 `json:"surge_multiplier"`		
		
}

type Price struct {
	// eg: "08f17084-23fd-4103-aa3e-9b660223934b"
	ProductID string `json:"product_id"`

	// ISO 4217 currency code for situations requiring currency conversion
	// eg: "USD"
	CurrencyCode string `json:"currency_code"`

	// eg: "UberBLACK"
	DisplayName string `json:"display_name"`

	// Formatted string of estimate in local currency of the start location. Estimate
	// could be a range, a single number (flat rate) or "Metered" for TAXI.
	// eg: "$23-29"
	Estimate string `json:"estimate"`

	// The lowest value in the estimate for the given currency
	// eg: 23
	LowEstimate int `json:"low_estimate"`

	// The highest value in the estimate for the given currency
	// eg: 29
	HighEstimate int `json:"high_estimate"`

	// Uber price gouging factor
	// http://www.technologyreview.com/review/529961/in-praise-of-efficient-price-gouging/
	// eg: 1
	SurgeMultiplier float64 `json:"surge_multiplier"`
	Duration float64 `json:"duration"`

	Distance float64 `json:"distance"`

}