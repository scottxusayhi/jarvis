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
    fetchRegisteredHosts,
    updateRegHost
} from '../../../states/actions'


import { Row, Col } from 'antd';

import { Table, Input, Popconfirm } from 'antd';

import Pager from "../../../components/Pager/Pager";


class EditableCell extends React.Component {
  state = {
    value: this.props.value,
    editable: this.props.editable || false,
  }
  componentWillReceiveProps(nextProps) {
    if (nextProps.editable !== this.state.editable) {
      this.setState({ editable: nextProps.editable });
      if (nextProps.editable) {
        this.cacheValue = this.state.value;
      }
    }
    if (nextProps.status && nextProps.status !== this.props.status) {
      if (nextProps.status === 'save') {
        this.props.onChange(this.state.value);
      } else if (nextProps.status === 'cancel') {
        this.setState({ value: this.cacheValue });
        this.props.onChange(this.cacheValue);
      }
    }
  }
  shouldComponentUpdate(nextProps, nextState) {
    return nextProps.editable !== this.state.editable ||
           nextState.value !== this.state.value;
  }
  handleChange(e) {
    const value = e.target.value;
    console.log(value)
    this.setState({ value });
  }
  render() {
    const { value, editable } = this.state;
    return (
      <div>
        {
          editable ?
            <div>
              <Input
                value={value}
                onChange={e => this.handleChange(e)}
              />
            </div>
            :
            <div className="editable-row-text">
              {value.toString() || ' '}
            </div>
        }
      </div>
    );
  }

  getValue() {
      return this.state.value
  }
}


// rowSelection object indicates the need for row selection
const rowSelection = {
  onChange: (selectedRowKeys, selectedRows) => {
    console.log(`selectedRowKeys: ${selectedRowKeys}`, 'selectedRows: ', selectedRows);
  },
};

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
        },
        updateRegHost: (id, data) => {
            dispatch(updateRegHost(id, data))
        }
    }
}

class RegisteredHosts extends Component {

  constructor (props) {
    super(props);
    this.state = {
        filter: {
            registered: 1
        },
        data: []
    }
    this.datacenterInput = []
      this.rackInput=[]
      this.slotInput=[]
      this.ownerInput=[]
      this.cpuInput=[]
      this.memInput=[]
      this.networkInput=[]
this.columns = [
{
  title: 'ID',
  dataIndex: 'id',
  key: 'id',
},
{
  title: '数据中心',
  dataIndex: 'datacenter',
  key: 'age',
  render: (text, record, index) => this.viewDatacenter(text, record, index),
},
{
  title: '机架',
  dataIndex: 'rack',
  key: 'rack',
    width: '10%',
    render: (text, record, index) => this.viewRack(text, record, index),
},
{
  title: '位置',
  dataIndex: 'slot',
  key: 'slot',
    render: (text, record, index) => this.viewSlot(text, record, index),
},
{
  title: '拥有人',
  dataIndex: 'owner',
  key: 'owner',
    render: (text, record, index) => this.viewOwner(text, record, index),
},
{
  title: '配置审计',
  dataIndex: 'matched',
  key: 'matched',
},
{
  title: '在线状态',
  dataIndex: 'online',
  key: 'online',
},
{
  title: 'VCPU',
  dataIndex: 'cpu',
  key: 'cpu',
    width: '4%',
    render: (text, record, index) => this.viewCpuInfo(text, record, index),
},
{
  title: '内存',
  dataIndex: 'memory',
  key: 'memory',
    width: '8%',
    render: (text, record, index) => this.viewMemInfo(text, record, index),
},
{
  title: '硬盘',
  dataIndex: 'disk',
  key: 'disk',
},
{
  title: '网络',
  dataIndex: 'network',
  key: 'network',
    render: (text, record, index) => this.viewNetworkInfo(text, record, index),
},
{
  title: '备注',
  dataIndex: 'comments',
  key: 'comments',
    width: '4%',
},
    {
        title: '操作',
        dataIndex: '',
        key: 'x',
        render: (text, record, index) => {
            const {editable} = this.state.data[index].datacenter
            return (
                <div className="editable-row-operations">
                    {
                        editable ? <span>
                  <a onClick={() => this.editDone(index, 'save')}>保存</a>
                  <Popconfirm title="确定取消？" onConfirm={() => this.editDone(index, 'cancel')}>
                    <a>取消</a>
                  </Popconfirm>
                </span>
                            :
                            <span>
                  <a onClick={() => this.edit(index)}>编辑</a>
                </span>
                    }
                </div>
            );
        },
    }
];
  }

  componentDidMount() {
      this.props.fetchRegisteredHosts(this.state.filter)
  }

