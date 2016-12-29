import * as React from 'react';
import { AuthenticationForm } from './types/authentication-form';

interface Props {
  form: AuthenticationForm
  authenticate(username: string, password: string): void;
  change(form: AuthenticationForm): void;
}

export class LoginForm extends React.Component<Props, {}> {

  usernameInput: HTMLInputElement;

  constructor(props: Props) {
    super(props);

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  componentDidMount() {
    this.usernameInput.focus();
  }

  handleChange(e: any) {
    let object: any = {};
    object[e.target.name] = e.target.value;
    this.props.change(Object.assign({}, this.props.form, object));
  }

  handleSubmit(e: any) {
    e.preventDefault();
    this.props.authenticate(this.props.form.username, this.props.form.password);
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit}>
        <input ref={(input) => { this.usernameInput = input; }} 
               type="text" 
               name="username" 
               placeholder="username" 
               value={this.props.form.username} 
               onChange={this.handleChange} />
               
        <input type="password" 
               name="password" 
               placeholder="password" 
               value={this.props.form.password} 
               onChange={this.handleChange} />
        <button>login</button>
      </form>
    )
  }
}
