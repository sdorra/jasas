import * as React from "react";

interface Props {
  authenticate(username: string, password: string): void;
}

interface State {
  username: string
  password: string
}

export class LoginForm extends React.Component<Props, State> {

  usernameInput: HTMLInputElement;

  constructor(props: Props) {
    super(props);
    this.state = {
      username: '',
      password: ''
    };

    // This binding is necessary to make `this` work in the callback
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  componentDidMount() {
    this.usernameInput.focus();
  }

  handleChange(e: any) {
    let obj: any = {};
    obj[e.target.name] = e.target.value;
    this.setState(Object.assign({}, this.state, obj));
  }

  handleSubmit(e: any) {
    e.preventDefault();
    this.props.authenticate(this.state.username, this.state.password);
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit}>
        <input ref={(input) => { this.usernameInput = input; }} 
               type="text" 
               name="username" 
               placeholder="username" 
               value={this.state.username} 
               onChange={this.handleChange} />
        <input type="password" 
               name="password" 
               placeholder="password" 
               value={this.state.password} 
               onChange={this.handleChange} />
        <button>login</button>
      </form>
    )
  }
}
