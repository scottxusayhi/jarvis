import React, { Component } from 'react'
import { TabContent, TabPane, Nav, NavItem, NavLink, Card, Button, CardTitle, CardText, Row, Col } from 'reactstrap';
import classnames from 'classnames';

import HotTable from 'react-handsontable';

import Collapsible from 'react-collapsible';
import EditCell from "./editcell";

class MemInfo extends Component {

  constructor (props) {
      super(props);
  }



  render() {
    return (
        <Collapsible trigger="配置：内存" open={true} transitionTime="200">
                <table className="table table-sm table-bordered">
                    <thead>
                    <tr>
                        <th width="20%"></th>
                        <th width="40%">注册信息</th>
                        <th width="40%">检测信息</th>
                    </tr>
                    </thead>                    
                  <tbody>
                  <tr>
                      <td>Total</td>
                      <td><EditCell>32 GB</EditCell></td>
                      <td>32 GB</td>
                  </tr>
                  <tr>
                      <td>Available</td>
                      <td>-</td>
                      <td>16 GB</td>
                  </tr>
                  <tr>
                      <td>Used</td>
                      <td>-</td>
                      <td>16 GB</td>
                  </tr>
                  </tbody>
                </table>
        </Collapsible>
    )
  }



}

export default MemInfo
