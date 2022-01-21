package plugin

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

//ExampleTest: https://github.com/grafana/grafana-plugin-sdk-go/blob/main/data/frame_test.go

func TestNoRows(t *testing.T) {
	skipIfIsShort(t)
	expectedFrame := data.NewFrame("response",
		data.NewField("m", nil, []string{}),
	)

	cypher := "Match(m) return m limit 0"

	runNeo4JIntegrationTest(t, cypher, expectedFrame)
}

func TestReturnStringColumn(t *testing.T) {
	skipIfIsShort(t)
	expectedFrame := data.NewFrame("response",
		data.NewField("m.title", nil, []string{
			"The Matrix",
		}),
	)

	cypher := "Match(m:Movie) where m.title = 'The Matrix' return m.title limit 1"

	runNeo4JIntegrationTest(t, cypher, expectedFrame)
}

func TestReturnNodeColumn(t *testing.T) {
	skipIfIsShort(t)
	expectedFrame := data.NewFrame("response",
		data.NewField("m", nil, []string{
			"{\"Id\":0,\"Labels\":[\"Movie\"],\"Props\":{\"released\":1999,\"tagline\":\"Welcome to the Real World\",\"title\":\"The Matrix\"}}",
		}),
	)

	cypher := "Match(m:Movie) where m.title = 'The Matrix' return m limit 1"

	runNeo4JIntegrationTest(t, cypher, expectedFrame)
}

func runNeo4JIntegrationTest(t *testing.T, cypher string, expected *data.Frame) {
	neo4JSettings := neo4JSettings{
		Url:      "neo4j://localhost:7687",
		Database: "",
		Username: "neo4j",
		Password: "Password123",
	}

	neo4JQuery := neo4JQuery{
		CypherQuery: cypher,
	}

	res, err := query(neo4JSettings, neo4JQuery)
	if err != nil {
		t.Fatal(err)
	}

	if res.Error != nil {
		t.Error(res.Error)
	}

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

func skipIfIsShort(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping test in short mode.")
    }
}
