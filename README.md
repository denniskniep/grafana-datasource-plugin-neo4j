Introduction: https://grafana.com/tutorials/build-a-data-source-plugin/

Using Grafana Toolkit: https://github.com/grafana/grafana/tree/main/packages/grafana-toolkit#usage


Build (-w = watch for changes)
```
cd neo4j-datasource-plugin
npx @grafana/toolkit plugin:dev -w
```

```
sudo docker-compose up
```
Grafana is started in development mode therefore no restart of grafana is required when source code changed

