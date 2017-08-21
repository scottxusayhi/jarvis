import React, { Component } from 'react'
import { connect } from 'react-redux'
import { TabContent, TabPane, Nav, NavItem, NavLink, Card, Button, CardTitle, CardText, Row, Col } from 'reactstrap';
import classnames from 'classnames';

import HotTable from 'react-handsontable';
import Collapsible from 'react-collapsible';

import {
    updateRegHost
} from '../../../../states/actions'

// subscribe
const mapStateToProps = state => {
    return {
        data: state.hostDetail.data
    }
}

// dispatch actions
const mapDispatchToProps = dispatch => {
    return {
        updateRegHost: (id, data) => {
            dispatch(updateRegHost(id, data))
        }
    }
}

class Times extends Component {

  constructor (props) {
      super(props);
  }



  render() {
    return (
                <table className="table table-sm table-bordered">
                  <tbody>
                  <tr>
                      <td width="30%">注册时间</td>
                      <td width="35%">{this.props.data && this.props.data.createdAt}</td>
                      <td width="35%">{this.props.data && viewTimeDiff(new Date()-new Date(this.props.data.createdAt))}</td>
                  </tr>
                  <tr>
                      <td>最后一次修改时间</td>
                      <td>{this.props.data && this.props.data.updatedAt}</td>
                      <td>{this.props.data && viewTimeDiff(new Date()-new Date(this.props.data.updatedAt))}</td>
                  </tr>
                  <tr>
                      <td>第一次心跳时间</td>
                      <td>{this.props.data && this.props.data.firstSeenAt}</td>
                      <td>{this.props.data && viewTimeDiff(new Date()-new Date(this.props.data.firstSeenAt))}</td>
                  </tr>
                  <tr>
                      <td>最后一次心跳时间</td>
                      <td>{this.props.data && this.props.data.lastSeenAt}</td>
                      <td>{this.props.data && viewTimeDiff(new Date()-new Date(this.props.data.lastSeenAt))}</td>
                  </tr>
                  </tbody>
                </table>
    )
  }

}


function viewTimeDiff(diff) {
    var msec = diff;
    var dd = Math.floor(msec / 1000 / 60 / 60 / 24)
    if (dd<0) {
        dd=0
    }
    msec -= dd * 1000 * 60 * 60 * 24
    var hh = Math.floor(msec / 1000 / 60 / 60);
    if (hh<0) {
        hh=0
    }
    msec -= hh * 1000 * 60 * 60;
    var mm = Math.floor(msec / 1000 / 60);
    if (mm<0) {
        mm=0
    }
    msec -= mm * 1000 * 60;
    var ss = Math.floor(msec / 1000);
    if (ss<0) {
        ss=0
    }
    msec -= ss * 1000;
    return dd+"天 " + hh+"时 "+mm+"分 " + ss + "秒 前"
}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(Times)
