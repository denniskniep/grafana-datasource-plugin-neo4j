import {
  DataQueryRequest,
  DataQueryResponse,
  DataSourceApi,
  DataSourceInstanceSettings,
  MutableDataFrame,
  FieldType
} from '@grafana/data';

import neo4j from 'neo4j-driver'
import { MyQuery, MyDataSourceOptions } from './types';

export class DataSource extends DataSourceApi<MyQuery, MyDataSourceOptions> {

  dataSourceOptions: MyDataSourceOptions;

  constructor(instanceSettings: DataSourceInstanceSettings<MyDataSourceOptions>) {
    super(instanceSettings);
    this.dataSourceOptions = instanceSettings.jsonData;    
  }

  async query(options: DataQueryRequest<MyQuery>): Promise<DataQueryResponse> {    
    //const { range } = options;    
    //const from = range!.from.valueOf();
    //const to = range!.to.valueOf();

    let data: MutableDataFrame[] = []
    for (const element in options.targets) {
        const query = options.targets[element];        
        let dataFrame = await cypherQuery(query, this.dataSourceOptions);        
        data.push(dataFrame);
    }
    return { data };       
  }

  async testDatasource() {
    // Implement a health check for your data source.
    return {
      status: 'success',
      message: 'Success',
    };
  }
}

async function  cypherQuery(query : MyQuery, dataSourceOptions: MyDataSourceOptions): Promise<MutableDataFrame>  {
  const driver = neo4j.driver(dataSourceOptions.url)
  const session = driver.session()
  
  let result;
  try {
    result = await session.run(query.cypherQuery)

    console.log(result)
  } finally {
    await session.close()
  }

  // on application exit:
  await driver.close()

  let dataFrame = new MutableDataFrame({
    refId: query.refId,
    fields: []
  });
  console.log("------")
  for (const columnName of result.records[0].keys) { 
    console.log(columnName.toString())
    dataFrame.addField({ name: columnName.toString(), type: FieldType.string })     
  }  
  console.log("------")
  for (const record of result.records) {     
    let row : any[] = [];
    record.map((value, entries, key) => {
      console.log(value,entries, key )
      row.push(value);
    });
    dataFrame.appendRow(row)   
  }
  console.log("------")
  return dataFrame;
}