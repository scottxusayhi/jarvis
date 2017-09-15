import React, {Component} from 'react'
import {connect} from 'react-redux'
import Overview from './Overview'
import Comments from './Comments'
import UselessComponent from './UselessComponent'
import {
    fetchHostDetail
} from '../../states/actions'
import { Tabs } from 'antd';
const TabPane = Tabs.TabPane;

// subscribe state
const mapStateToProps = state => {
    return {
        state: state
    }
}

// dispatch actions
const mapDispatchToProps = dispatch => {
    return {
        fetchHostDetail: (id) => {
            dispatch(fetchHostDetail(id))
        }
    }
}

class HostDetail extends Component {

    tabChange(key) {
        console.log("will go to " + key)
    }

    componentDidMount() {
        this.props.fetchHostDetail(this.props.match.params.hostId)
    }

    render() {
        console.log("rendering" + this.props.match.params.hostId);
        var hostId = this.props.match.params.hostId
        return (
            <Tabs onChange={this.tabChange} defaultActiveKey="1">
                <TabPane tab="总览" key="1"><Overview {...this.props}/></TabPane>
                <TabPane tab="备注" key="2"><Comments hostId={hostId} comments="temp comments"/></TabPane>
                <TabPane tab="其它" key="3"><UselessComponent a="b"/></TabPane>
            </Tabs>
        )
    }


}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(HostDetail)
