import { DataQuery, DataSourceJsonData } from '@grafana/data';

export interface MyQuery extends DataQuery {
  cypherQuery: string;
}

/**
 * These are options configured for each DataSource instance
 */
export interface MyDataSourceOptions extends DataSourceJsonData {
  url: string;
  database?: string;
  username?: string;  
}

export interface MySecureDataSourceOptions {
  password?: string;
}
