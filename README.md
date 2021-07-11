# Neo4J DataSource for Grafana
Allows Neo4J to be used as a DataSource for Grafana

## Installation


## Development
### Overview
Introduction: https://grafana.com/tutorials/build-a-data-source-plugin/

Using Grafana Toolkit: https://github.com/grafana/grafana/tree/main/packages/grafana-toolkit#usage


### Getting started
Install dependencies
```
cd neo4j-datasource-plugin
yarn install
```

Build Plugin in watch mode
```
cd neo4j-datasource-plugin
yarn watch
```

Starts preprovisioned Grafana and Neo4J 
```
sudo docker-compose up
```
Grafana: http://localhost:3000

Neo4J: http://localhost:7474

Grafana is started by docker-compose in development mode therefore no restart of grafana is required when source code changed.

