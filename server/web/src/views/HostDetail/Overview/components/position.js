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

class Position extends Component {

  constructor (props) {
      super(props);
      this.data=[
          ["数据中心", ""],
          ["机架", ""],
          ["槽位", ""],
          ["拥有人", ""],
      ]
  }



  render() {
    return (
                <table className="table table-sm table-bordered">
                  <tbody>
                  <tr>
                      <td>ID</td>
                      <td>20</td>
                  </tr>
                  <tr>
                      <td width="50%">数据中心</td>
                      <td><EditCell>{this.props.data.datacenter}</EditCell></td>
                  </tr>
                  <tr>
                      <td>机架</td>
                      <td><EditCell>{this.props.data.rack}</EditCell></td>
                  </tr>
                  <tr>
                      <td>槽位</td>
                      <td><EditCell>{this.props.data.slot}</EditCell></td>
                  </tr>
                  <tr>
                      <td>拥有人</td>
                      <td><EditCell>{this.props.data.owner}</EditCell></td>
                  </tr>
                  </tbody>
                </table>
    )
  }



}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(Position)
