import React, { Component } from 'react'
import { TabContent, TabPane, Nav, NavItem, NavLink, Card, Button, CardTitle, CardText, Row, Col } from 'reactstrap';
import classnames from 'classnames';

import HotTable from 'react-handsontable';

import Collapsible from 'react-collapsible';

class EditCell extends Component {

  constructor (props) {
      super(props);
  }



  render() {
    return (
        <Row>
            <Col>{this.props.children}</Col>
            <Col><button className="button btn-link"><i className="fa fa-pencil"/></button></Col>
        </Row>
    )
  }



}

export default EditCell
