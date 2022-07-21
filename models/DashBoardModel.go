package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type StockData struct {
	Type           string
	TotalCount     int
	ActiveCount    int
	AssignedCount  int
	Decommissioned int
	Maintenance    int
}

func GetDailyBus(date string) int {

	db := dbConn()
	var count int
	busData := db.QueryRow("select count(DISTINCT sync_tkt_vehicle_id) from intermediate_sync_ticket where date(sync_tkt_date) =?", date)
	busData.Scan(&count)
	db.Close()
	return count
}

func GetStockDashboardData(loginType string) (string, []StockData) {

	db := dbConn()
	stockDataMap := []StockData{}

	var rollData StockData
	rollData.Type = "Roll"
	assignRollData := db.QueryRow("select sum(no_of_rolls) from Mac_Service_transaction where type='roll'")
	_ = assignRollData.Scan(&rollData.AssignedCount)

	totalrollsData := db.QueryRow("select sum(stock_no_of_rolls) from rollstock")
	_ = totalrollsData.Scan(&rollData.TotalCount)

	rollData.ActiveCount = rollData.TotalCount - rollData.AssignedCount
	stockDataMap = append(stockDataMap, rollData)

	dataById, _ := db.Query("SELECT fitment_type, COUNT(screen_id), SUM(status = 1) active, SUM(status = 2) assigned,SUM(status = 3) Maintenance ,SUM(status = 4) Decomissioned FROM screen_fitment group by fitment_type;")
	for dataById.Next() {
		var stockData StockData

		_ = dataById.Scan(&stockData.Type, &stockData.TotalCount, &stockData.ActiveCount, &stockData.AssignedCount, &stockData.Maintenance, &stockData.Decommissioned)
		stockDataMap = append(stockDataMap, stockData)
	}

	var groupData StockData
	groupData.Type = "Fitment Group"
	GroupFitmentData := db.QueryRow("select COUNT(id_fitment_group),SUM(status = 1) active, SUM(status = 2) assigned,SUM(status = 3) Decomissioned from fitment_group")
	_ = GroupFitmentData.Scan(&groupData.TotalCount, &groupData.ActiveCount, &groupData.AssignedCount, &groupData.Decommissioned)
	stockDataMap = append(stockDataMap, groupData)

	db.Close()
	return "true", stockDataMap
}

func GetStockDashboardData_20_20(loginType string) (string, []StockData) {

	db := dbConn()
	stockDataMap := []StockData{}

	var rollData StockData
	rollData.Type = "Roll"
	assignRollData := db.QueryRow("select sum(no_of_rolls) from Mac_Service_transaction where type='roll'")
	_ = assignRollData.Scan(&rollData.AssignedCount)

	totalrollsData := db.QueryRow("select sum(stock_no_of_rolls) from rollstock")
	_ = totalrollsData.Scan(&rollData.TotalCount)

	rollData.ActiveCount = rollData.TotalCount - rollData.AssignedCount
	stockDataMap = append(stockDataMap, rollData)

	dataByID, _ := db.Query("SELECT fitment_type, COUNT(screen_id), SUM(status = 1) active, SUM(status = 2) assigned, SUM(status = 3) Maintenance ,SUM(status = 4) Decomissioned FROM screen_fitment_2020 group by fitment_type")
	for dataByID.Next() {
		var stockData StockData
		_ = dataByID.Scan(&stockData.Type, &stockData.TotalCount, &stockData.ActiveCount, &stockData.AssignedCount, &stockData.Maintenance, &stockData.Decommissioned)
		stockDataMap = append(stockDataMap, stockData)
	}
	dataByID.Close()
	var groupData StockData
	groupData.Type = "Fitment Group"
	GroupFitmentData := db.QueryRow("select COUNT(id_fitment_group),SUM(status = 1) active, SUM(status = 2) assigned,SUM(status = 3) Decomissioned from fitment_group_2020")
	_ = GroupFitmentData.Scan(&groupData.TotalCount, &groupData.ActiveCount, &groupData.AssignedCount, &groupData.Decommissioned)
	stockDataMap = append(stockDataMap, groupData)

	db.Close()
	return "true", stockDataMap
}
