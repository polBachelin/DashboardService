package query

import "go.mongodb.org/mongo-driver/bson"

func generateOrderStage(order Order) bson.M {
	sort := bson.M{}
	for i, dimension := range order.DimensionName {
		sort[dimension] = getOrderType(getMemberName(order.DimensionOrder[i]))
	}
	for i, measure := range order.MeasureName {
		sort[measure] = getOrderType(getMemberName(order.MeasureOrder[i]))
	}
	return bson.M{"$sort": sort}
}

func getOrderType(orderType string) int {
	switch orderType {
	case "asc":
		return 1
	case "desc":
		return -1
	default:
		return 1
	}
}
