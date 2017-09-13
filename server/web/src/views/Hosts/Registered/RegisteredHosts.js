import React, {Component} from 'react'
import {
    BrowserRouter as Router,
    Route,
    Link
} from 'react-router-dom'

import NewHostPopup from '../NewHost/NewHost'
import HostActions from "./HostActions/HostActions"
import {connect} from 'react-redux'
import {
    fetchHosts,
    fetchRegisteredHosts,
    updateRegHost,
    listItems,
} from '../../../states/actions'

import {Button} from 'antd'
import {Row, Col} from 'antd';
import {Table, Input, Popconfirm} from 'antd';
import {Select} from 'antd'

import EditableCell from './editablecell'


const ButtonGroup = Button.Group;
const Option = Select.Option;

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
        list: state.list,
    }
}

// dispatch actions
const mapDispatchToProps = dispatch => {
    return {
        fetchRegisteredHosts: query => {
            dispatch(fetchRegisteredHosts(query))
            dispatch(listItems())
        },
        updateRegHost: (id, data) => {
            dispatch(updateRegHost(id, data))
        },
        listItems: ()=> {
            dispatch(listItems())
        }
    }
}

class RegisteredHosts extends Component {

    constructor(props) {
        super(props);
        this.state = {
            filter: {
                registered: 1,
            },
            pagination: {
                showSizeChanger: true,
                defaultPageSize: 20,
                pageSizeOptions: ['20', '50', '100'],
                pageSize: 20,
                current: 1,
                showTotal: (total, range) => `${range[0]}-${range[1]} of ${total} items`,
            },
            sorter: {

            },
            data: [],
            list: {},
            tags: [],
        }
        this.datacenterInput = []
        this.rackInput = []
        this.slotInput = []
        this.ownerInput = []
        this.cpuInput = []
        this.memInput = []
        this.networkInput = []
        this.columns = [
            {
                title: 'ID',
                dataIndex: 'id',
                key: 'id',
                width: '4%'
            },
            {
                title: '数据中心',
                dataIndex: 'datacenter',
                key: 'datacenter',
                filters: [],
                width: '10%',
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
                width: '10%',
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
                filters: [{text: "匹配", value: 1}, {text: "不匹配", value: 0}]
            },
            {
                title: '在线状态',
                dataIndex: 'online',
                key: 'online',
                filters: [{text: "在线", value: 1}, {text: "离线", value: 0}]
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
                key: 'network-key',
                sorter: true,
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
                  <Popconfirm title="确定取消？" onConfirm={() => this.editDone(index, 'cancel')}>
                    <a>取消</a>
                  </Popconfirm>
                  <a onClick={() => this.editDone(index, 'save')}>保存</a>
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

    makeApiQuery(pagination, filter, sorter) {
        // from pagination
        var result = {
            registered: 1,
            page: pagination.current,
            perPage: pagination.pageSize,
        }

        // iterate all filter fields
        for (var key in filter) {
            if (filter.hasOwnProperty(key)) {
                var value=filter[key]
                if (Array.isArray(value) && value.length>0) {
                    result[key] = value.join(",")
                }
            }
        }

        // from sorter
        if (sorter.hasOwnProperty('field')) {
            var field = sorter['field']
            if (field==='network') {
                field = 'INET_ATON(networkExpected->>\'$.ip\')'
            }
            if (sorter.hasOwnProperty('order')) {
                var order = sorter['order']
                if (order==='descend') {
                    order = '-'
                } else {
                    order = '+'
                }
            }

            result['order'] = order+field
        }

        return result
    }

    handleTableChange(pagination, filter, sorter) {
        console.log(pagination)
        console.log(filter)
        console.log(sorter)
        this.props.fetchRegisteredHosts(this.makeApiQuery(pagination, filter, sorter))
    }

    componentDidMount() {
        this.props.fetchRegisteredHosts(this.makeApiQuery(this.state.pagination, this.state.filter, this.state.sorter))
    }

    componentWillReceiveProps(nextProps) {
        console.log("RegisteredHosts will receive props: ", nextProps)
        // update host list
        nextProps.items.data.list && this.setState({
            data: nextProps.items.data.list.map(host => {
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
            })
        })
        // update page info
        nextProps.items.data.pageInfo && this.setState({
            pagination: Object.assign({}, this.state.pagination, {
                total: nextProps.items.data.pageInfo.totalSize,
                current: nextProps.items.data.pageInfo.page,
                pageSize: nextProps.items.data.pageInfo.perPage,
            })
        })
        // update lists for dynamic filters
        if (nextProps.list.data.datacenters) {
            this.setState({
                list: nextProps.list.data
            })
            // filter in column headers
            for (var index in this.columns) {
                var column = this.columns[index]
                if (column.hasOwnProperty('key') && column['key']==='datacenter') {
                    column['filters'] = nextProps.list.data.datacenters.map((dc)=>{return {text: dc, value: dc}})
                }
                if (column.hasOwnProperty('key') && column['key']==='rack') {
                    column['filters'] = nextProps.list.data.racks.map((rack)=>{return {text: rack, value: rack}})
                }
                if (column.hasOwnProperty('key') && column['key']==='owner') {
                    column['filters'] = nextProps.list.data.owners.map((owner)=>{return {text: owner, value: owner}})
                }
            }
            // tag filter
            var tags = nextProps.list.data.tags
            var tagOptions = []
            for (var tag in tags) {
                var count = tags[tag].length
                tagOptions.push(<Option key={tag}>{tag} ({count})</Option>)
            }
            this.setState({
                tags: tagOptions
            })
        }

    }

    render() {
        console.log("RegisteredHosts rendering, state=", this.state);
        return (
            <div>
                <ButtonGroup>
                    <Button onClick={() => this.props.fetchRegisteredHosts(this.makeApiQuery(this.state.pagination, this.state.filter, this.state.sorter))}><i className="fa fa-refresh"></i></Button>
                </ButtonGroup>
                <ButtonGroup>
                    <NewHostPopup/>
                </ButtonGroup>
                <Row>
                        {/*<HostActions/>*/}
                    <Col span={10}>
                        <Select mode="tags" style={{ width: '100%' }} placeholder="按 Tag 过滤" onChange={(value)=>{console.log(value)}}>{this.state.tags}</Select>
                    </Col>
                </Row>

                <Table rowSelection={rowSelection} columns={this.columns} dataSource={this.state.data} size="middle"
                       pagination={this.state.pagination} onChange={(pagination, filter, sorter)=>this.handleTableChange(pagination, filter, sorter)}/>

            </div>
        )
    }

    edit(index) {
        const {data} = this.state;
        Object.keys(data[index]).forEach((item) => {
            if (data[index][item] && typeof data[index][item].editable !== 'undefined') {
                data[index][item].editable = true;
            }
        });
        this.setState({data});
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
                total: parseInt(this.memInput[index].getValue()) * 1024 * 1024 * 1024
            }),
            networkExpected: Object.assign({}, this.state.data[index].network.value, {
                ip: this.networkInput[index].getValue()
            })
        }
        console.log("to update:", updateData)

        const {data} = this.state;
        Object.keys(data[index]).forEach((item) => {
            if (data[index][item] && typeof data[index][item].editable !== 'undefined') {
                data[index][item].editable = false;
                data[index][item].status = type;
            }
        });
        this.setState({data}, () => {
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
            return (<EditableCell editable={text.editable} value={text.value}
                                  ref={(me) => this.datacenterInput[index] = me}/>)
        } else {
            return "N/A"
        }
    }


    viewRack(text, record, index) {
        if (record.registered) {
            return <EditableCell editable={text.editable} value={text.value} ref={(me) => {
                this.rackInput[index] = me
            }}/>
        } else {
            return "N/A"
        }
    }


    viewSlot(text, record, index) {
        if (record.registered) {
            return <EditableCell editable={text.editable} value={text.value} ref={(me) => {
                this.slotInput[index] = me
            }}/>
        } else {
            return "N/A"
        }
    }

    viewOwner(text, record, index) {
        if (record.registered) {
            return <EditableCell editable={text.editable} value={text.value} ref={(me) => {
                this.ownerInput[index] = me
            }}/>
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
        return <EditableCell editable={text.editable} value={text.value.vcpu} ref={(me) => {
            this.cpuInput[index] = me
        }}/>
    }

    viewMemInfo(text, record, index) {
        return <div>
            <Row type="flex" justify="space-around" align="middle">
                <Col span={12}><EditableCell editable={text.editable}
                                             value={Math.ceil(text.value.total / 1024 / 1024 / 1024)} ref={(me) => {
                    this.memInput[index] = me
                }}/></Col>
                <Col span={12}>GB</Col>
            </Row>
        </div>
    }

    viewDiskInfo(diskInfo) {
        return diskInfo.length
    }

    viewNetworkInfo(text, record, index) {
        return <EditableCell editable={text.editable} value={text.value.ip} ref={(me) => {
            this.networkInput[index] = me
        }}/>
    }

    viewOsInfo(osInfo) {
        return osInfo.type
    }


}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(RegisteredHosts)
