import React, { Component } from 'react'
import { connect } from 'react-redux'
import { TabContent, TabPane, Nav, NavItem, NavLink, Card, Button, CardTitle, CardText, Row, Col } from 'reactstrap';
import classnames from 'classnames';
import Overview from './Overview'
import Comments from './Comments'

import {
  fetchHostDetail
} from '../../states/actions'


// subscribe state
const mapStateToProps = state => {
  return {}
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

  constructor (props) {
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
        this.props.fetchHostDetail(this.props.match.params.hostId)
    }

    componentWillUnmount() {
      console.log("host detail (parent) page will unmount")
    }

    componentDidUpdate() {
      console.log("host detail (parent) page did update")
    }

    render() {
    console.log("rendering" + this.props.match.params.hostId);
    return (
      <div>
        <Nav tabs>
          <NavItem>
            <NavLink
              className={classnames({ active: this.state.activeTab === '1' })}
              onClick={() => { this.toggle('1'); }}>
              总览
            </NavLink>
          </NavItem>
          <NavItem>
            <NavLink
              className={classnames({ active: this.state.activeTab === '2' })}
              onClick={() => { this.toggle('2'); }}>
              备注
            </NavLink>
          </NavItem>
          <NavItem>
            <NavLink
              className={classnames({ active: this.state.activeTab === '3' })}
              onClick={() => { this.toggle('3'); }}>
              其它
            </NavLink>
          </NavItem>
        </Nav>
        <TabContent activeTab={this.state.activeTab}>
          <TabPane tabId="1">
              <Overview {...this.props}/>
          </TabPane>
          <TabPane tabId="2">
                <Comments/>
          </TabPane>
          <TabPane tabId="3">
                others
          </TabPane>
        </TabContent>
      </div>
    )
  }



}

export default connect(
    mapStateToProps,
    mapDispatchToProps
) (HostDetail)
