package plugin

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

//ExampleTest: https://github.com/grafana/grafana-plugin-sdk-go/blob/main/data/frame_test.go

func TestHealthcheckIsOk(t *testing.T) {
	skipIfIsShort(t)

	settings := neo4JSettings{
		Url:      "neo4j://localhost:7687",
		Database: "",
		Username: "neo4j",
		Password: "Password123",
	}

	const OK_STATUS backend.HealthStatus = 1
	testCheckHealth(t, settings, OK_STATUS)
}

func TestHealthcheckIsErrorDueToInvalidHost(t *testing.T) {
	skipIfIsShort(t)

	settings := neo4JSettings{
		Url:      "neo4j://invalid:7687",
		Database: "",
		Username: "neo4j",
		Password: "Password123",
	}

	const ERROR_STATUS backend.HealthStatus = 2
	testCheckHealthAndMessage(t, settings, ERROR_STATUS, "ConnectivityError")
}

func TestHealthcheckIsErrorDueToDeserialize(t *testing.T) {
	skipIfIsShort(t)

	settings := backend.DataSourceInstanceSettings{}
	settings.JSONData = []byte{1}

	_, err := NewNeo4JDatasource(settings)

	expectedMsg := "can not deserialize DataSource settings"
	if !strings.Contains(err.Error(), expectedMsg) {
		t.Error("Expected Message " + expectedMsg + ", but was " + err.Error())
	}
}

func TestHealthcheckIsErrorDueToInvalidPort(t *testing.T) {
	skipIfIsShort(t)
	settings := neo4JSettings{
		Url:      "neo4j://localhost:1234",
		Database: "",
		Username: "neo4j",
		Password: "Password123",
	}

	const ERROR_STATUS backend.HealthStatus = 2
	testCheckHealthAndMessage(t, settings, ERROR_STATUS, "ConnectivityError")
}

func TestHealthcheckIsErrorDueToInvalidUsername(t *testing.T) {
	skipIfIsShort(t)
	settings := neo4JSettings{
		Url:      "neo4j://localhost:7687",
		Database: "",
		Username: "doesNotExist",
		Password: "Password123",
	}

	const ERROR_STATUS backend.HealthStatus = 2
	testCheckHealthAndMessage(t, settings, ERROR_STATUS, "unauthorized due to authentication failure")
}

func TestHealthcheckIsErrorDueToInvalidPassword(t *testing.T) {
	skipIfIsShort(t)
	settings := neo4JSettings{
		Url:      "neo4j://localhost:7687",
		Database: "",
		Username: "neo4j",
		Password: "NotValid",
	}

	const ERROR_STATUS backend.HealthStatus = 2
	testCheckHealthAndMessage(t, settings, ERROR_STATUS, "unauthorized due to authentication failure")
}

func TestHealthcheckIsErrorDueToNoAuth(t *testing.T) {
	skipIfIsShort(t)
	settings := neo4JSettings{
		Url:      "neo4j://localhost:7687",
		Database: "",
		Username: "",
		Password: "",
	}

	const ERROR_STATUS backend.HealthStatus = 2
	testCheckHealthAndMessage(t, settings, ERROR_STATUS, "Unsupported authentication token")
}

func TestHealthcheckIsErrorDueToInvalidDatabase(t *testing.T) {
	skipIfIsShort(t)
	settings := neo4JSettings{
		Url:      "neo4j://localhost:7687",
		Database: "NotValid",
		Username: "neo4j",
		Password: "Password123",
	}

	const ERROR_STATUS backend.HealthStatus = 2
	testCheckHealthAndMessage(t, settings, ERROR_STATUS, "this database does not exist")
}

func testCheckHealth(t *testing.T, neo4JSettings neo4JSettings, expectedStatus backend.HealthStatus) {
	testCheckHealthAndMessage(t, neo4JSettings, expectedStatus, "")
}

func testCheckHealthAndMessage(t *testing.T, neo4JSettings neo4JSettings, expectedStatus backend.HealthStatus, expectedMessagePart string) {
	settings := backend.DataSourceInstanceSettings{}
	settings.JSONData = asJsonBytes(t, neo4JSettings)

	testCheckHealthAndMessageWithSettings(t, settings, expectedStatus, expectedMessagePart)
}

