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
  }

  getRegData() {
      if (this.props.regType==="postReg") {
          return this.props.state.regHost.postRegData
      }
      if (this.props.regType==="newReg") {
          return this.props.state.regHost.newRegData
      }
      return "did not load"
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
            <Input
                type="textarea"
                name="text"
                id="exampleText"
                rows="20"
                value={JSON.stringify(this.getRegData(), null, 2)}
            />
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
