import React, {Component} from 'react'

import{Row, Col} from 'antd'


import Image from "./components/image"
import Position from "./components/position";
import Times from './components/times'
import HostTags from './components/tags'
import Status from './components/status'
import OsInfo from './components/osinfo'
import CpuInfo from './components/cpuinfo'
import MemInfo from './components/meminfo'
import DiskInfo from './components/diskinfo'
import NetInfo from './components/netinfo'


class Overview extends Component {

    constructor(props) {
        super(props);
    }


    render() {
        console.log("rendering host pane" + this.props.match.params.hostId);
        return (
            <div>
                <Row>
                    <Col span={10}>
                        <Image/>
                    </Col>
                    <Col span={14}>
                        <Position/>
                        <Times/>
                        <HostTags/>
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
