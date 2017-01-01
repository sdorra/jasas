import * as urls from './urls';

describe('getQueryParameters', () => {

  it('should return map of parameters', () => {
    Object.defineProperty(window.location, 'search', {
      writable: true,
      value: '?test=abc&a=b&b=c'
    });

    let params = urls.getQueryParameters();

    expect(params['test']).toBe('abc');
    expect(params['a']).toBe('b');
    expect(params['b']).toBe('c');
  });

});

describe('getQueryParameter', () => {

  it('should return value of param a', () => {
    Object.defineProperty(window.location, 'search', {
      writable: true,
      value: '?a=b'
    });

    expect(urls.getQueryParameter('a')).toBe('b');
  });

});
