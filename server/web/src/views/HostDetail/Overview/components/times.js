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
                      <td width="50%">注册时间</td>
                      <td width="25%">goldwind</td>
                      <td width="25%">goldwind</td>
                  </tr>
                  <tr>
                      <td>最后一次修改时间</td>
                      <td>01</td>
                      <td>goldwind</td>
                  </tr>
                  <tr>
                      <td>第一次心跳时间</td>
                      <td>010203</td>
                      <td>goldwind</td>
                  </tr>
                  <tr>
                      <td>最后一次心跳时间</td>
                      <td>油油</td>
                      <td>goldwind</td>
                  </tr>
                  </tbody>
                </table>
    )
  }



}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(Times)
