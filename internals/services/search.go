package search

// type result struct {
// 	success bool
// 	data    string
// }

type SearchService interface {
	SearchForSomething() int
}

func SearchForSomething() int {
	// return &result{
	// 	success: true,
	// 	data:    "Somedata",
	// }
	return 1
}
