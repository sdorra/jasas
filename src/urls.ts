export function getQueryParameters(): {[key:string]:string} {
  let paramsMap: {[key:string]:string} = {};
  let queryString = window.location.search;
  if (queryString) {
    if (queryString.indexOf('?') === 0) {
      queryString = queryString.substring(1);
    }
    let params = queryString.split("&");
    params.forEach((p) => {
        var v = p.split("=");
        paramsMap[v[0]] = decodeURIComponent(v[1]);
    });
  }
  return paramsMap;
};

export function getQueryParameter(key: string) {
  return getQueryParameters()[key];
}
