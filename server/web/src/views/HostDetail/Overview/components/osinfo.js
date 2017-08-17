import React, { Component } from 'react'
import { TabContent, TabPane, Nav, NavItem, NavLink, Card, Button, CardTitle, CardText, Row, Col } from 'reactstrap';
import classnames from 'classnames';

import HotTable from 'react-handsontable';

import Collapsible from 'react-collapsible';

class OsInfo extends Component {

  constructor (props) {
      super(props);
  }


  onCellClick() {
      console.log("cell clicked")
  }

  render() {
    return (
            <Collapsible trigger="配置：OS" open={true} transitionTime={200}>

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
                      <td>Type</td>
                      <td width="40%">
                          <div className="row">
                              <div className="col">linux</div>
                                  <div className="col">
                                      <button className="button btn-link"><i className="fa fa-pencil"/></button>
                              </div>
                          </div>
                      </td>
                      <td>linux</td>
                  </tr>
                  <tr>
                      <td>Arch</td>
                      <td>
                          <Row>
                              <Col>amd64</Col>
                              <Col><button className="button btn-link"><i className="fa fa-pencil"/></button></Col>
                          </Row>
                      </td>
                      <td>amd64</td>
                  </tr>
                  <tr>
                      <td>Hostname</td>
                      <td>
                          <Row>
                              <Col>k2data-1</Col>
                              <Col><button className="button btn-link"><i className="fa fa-pencil"/></button></Col>
                          </Row>
                      </td>
                      <td>k2data-1</td>
                  </tr>
                  <tr>
                      <td>Distribution</td>
                      <td>-</td>
                      <td>Ubuntu-14.04</td>
                  </tr>
                  <tr>
                      <td>Uptime</td>
                      <td>-</td>
                      <td>1d 2h 3s</td>
                  </tr>
                  </tbody>
                </table>
            </Collapsible>
    )
  }



}

export default OsInfo
