import * as React from 'react';

import { Container } from './container';

export class App extends React.Component<any, any> {
  render() {
    return (
      <Container {...this.props} />
    )
  }
}
