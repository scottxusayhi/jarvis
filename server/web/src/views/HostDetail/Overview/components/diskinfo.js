import React, { Component } from 'react'
import { connect } from 'react-redux'
import { TabContent, TabPane, Nav, NavItem, NavLink, Card, Button, CardTitle, CardText, Row, Col } from 'reactstrap';
import classnames from 'classnames';
import EditCell from './editcell'
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

class DiskInfo extends Component {

  constructor (props) {
      super(props);
      this.refDiskDevice = []
      this.refDiskCap = []
  }

  diskIndex() {
      var num = 0
      if (this.props.data.diskExpected) {
          var a = this.props.data.diskExpected.length
          var b = this.props.data.diskDetected.length
          if (a>b) {num=a} else {num=b}
      }
      return Array.apply(null, {length: num})
  }


  render() {
    return (
        <Collapsible trigger="配置：磁盘" open={true} transitionTime={200}>
                <table className="table table-sm table-bordered">
                    <thead>
                    <tr>
                        <th width="10%"></th>
                        <th width="10%"></th>
                        <th width="40%">注册信息</th>
                        <th width="40%">检测信息</th>
                    </tr>
                    </thead>                    
                  {/*<tbody>*/}

                  {this.diskIndex().map((o, index)=>{
                      return (
                          <tbody>
                          <tr>
                              <td>disk-{index}</td>
                              <td>设备</td>
                              <td>{this.props.data.diskExpected[index] && this.props.data.registered && <EditCell ref={(me)=>this.refDiskDevice[index]=me} onEnter={()=>this.updateDiskDevice(index)}>{this.props.data.diskExpected[index].device}</EditCell> || "-"}</td>
                              <td>{this.props.data.diskDetected[index] && this.props.data.diskDetected[index].device}</td>
                          </tr>
                          <tr>
                              <td></td>
                              <td>型号</td>
                              <td>-</td>
                              <td>{this.props.data.diskDetected[index] && this.props.data.diskDetected[index].model}</td>
                          </tr>
                          <tr>
                              <td></td>
                              <td>容量</td>
                              <td>{this.props.data.diskExpected[index] && this.props.data.registered && <EditCell ref={(me)=>this.refDiskCap[index]=me} onEnter={()=>this.updateDiskCap(index)}>{this.props.data.diskExpected[index].capacity}</EditCell> || "-"}</td>
                              <td>{this.props.data.diskDetected[index] && this.props.data.diskDetected[index].capacity}</td>
                          </tr>
                          <tr>
                              <td></td>
                              <td>已使用</td>
                              <td>-</td>
                              <td>{this.props.data.diskDetected[index] && this.props.data.diskDetected[index].used}</td>
                          </tr>
                          </tbody>
                      )
                  })}

                </table>
        </Collapsible>
    )
  }


  updateDiskDevice(index) {
      var data = {
          diskExpected: this.props.data.diskExpected
      }
      data.diskExpected[index].device = this.refDiskDevice[index].getWrappedInstance().getInput()
      this.props.updateRegHost(this.props.data.systemId, data)
  }

  updateDiskCap(index) {
      var data = {
          diskExpected: this.props.data.diskExpected
      }
      data.diskExpected[index].capability = this.refDiskCap[index].getWrappedInstance().getInput()
      this.props.updateRegHost(this.props.data.systemId, data)
  }

}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(DiskInfo)
