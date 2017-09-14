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
import {Tag} from 'antd'

import EditableCell from './editablecell'
import CommentEditor from "../commenteditor";


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
        fetchHosts: query => {
            dispatch(fetchHosts(query))
            dispatch(listItems())
        },
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
            // the final api query includes 5 parts: fixedFilter, pagination, sorter, fieldFilter, and tagFilter
            fixedFilter: {
                registered: 1,
            },
            pageFilter: {
                page: 1,
                perPage: 20,
            },
            sortFilter: {

            },
            fieldFilter: {

            },
            tagFilter: {

            },

            // the following states are used to control the view
            pagination: {
                showSizeChanger: true,
                defaultPageSize: 20,
                pageSizeOptions: ['20', '50', '100'],
                pageSize: 20,
                current: 1,
                showTotal: (total, range) => `${range[0]}-${range[1]} of ${total} items`,
            },
            sorter: {},
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
                render: (text, record, index) => this.viewHostId(text, record, index),
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
            // {
            //     title: '健康状态',
            //     dataIndex: 'healthStatus',
            //     key: 'healthStatus',
            //     render: (text, record, index) => this.viewHealthStatus(text, record, index),
            // },
            // {
            //     title: '注册状态',
            //     dataIndex: 'registered',
            //     key: 'registered',
            //     render: (text, record, index) => this.viewRegisterStatus(text, record, index),
            // },
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
            // {
            //     title: 'OS',
            //     dataIndex: 'os',
            //     key: 'os',
            //     render: (text, record, index) => this.viewOsInfo(text, record, index),
            // },
            {
                title: '备注',
                dataIndex: 'comments',
                key: 'comments',
                width: '4%',
                render: (text, record, index) => this.viewComments(text, record, index)
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

    makeQueryOnTableChange(pagination, filter, sorter) {
        // pageFilter
        var pageFilter = {
            page: pagination.current,
            perPage: pagination.pageSize,
        }
        console.log(pageFilter)

        // fieldFilter
        var fieldFilter = {}
        for (var key in filter) {
            if (filter.hasOwnProperty(key)) {
                var value=filter[key]
                if (Array.isArray(value) && value.length>0) {
                    fieldFilter[key] = value.join(",")
                }
            }
        }
        console.log(fieldFilter)

        // sortFilter
        var sortFilter = {}
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

            sortFilter['order'] = order+field
        }
        console.log(sortFilter)

        this.setState({
            pageFilter: pageFilter,
            fieldFilter: fieldFilter,
            sortFilter: sortFilter,
        })

        return {
            ...this.state.fixedFilter,
            ...this.state.tagFilter,
            ...pageFilter,
            ...fieldFilter,
            ...sortFilter,

        }
    }

    makeQueryOnTagFilterChange(selectedTags) {
        console.log("selectedTags: " + selectedTags)
        var tagFilter
        if (selectedTags.length===0) {
            tagFilter = {}
        } else {
            var tagFilter = {
                tags: selectedTags.join(",")
            }
        }

        this.setState({
            tagFilter: tagFilter
        })
        return {
            ...this.state.fixedFilter,
            ...tagFilter,
            ...this.state.pageFilter,
            ...this.state.fieldFilter,
            ...this.state.sortFilter,
        }
    }

    makeQueryNow() {
        return {
            ...this.state.fixedFilter,
            ...this.state.tagFilter,
            ...this.state.pageFilter,
            ...this.state.fieldFilter,
            ...this.state.sortFilter
        }
    }

    handleTableChange(pagination, filter, sorter) {
        console.log(pagination)
        console.log(filter)
        console.log(sorter)
        this.props.fetchRegisteredHosts(this.makeQueryOnTableChange(pagination, filter, sorter))
    }

    handleTagFilterChange(selectedTags) {
        console.log(selectedTags)
        this.props.fetchRegisteredHosts(this.makeQueryOnTagFilterChange(selectedTags))
    }

    componentDidMount() {
        this.props.fetchRegisteredHosts(this.makeQueryNow())
    }

    componentWillReceiveProps(nextProps) {
        console.log("RegisteredHosts will receive props: ", nextProps)
        // update host list
        nextProps.items.data.list && this.setState({
            data: nextProps.items.data.list.map(host => {
                return {
                    key: host.systemId,
                    id: host.systemId,
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
                    tags: host.tags,
                    comments: host.comments,
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
                <Row>
                    <Col span={1}>
                        <Button onClick={() => this.props.fetchRegisteredHosts(this.makeQueryNow())}><i className="fa fa-refresh"></i></Button>
                    </Col>
                    <Col span={2}>
                        <NewHostPopup/>
                    </Col>
                    <Col span={8}>
                        <Select mode="multiple" style={{ width: '100%' }} placeholder="按 Tag 过滤" onChange={(value)=>{this.props.fetchRegisteredHosts(this.makeQueryOnTagFilterChange(value))}}>{this.state.tags}</Select>
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

    viewHostId(text, record, index) {
        var link = "/hosts/" + record.id

        return (
            <div>
                <Row><Link to={link}>{record.id}</Link></Row>
                <Row>
                    {
                        record.tags.map((tag)=> {
                            return <Tag key={tag} color="blue">{tag}</Tag>
                        })
                    }
                </Row>
            </div>
        )
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

    viewComments(text, record, index) {
        var commentLine
        if (record.comments.length===0) {
            commentLine = 0
        } else {
            commentLine = record.comments.split(/\r\n|\r|\n/).length
        }
        return <CommentEditor host={record.id} comments={record.comments}>{commentLine}</CommentEditor>
        // return <Tag color="cyan" value={commentLine}>{commentLine}</Tag>

    }

    viewOsInfo(text, record, index) {
        return text.type + "-" + text.dist + "-" + text.version + "-" + text.arch
    }



}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(RegisteredHosts)
