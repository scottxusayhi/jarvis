import React, { Component } from 'react'
import { connect } from 'react-redux'
import { TabContent, TabPane, Nav, NavItem, NavLink, Card, Button, CardTitle, CardText, Row, Col } from 'reactstrap';
import classnames from 'classnames';

import HotTable from 'react-handsontable';

import Collapsible from 'react-collapsible';
import EditCell from './editcell'

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

class CpuInfo extends Component {

  constructor (props) {
      super(props);
  }



  render() {
    return (
        <Collapsible trigger="配置：CPU" open={true} transitionTime={200}>
                <table className="table table-sm table-bordered table-responsive">
                    <thead>
                    <tr>
                        <th width="20%"></th>
                        <th width="40%">注册信息</th>
                        <th width="40%">检测信息</th>
                    </tr>
                    </thead>
                  <tbody>
                  <tr>
                      <td>Socket(s)</td>
                      <td>{this.props.data.cpuExpected && this.props.data.registered && (<EditCell ref={(me)=>this.refCpuSocket=me} onEnter={()=>this.updateCpuSocket()}>{this.props.data.cpuExpected.socket}</EditCell>) || "-"}</td>
                      <td>{this.props.data.cpuDetected && this.props.data.cpuDetected.socket}</td>
                  </tr>
                  <tr>
                      <td>VCPU</td>
                      <td>{this.props.data.cpuExpected && this.props.data.registered && <EditCell ref={(me)=>this.refCpuVcpu=me} onEnter={()=>this.updateCpuVcpu()}>{this.props.data.cpuExpected.vcpu}</EditCell> || "-"}</td>
                      <td>{this.props.data.cpuDetected && this.props.data.cpuDetected.vcpu}</td>
                  </tr>
                  <tr>
                      <td>Model</td>
                      <td>{this.props.data.cpuExpected && this.props.data.registered && <EditCell ref={(me)=>this.refCpuModel=me} onEnter={()=>this.updateCpuModel()}>{this.props.data.cpuExpected.model}</EditCell> || "-"}</td>
                      <td>{this.props.data.cpuDetected && this.props.data.cpuDetected.model}</td>
                  </tr>
                  </tbody>
                </table>
        </Collapsible>
    )
  }

  updateCpuSocket() {
      var data = {
          cpuExpected: {
              socket: Number(this.refCpuSocket.getWrappedInstance().getInput()),
              vcpu: this.props.data.cpuExpected.vcpu,
              model: this.props.data.cpuExpected.model
          }
      }
      this.props.updateRegHost(this.props.data.systemId, data)
  }

  updateCpuVcpu() {
      var data = {
          cpuExpected: {
              socket: this.props.data.cpuExpected.socket,
              vcpu: Number(this.refCpuVcpu.getWrappedInstance().getInput()),
              model: this.props.data.cpuExpected.model
          }
      }
      this.props.updateRegHost(this.props.data.systemId, data)
  }

  updateCpuModel() {
      var data = {
          cpuExpected: {
              socket: this.props.data.cpuExpected.socket,
              vcpu: this.props.data.cpuExpected.vcpu,
              model: Number(this.refCpuModel.getWrappedInstance().getInput()),
          }
      }
      this.props.updateRegHost(this.props.data.systemId, data)
  }

}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(CpuInfo)
