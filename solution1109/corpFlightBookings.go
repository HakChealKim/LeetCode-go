package main

func corpFlightBookings(bookings [][]int, n int) []int {
	presum := make([]int, n)

	for _, book := range bookings {
		presum[book[0]-1] += book[2]
		if book[1] < n {
			presum[book[1]] -= book[2]
		}
	}
	for i := 1; i < len(presum); i++ {
		presum[i] += presum[i-1]
	}

	return presum

}
