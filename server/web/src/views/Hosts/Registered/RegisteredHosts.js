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
      <div className="animated fadeIn">

                <div className="btn-toolbar mb-3" role="toolbar" aria-label="Toolbar with button groups">
                  <div className="btn-group mr-2" role="group" aria-label="1 group">
                    <button type="button" className="btn btn-secondary" onClick={(filter) => this.props.fetchHosts(filter)}><i className="fa fa-refresh"></i></button>
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
                      <th>数据中心</th>
                      <th>机架</th>
                      <th>槽位</th>
                      <th>在线状态</th>
                      <th>健康状态</th>
                      <th>配置审计</th>
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
                            <td>{host.datacenter}</td>
                            <td>{host.rack}</td>
                            <td>{host.slot}</td>
                            <td>
                              <span className="badge badge-success">在线</span>
                            </td>
                            <td>
                              <span className="badge badge-success">正常</span>
                            </td>
                            <td>
                              <span className="badge badge-success">匹配</span>
                            </td>
                            <td>8</td>
                            <td>128G</td>
                            <td>4*2T 1*1T</td>
                            <td>192.168.130.100</td>
                            <td>Linux-ubuntu14-amd64</td>
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
}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(RegisteredHosts)
