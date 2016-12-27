import * as React from 'react';
import { LoginForm } from './login-form';
import { Welcome } from './welcome';

import { connect } from "react-redux";
import { bindActionCreators } from "redux";

import { authenticate, logout } from './actions/authentication'

interface StateProps {
  loading: boolean;
  authenticated: boolean;
  error: boolean;
  username: string;
  message: string;
}

interface DispatchProps {
  authenticate(username: string, password: string): void;
  logout(): void;
}

type ContainerProps = StateProps & DispatchProps;

function mapStateToProps(state: any) {
  return {
    error: state.error,
    message: state.message,
    loading: state.loading,
    authenticated: state.authenticated,
    username: state.username
  }
}

function mapDispatchToProps(dispatch: any) {
  return {
    authenticate: (username: string, password: string): void => dispatch(authenticate(username, password)),
    logout: (): void => dispatch(logout()),
  }
}

@connect<StateProps, DispatchProps, ContainerProps>(mapStateToProps, mapDispatchToProps)
export class Container extends React.Component<ContainerProps, any> {

  render() {
    let body;
    if (this.props.loading) {
      body = <p>Loading ...</p>;
    } else if (this.props.authenticated) {
      body = <Welcome username={this.props.username} logout={this.props.logout} />;
    } else {
      body = <LoginForm authenticate={this.props.authenticate} />;
    }

    let error = <p></p>; 
    if (this.props.error) {
      error = <p className="error">{this.props.message}</p>;
    }

    return (
      <div className="form">
        <h1>Jasas</h1>
        <p>Just another simple authentication service</p>
        {error}
        {body}
      </div>
    )
  }
}
