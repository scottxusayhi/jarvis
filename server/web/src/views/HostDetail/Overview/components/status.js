import React, { Component } from 'react'
import { TabContent, TabPane, Nav, NavItem, NavLink, Card, Button, CardTitle, CardText, Row, Col } from 'reactstrap';
import classnames from 'classnames';

import HotTable from 'react-handsontable';
import Collapsible from 'react-collapsible';
class Status extends Component {

  constructor (props) {
      super(props);
  }



  render() {
    return (
        <Collapsible trigger="状态" open={true} transitionTime="200">
                <table className="table table-sm table-bordered">
                    <thead>
                    <tr>
                        <th width="20%">已注册</th>
                        <th width="20%">连接过</th>
                        <th width="20%">配置审计</th>
                        <th width="20%">在线</th>
                        <th width="20%">健康</th>
                    </tr>
                  </thead>
                  <tbody>
                  <tr>
                      <td>YES</td>
                      <td>YES</td>
                      <td>YES</td>
                      <td>YES</td>
                      <td>YES</td>
                  </tr>
                  </tbody>
                </table>
        </Collapsible>
    )
  }



}

export default Status