func testCheckHealthAndMessageWithSettings(t *testing.T, settings backend.DataSourceInstanceSettings, expectedStatus backend.HealthStatus, expectedMessagePart string) {

	instance, err := NewNeo4JDatasource(settings)
	if err != nil {
		t.Fatal(err)
	}

	neo4JDatasource := instance.(*Neo4JDatasource)
	res, err := neo4JDatasource.checkHealth()

	if err != nil {
		t.Fatal(err)
	}

	if res.Status != expectedStatus {
		t.Error("Expected Status " + expectedStatus.String() + ", but was " + res.Status.String())
	}

	if expectedMessagePart != "" && !strings.Contains(res.Message, expectedMessagePart) {
		t.Error("Expected Message contains " + expectedMessagePart + ", but was " + res.Message)
	}
	fmt.Println("Status:" + res.Status.String())
	fmt.Println("Message:" + res.Message)
}

func TestNoRows(t *testing.T) {
	skipIfIsShort(t)
	expectedFrame := data.NewFrame("response",
		data.NewField("m", nil, []*string{}),
	)

	cypher := "Match(m) return m limit 0"

	runNeo4JIntegrationTableTest(t, cypher, expectedFrame)
}

func TestStringColumn(t *testing.T) {
	skipIfIsShort(t)
	expectedFrame := data.NewFrame("response",
		data.NewField("A", nil, []*string{
			ptrS("One"),
		}),
	)

	cypher := "With 'One' as A return A"

	runNeo4JIntegrationTableTest(t, cypher, expectedFrame)
}

func TestIntColumn(t *testing.T) {
	skipIfIsShort(t)
	expectedFrame := data.NewFrame("response",
		data.NewField("A", nil, []*int64{
			ptrI(1),
		}),
	)

	cypher := "With 1 as A return A"

	runNeo4JIntegrationTableTest(t, cypher, expectedFrame)
}

func TestBooleanColumn(t *testing.T) {
	skipIfIsShort(t)
	expectedFrame := data.NewFrame("response",
		data.NewField("A", nil, []*bool{
			ptrB(true),
			ptrB(false),
		}),
	)

	cypher := "With TRUE as A return A UNION ALL With FALSE as A return A"

	runNeo4JIntegrationTableTest(t, cypher, expectedFrame)
}

func TestFloatColumn(t *testing.T) {
	skipIfIsShort(t)
	expectedFrame := data.NewFrame("response",
		data.NewField("A", nil, []*float64{
			ptrF(0.81234),
		}),
	)

	cypher := "With 0.81234 as A return A"

	runNeo4JIntegrationTableTest(t, cypher, expectedFrame)
}

func TestByteArrayColumn(t *testing.T) {
	skipIfIsShort(t)
	expectedFrame := data.NewFrame("response",
		data.NewField("output", nil, []*string{
			ptrS("[78,101,111,52,106]"),
		}),
	)

	cypher := "RETURN apoc.text.bytes('Neo4j') AS output"

	runNeo4JIntegrationTableTest(t, cypher, expectedFrame)
}

func TestStringListColumn(t *testing.T) {
	skipIfIsShort(t)
	expectedFrame := data.NewFrame("response",
		data.NewField("A", nil, []*string{
			ptrS("[\"a\",\"b\",\"c\"]"),
		}),
	)

	cypher := "Return ['a', 'b',  'c'] as A"

	runNeo4JIntegrationTableTest(t, cypher, expectedFrame)
}

func TestIntListColumn(t *testing.T) {
	skipIfIsShort(t)
	expectedFrame := data.NewFrame("response",
		data.NewField("A", nil, []*string{
			ptrS("[1,2,3]"),
		}),
	)

	cypher := "Return [1,2,3] as A"

	runNeo4JIntegrationTableTest(t, cypher, expectedFrame)
}

