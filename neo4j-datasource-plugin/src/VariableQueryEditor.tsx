import React, { useState } from 'react';
import { MyQuery } from './types';
import { QueryField } from '@grafana/ui';

interface VariableQueryProps {
  query: MyQuery;
  onChange: (query: MyQuery, definition: string) => void;
}

export const VariableQueryEditor: React.FC<VariableQueryProps> = ({ onChange, query }) => {
  const [state, setState] = useState(query);

  const saveQuery = () => {
    onChange(state, `${state.cypherQuery}`);
  };

  const handleChange = (value: string) =>
    setState({
      ...state,
      cypherQuery: value
    });
    
  return (
    <>     
      <div className="gf-form">
        <span className="gf-form-label width-10">Query</span>

        <QueryField
          portalOrigin="mock-origin"
          onBlur={saveQuery}
          onChange={handleChange}
          query={state.cypherQuery || ''}
          placeholder="Enter a cypher query"
        />

      </div>
    </>
  );
};