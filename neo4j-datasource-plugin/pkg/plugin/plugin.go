package plugin

import (
	"context"
	"encoding/json"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// Make sure SampleDatasource implements required interfaces. This is important to do
// since otherwise we will only get a not implemented error response from plugin in
// runtime. In this example datasource instance implements backend.QueryDataHandler,
// backend.CheckHealthHandler. Plugin should not
// implement all these interfaces - only those which are required for a particular task.
// For example if plugin does not need streaming functionality then you are free to remove
// methods that implement backend.StreamHandler. Implementing instancemgmt.InstanceDisposer
// is useful to clean up resources used by previous datasource instance when a new datasource
// instance created upon datasource settings changed.
var (
	_ backend.QueryDataHandler   = (*SampleDatasource)(nil)
	_ backend.CheckHealthHandler = (*SampleDatasource)(nil)
	_ backend.DataSourceInstanceSettings
	_ instancemgmt.InstanceDisposer = (*SampleDatasource)(nil)
)

// NewSampleDatasource creates a new datasource instance.
func NewSampleDatasource(_ backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {
	return &SampleDatasource{}, nil
}

// SampleDatasource is an example datasource which can respond to data queries, reports
// its health and has streaming skills.
type SampleDatasource struct{}

// Dispose here tells plugin SDK that plugin wants to clean up resources when a new instance
// created. As soon as datasource settings change detected by SDK old datasource instance will
// be disposed and a new one will be created using NewSampleDatasource factory function.
func (d *SampleDatasource) Dispose() {
	// Clean up datasource instance resources.
}

// QueryData handles multiple queries and returns multiple responses.
// req contains the queries []DataQuery (where each query contains RefID as a unique identifier).
// The QueryDataResponse contains a map of RefID to the response for each query, and each response
// contains Frames ([]*Frame).
func (d *SampleDatasource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	log.DefaultLogger.Info("QueryData called")

	// create response struct
	response := backend.NewQueryDataResponse()

	neo4JSettings, err := unmarshalDataSourceSettings(req.PluginContext.DataSourceInstanceSettings)
	if err != nil {
		return response, err
	}

	// loop over queries and execute them individually.
	for _, q := range req.Queries {

		var res backend.DataResponse

		// Unmarshal the JSON into our queryModel.
		var neo4JQuery neo4JQuery
		err := json.Unmarshal(q.JSON, &neo4JQuery)
		if err != nil {
			res.Error = err
			response.Responses[q.RefID] = res
			continue
		}

		neo4JQuery.RefID = q.RefID
		neo4JQuery.QueryType = q.QueryType
		neo4JQuery.Interval = q.Interval
		neo4JQuery.MaxDataPoints = q.MaxDataPoints
		neo4JQuery.TimeRange = q.TimeRange

		res, err = query(neo4JSettings, neo4JQuery)
		if err != nil {
			res.Error = err
		}

		if res.Error != nil {
			log.DefaultLogger.Error("Error in query", res.Error)
		}

		response.Responses[q.RefID] = res
	}

	return response, nil
}

func query(settings neo4JSettings, query neo4JQuery) (backend.DataResponse, error) {
	log.DefaultLogger.Info("Execute Cypher Query: '" + query.CypherQuery + "'")

	response := backend.DataResponse{}

	// ToDo: Support other AuthTypes
	driver, err := neo4j.NewDriver(settings.Url, neo4j.BasicAuth(settings.Username, settings.Password, ""))
	if err != nil {
		return response, err
	}
	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	result, err := session.Run(query.CypherQuery, map[string]interface{}{})
	if err != nil {
		return response, err
	}

	return toDataResponse(result)
}

func toDataResponse(result neo4j.Result) (backend.DataResponse, error) {
	response := backend.DataResponse{}

	keys, err := result.Keys()
	if err != nil {
		return response, err
	}

	// create data frame response.
	frame := data.NewFrame("response")

	var currentRecord *neo4j.Record
	if result.Next() {
		currentRecord = result.Record()
	}

	for i, k := range keys {
		// infer datatypes of columns from first Row
		typ := getTypeArray(currentRecord, i)

		frame.Fields = append(frame.Fields,
			data.NewField(k, nil, typ),
		)
	}

	row := 0
	for currentRecord != nil {
		values := result.Record().Values
		for col, v := range values {
			f := frame.Fields[row]
			f.Extend(1)
			conValue := toValue(v)
			frame.Set(col, row, conValue)
		}
		if result.Next() {
			currentRecord = result.Record()
			row++
		} else {
			currentRecord = nil
		}
	}

	// add the frames to the response.
	response.Frames = append(response.Frames, frame)
	return response, nil
}

// CheckHealth handles health checks sent from Grafana to the plugin.
// The main use case for these health checks is the test button on the
// datasource configuration page which allows users to verify that
// a datasource is working as expected.
func (d *SampleDatasource) CheckHealth(ctx context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	log.DefaultLogger.Info("CheckHealth called")

	settings, err := unmarshalDataSourceSettings(req.PluginContext.DataSourceInstanceSettings)

	log.DefaultLogger.Info("Settings:", settings)

	if err != nil {
		return &backend.CheckHealthResult{
			Status:  backend.HealthStatusError,
			Message: err.Error(),
		}, nil
	}

	neo4JQuery := neo4JQuery{
		CypherQuery: "Match(a) return a limit 1",
	}

	_, err = query(settings, neo4JQuery)

	if err != nil {
		return &backend.CheckHealthResult{
			Status:  backend.HealthStatusError,
			Message: err.Error(),
		}, nil
	}

	return &backend.CheckHealthResult{
		Status:  backend.HealthStatusOk,
		Message: "Data source is working",
	}, nil
}

func unmarshalDataSourceSettings(dSIset *backend.DataSourceInstanceSettings) (neo4JSettings, error) {
	// Unmarshal the JSON into our settings Model.
	var neo4JSettings neo4JSettings
	err := json.Unmarshal(dSIset.JSONData, &neo4JSettings)
	if err != nil {
		return neo4JSettings, err

	}
	return neo4JSettings, nil
}

//https://github.com/neo4j/neo4j-go-driver#value-types
func getTypeArray(record *neo4j.Record, idx int) interface{} {
	if record == nil {
		return []string{}
	}

	typ := record.Values[idx]

	switch typ.(type) {
	case int64:
		return []int64{}
	default:
		return []string{}
	}
}

//https://github.com/neo4j/neo4j-go-driver#value-types
func toValue(val interface{}) interface{} {
	if val == nil {
		return nil
	}
	switch t := val.(type) {
	case string, int64:
		return t
	default:
		r, err := json.Marshal(val)
		if err != nil {
			log.DefaultLogger.Info("Marshalling failed ", "err", err)
		}
		return string(r)
	}
}

type neo4JQuery struct {
	RefID string

	// QueryType is an optional identifier for the type of query.
	// It can be used to distinguish different types of queries.
	QueryType string

	// MaxDataPoints is the maximum number of datapoints that should be returned from a time series query.
	MaxDataPoints int64

	// Interval is the suggested duration between time points in a time series query.
	Interval time.Duration

	// TimeRange is the Start and End of the query as sent by the frontend.
	TimeRange backend.TimeRange

	CypherQuery string `json:"cypherQuery"`
}

type neo4JSettings struct {
	Url      string `json:"url"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}
