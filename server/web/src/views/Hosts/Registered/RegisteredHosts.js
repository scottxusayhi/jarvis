import React, { Component } from 'react'
import {
  BrowserRouter as Router,
  Route,
  Link
} from 'react-router-dom'

import NewHostPopup from '../NewHost/NewHost'
import HostActions from "./HostActions/HostActions"
import { connect } from 'react-redux'
import {
    fetchHosts,
    fetchRegisteredHosts
} from '../../../states/actions'
import ApiAlert from "../../../components/ApiAlert/ApiAlert";
import { Container, Row, Col } from 'reactstrap';
import RightView from "../../../components/RightView/RightView";
import Pager from "../../../components/Pager/Pager";

// subscribe
const mapStateToProps = state => {
    return {
        items: state.registeredHosts,
    }
}

// dispatch actions
const mapDispatchToProps = dispatch => {
    return {
        fetchRegisteredHosts: filter => {
            dispatch(fetchRegisteredHosts(filter))
        }
    }
}

class RegisteredHosts extends Component {

  constructor (props) {
    super(props);
    this.state = {
        filter: {
            registered: 1
        }
    }
  }

  componentDidMount() {
      this.props.fetchRegisteredHosts(this.state.filter)
  }

  render() {
    console.log("rendering");
    return (
      <div>
      {/*<div className="animated fadeIn">*/}
      {/*<ApiAlert/>*/}
      <Container>
        <Row>
                <div className="btn-toolbar mb-3" role="toolbar" aria-label="Toolbar with button groups">

                  <div className="btn-group mr-2" role="group" aria-label="1 group">
                    <button type="button" className="btn btn-secondary" onClick={() => this.props.fetchRegisteredHosts(this.state.filter)}><i className="fa fa-refresh"></i></button>
                  </div>

                  <div className="btn-group mr-2" role="group" aria-label="2 group">
                      <NewHostPopup/>
                  </div>
                  <div className="btn-group mr-2" role="group" aria-label="2 group">
                    <HostActions/>
                  </div>

                </div>


          <Col/>
          <Col/>
          <Col/>
          <Col><Pager pageInfo={this.props.items.data.pageInfo} onPageChange={(page)=>this.props.fetchHosts({registered:1, page: page})}/></Col>
        </Row>
                    </Container>


                <table className="table table-sm table-hover">
                  <thead>
                    <tr>
                        <th> <input type="checkbox"/> </th>
                        <th>ID</th>
                        <th>数据中心</th>
                        <th>位置</th>
                        <th>拥有人</th>
                        <th>配置审计</th>
                        <th>在线状态</th>
                        <th>VCPU</th>
                        <th>内存</th>
                        <th>硬盘</th>
                        <th>网络</th>
                        <th>操作系统</th>
                    </tr>
                  </thead>

                  <tbody>

                  {
                    this.props.items.data.list &&
                        this.props.items.data.list.map(host=> {
                          return <tr>
                            <td><input type="checkbox"/></td>
                              <td>{this.viewHostId(host)}</td>
                            <td>{this.viewDatacenter(host)}</td>
                            <td>{this.viewPosition(host)}</td>
                              <td>{host.owner}</td>
                            <td>
                                {this.viewConfigAuditStatus(host.connected, host.matched)}
                            </td>
                            <td>
                                {this.viewOnlineStatus(host.online)}
                            </td>
                            <td>{this.viewCpuInfo(host.cpuExpected)}</td>
                            <td>{this.viewMemInfo(host.memExpected)}</td>
                              <td>{this.viewDiskInfo(host.diskExpected)}</td>
                            <td>{this.viewNetworkInfo(host.networkExpected)}</td>
                            <td>{this.viewOsInfo(host.osExpected)}</td>
                          </tr>
                        })
                  }

                  </tbody>
                </table>
      </div>
    )
  }

  viewHostId(host) {
      var link = "/hosts/" + host.systemId
      return <Link to={link}>{host.systemId}</Link>
  }

  viewDatacenter(host) {
      if (host.registered) {
          return host.datacenter
      } else {
          return "N/A"
      }
  }

  viewPosition(host) {
      if (host.registered) {
          return host.rack+"-"+host.slot
      } else {
          return "N/A"
      }
  }

  viewConfigAuditStatus(connected, matched) {
      if (!connected) {
          return <span className="badge badge-default">未曾连接</span>
      }
      else if (matched) {
          return <span className="badge badge-success">匹配</span>
      }
      else {
          return <span className="badge badge-danger">不匹配</span>
      }
  }

  viewOnlineStatus(online) {
      if (online) {
          return <span className="badge badge-success">在线</span>
      } else {
          return <span className="badge badge-danger">离线</span>
      }
  }

  viewCpuInfo(cpuInfo) {
      return cpuInfo.vcpu
  }

  viewMemInfo(memInfo) {
      return memInfo.total
  }

  viewDiskInfo(diskInfo) {
      return diskInfo.length
  }

  viewNetworkInfo(netInfo) {
      return netInfo.ip
  }

  viewOsInfo(osInfo) {
      return osInfo.type
  }


}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(RegisteredHosts)
