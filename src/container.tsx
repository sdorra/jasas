import * as React from 'react';
import { LoginForm } from './login-form';
import { Welcome } from './welcome';

import { connect } from "react-redux";
import { bindActionCreators } from "redux";

import { authenticate, logout, changeForm, AuthenticationForm } from './actions/authentication'

interface StateProps {
  loading: boolean;
  authenticated: boolean;
  error: boolean;
  username: string;
  message: string;
  form: AuthenticationForm;
  redirectUri?: string;
}

interface DispatchProps {
  authenticate(username: string, password: string): void;
  change(form: AuthenticationForm): void;
  logout(): void;
}

type ContainerProps = StateProps & DispatchProps;

function mapStateToProps(state: any) {
  return {
    error: state.auth.error,
    message: state.auth.message,
    loading: state.auth.loading,
    authenticated: state.auth.authenticated,
    username: state.auth.username,
    form: state.auth.form,
    redirectUri: state.auth.redirectUri
  }
}

function mapDispatchToProps(dispatch: any) {
  return {
    authenticate: (username: string, password: string): void => dispatch(authenticate(username, password)),
    change: (form: AuthenticationForm): void => dispatch(changeForm(form)),
    logout: (): void => dispatch(logout()),
  }
}

@connect<StateProps, DispatchProps, ContainerProps>(mapStateToProps, mapDispatchToProps)
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
