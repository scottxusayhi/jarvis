import React, { Component } from 'react'
import PropTypes from 'prop-types';
import { connect } from 'react-redux'
import { Button, Modal, ModalHeader, ModalBody, ModalFooter } from 'reactstrap';

import {
    registerHost
} from '../../../states/actions'

// subscribe
const mapStateToProps = state => {
    if (state.regHost.type=="newReg") {
        return {
            data: state.regHost.newRegData
        }
    }
    if (state.regHost.type=="postReg") {
        return {
            data: state.regHost.postRegData
        }
    }
    return {
        data: {}
    }
}

// dispatch actions
const mapDispatchToProps = dispatch => {
    return {

    }
}


class RegCpu extends Component {

  constructor (props) {
    super(props);
  }


  getInput() {
      return {
          cpuExpected: {
              socket: this.inputCpuSocket.value,
              vcpu: this.inputCpuVcpu.value,
              model: this.inputCpuModel.value,
          }
      }
  }


  render() {
    return (
        <div>
            <p className="h7">1/10 CPU信息</p>
            <div className="form-group row">
              <label htmlFor="example-text-input" className="col-3 col-form-label">Socket(s)</label>
              <div className="col-9">
                  <input
                    className="form-control"
                    type="text"
                    placeholder="search..."
                    defaultValue={this.props.data.cpuExpected.socket && this.props.data.cpuExpected.socket}
                    ref={(me)=> {this.inputCpuSocket = me}}
                    // name="datacenter"
                    key={this.inputCpuSocket}
                    id="example-text-input"
                    onChange={()=>{console.log(this.inputCpuSocket.value)}}
                />
              </div>
            </div>
            <div className="form-group row">
              <label htmlFor="example-text-input" className="col-3 col-form-label">VCPU</label>
              <div className="col-9">
                  <input
                    className="form-control"
                    type="text"
                    placeholder="search..."
                    defaultValue={this.props.data.cpuExpected.vcpu && this.props.data.cpuExpected.vcpu}
                    ref={(me)=> {this.inputCpuVcpu = me}}
                    // name="datacenter"
                    key={this.inputCpuVcpu}
                    id="example-text-input"
                    onChange={()=>{console.log(this.inputCpuVcpu.value)}}
                />
              </div>
            </div>
            <div className="form-group row">
              <label htmlFor="example-text-input" className="col-3 col-form-label">型号</label>
              <div className="col-9">
                  <input
                    className="form-control"
                    type="text"
                    placeholder="search..."
                    defaultValue={this.props.data.cpuExpected.model && this.props.data.cpuExpected.model}
                    ref={(me)=> {this.inputCpuModel = me}}
                    // name="datacenter"
                    key={this.inputCpuModel}
                    id="example-text-input"
                    onChange={()=>{console.log(this.inputCpuModel.value)}}
                />
              </div>
            </div>

        </div>

    );
  }

}

RegCpu.defaultProps = {
    btnColor: "secondary",
    btnText: "注册",
    method: "json",
}

// https://facebook.github.io/react/docs/typechecking-with-proptypes.html
// for more prop types
RegCpu.propTypes = {
    btnColor: PropTypes.string,
    btnText: PropTypes.string,
    method: PropTypes.string,
}

export default connect(
    mapStateToProps,
    mapDispatchToProps,
    null,
    {withRef: true}    
) (RegCpu)
