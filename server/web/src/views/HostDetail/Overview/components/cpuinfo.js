import React, { Component } from 'react'
import { TabContent, TabPane, Nav, NavItem, NavLink, Card, Button, CardTitle, CardText, Row, Col } from 'reactstrap';
import classnames from 'classnames';

import HotTable from 'react-handsontable';

import Collapsible from 'react-collapsible';
import EditCell from './editcell'

class CpuInfo extends Component {

  constructor (props) {
      super(props);
  }



  render() {
    return (
        <Collapsible trigger="配置：CPU" open={true} transitionTime={200}>
                <table className="table table-sm table-bordered table-responsive">
                    <thead>
                    <tr>
                        <th width="20%"></th>
                        <th width="40%">注册信息</th>
                        <th width="40%">检测信息</th>
                    </tr>
                    </thead>
                  <tbody>
                  <tr>
                      <td>Socket(s)</td>
                      <td><EditCell>1</EditCell></td>
                      <td>1</td>
                  </tr>
                  <tr>
                      <td>VCPU</td>
                      <td><EditCell>8</EditCell></td>
                      <td>8</td>
                  </tr>
                  <tr>
                      <td>Model</td>
                      <td><EditCell>GenuineIntel x86_64 family 6 model 60 stepping 3</EditCell></td>
                      <td>GenuineIntel x86_64 family 6 model 60 stepping 3</td>
                  </tr>
                  </tbody>
                </table>
        </Collapsible>
    )
  }



}

export default CpuInfo
