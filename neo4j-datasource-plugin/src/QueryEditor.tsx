import React, { PureComponent } from 'react';
import { QueryField } from '@grafana/ui';
import { QueryEditorProps } from '@grafana/data';
import { DataSource } from './datasource';
import { MyDataSourceOptions, MyQuery } from './types';

type Props = QueryEditorProps<DataSource, MyQuery, MyDataSourceOptions>;

export class QueryEditor extends PureComponent<Props> {
  onCypherQueryChange = (value: string) => {
    const { onChange, query } = this.props;
    onChange({ ...query, cypherQuery: value });
  };

  render() {
    return (
      <div>
        <QueryField
          portalOrigin="mock-origin"
          onChange={this.onCypherQueryChange}
          query={this.props.query.cypherQuery || ''}
          placeholder="Enter a cypher query"
        />
      </div>
    );
  }
}
