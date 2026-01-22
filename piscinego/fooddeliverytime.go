package piscine

func FoodDeliveryTime(order string) int {
	switch order {
	case "burger":
		return 15
	case "chips":
		return 10
	case "nuggets":
		return 12
	default:
		return 404
	}
}
