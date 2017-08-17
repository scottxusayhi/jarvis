import React, { Component } from 'react'
import { TabContent, TabPane, Nav, NavItem, NavLink, Card, Button, CardTitle, CardText, Row, Col } from 'reactstrap';
import classnames from 'classnames';

import HotTable from 'react-handsontable';

import Collapsible from 'react-collapsible';

class Image extends Component {

  constructor (props) {
      super(props);
  }



  render() {
    return (
    <div>
        <img src="img/default-host.png" className="rounded mx-auto d-block" alt="A generic square placeholder image with rounded corners in a figure."/>
    </div>


    )
  }
}

export default Image
