import React, { PureComponent } from 'react';
import { CodeEditor } from '@grafana/ui';
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
      <div className="gf-form">
        <CodeEditor
          language={""}
          showLineNumbers={true}
          showMiniMap={true}
          width={1000}
          height={300}
          value={this.props.query.cypherQuery || ''}
          onBlur={this.onCypherQueryChange}
          onSave={this.onCypherQueryChange}
        />
      </div>
    );
  }
}
