package main

func main() {
}
func StoreRevenue(firstStore, secondStore, thirdStore []int) (IndexOfMaxRevenue, IndexOfAverageDailyRevenue, indexTheBestDailyRevenue, indexTheWorstDailyRevenue []int){
	maxBestDailyRevenue := firstStore[0]
	minWorstDailyRevenue := firstStore[0]
	maxRevenue := 0
	maxAverageDailyRevenue := 0
	maxBestDailyRevenue, minWorstDailyRevenue, maxRevenue, maxAverageDailyRevenue = searchRevenue(firstStore, maxBestDailyRevenue, minWorstDailyRevenue, maxRevenue, maxAverageDailyRevenue)
	maxBestDailyRevenue, minWorstDailyRevenue, maxRevenue, maxAverageDailyRevenue = searchRevenue(secondStore, maxBestDailyRevenue, minWorstDailyRevenue, maxRevenue, maxAverageDailyRevenue)
	maxBestDailyRevenue, minWorstDailyRevenue, maxRevenue, maxAverageDailyRevenue = searchRevenue(thirdStore, maxBestDailyRevenue, minWorstDailyRevenue, maxRevenue, maxAverageDailyRevenue)
	increment := 1
	//IndexOfMaxRevenue1, IndexOfAverageDailyRevenue1, indexTheBestDailyRevenue1, indexTheWorstDailyRevenue1 := []int{}
	IndexOfMaxRevenue1, IndexOfAverageDailyRevenue1, indexTheBestDailyRevenue1, indexTheWorstDailyRevenue1 := employeeComparison(increment, firstStore,maxBestDailyRevenue, minWorstDailyRevenue, maxRevenue, maxAverageDailyRevenue,IndexOfMaxRevenue, IndexOfAverageDailyRevenue, indexTheBestDailyRevenue, indexTheWorstDailyRevenue)
	increment++
	IndexOfMaxRevenue, IndexOfAverageDailyRevenue, indexTheBestDailyRevenue, indexTheWorstDailyRevenue = assignment(IndexOfMaxRevenue1,IndexOfMaxRevenue), assignment(IndexOfAverageDailyRevenue1,IndexOfAverageDailyRevenue), assignment(indexTheBestDailyRevenue1, indexTheBestDailyRevenue), assignment(indexTheWorstDailyRevenue1, indexTheWorstDailyRevenue)
	IndexOfMaxRevenue1, IndexOfAverageDailyRevenue1, indexTheBestDailyRevenue1, indexTheWorstDailyRevenue1 = employeeComparison(increment, secondStore,maxBestDailyRevenue, minWorstDailyRevenue, maxRevenue, maxAverageDailyRevenue, IndexOfMaxRevenue, IndexOfAverageDailyRevenue, indexTheBestDailyRevenue, indexTheWorstDailyRevenue)
	increment++
	IndexOfMaxRevenue, IndexOfAverageDailyRevenue, indexTheBestDailyRevenue, indexTheWorstDailyRevenue = assignment(IndexOfMaxRevenue1,IndexOfMaxRevenue), assignment(IndexOfAverageDailyRevenue1,IndexOfAverageDailyRevenue), assignment(indexTheBestDailyRevenue1, indexTheBestDailyRevenue), assignment(indexTheWorstDailyRevenue1, indexTheWorstDailyRevenue)
	IndexOfMaxRevenue1, IndexOfAverageDailyRevenue1, indexTheBestDailyRevenue1, indexTheWorstDailyRevenue1 = employeeComparison(increment, thirdStore,maxBestDailyRevenue, minWorstDailyRevenue, maxRevenue, maxAverageDailyRevenue, IndexOfMaxRevenue, IndexOfAverageDailyRevenue, indexTheBestDailyRevenue, indexTheWorstDailyRevenue)
	IndexOfMaxRevenue, IndexOfAverageDailyRevenue, indexTheBestDailyRevenue, indexTheWorstDailyRevenue = assignment(IndexOfMaxRevenue1,IndexOfMaxRevenue), assignment(IndexOfAverageDailyRevenue1,IndexOfAverageDailyRevenue), assignment(indexTheBestDailyRevenue1, indexTheBestDailyRevenue), assignment(indexTheWorstDailyRevenue1, indexTheWorstDailyRevenue)
	return
}
func searchRevenue(openingHours []int, maxBestDailyRevenue1, minWorstDailyRevenue1, maxRevenue1, maxAverageDailyRevenue1  int) (maxBestDailyRevenue, minWorstDailyRevenue, maxRevenue, maxAverageDailyRevenue int){
	revenue := 0
	maxBestDailyRevenue = maxBestDailyRevenue1
	minWorstDailyRevenue = minWorstDailyRevenue1
	maxRevenue = maxRevenue1
	maxAverageDailyRevenue = maxAverageDailyRevenue1
	for _, value := range openingHours{
		if maxBestDailyRevenue < value{
			maxBestDailyRevenue = value
		}
		if minWorstDailyRevenue > value {
			minWorstDailyRevenue = value
		}
		revenue += value
	}
	if maxRevenue < revenue{
		maxRevenue = revenue
	}
	if revenue / len(openingHours) > maxAverageDailyRevenue {
		maxAverageDailyRevenue = revenue / len(openingHours)
	}
	return
}
func employeeComparison(increment int, comparisonOfValue []int, maxBestDailyRevenue, minWorstDailyRevenue, maxRevenue, maxAverageDailyRevenue int, IndexOfMaxRevenue, IndexOfAverageDailyRevenue, indexTheBestDailyRevenue, indexTheWorstDailyRevenue []int) (IndexOfMaxRevenue1, IndexOfAverageDailyRevenue1, indexTheBestDailyRevenue1, indexTheWorstDailyRevenue1 int){
	revenue := 0
	quantityInOneEmployee := 0
	quantityInOneEmployee1 := 0
	//Для первого работника
	for _, search := range comparisonOfValue{
		revenue += search
		if maxBestDailyRevenue == search && quantityInOneEmployee1 == 0{
			indexTheBestDailyRevenue1 = increment
			quantityInOneEmployee1++
		}
		if minWorstDailyRevenue == search && quantityInOneEmployee == 0{
			indexTheWorstDailyRevenue1 =  increment
			quantityInOneEmployee++
		}
	}
	if revenue == maxRevenue{
		IndexOfMaxRevenue1 =increment
	}
	maxAverage := revenue / len(comparisonOfValue)
	if maxAverage == maxAverageDailyRevenue{
		IndexOfAverageDailyRevenue1 =  increment
	}
	return
}
func assignment( value int, values []int) []int{
	if value !=0 {
	values = append(values, value)
	return values
	}
	return values
}