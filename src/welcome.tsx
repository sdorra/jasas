import * as React from "react";

interface Props {
  username: string;
  logout(): void;
}

export class Welcome extends React.Component<Props, any> {

  constructor(props: Props) {
    super(props);
    this.handleLogoutClick = this.handleLogoutClick.bind(this);
  }

  handleLogoutClick(e: any) {
    e.preventDefault();
    this.props.logout();
  }

  render() {
    return (
      <div className="welcome">
        <p>Welcome {this.props.username}</p>
        <button onClick={this.handleLogoutClick}>Logout</button>
      </div>
    )
  }
}
