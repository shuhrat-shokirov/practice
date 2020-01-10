package main

import (
	"reflect"
	"testing"
)

func TestStoreRevenue(t *testing.T) {
	tests := []struct {
		name                           string
		firstRevenue                   []int
		secondRevenue				   []int
		thirdRevenue				   []int
		wantIndexOfMaxRevenue          []int
		wantIndexOfAverageDailyRevenue []int
		wantIndexTheBestDailyRevenue   []int
		wantIndexTheWorstDailyRevenue  []int
	}{
		{"If all employees work equally",[]int{10, 10, 10, 15},[]int{10, 10, 10, 15}, []int{10, 10, 10, 15},[]int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3},[]int{1, 2, 3}},
		{"Workers worked differently",[]int{7, 1, 2, 3},[]int{10, 10, 10, 15}, []int{7, 20, 2},[]int{2}, []int{2}, []int{3},[]int{1}},
		{"Some days workers converge",[]int{7, 1, 2, 3},[]int{10, 10, 10, 15}, []int{7, 1, 2},[]int{2}, []int{2}, []int{2},[]int{1, 3}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
			gotIndexOfMaxRevenue, gotIndexOfAverageDailyRevenue, gotIndexTheBestDailyRevenue, gotIndexTheWorstDailyRevenue := StoreRevenue(tt.firstRevenue, tt.secondRevenue, tt.thirdRevenue)
			if !reflect.DeepEqual(gotIndexOfMaxRevenue, tt.wantIndexOfMaxRevenue) {
				t.Error("convert:", tt.name, "want:", tt.wantIndexOfMaxRevenue, "got:", gotIndexOfMaxRevenue)
			}
			if !reflect.DeepEqual(gotIndexOfAverageDailyRevenue, tt.wantIndexOfAverageDailyRevenue) {
				t.Error("convert:", tt.name, "want:", tt.wantIndexOfAverageDailyRevenue, "got:", gotIndexOfAverageDailyRevenue)
			}
			if !reflect.DeepEqual(gotIndexTheBestDailyRevenue, tt.wantIndexTheBestDailyRevenue) {
				t.Error("convert:", tt.name, "want:", tt.wantIndexTheBestDailyRevenue, "got:", gotIndexTheBestDailyRevenue)
			}
			if !reflect.DeepEqual(gotIndexTheWorstDailyRevenue, tt.wantIndexTheWorstDailyRevenue) {
				t.Error("convert:", tt.name, "want:", tt.wantIndexTheWorstDailyRevenue, "got:", gotIndexTheWorstDailyRevenue)
			}
	}
}