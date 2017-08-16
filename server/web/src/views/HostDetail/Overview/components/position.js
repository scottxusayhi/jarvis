import React, { Component } from 'react'
import { TabContent, TabPane, Nav, NavItem, NavLink, Card, Button, CardTitle, CardText, Row, Col } from 'reactstrap';
import classnames from 'classnames';

import HotTable from 'react-handsontable';
import Collapsible from 'react-collapsible';
import EditCell from "./editcell";
class Position extends Component {

  constructor (props) {
      super(props);
      this.data=[
          ["数据中心", ""],
          ["机架", ""],
          ["槽位", ""],
          ["拥有人", ""],
      ]
  }



  render() {
    return (
                <table className="table table-sm table-bordered">
                  <tbody>
                  <tr>
                      <td>ID</td>
                      <td>20</td>
                  </tr>
                  <tr>
                      <td width="50%">数据中心</td>
                      <td><EditCell>goldwind</EditCell></td>
                  </tr>
                  <tr>
                      <td>机架</td>
                      <td><EditCell>01</EditCell></td>
                  </tr>
                  <tr>
                      <td>槽位</td>
                      <td><EditCell>010203</EditCell></td>
                  </tr>
                  <tr>
                      <td>拥有人</td>
                      <td><EditCell>油油</EditCell></td>
                  </tr>
                  </tbody>
                </table>
    )
  }



}

export default Position
