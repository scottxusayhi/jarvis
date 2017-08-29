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

class Image extends Component {

  constructor (props) {
      super(props);
  }



  render() {
    return (
    <div>
        <img src="img/default-host.png" className="rounded mx-auto d-block" alt="image not found :("/>
    </div>


    )
  }
}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(Image)
