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
* [Install](https://grafana.com/docs/grafana/latest/administration/plugin-management/#install-grafana-plugins)

## Development -  Getting started

### Frontend
[Grafana Ui - Component library](https://developers.grafana.com/ui/latest/index.html)


Optional: Use Docker for building with yarn
```bash
sudo docker run --rm -w /app -it -v $(pwd):/app node:18.17.0-alpine ash
```

Change into plugin directory
```bash
cd neo4j-datasource-plugin
```

Install dependencies
```bash
npm install
```

Build Plugin in development mode (incl. auto build on change)
```bash
npm run dev
```


Build plugin in production mode
```bash
yanpm run build
```

Execute Prettier
```bash
yarn prettier --write .
```

### Backend
> **_NOTE:_**  Depending on your version of go you might need `go get -u` instead of `go install`.

Optional: Use Docker for building with go
```bash
sudo docker run --rm -w /app -it -v $(pwd):/app golang:1.20 bash
cd neo4j-datasource-plugin
go install github.com/magefile/mage
```

1. Update [Grafana plugin SDK for Go](https://grafana.github.io/plugin-tools/docs/development/backend#update-the-go-sdk) dependency to the latest minor version:

   ```bash
   go get -u github.com/grafana/grafana-plugin-sdk-go
   go mod tidy
   ```


1. Build backend plugin binaries for Linux, Windows and Darwin:

   ```bash
   export GOFLAGS=-buildvcs=false
   mage -v
   ```

1. List all available Mage targets for additional commands:

   ```bash
   mage -l
   ```

1. Build with go:

   ```bash
   export GOFLAGS=-buildvcs=false

   go install  ./...
   go build  ./...
   ```

1. Run all tests & coverage:

   Important: Start Docker-Compose environment first!

   ```bash
   sudo docker-compose -f docker-compose.dev.yaml up
   ``` 
   ```bash
   mage coverage
   ```

1. Run tests which are independet of Docker-Compose environment
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

Grafana is started by docker-compose in development mode therefore no restart of grafana is required when frontend source code changed.

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
[Docs](https://grafana.github.io/plugin-tools/docs/distribution/signing-your-plugin)

### Sign plugin as private 

```bash
cd neo4j-datasource-plugin
export GRAFANA_API_KEY=<GRAFANA_API_KEY>
npx @grafana/sign-plugin@latest --rootUrls http://localhost:3000/
```

### Sign plugin as community 

```bash
cd neo4j-datasource-plugin
export GRAFANA_API_KEY=<GRAFANA_API_KEY>
npx @grafana/sign-plugin@latest
```


## Learn more
- [Grafana Plugin development - Getting Started](https://grafana.github.io/plugin-tools/docs/getting-started/)
- [Create a grafana plugin](https://grafana.com/docs/grafana/latest/developers/plugins/create-a-grafana-plugin/)
- [Update Create Plugin Tool](https://grafana.github.io/plugin-tools/docs/getting-started/updating-to-new-releases)
- [Build a data source backend plugin tutorial](https://grafana.com/tutorials/build-a-data-source-backend-plugin/)
- [Example data source backend plugin](https://github.com/grafana/grafana-plugin-examples)
- [Grafana documentation](https://grafana.com/docs/)
- [Grafana Tutorials](https://grafana.com/tutorials/) - Grafana Tutorials are step-by-step guides that help you make the most of Grafana
- [Grafana UI Library](https://developers.grafana.com/ui) - UI components to help you build interfaces using Grafana Design System
