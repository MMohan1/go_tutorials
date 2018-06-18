package avarage

func Avarage(x []float64) float64 {
	total_items := len(x)
	var total float64 = 0
	for _,item := range x {
		total += item
	}
	return float64(total/float64(total_items))

}


