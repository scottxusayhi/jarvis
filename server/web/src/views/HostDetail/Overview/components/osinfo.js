import React, { Component } from 'react'
import { connect } from 'react-redux'
import { TabContent, TabPane, Nav, NavItem, NavLink, Card, Button, CardTitle, CardText, Row, Col } from 'reactstrap';
import classnames from 'classnames';

import HotTable from 'react-handsontable';

import Collapsible from 'react-collapsible';
import EditCell from "./editcell";

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

class OsInfo extends Component {

  constructor (props) {
      super(props);
  }


  onCellClick() {
      console.log("cell clicked")
  }

  render() {
    return (
            <Collapsible trigger="配置：OS" open={true} transitionTime={200}>

                <table className="table table-sm table-bordered">
                    <thead>
                    <tr>
                        <th width="20%"></th>
                        <th width="40%">注册信息</th>
                        <th width="40%">检测信息</th>
                    </tr>
                    </thead>
                  <tbody>
                  <tr>
                      <td>Type</td>
                      <td><EditCell>{this.props.data.osExpected && this.props.data.osExpected.type}</EditCell></td>
                      <td>{this.props.data.osDetected && this.props.data.osDetected.type}</td>
                  </tr>
                  <tr>
                      <td>Arch</td>
                      <td><EditCell>{this.props.data.osExpected && this.props.data.osExpected.arch}</EditCell></td>
                      <td>{this.props.data.osDetected && this.props.data.osDetected.arch}</td>
                  </tr>
                  <tr>
                      <td>Hostname</td>
                      <td><EditCell>{this.props.data.osExpected && this.props.data.osExpected.hostname}</EditCell></td>
                      <td>{this.props.data.osDetected && this.props.data.osDetected.hostname}</td>
                  </tr>
                  <tr>
                      <td>Distribution</td>
                      <td>-</td>
                      <td>{this.props.data.osDetected && this.props.data.osDetected.dist}</td>
                  </tr>
                  <tr>
                      <td>Uptime</td>
                      <td>-</td>
                      <td>{this.props.data.osDetected && this.props.data.osDetected.uptime}</td>
                  </tr>
                  </tbody>
                </table>
            </Collapsible>
    )
  }



}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(OsInfo)
