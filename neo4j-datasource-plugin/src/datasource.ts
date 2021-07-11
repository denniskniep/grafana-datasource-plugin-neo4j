import {
  DataQueryRequest,
  DataQueryResponse,
  DataSourceApi,
  DataSourceInstanceSettings,
  MutableDataFrame,
  FieldType
} from '@grafana/data';

import { getTemplateSrv } from '@grafana/runtime';
import neo4j, { Driver } from 'neo4j-driver';
import { MyQuery, MyDataSourceOptions } from './types';

export class DataSource extends DataSourceApi<MyQuery, MyDataSourceOptions> {

  dataSourceOptions: MyDataSourceOptions;

  constructor(instanceSettings: DataSourceInstanceSettings<MyDataSourceOptions>) {
    super(instanceSettings);
    this.dataSourceOptions = instanceSettings.jsonData;    
  }

  async query(options: DataQueryRequest<MyQuery>): Promise<DataQueryResponse> {  

    let data: MutableDataFrame[] = []
    for (const element in options.targets) {
        const query = options.targets[element];   
        let cypherQuery = await buildCypherQuery(query, options, this.dataSourceOptions);     
        let dataFrame = await executeCypherQuery(query.refId, cypherQuery, this.dataSourceOptions);        
        data.push(dataFrame);
    }
    return { data };       
  }

  async testDatasource() {
    // Implement a health check for your data source.
   
    await executeCypherQuery("A", "Match(a) return a limit 1", this.dataSourceOptions);        
   
    return {
      status: 'success',
      message: 'Success',
    };
  }
}

async function buildCypherQuery(query : MyQuery, options: DataQueryRequest<MyQuery>, dataSourceOptions: MyDataSourceOptions): string  {
  //const { range } = options;    
  //const from = range!.from.valueOf();
  //const to = range!.to.valueOf();
  return getTemplateSrv().replace(query.cypherQuery, options.scopedVars);
}

async function executeCypherQuery(refId : string, query : string, dataSourceOptions: MyDataSourceOptions): Promise<MutableDataFrame>  {
  let result;
  
  let driver : Driver;
  if(dataSourceOptions.username && dataSourceOptions.password){
    driver = neo4j.driver(dataSourceOptions.url, neo4j.auth.basic(dataSourceOptions.username, dataSourceOptions.password))
  }else{
    driver = neo4j.driver(dataSourceOptions.url)
  }
  
  const session = driver.session({
    database: dataSourceOptions.database,
    defaultAccessMode: neo4j.session.READ
  })
  
  try {
    result = await session.run(query)
  } finally {
    await session.close()
  }

  // on application exit:
  await driver.close()

  let dataFrame = new MutableDataFrame({
    refId: refId,
    fields: [],
    meta: {
      preferredVisualisationType: 'table',
    }
  });

  if(result.records.length == 0){
    return dataFrame;
  }

  for (const columnName of result.records[0].keys) { 
    dataFrame.addField({ name: columnName.toString(), type: FieldType.string })     
  }  

  for (const record of result.records) {     
    let row : any[] = [];
    record.map((value, entries, key) => {
      row.push(value);
    });
    dataFrame.appendRow(row)   
  }

  return dataFrame;
}
