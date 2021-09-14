[![CI](https://github.com/denniskniep/grafana-datasource-plugin-neo4j/actions/workflows/ci.yml/badge.svg)](https://github.com/denniskniep/grafana-datasource-plugin-neo4j/actions/workflows/ci.yml) [![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/denniskniep/grafana-datasource-plugin-neo4j?sort=semver)](https://github.com/denniskniep/grafana-datasource-plugin-neo4j/releases/latest)
# Neo4J DataSource for Grafana
Allows Neo4J to be used as a DataSource for Grafana

## Installation
* Download [Release](https://github.com/denniskniep/grafana-datasource-plugin-neo4j/releases)

## Manual and Screenshots
[Plugin Manual](https://github.com/denniskniep/grafana-datasource-plugin-neo4j/blob/main/neo4j-datasource-plugin/README.md)

## Showcase
Starts preprovisioned Grafana and Neo4J 
```
sudo docker-compose up
```

## Development

### Getting started
Optional: Use Docker for building with yarn
```bash
sudo docker run --rm -it -v $(pwd):/app node:14.17.3-alpine ash
cd /app
```

Install dependencies
```bash
cd neo4j-datasource-plugin
yarn install
```

Build Plugin in development mode
```bash
cd neo4j-datasource-plugin
yarn dev
```

Build Plugin in watch mode (development mode with auto build on change)
```bash
cd neo4j-datasource-plugin
yarn watch
```

Build plugin in production mode
```bash
cd neo4j-datasource-plugin
yarn build
```

Sign plugin

```bash
cd neo4j-datasource-plugin
export GRAFANA_API_KEY=<GRAFANA_API_KEY>
yarn sign --rootUrls http://localhost:3000/
```

Starts preprovisioned Grafana and Neo4J for development
```
sudo docker-compose -f docker-compose.dev.yaml up
```
Grafana: http://localhost:3000

Neo4J: http://localhost:7474

Grafana is started by docker-compose in development mode therefore no restart of grafana is required when source code changed.


### Learn more

- [Build a data source plugin tutorial](https://grafana.com/tutorials/build-a-data-source-plugin)
- [Grafana documentation](https://grafana.com/docs/)
- [Grafana Tutorials](https://grafana.com/tutorials/) - Grafana Tutorials are step-by-step guides that help you make the most of Grafana
- [Grafana UI Library](https://developers.grafana.com/ui) - UI components to help you build interfaces using Grafana Design System
- [Grafana Toolkit](https://github.com/grafana/grafana/tree/main/packages/grafana-toolkit#usage)
