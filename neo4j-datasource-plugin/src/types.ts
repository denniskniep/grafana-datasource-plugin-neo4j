import { DataQuery, DataSourceJsonData } from '@grafana/data';

export interface MyQuery extends DataQuery {
  cypherQuery: string;
  Format: Format;
}

// Define Format enum for visualization format in the Query Editor
export enum Format {
  Table = 'table',
  NodeGraph = 'nodegraph',
}

export type FormatInterface = {
  [key in Format]: string;
};

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
