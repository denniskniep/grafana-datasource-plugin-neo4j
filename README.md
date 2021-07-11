# Neo4J DataSource for Grafana
Allows Neo4J to be used as a DataSource for Grafana

## Installation


## Development
### Overview
Introduction: https://grafana.com/tutorials/build-a-data-source-plugin/

Using Grafana Toolkit: https://github.com/grafana/grafana/tree/main/packages/grafana-toolkit#usage


### Build and run

Build (-w = watch for changes)
```
cd neo4j-datasource-plugin
npx @grafana/toolkit plugin:dev -w
```

Execute from repository root
```
sudo docker-compose up
```

Grafana is started by docker-compose in development mode therefore no restart of grafana is required when source code changed. Furthermore neo4j is also started by docker-compose.

