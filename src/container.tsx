import * as React from 'react';
import { LoginForm } from './login-form';
import { Welcome } from './welcome';

import { connect } from "react-redux";
import { bindActionCreators } from "redux";
import { State } from './types/state';
import { AuthenticationForm } from './types/authentication-form';
import { AuthenticationState } from './types/authentication-state';
import { authenticate, logout, changeForm } from './actions/authentication'

interface DispatchProps {
  authenticate(username: string, password: string): void;
  change(form: AuthenticationForm): void;
  logout(): void;
}

type ContainerProps = AuthenticationState & DispatchProps;

function mapDispatchToProps(dispatch: any) {
  return {
    authenticate: (username: string, password: string): void => dispatch(authenticate(username, password)),
    change: (form: AuthenticationForm): void => dispatch(changeForm(form)),
    logout: (): void => dispatch(logout()),
  }
}

@connect<AuthenticationState, DispatchProps, ContainerProps>(
  (state: State) => state.auth, mapDispatchToProps
)
export class Container extends React.Component<ContainerProps, any> {

  render() {
    if (this.props.authenticated && this.props.redirectUri) {
      window.location.href = this.props.redirectUri;
    }

    let body;
    if (this.props.loading) {
      body = <p>Loading ...</p>;
    } else if (this.props.authenticated) {
      body = <Welcome username={this.props.username} logout={this.props.logout} />;
    } else {
      body = <LoginForm form={this.props.form} change={this.props.change} authenticate={this.props.authenticate} />;
    }

    let error = <p></p>; 
    if (this.props.error) {
      error = <p className="error">{this.props.message}</p>;
    }

    return (
      <div className="container">
        <h1>Jasas</h1>
        <p>Just another simple authentication service</p>
        {error}
        {body}
      </div>
    )
  }
}
