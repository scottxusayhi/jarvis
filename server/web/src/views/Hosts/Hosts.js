import React, { Component } from 'react';
import HostActions from "./HostActions/HostActions";

class Hosts extends Component {

  constructor () {
    super();
    this.state = {
      items: {}
    }
  }

  componentDidMount() {
    fetch('http://localhost:2999/api/v1/hosts')
        .then(result=>{
          return result.json();
        }).then(json=>{
          console.log('parsed json', json);
          this.setState({items: json});
        }).catch(ex=>{
          console.error('parse json error', ex);
    })
  }

  render() {
    console.log("rendering")
    return (
      <div className="animated fadeIn">

        <div className="row">
          <div className="col-lg-12">
            <div className="card">

              <div className="card-header">
                <i className="fa fa-align-justify"></i> Hosts
              </div>

              <div className="card-block">


                <div className="btn-toolbar mb-3" role="toolbar" aria-label="Toolbar with button groups">
                  <div className="btn-group mr-2" role="group" aria-label="1 group">
                    <button type="button" className="btn btn-secondary"><i className="fa fa-refresh"></i></button>
                  </div>

                  <div className="btn-group mr-2" role="group" aria-label="2 group">
                    <button type="button" className="btn btn-secondary"><i className="fa fa-plus"></i>&nbsp; 创建</button>
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
                    this.state.items.list &&
                        this.state.items.list.map(host=> {
                          console.log(host);
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


                    <tr>
                      <td><input type="checkbox"/></td>
                      <td>goldwind</td>
                      <td>01</td>
                      <td>010203</td>
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

                    <tr>
                      <td><input type="checkbox"/></td>
                      <td>goldwind</td>
                      <td>01</td>
                      <td>010203</td>
                      <td>
                        <span className="badge badge-danger">离线</span>
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

                    <tr>
                      <td><input type="checkbox"/></td>
                      <td>goldwind</td>
                      <td>01</td>
                      <td>010203</td>
                      <td>
                        <span className="badge badge-success">在线</span>
                      </td>
                      <td>
                        <span className="badge badge-warning">警报</span>
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

                    <tr>
                      <td><input type="checkbox"/></td>
                      <td>goldwind</td>
                      <td>01</td>
                      <td>010203</td>
                      <td>
                        <span className="badge badge-success">在线</span>
                      </td>
                      <td>
                        <span className="badge badge-danger">错误</span>
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

                    <tr>
                      <td><input type="checkbox"/></td>
                      <td>goldwind</td>
                      <td>01</td>
                      <td>010203</td>
                      <td>
                        <span className="badge badge-success">在线</span>
                      </td>
                      <td>
                        <span className="badge badge-success">正常</span>
                      </td>
                      <td>
                        <span className="badge badge-danger">不匹配</span>
                      </td>
                      <td>8</td>
                      <td>128G</td>
                      <td>4*2T 1*1T</td>
                      <td>192.168.130.100</td>
                      <td>Linux-ubuntu14-amd64</td>
                    </tr>

                    <tr>
                      <td><input type="checkbox"/></td>
                      <td>goldwind</td>
                      <td>01</td>
                      <td>010203</td>
                      <td>
                        <span className="badge badge-success">在线</span>
                      </td>
                      <td>
                        <span className="badge badge-success">正常</span>
                      </td>
                      <td>
                        <span className="badge badge-default">未连接</span>
                      </td>
                      <td>8</td>
                      <td>128G</td>
                      <td>4*2T 1*1T</td>
                      <td>192.168.130.100</td>
                      <td>Linux-ubuntu14-amd64</td>
                    </tr>

                    <tr>
                      <td><input type="checkbox"/></td>
                      <td>goldwind</td>
                      <td>01</td>
                      <td>010203</td>
                      <td>
                        <span className="badge badge-success">在线</span>
                      </td>
                      <td>
                        <span className="badge badge-success">正常</span>
                      </td>
                      <td>
                        <span className="badge badge-info">新连接</span>
                      </td>
                      <td>8</td>
                      <td>128G</td>
                      <td>4*2T 1*1T</td>
                      <td>192.168.130.100</td>
                      <td>Linux-ubuntu14-amd64</td>
                    </tr>

                    <tr>
                      <td><input type="checkbox"/></td>
                      <td>goldwind</td>
                      <td>01</td>
                      <td>010203</td>
                      <td>
                        <span className="badge badge-success">在线</span>
                      </td>
                      <td>
                        <span className="badge badge-success">正常</span>
                      </td>
                      <td>
                        <span className="badge badge-pill">匹配</span>
                      </td>
                      <td>8</td>
                      <td>128G</td>
                      <td>4*2T 1*1T</td>
                      <td>192.168.130.100</td>
                      <td>Linux-ubuntu14-amd64</td>
                    </tr>

                    <tr>
                      <td><input type="checkbox"/></td>
                      <td>goldwind</td>
                      <td>01</td>
                      <td>010203</td>
                      <td>
                        <span className="badge badge-success">在线</span>
                      </td>
                      <td>
                        <span className="badge badge-success">正常</span>
                      </td>
                      <td>
                        <span className="badge badge-primary">匹配</span>
                      </td>
                      <td>8</td>
                      <td>128G</td>
                      <td>4*2T 1*1T</td>
                      <td>192.168.130.100</td>
                      <td>Linux-ubuntu14-amd64</td>
                    </tr>

                    <tr>
                      <td><input type="checkbox"/></td>
                      <td>goldwind</td>
                      <td>01</td>
                      <td>010203</td>
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

                    <tr>
                      <td><input type="checkbox"/></td>
                      <td>goldwind</td>
                      <td>01</td>
                      <td>010203</td>
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

                    <tr>
                      <td><input type="checkbox"/></td>
                      <td>goldwind</td>
                      <td>01</td>
                      <td>010203</td>
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

                    <tr>
                      <td><input type="checkbox"/></td>
                      <td>goldwind</td>
                      <td>01</td>
                      <td>010203</td>
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

                    <tr>
                      <td><input type="checkbox"/></td>
                      <td>goldwind</td>
                      <td>01</td>
                      <td>010203</td>
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

                    <tr>
                      <td><input type="checkbox"/></td>
                      <td>goldwind</td>
                      <td>01</td>
                      <td>010203</td>
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

                    <tr>
                      <td><input type="checkbox"/></td>
                      <td>goldwind</td>
                      <td>01</td>
                      <td>010203</td>
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

                    <tr>
                      <td><input type="checkbox"/></td>
                      <td>Zbyněk Phoibos</td>
                      <td>20</td>
                      <td>Staff</td>
                      <td>
                        <span className="badge badge-danger">Banned</span>
                      </td>
                    </tr>

                    <tr>
                      <td><input type="checkbox"/></td>
                      <td>Einar Randall</td>
                      <td>20</td>
                      <td>Admin</td>
                      <td>
                        <span className="badge badge-default">Inactive</span>
                      </td>
                    </tr>

                    <tr>
                      <td><input type="checkbox"/></td>
                      <td>Félix Troels</td>
                      <td>20</td>
                      <td>Member</td>
                      <td>
                        <span className="badge badge-warning">Pending</span>
                      </td>
                    </tr>

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
            </div>
          </div>
        </div>
      </div>

    )
  }
}

export default Hosts;
