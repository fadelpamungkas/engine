package models

type Restaurant struct {
	ID           string `json:"id" bson:"_id"`
	RestaurantId string `json:"restaurant_id" bson:"restaurant_id"`
	Name         string `json:"name" bson:"name"`
	Cuisine      string `json:"cuisine" bson:"cuisine"`
	Borough      string `json:"borough" bson:"borough"`
}
