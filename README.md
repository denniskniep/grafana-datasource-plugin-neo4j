[![CI](https://github.com/denniskniep/grafana-datasource-plugin-neo4j/actions/workflows/ci.yml/badge.svg)](https://github.com/denniskniep/grafana-datasource-plugin-neo4j/actions/workflows/ci.yml) 
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/denniskniep/grafana-datasource-plugin-neo4j?sort=semver)](https://github.com/denniskniep/grafana-datasource-plugin-neo4j/releases/latest)



# Neo4j DataSource for Grafana
Allows Neo4j to be used as a DataSource for Grafana

## Showcase/Quickstart
Starts preprovisioned Grafana and Neo4j 
```
sudo docker-compose up
```

## Screenshots and Plugin Manual
[Plugin Manual](https://github.com/denniskniep/grafana-datasource-plugin-neo4j/blob/main/neo4j-datasource-plugin/README.md)

[Changelog](https://github.com/denniskniep/grafana-datasource-plugin-neo4j/blob/main/neo4j-datasource-plugin/CHANGELOG.md)

## Installation
* Download [Release](https://github.com/denniskniep/grafana-datasource-plugin-neo4j/releases)
* [Install](https://grafana.com/docs/grafana/latest/plugins/installation/#install-a-packaged-plugin)

## Development -  Getting started

### Frontend
Optional: Use Docker for building with yarn
```bash
sudo docker run --rm -w /app -it -v $(pwd):/app node:14.17.3-alpine ash
```

Change into plugin directory
```bash
cd neo4j-datasource-plugin
```

Install dependencies
```bash
yarn install
```

Build Plugin in development mode
```bash
yarn dev
```

Build Plugin in watch mode (development mode with auto build on change)
You need to hit refresh in browser to load the changes!
```bash
yarn watch
```

Build plugin in production mode
```bash
yarn build
```

Execute Prettier
```bash
yarn prettier --write .
```

### Backend
Optional: Use Docker for building with go
```bash
sudo docker run --rm -w /app -it -v $(pwd):/app golang:1.16 bash
go get -u github.com/magefile/mage
```

1. Update [Grafana plugin SDK for Go](https://grafana.com/docs/grafana/latest/developers/plugins/backend/grafana-plugin-sdk-for-go/) dependency to the latest minor version:

   ```bash
   go get -u github.com/grafana/grafana-plugin-sdk-go
   go mod tidy
   ```

2. Build with go:

   ```bash
   go get ./...
   go build ./...
   ```

2. Build backend plugin binaries for Linux, Windows and Darwin:

   ```bash
   mage -v
   ```

3. List all available Mage targets for additional commands:

   ```bash
   mage -l
   ```

4. Run all tests & coverage:

   Important: Start Docker-Compose environment first!

   ```bash
   sudo docker-compose -f docker-compose.dev.yaml up
   ``` 
   ```bash
   mage coverage
   ```

5. Run tests which are independet of Docker-Compose environment
   ```bash
   mage coverageShort
   ```


### Test with Grafana
Starts preprovisioned Grafana and Neo4J for development
```
sudo docker-compose -f docker-compose.dev.yaml up
```
Grafana: http://localhost:3000

Neo4J: http://localhost:7474

Grafana is started by docker-compose in development mode therefore no restart of grafana is required when source code changed.

### Run example queries

Nodes
```
Match(m:Movie) return m
```

Tabledata
```
Match(m:Movie) return m.title, m.tagline
```

Timeseriesdata
```
return datetime() - duration({minutes: 1})  as Time, 99 as Test
UNION ALL
return datetime() - duration({minutes: 2})  as Time, 85 as Test
UNION ALL
return datetime() - duration({minutes: 3})  as Time, 86 as Test
UNION ALL
return datetime() - duration({minutes: 4})  as Time, 100 as Test
UNION ALL
return datetime() - duration({minutes: 5})  as Time, 32 as Test
```

## Signing

Sign plugin as private 

```bash
cd neo4j-datasource-plugin
export GRAFANA_API_KEY=<GRAFANA_API_KEY>
yarn sign --rootUrls http://localhost:3000/
```

Sign plugin as community 

```bash
cd neo4j-datasource-plugin
export GRAFANA_API_KEY=<GRAFANA_API_KEY>
yarn sign
```


## Learn more
- [Build a data source plugin tutorial](https://grafana.com/tutorials/build-a-data-source-plugin)
- [Build a data source backend plugin tutorial](https://grafana.com/tutorials/build-a-data-source-backend-plugin/)
- [Example data source backend plugin](https://github.com/grafana/grafana-starter-datasource-backend)
- [Grafana documentation](https://grafana.com/docs/)
- [Grafana Tutorials](https://grafana.com/tutorials/) - Grafana Tutorials are step-by-step guides that help you make the most of Grafana
- [Grafana UI Library](https://developers.grafana.com/ui) - UI components to help you build interfaces using Grafana Design System
- [Grafana Toolkit](https://github.com/grafana/grafana/tree/main/packages/grafana-toolkit#usage)
- [Grafana plugin SDK for Go](https://grafana.com/docs/grafana/latest/developers/plugins/backend/grafana-plugin-sdk-for-go/)
- [Roadmap: Grafana plugins platform](https://github.com/grafana/grafana/issues/36228)
