import React, {Component} from 'react'
import {
    BrowserRouter as Router,
    Route,
    Link
} from 'react-router-dom'
import HostActions from "./HostActions/HostActions"
import {connect} from 'react-redux'
import {
    fetchHosts
} from '../../../states/actions'
import Pager from "../../../components/Pager/Pager";

import {Row, Col} from 'antd';
import {Table, Input, Popconfirm} from 'antd';
import {Pagination} from 'antd'

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

// rowSelection object indicates the need for row selection
const rowSelection = {
    onChange: (selectedRowKeys, selectedRows) => {
        console.log(`selectedRowKeys: ${selectedRowKeys}`, 'selectedRows: ', selectedRows);
    },
};

class ConnectedHosts extends Component {
    constructor(props) {
        super(props);
        this.state = {
            filter: {
                connected: 1
            },
            data: [],
            pagination: {
                showSizeChanger: true,
                defaultPageSize: 20,
                pageSizeOptions: ['20', '50', '100'],
                showTotal: (total, range) => `${range[0]}-${range[1]} of ${total} items`,
                onChange: (page, pageSize) => this.switchPage(page, pageSize),
                onShowSizeChange: (current, size) => this.switchPage(current, size)
            }
        }

        this.columns = [
            {
                title: 'ID',
                dataIndex: 'id',
                key: 'id',
                width: '4%',
                render: (text, record, index) => this.viewHostId(text, record, index)
            },
            {
                title: '数据中心',
                dataIndex: 'datacenter',
                key: 'age',
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
                title: '在线状态',
                dataIndex: 'online',
                key: 'online',
                render: (text, record, index) => this.viewOnlineStatus(text, record, index),
            },
            {
                title: '健康状态',
                dataIndex: 'healthStatus',
                key: 'healthStatus',
                render: (text, record, index) => this.viewHealthStatus(text, record, index),
            },
            {
                title: '注册状态',
                dataIndex: 'registered',
                key: 'registered',
                render: (text, record, index) => this.viewRegisterStatus(text, record, index),
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
                render: (text, record, index) => this.viewDiskInfo(text, record, index),
            },
            {
                title: '网络',
                dataIndex: 'network',
                key: 'network',
                render: (text, record, index) => this.viewNetworkInfo(text, record, index),
            },
            {
                title: 'OS',
                dataIndex: 'os',
                key: 'os',
                render: (text, record, index) => this.viewOsInfo(text, record, index),
            },
            {
                title: '备注',
                dataIndex: 'comments',
                key: 'comments',
                width: '4%',
            }
        ];

    }

    switchPage(page, pageSize) {
        var {filter} = this.state
        filter = Object.assign({}, filter, {
            page: page,
            perPage: pageSize
        })
        this.setState({filter})
        this.props.fetchHosts(filter)
    }

    componentDidMount() {
        this.props.fetchHosts(this.state.filter)
    }

    componentWillReceiveProps(nextProps) {
        console.log("ConnectedHosts will receive props: ", nextProps)
        nextProps.items.data.list && this.setState({
            data: nextProps.items.data.list.map(host => {
                return {
                    key: host.systemId,
                    id: host.systemId,
                    datacenter: host.datacenter,
                    rack: host.rack,
                    slot: host.slot,
                    owner: host.owner,
                    matched: host.matched,
                    online: host.online,
                    healthStatus: host.healthStatus,
                    registered: host.registered,
                    cpu: host.cpuDetected,
                    memory: host.memDetected,
                    disk: host.diskDetected,
                    os: host.osDetected,
                    network: host.networkDetected,
                }
            })
        })

        nextProps.items.data.pageInfo && this.setState({
            pagination: Object.assign({}, this.state.pagination, {
                total: nextProps.items.data.pageInfo.totalSize,
                current: nextProps.items.data.pageInfo.page,
            })
        })

    }

    render() {
        console.log("ConnectedHost rendering");
        return (
            <div>
                <Row>
                    <div className="btn-toolbar mb-3" role="toolbar" aria-label="Toolbar with button groups">
                        <div className="btn-group mr-2" role="group" aria-label="1 group">
                            <button type="button" className="btn btn-secondary"
                                    onClick={() => this.props.fetchHosts(this.state.filter)}><i className="fa fa-refresh"></i></button>
                        </div>
                        <div className="btn-group mr-2" role="group" aria-label="2 group">
                            <HostActions/>
                        </div>
                    </div>
                    <Col/>
                </Row>

                <Table rowSelection={rowSelection} columns={this.columns} dataSource={this.state.data} size="middle"
                       pagination={this.state.pagination}/>
            </div>
        )
    }

    viewHostId(text, record, index) {
        var link = "/hosts/" + text
        return <Link to={link}>{text}</Link>
    }


    viewDatacenter(text, record, index) {
        if (record.registered) {
            return record.datacenter
        } else {
            return "N/A"
        }
    }

    viewRack(text, record, index) {
        if (record.registered) {
            return record.rack
        } else {
            return "N/A"
        }
    }

    viewSlot(text, record, index) {
        if (record.registered) {
            return record.slot
        } else {
            return "N/A"
        }
    }

    viewPosition(text, record, index) {
        if (record.registered) {
            return record.rack + "-" + record.slot
        } else {
            return "N/A"
        }
    }

    viewOnlineStatus(text, record, index) {
        if (record.online) {
            return <span className="badge badge-success">在线</span>
        } else {
            return <span className="badge badge-danger">离线</span>
        }
    }

    viewHealthStatus(text, record, index) {
        switch (record.healthStatus) {
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

    viewRegisterStatus(text, record, index) {
        if (record.registered) {
            return <span className="badge badge-success">已注册</span>
        } else {
            return <span className="badge badge-info">未注册</span>
        }
    }

    viewCpuInfo(text, record, index) {
        return text.vcpu
    }

    viewMemInfo(text, record, index) {
        return Math.ceil(text.total / 1024 / 1024 / 1024) + " GB"
    }

    viewDiskInfo(text, record, index) {
        return text.length
    }

    viewNetworkInfo(text, record, index) {
        return text.ip
    }

    viewOsInfo(text, record, index) {
        return text.type + "-" + text.dist + "-" + text.version + "-" + text.arch
    }


}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(ConnectedHosts)
