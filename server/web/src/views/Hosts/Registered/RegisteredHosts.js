import React, { Component } from 'react'
import HostActions from "./HostActions/HostActions"
import NewHostPopup from '../NewHost/NewHost'
import { connect } from 'react-redux'
import {
    fetchHosts
} from '../../../states/actions'

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

class RegisteredHosts extends Component {

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
                <div className="btn-toolbar mb-3" role="toolbar" aria-label="Toolbar with button groups">
                  <div className="btn-group mr-2" role="group" aria-label="1 group">
                    <button type="button" className="btn btn-secondary" onClick={() => this.props.fetchHosts({})}><i className="fa fa-refresh"></i></button>
                  </div>

                  <div className="btn-group mr-2" role="group" aria-label="2 group">
                    {/*<button type="button" className="btn btn-secondary"><i className="fa fa-plus"></i>&nbsp; 创建</button>*/}
                      <NewHostPopup/>
                  </div>


                  <div className="btn-group mr-2" role="group" aria-label="2 group">
                    <HostActions/>
                  </div>

                </div>


                <table className="table table-sm">
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
                            console.log(host)
                          return <tr>
                            <td><input type="checkbox"/></td>
                              <td>{host.systemId}</td>
                            <td>{host.datacenter}</td>
                            <td>{host.rack}-{host.slot}</td>
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


                <nav>
                  <ul className="pagination">
                    <li className="page-item"><a className="page-link" href="#">前一页</a></li>
                    <li className="page-item active">
                      <a className="page-link" href="#">1</a>
                    </li>
                    <li className="page-item"><a className="page-link" href="#">2</a></li>
                    <li className="page-item"><a className="page-link" href="#">3</a></li>
                    <li className="page-item"><a className="page-link" href="#">4</a></li>
                    <li className="page-item"><a className="page-link" href="#">后一页</a></li>
                  </ul>
                </nav>

      </div>

    )
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
