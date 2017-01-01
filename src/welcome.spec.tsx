import * as React from 'react';
import { create, ReactTestRendererJSON } from 'react-test-renderer';

import { Welcome } from './welcome';

describe('welcome', () => {

  it('should render text and button', () => {
    const logout = () => {};
    const tree = create(
      <Welcome username="tricia" logout={logout} />
    ).toJSON();

    expect(tree).toMatchSnapshot();
  });

  it('should call logout, after button click', () => {
    let pressed: number = 0;
    const logout = () => {
      pressed++;
    };
    const tree = create(
      <Welcome username="tricia" logout={logout} />
    ).toJSON();

    let button: any = tree.children.find((el: any) => {
      return el.type === 'button';
    });

    let e = {
      preventDefault: () => {}
    };
    button.props.onClick(e);
    expect(pressed).toBe(1);
    
    button.props.onClick(e);
    expect(pressed).toBe(2);
  });

});
