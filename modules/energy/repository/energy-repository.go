package repository

func FindAllMeterRead() string {
	return "SELECT meter_information_id, " +
		"verb, " +
		"noun, " +
		"revision, " +
		"date_time, " +
		"source, " +
		"meter_code " +
		"FROM meter_information LIMIT ? OFFSET ?"
}

func FindMeterReadByMeterReadId() string {
	return "SELECT meter_information_id, " +
		"verb, " +
		"noun, " +
		"revision, " +
		"date_time, " +
		"source, " +
		"meter_code " +
		"FROM meter_information WHERE meter_information_id = ?"
}

func FindAllMeter() string {
	return "SELECT meter_code, meter_type FROM meter " +
		"WHERE meter.meter_code = ?"
}

func FindAllIntervalMeter() string {
	return "SELECT interval_readin_id, " +
		"end_time, " +
		"interval_length, " +
		"interval_value, " +
		"interval_flags, " +
		"collection_time FROM interval_block " +
		"WHERE interval_block.meter_code = ?"
}

func SaveMeterRead() string {
	return "INSERT INTO meter_information(verb, noun, " +
		"revision, date_time, " +
		"source, meter_code) VALUES(?, ?, ?, ?, ?, ?)"
}

func SaveMeterData() string {
	return "INSERT INTO meter(meter_code, meter_type) " +
		"VALUES(?, ?)"
}

func SaveIntervalMeter() string {
	return "INSERT INTO interval_block(interval_readin_id, end_time, " +
		"interval_length, interval_value, " +
		"interval_flags, collection_time, meter_code) " +
		"VALUES(?, ?, ?, ?, ?, ?, ?)"
}

func UpdateMeterRead() string {
	return "UPDATE meter_information SET verb = ?, " +
		"noun = ?, " +
		"revision = ?, " +
		"date_time = ?, " +
		"source = ?, " +
		"meter_code = ? WHERE meter_information_id = ?"
}

func DeleteMeterRead() string {
	return "DELETE FROM meter_information WHERE meter_information_id = ?"
}
