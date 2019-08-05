package model

type ListMasterRead struct {
	MasterReadID int    `json:"meter_read_id"`
	Verb         string `json:"verb"`
	Noun         string `json:"noun"`
	Revision     int    `json:"revision"`
	Datetime     string `json:"date_time"`
	Source       string `json:"source"`
	MeterCode    string
	MeterData    []MeterData
}

type DetailMasterRead struct {
	MasterReadID int    `json:"meter_read_id"`
	Verb         string `json:"verb"`
	Noun         string `json:"noun"`
	Revision     int    `json:"revision"`
	Datetime     string `json:"date_time"`
	Source       string `json:"source"`
	MeterCode    string
	MeterData    []MeterData
}

type MeterData struct {
	MeterCode      string `json:"meter_code"`
	MeterType      string `json:"meter_type"`
	IntervalBlocks []Interval
}

type Interval struct {
	IntervalReadIn string `json:"interval_read_in"`
	EndTime        string `json:"end_time"`
	IntervalLength int    `json:"interval_length"`
	IntervalValue  int    `json:"interval_value"`
	IntervalFlag   int    `json:"interval_flag"`
	CollectionTime string `json:"collection_time"`
}
