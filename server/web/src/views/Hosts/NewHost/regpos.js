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


class RegPosition extends Component {

  constructor (props) {
    super(props);
    this.state = {
    }
  }


  getInput() {
      return {
          datacenter: this.inputDatacenter.value,
          rack: this.inputRack.value,
          slot: this.inputSlot.value,
          owner: this.inputOwner.value,
      }
  }

  onShow() {
      console.log("on show")
  }

    componentWillUnmount() {
        console.log("regos table will un-mount")
    }

    componentWillReceiveProps(nextProps) {
      console.log(nextProps)
    }

    componentDidUpdate() {
        console.log("did update")
    }

    render() {
      console.log(this.state)

    return (
        <div>
            <p className="h7">1/10 位置信息</p>
            <div className="form-group row">
              <label htmlFor="example-text-input" className="col-3 col-form-label">数据中心</label>
              <div className="col-9">
                <input
                    className="form-control"
                    type="text"
                    placeholder="datacenter..."
                    defaultValue={this.props.data.datacenter && this.props.data.datacenter}
                    ref={(me)=> {this.inputDatacenter = me}}
                    key={this.inputDatacenter}
                    id="example-text-input"
                    onChange={()=>{console.log(this.inputDatacenter.value)}}
                />
              </div>
            </div>
            <div className="form-group row">
              <label htmlFor="example-text-input" className="col-3 col-form-label">机架</label>
              <div className="col-9">
                  <input
                    className="form-control"
                    type="text"
                    placeholder="rack..."
                    defaultValue={this.props.data.rack && this.props.data.rack}
                    ref={(me)=>this.inputRack = me}
                    // name="datacenter"
                    key={this.inputRack}
                    id="example-text-input"
                    onChange={()=>{console.log(this.inputRack.value)}}
                />
              </div>
            </div>
            <div className="form-group row">
              <label htmlFor="example-text-input" className="col-3 col-form-label">槽位</label>
              <div className="col-9">
                  <input
                    className="form-control"
                    type="text"
                    placeholder="slot..."
                    defaultValue={this.props.data.slot && this.props.data.slot}
                    ref={(me)=>this.inputSlot = me}
                    // name="datacenter"
                    key={this.inputSlot}
                    id="example-text-input"
                    onChange={()=>{console.log(this.inputSlot.value)}}
                />
              </div>
            </div>
            <div className="form-group row">
              <label htmlFor="example-text-input" className="col-3 col-form-label">拥有人</label>
              <div className="col-9">
                  <input
                    className="form-control"
                    type="text"
                    placeholder="owner..."
                    defaultValue={this.props.data.owner && this.props.data.owner}
                    ref={(me)=>this.inputOwner = me}
                    key={this.inputOwner}
                    id="example-text-input"
                    onChange={()=>{console.log(this.inputOwner.value)}}
                />
              </div>
            </div>

        </div>

    );
  }

}

RegPosition.defaultProps = {
    btnColor: "secondary",
    btnText: "注册",
    method: "json",
}

// https://facebook.github.io/react/docs/typechecking-with-proptypes.html
// for more prop types
RegPosition.propTypes = {
    btnColor: PropTypes.string,
    btnText: PropTypes.string,
    method: PropTypes.string,
}

export default connect(
    mapStateToProps,
    mapDispatchToProps,
    null,
    {withRef: true}
) (RegPosition)
