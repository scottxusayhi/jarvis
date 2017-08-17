import React, { Component } from 'react'
import { connect } from 'react-redux'

import { TabContent, TabPane, Nav, NavItem, NavLink, Card, Button, CardTitle, CardText, Row, Col } from 'reactstrap';
import { InputGroup, InputGroupAddon, Input } from 'reactstrap';
import classnames from 'classnames';

import Registration from "./registration"


import {
    registerHost,
    postRegStart,
    postRegDataSaved
} from '../../../states/actions'

// subscribe
const mapStateToProps = state => {
    return {
        state: state
    }
}

// dispatch actions
const mapDispatchToProps = dispatch => {
    return {
        registerHost: (payload) => {
          dispatch(registerHost(payload))
        },
        postRegStart: (id, initData) => {
            dispatch(postRegStart(id, initData))
        },
        postRegDataSaved: (data) => {
            dispatch(postRegDataSaved(data))
        }
    }
}

function makeRegDataFromDetected(detected) {
    return {
        datacenter: detected.datacenter,
        rack: detected.rack,
        slot: detected.slot,
        tags: detected.tags,
        owner: detected.owner,
        osExpected: detected.osDetected,
        cpuExpected: detected.cpuDetected,
        memExpected: detected.memDetected,
        diskExpected: detected.diskDetected,
        networkExpected: detected.networkDetected
    }
}


class NewHostTab extends Component {
  constructor(props) {
    super(props);

    this.toggle = this.toggle.bind(this);
    this.state = {
      activeTab: '1'
    };

  }

  toggle(tab) {
    if (this.state.activeTab !== tab) {
      this.setState({
        activeTab: tab
      });
    }
  }

  componentDidMount() {
      if (this.props.regType=="postReg") {
        this.props.postRegStart(this.props.state.hostDetail.id, makeRegDataFromDetected(this.props.state.hostDetail.data))
      } else {
          // TODO
      }
  }

  render() {
    return (
      <div>
        <Nav tabs>
          <NavItem>
            <NavLink className={classnames({ active: this.state.activeTab === '1' })}
              onClick={() => { this.toggle('1'); }}>
              表格
            </NavLink>
          </NavItem>
          <NavItem>
            <NavLink
              className={classnames({ active: this.state.activeTab === '2' })}
              onClick={() => { this.toggle('2'); }}>
              JSON
            </NavLink>
          </NavItem>
        </Nav>
        <TabContent activeTab={this.state.activeTab}>
          <TabPane tabId="1">
            <Registration/>
          </TabPane>
          <TabPane tabId="2">
            <Input type="textarea" name="text" id="exampleText" rows="20"/>
          </TabPane>
        </TabContent>
      </div>
    );
  }
}

export default connect(
    mapStateToProps,
    mapDispatchToProps
) (NewHostTab)
