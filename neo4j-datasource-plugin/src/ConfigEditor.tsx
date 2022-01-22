import React, { ChangeEvent, PureComponent } from 'react';
import { LegacyForms } from '@grafana/ui';
import { DataSourcePluginOptionsEditorProps } from '@grafana/data';
import { MyDataSourceOptions, MySecureDataSourceOptions } from './types';

const { SecretFormField, FormField } = LegacyForms;

interface Props extends DataSourcePluginOptionsEditorProps<MyDataSourceOptions, MySecureDataSourceOptions> {}

interface State {}

export class ConfigEditor extends PureComponent<Props, State> {
  onUrlChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { onOptionsChange, options } = this.props;
    const jsonData = {
      ...options.jsonData,
      url: event.target.value,
    };
    onOptionsChange({ ...options, jsonData });
  };

  onDatabaseChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { onOptionsChange, options } = this.props;
    const jsonData = {
      ...options.jsonData,
      database: event.target.value,
    };
    onOptionsChange({ ...options, jsonData });
  };

  onUsernameChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { onOptionsChange, options } = this.props;
    const jsonData = {
      ...options.jsonData,
      username: event.target.value,
    };
    onOptionsChange({ ...options, jsonData });
  };

  onPasswordChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { onOptionsChange, options } = this.props;
    const secureJsonData = {
      ...options.secureJsonData,
      password: event.target.value,
    };

    onOptionsChange({ ...options,  secureJsonData });
  };

  onResetPassword = () => {
    const { onOptionsChange, options } = this.props;
    const secureJsonData = {
      ...options.secureJsonData,
      password: '',
    };

    const secureJsonFields= {
      ...options.secureJsonFields,
      password: false,
    }

    onOptionsChange({ ...options, secureJsonFields, secureJsonData });
  };

  render() {    
    const { options } = this.props;
    const { jsonData, secureJsonData, secureJsonFields } = options;    

    return (
      <div className="gf-form-group">
        <div className="gf-form">
          <FormField
            label="Url"
            labelWidth={6}
            inputWidth={20}
            onChange={this.onUrlChange}
            value={jsonData.url || ''}
            placeholder="e.g. neo4j://localhost:7687"
          />
        </div>

        <div className="gf-form">
          <FormField
            label="Database"
            labelWidth={6}
            inputWidth={20}
            onChange={this.onDatabaseChange}
            value={jsonData.database || ''}
            placeholder="leave empty for default"
          />
        </div>

        <div className="gf-form">
          <FormField
            label="Username"
            labelWidth={6}
            inputWidth={20}
            onChange={this.onUsernameChange}
            value={jsonData.username || ''}
            placeholder="leave empty for no authentication"
          />
        </div>

        <div className="gf-form-inline">
          <div className="gf-form">
            <SecretFormField
              isConfigured={(secureJsonFields && secureJsonFields.password) as boolean}
              value={secureJsonData && secureJsonData.password || ''}
              label="Password"
              placeholder="leave empty for no authentication"
              labelWidth={6}
              inputWidth={20}
              onReset={this.onResetPassword}
              onChange={this.onPasswordChange}
            />
          </div>
        </div>
      </div>
    );
  }
}
