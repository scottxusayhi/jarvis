import React, { Component } from 'react'
import { TabContent, TabPane, Nav, NavItem, NavLink, Card, Button, CardTitle, CardText, Row, Col } from 'reactstrap';
import classnames from 'classnames';

import HotTable from 'react-handsontable';
import Collapsible from 'react-collapsible';
class Times extends Component {

  constructor (props) {
      super(props);
  }



  render() {
    return (
                <table className="table table-sm table-bordered">
                  <tbody>
                  <tr>
                      <td width="50%">注册时间</td>
                      <td width="25%">goldwind</td>
                      <td width="25%">goldwind</td>
                  </tr>
                  <tr>
                      <td>最后一次修改时间</td>
                      <td>01</td>
                      <td>goldwind</td>
                  </tr>
                  <tr>
                      <td>第一次心跳时间</td>
                      <td>010203</td>
                      <td>goldwind</td>
                  </tr>
                  <tr>
                      <td>最后一次心跳时间</td>
                      <td>油油</td>
                      <td>goldwind</td>
                  </tr>
                  </tbody>
                </table>
    )
  }



}

export default Times
