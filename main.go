package main

func main() {
}
func storeRevenue(stores [][]int) (indexOfMaxRevenue, indexOfAverageDailyRevenue, indexTheBestDailyRevenue, indexTheWorstDailyRevenue []int) {
	maxBestDailyRevenue := stores[0][0]
	minWorstDailyRevenue := stores[0][0]
	maxRevenue := 0
	maxAverageDailyRevenue := 0
	indexMax :=0
	indexAverageDaily := 0
	indexBestBaily := 0
	indexWorstDaily := 0
	for _, store := range stores {
		searchRevenue(store, &maxBestDailyRevenue, &minWorstDailyRevenue, &maxRevenue, &maxAverageDailyRevenue)
	}
	for index, store := range stores {
		indexMax, indexAverageDaily, indexBestBaily, indexWorstDaily = employeeComparison(index, store, maxBestDailyRevenue, minWorstDailyRevenue, maxRevenue, maxAverageDailyRevenue)
		assignment(indexMax, &indexOfMaxRevenue)
		assignment(indexAverageDaily, &indexOfAverageDailyRevenue)
		assignment(indexBestBaily, &indexTheBestDailyRevenue)
		assignment(indexWorstDaily, &indexTheWorstDailyRevenue)
	}
	return
}
func searchRevenue(openingHours []int, maxBestDailyRevenue, minWorstDailyRevenue, maxRevenue, maxAverageDailyRevenue *int)  {
	revenue := 0
	for _, value := range openingHours {
		if *maxBestDailyRevenue < value {
			*maxBestDailyRevenue = value
		}
		if *minWorstDailyRevenue > value {
			*minWorstDailyRevenue = value
		}
		revenue += value
	}
	if *maxRevenue < revenue {
		*maxRevenue = revenue
	}
	if revenue/len(openingHours) > *maxAverageDailyRevenue {
		*maxAverageDailyRevenue = revenue / len(openingHours)
	}
}
func employeeComparison(index int, comparisonOfValue []int, maxBestDailyRevenue, minWorstDailyRevenue, maxRevenue, maxAverageDailyRevenue int) (indexMax, indexAverageDaily, indexBestBaily, indexWorstDaily int) {
	revenue := 0
	quantityInOneEmployeeForMax := 0
	quantityInOneEmployeeForMin := 0
	for _, search := range comparisonOfValue {
		revenue += search
		if maxBestDailyRevenue == search && quantityInOneEmployeeForMax == 0 {
			indexBestBaily = index + 1
			quantityInOneEmployeeForMax++
		}
		if minWorstDailyRevenue == search && quantityInOneEmployeeForMin == 0 {
			indexWorstDaily = index + 1
			quantityInOneEmployeeForMin++
		}
	}
	if revenue == maxRevenue {
		indexMax = index + 1
	}
	maxAverage := revenue / len(comparisonOfValue)
	if maxAverage == maxAverageDailyRevenue {
		indexAverageDaily = index + 1
	}
	return
}
func assignment(value int, values *[]int) {
	if value != 0 {
		*values = append(*values, value )
	}
}
