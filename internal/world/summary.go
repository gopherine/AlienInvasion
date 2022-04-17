package world

func TotalCityDestroyed(cities []*City) (count int) {
	for _, v := range cities {
		if v.Destroyed {
			count += 1
		}
	}

	return
}
