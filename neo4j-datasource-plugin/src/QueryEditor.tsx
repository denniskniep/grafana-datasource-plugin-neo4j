import React, { PureComponent } from 'react';
import { CodeEditor, Monaco, MonacoEditor } from '@grafana/ui';
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
        <CodeEditor
          language={""}
          showMiniMap={false}
          showLineNumbers={true}
          height={250}
          value={this.props.query.cypherQuery || ''}
          onBlur={this.onCypherQueryChange}
          onSave={this.onCypherQueryChange}         
        />
      </div>
    );
  }
}
