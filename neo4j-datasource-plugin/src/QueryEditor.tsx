import React, { PureComponent } from 'react';
import { InlineFieldRow, InlineFormLabel, QueryField, Select } from '@grafana/ui';
import { QueryEditorProps, SelectableValue } from '@grafana/data';
import { DataSource } from './datasource';
import { MyDataSourceOptions, MyQuery, Format } from './types';

type Props = QueryEditorProps<DataSource, MyQuery, MyDataSourceOptions>;

const Formats = [
  {
    label: 'Table',
    value: Format.Table,
    description: 'Table View',
  },
  {
    label: 'Node Graph',
    value: Format.NodeGraph,
    description: 'Node Graph View',
  },
] as Array<SelectableValue<Format>>;

export class QueryEditor extends PureComponent<Props> {
  onCypherQueryChange = (value: string) => {
    const { onChange, query } = this.props;
    onChange({ ...query, cypherQuery: value });
  };

  onFormatChanged = (selected: SelectableValue<Format>) => {
    const { onChange, query, onRunQuery } = this.props;
    onChange({ ...query, Format: selected.value || Format.Table });
    onRunQuery();
  };

  resolveFormat = (value: string | undefined) => {
    if (value === Format.NodeGraph) {
      return Formats[1];
    }
    return Formats[0];
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
        <InlineFieldRow>
          <InlineFormLabel width={5}>Format</InlineFormLabel>
          <Select
            className="width-14"
            value={this.resolveFormat(this.props.query.Format)}
            options={Formats}
            defaultValue={Formats[0]}
            onChange={this.onFormatChanged}
            width="auto"
          />
        </InlineFieldRow>
      </div>
    );
  }
}