  componentWillReceiveProps(nextProps) {
      console.log("RegisteredHosts will receive props: ", nextProps)
      nextProps.items.data.list && this.setState({data: nextProps.items.data.list.map(host => {
          return {
              key: host.systemId,
              id: this.viewHostId(host.systemId),
              datacenter: {editable: false, value: host.datacenter},
              rack: {editable: false, value: host.rack},
              slot: {editable: false, value: host.slot},
              owner: {editable: false, value: host.owner},
              matched: this.viewConfigAuditStatus(host.connected, host.matched),
              online: this.viewOnlineStatus(host.online),
              cpu: {editable: false, value: host.cpuExpected},
              memory: {editable: false, value: host.memExpected},
              disk: this.viewDiskInfo(host.diskExpected),
              network: {editable: false, value: host.networkExpected},
              registered: host.registered,
          }
      })})
  }

  render() {
    console.log("RegisteredHosts rendering, state.data=", this.state.data);
    return (
      <div>
        <Row>
            <Col>
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
            </Col>

          <Col><Pager pageInfo={this.props.items.data.pageInfo} onPageChange={(page)=>this.props.fetchHosts({registered:1, page: page})}/></Col>
        </Row>

<Table rowSelection={rowSelection} columns={this.columns} dataSource={this.state.data} size="middle"/>

      </div>
    )
  }

  edit(index) {
    const { data } = this.state;
    Object.keys(data[index]).forEach((item) => {
      if (data[index][item] && typeof data[index][item].editable !== 'undefined') {
        data[index][item].editable = true;
      }
    });
    this.setState({ data });
  }
  editDone(index, type) {
      var updateData = {
          datacenter: this.datacenterInput[index].getValue(),
          rack: this.rackInput[index].getValue(),
          slot: this.slotInput[index].getValue(),
          owner: this.ownerInput[index].getValue(),
          cpuExpected: Object.assign({}, this.state.data[index].cpu.value, {
              vcpu: parseInt(this.cpuInput[index].getValue())
          }),
          memExpected: Object.assign({}, this.state.data[index].memory.value, {
             total: parseInt(this.memInput[index].getValue())*1024*1024*1024
          }),
          networkExpected: Object.assign({}, this.state.data[index].network.value, {
              ip: this.networkInput[index].getValue()
          })
      }
      console.log("to update:", updateData)

    const { data } = this.state;
    Object.keys(data[index]).forEach((item) => {
      if (data[index][item] && typeof data[index][item].editable !== 'undefined') {
        data[index][item].editable = false;
        data[index][item].status = type;
      }
    });
    this.setState({ data }, () => {
      Object.keys(data[index]).forEach((item) => {
        if (data[index][item] && typeof data[index][item].editable !== 'undefined') {
          delete data[index][item].status;
        }
      });
    });

    this.props.updateRegHost(this.state.data[index].key, updateData)

  }

  viewHostId(id) {
      var link = "/hosts/" + id
      return <Link to={link}>{id}</Link>
  }

  viewDatacenter(text, record, index) {
      if (record.registered) {
          return (<EditableCell editable={text.editable} value={text.value} ref={(me)=>this.datacenterInput[index]=me} />)
      } else {
          return "N/A"
      }
  }


  viewRack(text, record, index) {
      if (record.registered) {
          return <EditableCell editable={text.editable} value={text.value} ref={(me)=>{this.rackInput[index]=me}} />
      } else {
          return "N/A"
      }
  }


  viewSlot(text, record, index) {
      if (record.registered) {
          return <EditableCell editable={text.editable} value={text.value} ref={(me)=>{this.slotInput[index]=me}} />
      } else {
          return "N/A"
      }
  }

  viewOwner(text, record, index) {
      if (record.registered) {
          return <EditableCell editable={text.editable} value={text.value} ref={(me)=>{this.ownerInput[index]=me}} />
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

  viewCpuInfo(text, record, index) {
      return <EditableCell editable={text.editable} value={text.value.vcpu} ref={(me)=>{this.cpuInput[index]=me}} />
  }

  viewMemInfo(text, record, index) {
      return <div>
          <Row>
          <Col span={12}><EditableCell editable={text.editable} value={Math.ceil(text.value.total/1024/1024/1024)} ref={(me)=>{this.memInput[index]=me}} /></Col>
              <Col span={12}>GB</Col>
          </Row>
          </div>
  }

  viewDiskInfo(diskInfo) {
    return diskInfo.length
  }

  viewNetworkInfo(text, record, index) {
      return <EditableCell editable={text.editable} value={text.value.ip} ref={(me)=>{this.networkInput[index]=me}} />
  }

  viewOsInfo(osInfo) {
      return osInfo.type
  }


}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(RegisteredHosts)
