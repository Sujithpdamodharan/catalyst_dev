package viewmodels

type DetailedDashboard struct {
	UserName               string
	LoginType              string
	SessionProject         int
	SessionProjectName     string
	SerialNo               string
	BusCountArry           []int
	AssociationArry        []string
	TotalTicketCount       []int
	TotalTransactionAmount []float64
	TransactionDate        []string
	TotalCardCount         []int
	TotalCardAmount        []float64
	CardDate               []string
	LiveBusArray           []int
	LiveBUsDate            []string
	TileValues             []string
	TotalBusArry           []int
	TotalBusName           []string
	CountBus               int
	AgentData              [][]string
	Count                  string
	Project                []string
	IDArray                []string
}