func TestMapColumn(t *testing.T) {
	skipIfIsShort(t)
	expectedFrame := data.NewFrame("response",
		data.NewField("A", nil, []*string{
			ptrS("{\"key\":\"Value\",\"listKey\":[{\"inner\":\"Map1\"},{\"inner\":\"Map2\"}]}"),
		}),
	)

	cypher := "RETURN {key: 'Value', listKey: [{inner: 'Map1'}, {inner: 'Map2'}]} as A"

	runNeo4JIntegrationTableTest(t, cypher, expectedFrame)
}

func TestUTCDateTimeColumn(t *testing.T) {
	skipIfIsShort(t)
	expectedFrame := data.NewFrame("response",
		data.NewField("A", nil, []*time.Time{
			ptrT(time.Date(2022, time.Month(3), 2, 13, 14, 15, 144000000, time.UTC)),
		}),
	)

	cypher := "return datetime(\"2022-03-02T13:14:15.144Z\") as A"

	runNeo4JIntegrationTableTest(t, cypher, expectedFrame)
}

func TestDateTimeColumn(t *testing.T) {
	skipIfIsShort(t)
	expectedFrame := data.NewFrame("response",
		data.NewField("A", nil, []*time.Time{
			ptrT(time.Date(2022, time.Month(3), 2, 13, 14, 15, 144000000, time.FixedZone("TEST", 3600))),
		}),
	)

	cypher := "return datetime(\"2022-03-02T13:14:15.144+0100\") as A"

	runNeo4JIntegrationTableTest(t, cypher, expectedFrame)
}

func TestDateColumn(t *testing.T) {
	skipIfIsShort(t)
	expectedFrame := data.NewFrame("response",
		data.NewField("A", nil, []*time.Time{
			ptrT(time.Date(2019, time.Month(6), 1, 0, 0, 0, 0, time.UTC)),
		}),
	)

	cypher := "return date(\"2019-06-01\") as A"

	runNeo4JIntegrationTableTest(t, cypher, expectedFrame)
}

func TestTimeColumn(t *testing.T) {
	skipIfIsShort(t)
	expectedFrame := data.NewFrame("response",
		data.NewField("A", nil, []*time.Time{
			ptrT(time.Date(-1, time.Month(11), 30, 19, 15, 30, 0, time.UTC)),
		}),
	)

	cypher := "return time(\"19:15:30\") as A"

	runNeo4JIntegrationTableTest(t, cypher, expectedFrame)
}

func TestDurationColumn(t *testing.T) {
	skipIfIsShort(t)
	expectedFrame := data.NewFrame("response",
		data.NewField("A", nil, []*string{
			ptrS("P0M0DT180S"),
		}),
	)

	cypher := "return duration(\"PT3M\") as A"

	runNeo4JIntegrationTableTest(t, cypher, expectedFrame)
}

func TestNodeColumn(t *testing.T) {
	skipIfIsShort(t)
	expectedFrame := data.NewFrame("response",
		data.NewField("m", nil, []*string{
			ptrS("{\"Id\":0,\"ElementId\":\"0\",\"Labels\":[\"Movie\"],\"Props\":{\"released\":1999,\"tagline\":\"Welcome to the Real World\",\"title\":\"The Matrix\"}}"),
		}),
	)

	cypher := "Match(m:Movie) where m.title = 'The Matrix' return m limit 1"

	runNeo4JIntegrationTableTest(t, cypher, expectedFrame)
}

func TestRelationshipColumn(t *testing.T) {
	skipIfIsShort(t)
	expectedFrame := data.NewFrame("response",
		data.NewField("r", nil, []*string{
			ptrS("{\"Id\":0,\"ElementId\":\"0\",\"StartId\":1,\"StartElementId\":\"1\",\"EndId\":0,\"EndElementId\":\"0\",\"Type\":\"ACTED_IN\",\"Props\":{\"roles\":[\"Neo\"]}}"),
		}),
	)

	cypher := "MATCH (p:Person)-[r:ACTED_IN]->(m:Movie) where m.title = 'The Matrix' AND p.name = 'Keanu Reeves' RETURN r LIMIT 1"

	runNeo4JIntegrationTableTest(t, cypher, expectedFrame)
}

