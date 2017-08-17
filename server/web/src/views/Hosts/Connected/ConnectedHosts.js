import React, { Component } from 'react'
import {
  BrowserRouter as Router,
  Route,
  Link
} from 'react-router-dom'
import HostActions from "./HostActions/HostActions"
import { connect } from 'react-redux'
import {
    fetchHosts
} from '../../../states/actions'
import ApiAlert from "../../../components/ApiAlert/ApiAlert";
import { Container, Row, Col } from 'reactstrap';
import RightView from "../../../components/RightView/RightView";
import Pager from "../../../components/Pager/Pager";

// subscribe
const mapStateToProps = state => {
    return {
        items: state.hosts,
    }
}

// dispatch actions
const mapDispatchToProps = dispatch => {
    return {
        fetchHosts: filter => {
            dispatch(fetchHosts(filter))
        }
    }
}

class ConnectedHosts extends Component {

  constructor (props) {
    super(props);
    this.state = {}
  }

  componentDidMount() {
      this.props.fetchHosts({})
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
                    <button type="button" className="btn btn-secondary" onClick={(filter) => this.props.fetchHosts(filter)}><i className="fa fa-refresh"></i></button>
                  </div>


                  <div className="btn-group mr-2" role="group" aria-label="2 group">
                    <HostActions/>
                  </div>

                </div>


          <Col></Col>
          <Col></Col>
          <Col></Col>
          <Col><Pager></Pager></Col>
        </Row>
                    </Container>


                <table className="table table-sm table-hover">
                  <thead>
                    <tr>
                      <th> <input type="checkbox"/> </th>
                        <th>ID</th>
                      <th>数据中心</th>
                      <th>位置</th>
                      <th>在线状态</th>
                      <th>健康状态</th>
                      <th>注册状态</th>
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
                            <td>
                                {this.viewOnlineStatus(host.online)}
                            </td>
                            <td>
                                {this.viewHealthStatus(host.healthStatus)}
                            </td>
                            <td>
                                {this.viewRegisterStatus(host.registerec)}
                            </td>
                            <td>{this.viewCpuInfo(host.cpuDetected)}</td>
                            <td>{this.viewMemInfo(host.memDetected)}</td>
                            <td>{this.viewDiskInfo(host.diskDetected)}</td>
                            <td>{this.viewNetworkInfo(host.networkDetected)}</td>
                            <td>{this.viewOsInfo(host.osDetected)}</td>
                          </tr>
                        })
                  }

                  </tbody>
                </table>

    {/*<RightView>*/}
        {/*<Pager/>*/}
    {/*</RightView>*/}

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

  viewOnlineStatus(online){
      if (online) {
          return <span className="badge badge-success">在线</span>
      } else {
          return <span className="badge badge-danger">离线</span>
      }
  }

  viewHealthStatus(health) {
      switch (health) {
          case "unknown": {
              return <span className="badge badge-default">未知</span>
          }
          case "ok": {
              return <span className="badge badge-success">正常</span>
          }
          case "warning": {
              return <span className="badge badge-warning">告警</span>
          }
          case "error": {
              return <span className="badge badge-danger">错误</span>
          }
          default: {
              return <span className="badge badge-default">未定义</span>
          }
      }
  }

  viewRegisterStatus(registered) {
      if (registered) {
          return <span className="badge badge-success">已注册</span>
      } else {
          return <span className="badge badge-info">未注册</span>
      }
  }

  viewCpuInfo(cpuInfo) {
      return cpuInfo.vcpu
  }

  viewMemInfo(memInfo) {
      return Math.ceil(memInfo.total/1024/1024/1024)+" GB"
  }

  viewDiskInfo(diskInfo) {
      return diskInfo.length
  }

  viewNetworkInfo(netInfo) {
      return netInfo.ip
  }

  viewOsInfo(osInfo) {
      return osInfo.type+"-"+osInfo.dist+"-"+osInfo.version+"-"+osInfo.arch
  }


}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(ConnectedHosts)
