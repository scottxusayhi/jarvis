import React, { Component } from 'react'
import { TabContent, TabPane, Nav, NavItem, NavLink, Card, Button, CardTitle, CardText, Row, Col } from 'reactstrap';
import classnames from 'classnames';

import HotTable from 'react-handsontable';

import Image from "./components/image"
import Position from "./components/position";
import Times from './components/times'
import Status from './components/status'
import OsInfo from './components/osinfo'
import CpuInfo from './components/cpuinfo'
import MemInfo from './components/meminfo'
import DiskInfo from './components/diskinfo'
import NetInfo from './components/netinfo'

class Overview extends Component {

  constructor (props) {
      super(props);
  }


  render() {
    console.log("rendering host pane" + this.props.match.params.hostId);
    return (
      <div>
        <Row>
          <Col>
            <Image/>
          </Col>
          <Col>
              <Position/>
              <Times/>
          </Col>
        </Row>

        <Row>
          <Col>
              <Status/>
          </Col>
        </Row>

        <Row>
          <Col>
              <OsInfo/>
          </Col>
        </Row>

        <Row>
          <Col>
            <CpuInfo/>
          </Col>
        </Row>

        <Row>
          <Col>
            <MemInfo/>
          </Col>
        </Row>

        <Row>
          <Col>
            <DiskInfo/>
          </Col>
        </Row>
        <Row>
          <Col>
            <NetInfo/>
          </Col>
        </Row>
      </div>
    )
  }



}

export default Overview