func TestMultipleRows(t *testing.T) {
	skipIfIsShort(t)
	expectedFrame := data.NewFrame("response",
		data.NewField("A", nil, []*string{
			ptrS("One"),
			ptrS("Two"),
		}),
	)

	cypher := "With 'One' as A return A UNION ALL With 'Two' as A return A"

	runNeo4JIntegrationTableTest(t, cypher, expectedFrame)
}

func TestMultipleRowsAndColumns(t *testing.T) {
	skipIfIsShort(t)
	expectedFrame := data.NewFrame("response",
		data.NewField("A", nil, []*string{
			ptrS("One"),
			ptrS("Three"),
		}),
		data.NewField("B", nil, []*string{
			ptrS("Two"),
			ptrS("Four"),
		}),
		data.NewField("C", nil, []*int64{
			ptrI(1),
			ptrI(2),
		}),
	)

	cypher := "With 'One' as A, 'Two' as B, 1 as C return A, B, C UNION ALL With 'Three' as A, 'Four' as B, 2 as C return A, B, C"

	runNeo4JIntegrationTableTest(t, cypher, expectedFrame)
}

func TestNullValue(t *testing.T) {
	skipIfIsShort(t)
	expectedFrame := data.NewFrame("response",
		data.NewField("A", nil, []*string{
			ptrS("abc"),
			nil,
		}),
	)

	cypher := "Return 'abc' as A UNION ALL Return null as A"

	runNeo4JIntegrationTableTest(t, cypher, expectedFrame)
}

func TestColumnNameWithDot(t *testing.T) {
	skipIfIsShort(t)
	expectedFrame := data.NewFrame("response",
		data.NewField("m.title", nil, []*string{
			ptrS("The Matrix"),
		}),
	)

	cypher := "Match(m:Movie) where m.title = 'The Matrix' return m.title limit 1"

	runNeo4JIntegrationTableTest(t, cypher, expectedFrame)
}

func TestNullInIntColumn(t *testing.T) {
	skipIfIsShort(t)
	expectedFrame := data.NewFrame("response",
		data.NewField("A", nil, []*int64{
			nil,
			ptrI(1),
		}))

	cypher := "With null as A return A UNION ALL With 1 as A return A"

	runNeo4JIntegrationTableTest(t, cypher, expectedFrame)
}

func TestAllNullInColumn(t *testing.T) {
	skipIfIsShort(t)
	expectedFrame := data.NewFrame("response",
		data.NewField("A", nil, []*string{
			nil,
			nil,
		}))

	cypher := "With null as A return A UNION ALL With null as A return A"

	runNeo4JIntegrationTableTest(t, cypher, expectedFrame)
}

func TestGraphFormat(t *testing.T) {
	skipIfIsShort(t)
	expectedNodesFrame := data.NewFrame("nodes",
		data.NewField("id", nil, []*string{
			ptrS("1"),
			ptrS("0"),
		}),
		data.NewField("title", nil, []*string{
			ptrS("Person"),
			ptrS("Movie"),
		}),
		data.NewField("detail__labels", nil, []*string{
			nil,
			nil,
		}),
		data.NewField("detail__born", nil, []*string{
			ptrS("1964"),
			nil,
		}),
		data.NewField("detail__name", nil, []*string{
			ptrS("\"Keanu Reeves\""),
			nil,
		}),
		data.NewField("detail__released", nil, []*string{
			nil,
			ptrS("1999"),
		}),
		data.NewField("detail__tagline", nil, []*string{
			nil,
			ptrS("\"Welcome to the Real World\""),
		}),
		data.NewField("detail__title", nil, []*string{
			nil,
			ptrS("\"The Matrix\""),
		}),
	)

	expectedEdgesFrame := data.NewFrame("edges",
		data.NewField("id", nil, []*string{
			ptrS("0"),
		}),
		data.NewField("source", nil, []*string{
			ptrS("1"),
		}),
		data.NewField("target", nil, []*string{
			ptrS("0"),
		}),
		data.NewField("mainStat", nil, []*string{
			ptrS("ACTED_IN"),
		}),
		data.NewField("detail__roles", nil, []*string{
			ptrS("[\"Neo\"]"),
		}),
	)

	m := data.FrameMeta{PreferredVisualization: "nodeGraph"}
	expectedNodesFrame = expectedNodesFrame.SetMeta(&m)
	expectedEdgesFrame = expectedEdgesFrame.SetMeta(&m)

	cypher := "Match(m:Movie)<-[a:ACTED_IN]-(f:Person) where m.title = 'The Matrix' and f.name = 'Keanu Reeves' return  m.title, f, a, m order by m.title"

	runNeo4JIntegrationGraphTest(t, cypher, expectedNodesFrame, expectedEdgesFrame)
}

