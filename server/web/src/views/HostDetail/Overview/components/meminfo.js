import React, { Component } from 'react'
import { connect } from 'react-redux'

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

class MemInfo extends Component {

  constructor (props) {
      super(props);
  }



  render() {
      console.log(this.props)
    return (
        <Collapsible trigger="配置：内存" open={true} transitionTime={200}>
                <table className="table table-sm table-bordered">
                    <thead>
                    <tr>
                        <th width="20%"></th>
                        <th width="40%">注册信息</th>
                        <th width="40%">检测信息</th>
                    </tr>
                    </thead>                    
                  <tbody>
                  <tr>
                      <td>Total</td>
                      <td>{this.props.data.memExpected && this.props.data.registered && <EditCell ref={(me)=>this.refMemTotal=me} onEnter={()=>this.updateMemTotal()}>{this.props.data.memExpected.total}</EditCell> || "-"}</td>
                      <td>{this.props.data.memDetected && this.props.data.memDetected.total}</td>
                  </tr>
                  <tr>
                      <td>Available</td>
                      <td>-</td>
                      <td>16 GB</td>
                  </tr>
                  <tr>
                      <td>Used</td>
                      <td>-</td>
                      <td>16 GB</td>
                  </tr>
                  </tbody>
                </table>
        </Collapsible>
    )
  }

  updateMemTotal() {
      var data = {
          memExpected: {
              total: Number(this.refMemTotal.getWrappedInstance().getInput()),
          }
      }
      this.props.updateRegHost(this.props.data.systemId, data)
  }


}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(MemInfo)
