import React, { PureComponent } from 'react';
import { ReactMonacoEditor } from '@grafana/ui';
import { QueryEditorProps } from '@grafana/data';
import { DataSource } from './datasource';
import { MyDataSourceOptions, MyQuery } from './types';

type Props = QueryEditorProps<DataSource, MyQuery, MyDataSourceOptions>;

export class QueryEditor extends PureComponent<Props> {
  onCypherQueryChange = (value: string | undefined) => {
    const { onChange, query } = this.props;
     onChange({ ...query, cypherQuery: value || '' });
  };

  render() {
    return (
      <div style={{height: "240px"}}>
        <ReactMonacoEditor options={{ minimap: {enabled : false}, automaticLayout: true}} value={this.props.query.cypherQuery || ''} language={'cypher'} onChange={this.onCypherQueryChange}/>
      </div>
    );
  }
}