func runNeo4JIntegrationGraphTest(t *testing.T, cypher string, expectedNodes *data.Frame, expectedEdges *data.Frame) {
	res := runNeo4JIntegrationTest(t, cypher, "nodegraph")
	if len(res.Frames) != 2 {
		t.Fatal("Frames len is not 2")
	}

	nodeFrame := res.Frames[0]
	edgeFrame := res.Frames[1]

	expectedNodesAsTable, _ := expectedNodes.StringTable(-1, -1)
	fmt.Println("Expected Nodes:\n", expectedNodesAsTable)

	nodesFrameAsTable, _ := nodeFrame.StringTable(-1, -1)
	fmt.Println("Actual Nodes:\n", nodesFrameAsTable)

	expectedEdgesAsTable, _ := expectedEdges.StringTable(-1, -1)
	fmt.Println("Expected Edges:\n", expectedEdgesAsTable)

	edgesFrameAsTable, _ := edgeFrame.StringTable(-1, -1)
	fmt.Println("Actual Edges:\n", edgesFrameAsTable)

	diffNodes := cmp.Diff(nodeFrame, expectedNodes, data.FrameTestCompareOptions()...)
	if diffNodes != "" {
		t.Fatal(diffNodes)
	}

	diffEdges := cmp.Diff(edgeFrame, expectedEdges, data.FrameTestCompareOptions()...)
	if diffEdges != "" {
		t.Fatal(diffEdges)
	}

}

func runNeo4JIntegrationTableTest(t *testing.T, cypher string, expected *data.Frame) {
	res := runNeo4JIntegrationTest(t, cypher, "table")

	if len(res.Frames) != 1 {
		t.Fatal("Frames len is not 1")
	}

	frame := res.Frames[0]

	expectedAsTable, _ := expected.StringTable(-1, -1)
	fmt.Println("Expected:\n", expectedAsTable)

	frameAsTable, _ := frame.StringTable(-1, -1)
	fmt.Println("Actual:\n", frameAsTable)

	diff := cmp.Diff(frame, expected, data.FrameTestCompareOptions()...)
	if diff != "" {
		t.Fatal(diff)
	}
}

func runNeo4JIntegrationTest(t *testing.T, cypher string, format string) backend.DataResponse {
	neo4JSettings := neo4JSettings{
		Url:      "neo4j://localhost:7687",
		Database: "",
		Username: "neo4j",
		Password: "Password123",
	}

	neo4JQuery := neo4JQuery{
		CypherQuery: cypher,
		Format:      format,
	}

	settings := backend.DataSourceInstanceSettings{}
	settings.JSONData = asJsonBytes(t, neo4JSettings)

	instance, _ := NewNeo4JDatasource(settings)
	neo4JDatasource := instance.(*Neo4JDatasource)

	res, err := neo4JDatasource.query(neo4JQuery)
	if err != nil {
		t.Fatal(err)
	}

	if res.Error != nil {
		t.Error(res.Error)
	}
	return res
}

func asJsonBytes(t *testing.T, obj interface{}) []byte {
	objAsBytes, err := json.Marshal(obj)
	if err != nil {
		t.Fatal(err)
	}
	return objAsBytes
}

func skipIfIsShort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
}

func ptrS(value string) *string {
	return &value
}

func ptrI(value int64) *int64 {
	return &value
}

func ptrT(value time.Time) *time.Time {
	return &value
}

func ptrF(value float64) *float64 {
	return &value
}

func ptrB(value bool) *bool {
	return &value
}
