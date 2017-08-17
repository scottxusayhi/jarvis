import React, { Component } from 'react'
import { TabContent, TabPane, Nav, NavItem, NavLink, Card, Button, CardTitle, CardText, Row, Col } from 'reactstrap';
import classnames from 'classnames';

import HotTable from 'react-handsontable';

import Collapsible from 'react-collapsible';

class DiskInfo extends Component {

  constructor (props) {
      super(props);
  }



  render() {
    return (
        <Collapsible trigger="配置：磁盘" open={true} transitionTime={200}>
                <table className="table table-sm table-bordered">
                    <thead>
                    <tr>
                        <th width="10%"></th>
                        <th width="10%"></th>
                        <th width="40%">注册信息</th>
                        <th width="40%">检测信息</th>
                    </tr>
                    </thead>                    
                  <tbody>
                  <tr>
                      <td>/dev/sda</td>
                      <td>Model</td>
                      <td>-</td>
                      <td>ST2000DM001-1ER1</td>
                  </tr>
                  <tr>
                      <td></td>
                      <td>Cap</td>
                      <td>-</td>
                      <td>2000 GB</td>
                  </tr>
                  <tr>
                      <td></td>
                      <td>Used</td>
                      <td>-</td>
                      <td>199 GB</td>
                  </tr>

                  <tr>
                      <td>/dev/sda</td>
                      <td>Model</td>
                      <td>-</td>
                      <td>ST2000DM001-1ER1</td>
                  </tr>
                  <tr>
                      <td></td>
                      <td>Cap</td>
                      <td>-</td>
                      <td>2000 GB</td>
                  </tr>
                  <tr>
                      <td></td>
                      <td>Used</td>
                      <td>-</td>
                      <td>199 GB</td>
                  </tr>

                  <tr>
                      <td>/dev/sda</td>
                      <td>Model</td>
                      <td>-</td>
                      <td>ST2000DM001-1ER1</td>
                  </tr>
                  <tr>
                      <td></td>
                      <td>Cap</td>
                      <td>-</td>
                      <td>2000 GB</td>
                  </tr>
                  <tr>
                      <td></td>
                      <td>Used</td>
                      <td>-</td>
                      <td>199 GB</td>
                  </tr>

                  </tbody>
                </table>
        </Collapsible>
    )
  }



}

export default DiskInfo
