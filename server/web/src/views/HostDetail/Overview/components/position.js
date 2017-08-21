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
                      <td>{this.props.data.systemId}</td>
                  </tr>
                  <tr>
                      <td width="30%">数据中心</td>
                      <td>{this.props.data.registered && <EditCell ref={(me)=>this.refDatacenter=me} onEnter={()=>this.updateDatacenter()}>{this.props.data.datacenter}</EditCell> || "-"}</td>
                  </tr>
                  <tr>
                      <td>机架</td>
                      <td>{this.props.data.registered && <EditCell ref={(me)=>this.refRack=me} onEnter={()=>this.updateRack()}>{this.props.data.rack}</EditCell> || "-"}</td>
                  </tr>
                  <tr>
                      <td>槽位</td>
                      <td>{this.props.data.registered && <EditCell ref={(me)=>this.refSlot=me} onEnter={()=>this.updateSlot()}>{this.props.data.slot}</EditCell> || "-"}</td>
                  </tr>
                  <tr>
                      <td>拥有人</td>
                      <td>{this.props.data.registered && <EditCell ref={(me)=>this.refOwner=me} onEnter={()=>this.updateOwner()}>{this.props.data.owner}</EditCell> || "-"}</td>
                  </tr>
                  </tbody>
                </table>
    )
  }

  updateDatacenter() {
      var data = {
          datacenter: this.refDatacenter.getWrappedInstance().getInput()
      }
      this.props.updateRegHost(this.props.data.systemId, data)
  }

  updateRack() {
      var data = {
          rack: this.refRack.getWrappedInstance().getInput()
      }
      this.props.updateRegHost(this.props.data.systemId, data)
  }

  updateSlot() {
      var data = {
          slot: this.refSlot.getWrappedInstance().getInput()
      }
      this.props.updateRegHost(this.props.data.systemId, data)
  }

  updateOwner() {
      var data = {
          owner: this.refOwner.getWrappedInstance().getInput()
      }
      this.props.updateRegHost(this.props.data.systemId, data)
  }


}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(Position)
