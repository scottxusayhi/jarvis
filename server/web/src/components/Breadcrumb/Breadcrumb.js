import React from 'react';
import { Route, Link } from 'react-router-dom';
import { Breadcrumb, BreadcrumbItem } from 'reactstrap';
import routes from '../../routes';

const findRouteName = url => routes[url];

const getPaths = (pathname) => {
  const paths = ['/'];

  if (pathname === '/') return paths;

  pathname.split('/').reduce((prev, curr, index) => {
    const currPath = `${prev}/${curr}`;
    paths.push(currPath);
    return currPath;
  });
  return paths;
};

const BreadcrumbsItem = ({ ...rest, match }) => {
  console.log(match.url, findRouteName(match.url))
  var routeName = findRouteName(match.url);
  ////// xudi if undefined, change to right-most path
  if (routeName==undefined) {
      match.url.split('/').map((o, index)=> {
          if(index==match.url.split('/').length-1) {
              routeName = o
          }
      })
  }
  ///////xudi////
  if (routeName) {
    return (
      match.isExact ?
      (
        <BreadcrumbItem active>{routeName}</BreadcrumbItem>
      ) :
      (
        <BreadcrumbItem>
          <Link to={match.url || ''}>
            {routeName}
          </Link>
        </BreadcrumbItem>
      )
    );
  }
  return null;
};

const Breadcrumbs = ({ ...rest, location : { pathname }, match }) => {
  const paths = getPaths(pathname);
  const items = paths.map((path, i) => <Route key={i++} path={path} component={BreadcrumbsItem} />);
  return (
    <Breadcrumb>
      {items}
    </Breadcrumb>

  );
};

export default props => (
  <div>
    <Route path="/:path" component={Breadcrumbs} {...props} />
  </div>
);
