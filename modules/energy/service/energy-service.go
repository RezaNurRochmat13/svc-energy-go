package service

import (
	"database/sql"
	"log"
	"svc-energy-go/modules/energy/model"
	"svc-energy-go/modules/energy/repository"
	"svc-energy-go/resources"
)

func DatabaseConnection() *sql.DB {
	databaseConfiguration, errorHandlerDatabaseConnection := resources.SetupDatabaseConnection()

	if errorHandlerDatabaseConnection != nil {
		log.Printf("Error when connecting db : %s", errorHandlerDatabaseConnection)
	}

	return databaseConfiguration
}

func FetchAllMeterReads(limit string, offset string) ([]model.ListMasterRead, error) {
	databaseInterface := DatabaseConnection()

	var (
		modelMeterRead   model.ListMasterRead
		resultMeterReads []model.ListMasterRead
	)

	meterReadAllRepository := repository.FindAllMeterRead()

	queryFindAllMeterReads, errorHandlerQuery := databaseInterface.Query(meterReadAllRepository, limit, offset)

	if errorHandlerQuery != nil {
		log.Printf("Error when execute query : %s", errorHandlerQuery)
		return nil, errorHandlerQuery
	}

	for queryFindAllMeterReads.Next() {
		errorScanHandler := queryFindAllMeterReads.Scan(
			&modelMeterRead.MasterReadID,
			&modelMeterRead.Verb,
			&modelMeterRead.Noun,
			&modelMeterRead.Revision,
			&modelMeterRead.Datetime,
			&modelMeterRead.Source,
			&modelMeterRead.MeterCode,
		)

		if errorScanHandler != nil {
			log.Printf("Error when scan values : %s", errorScanHandler)
			return nil, errorScanHandler
		}

		modelMeterRead.MeterData = FetchMeterData(modelMeterRead.MeterCode)

		resultMeterReads = append(resultMeterReads, modelMeterRead)

	}
	return resultMeterReads, nil
}

func FetchDetailMeterRead(meterReadId string) ([]model.DetailMasterRead, error) {
	databaseInterface := DatabaseConnection()

	var (
		modelDetailMeterRead  model.DetailMasterRead
		resultDetailMeterRead []model.DetailMasterRead
	)

	meterReadDetailRepository := repository.FindMeterReadByMeterReadId()

	queryFetchMeterRead, errorHandlerQuery := databaseInterface.Query(meterReadDetailRepository, meterReadId)

	if errorHandlerQuery != nil {
		log.Printf("Error when execute query : %s", errorHandlerQuery)
		return nil, errorHandlerQuery
	}

	for queryFetchMeterRead.Next() {
		errorScanHandler := queryFetchMeterRead.Scan(
			&modelDetailMeterRead.MasterReadID,
			&modelDetailMeterRead.Verb,
			&modelDetailMeterRead.Noun,
			&modelDetailMeterRead.Revision,
			&modelDetailMeterRead.Datetime,
			&modelDetailMeterRead.Source,
			&modelDetailMeterRead.MeterCode,
		)

		if errorScanHandler != nil {
			log.Printf("Error when scan values : %s", errorScanHandler)
			return nil, errorScanHandler
		}

		modelDetailMeterRead.MeterData = FetchMeterData(modelDetailMeterRead.MeterCode)

		resultDetailMeterRead = append(resultDetailMeterRead, modelDetailMeterRead)
	}

	return resultDetailMeterRead, nil

}

func FetchMeterData(meterCode string) []model.MeterData {
	databaseInterface := DatabaseConnection()

	var (
		modelMeterData  model.MeterData
		resultMeterData []model.MeterData
	)

	meterDataRepository := repository.FindAllMeter()

	queryFindAllMeter, errorHandlerQuery := databaseInterface.Query(meterDataRepository, meterCode)

	if errorHandlerQuery != nil {
		log.Printf("Error when execute query : %s", errorHandlerQuery)
		return nil
	}

	for queryFindAllMeter.Next() {
		errorHandlerScan := queryFindAllMeter.Scan(
			&modelMeterData.MeterCode,
			&modelMeterData.MeterType,
		)

		if errorHandlerScan != nil {
			log.Printf("Error when scan values : %s", errorHandlerScan)
		}

		modelMeterData.IntervalBlocks = FetchIntervalBlock(modelMeterData.MeterCode)

		resultMeterData = append(resultMeterData, modelMeterData)
	}

	return resultMeterData

}

func FetchIntervalBlock(meterCode string) []model.Interval {
	databaseInterface := DatabaseConnection()

	var (
		modelIntervalBlockData  model.Interval
		resultIntervalBlockData []model.Interval
	)

	intervalBlockRepository := repository.FindAllIntervalMeter()

	queryFetchIntervalBlock, errorHandlerQuery := databaseInterface.Query(intervalBlockRepository, meterCode)

	if errorHandlerQuery != nil {
		log.Printf("Error when execute query : %s", errorHandlerQuery)
		return nil
	}

	for queryFetchIntervalBlock.Next() {
		errorScanHandler := queryFetchIntervalBlock.Scan(
			&modelIntervalBlockData.IntervalReadIn,
			&modelIntervalBlockData.EndTime,
			&modelIntervalBlockData.IntervalLength,
			&modelIntervalBlockData.IntervalValue,
			&modelIntervalBlockData.IntervalFlag,
			&modelIntervalBlockData.CollectionTime,
		)

		if errorScanHandler != nil {
			log.Printf("Error when scan values : %s", errorScanHandler)
			return nil
		}

		resultIntervalBlockData = append(resultIntervalBlockData, modelIntervalBlockData)
	}

	return resultIntervalBlockData
}
