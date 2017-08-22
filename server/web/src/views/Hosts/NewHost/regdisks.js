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


class RegDisks extends Component {

  constructor (props) {
    super(props);
    this.inputDiskDevice = [];
    this.inputDiskModel = [];
    this.inputDiskCap = [];
  }

  getInput() {
      return {
          diskExpected: Array.apply(null, {length: this.inputDiskDevice.length}).map((o, index)=>{
              console.log(index)
              return {
                  device: this.inputDiskDevice[index].value,
                  // model: this.inputDiskModel[index].value,
                  capacity: Number(this.inputDiskCap[index].value),
              }
          })
      }
  }


  render() {
    return (
        <div>
            <p className="h7">1/10 磁盘信息</p>
            {this.props.data.diskExpected && this.props.data.diskExpected.map((disk, index)=>{return (
                <div>
                <div className="form-group row">
              <label htmlFor="example-text-input" className="col-3 col-form-label">设备{index+1}</label>
              <div className="col-9">
                  <input
                    className="form-control"
                    type="text"
                    placeholder="search..."
                    defaultValue={disk.device}
                    ref={(me)=> {this.inputDiskDevice.push(me)}}
                    // name="datacenter"
                    key={this.inputDiskDevice}
                    id="example-text-input"
                    onChange={()=>{console.log(this.inputDiskDevice[index].value)}}
                />
              </div>
                </div>
                <div className="form-group row">
              <label htmlFor="example-text-input" className="col-3 col-form-label">型号{index+1}</label>
              <div className="col-9">
                  <input
                    className="form-control"
                    type="text"
                    placeholder="search..."
                    defaultValue={disk.model}
                    ref={(me)=> {this.inputDiskModel.push(me)}}
                    // name="datacenter"
                    key={this.inputDiskModel}
                    id="example-text-input"
                    onChange={()=>{console.log(this.inputDiskModel[index].value)}}
                />
              </div>

            </div>
            <div className="form-group row">
              <label htmlFor="example-text-input" className="col-3 col-form-label">容量{index+1}</label>
              <div className="col-9">
                  <input
                    className="form-control"
                    type="text"
                    placeholder="search..."
                    defaultValue={disk.capacity}
                    ref={(me)=> {this.inputDiskCap.push(me)}}
                    // name="datacenter"
                    key={this.inputDiskCap}
                    id="example-text-input"
                    onChange={()=>{console.log(this.inputDiskCap[index].value)}}
                />
              </div>
            </div>
                </div>
            )})}

            <button type="button" className="btn btn-secondary">添加</button>
        </div>

    );
  }

}

RegDisks.defaultProps = {
    btnColor: "secondary",
    btnText: "注册",
    method: "json",
}

// https://facebook.github.io/react/docs/typechecking-with-proptypes.html
// for more prop types
RegDisks.propTypes = {
    btnColor: PropTypes.string,
    btnText: PropTypes.string,
    method: PropTypes.string,
}

export default connect(
    mapStateToProps,
    mapDispatchToProps,
    null,
    {withRef: true}    
) (RegDisks)
