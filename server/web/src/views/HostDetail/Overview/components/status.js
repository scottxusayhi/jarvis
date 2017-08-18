import React, { Component } from 'react'
import { connect } from 'react-redux'
import { TabContent, TabPane, Nav, NavItem, NavLink, Card, Button, CardTitle, CardText, Row, Col } from 'reactstrap';
import classnames from 'classnames';

import HotTable from 'react-handsontable';
import Collapsible from 'react-collapsible';

import NewHost from '../../../Hosts/NewHost/NewHost'

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

class Status extends Component {

  constructor (props) {
      super(props);
  }

  registerBtn() {
      return <Col><NewHost btnColor="link" btnText="少侠请注册" regType="postReg"/></Col>
  }


  render() {
      console.log("status=======", this.props.data)
    return (
        <Collapsible trigger="状态" open={true} transitionTime={200}>
                <table className="table table-sm table-bordered">
                    <thead>
                    <tr>
                        <th width="20%">已注册</th>
                        <th width="20%">连接过</th>
                        <th width="20%">配置审计</th>
                        <th width="20%">在线</th>
                        <th width="20%">健康</th>
                    </tr>
                  </thead>
                  <tbody>
                  <tr>
                      <td>
                          <Row>
                              <Col>{viewBool(this.props.data.registered)}</Col>
                              {!this.props.data.registered && this.registerBtn()}
                          </Row>
                      </td>
                      <td>{viewBool(this.props.data.connected)}</td>
                      <td>{viewBool(this.props.data.matched)}</td>
                      <td>{viewBool(this.props.data.online)}</td>
                      <td>{this.props.data.healthStatus}</td>
                  </tr>
                  </tbody>
                </table>
        </Collapsible>
    )
  }

}

function viewBool(b) {
    if (b) {
        return "YES"
    } else {
        return "NO"
    }
}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(Status)
