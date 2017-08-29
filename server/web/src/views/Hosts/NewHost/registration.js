import React, { Component } from 'react'
import PropTypes from 'prop-types';
import { connect } from 'react-redux'
import { Button, Modal, ModalHeader, ModalBody, ModalFooter } from 'reactstrap';

import RegPosition from "./regpos"
import RegOs from "./regos"
import RegCpu from "./regcpu"
import RegMem from "./regmem"
import RegDisks from "./regdisks"
import RegNet from "./regnet"

import {
    newRegDataSaved,
    registerHost,
    postRegDataSaved,
    postRegHost,
} from '../../../states/actions'

// subscribe
const mapStateToProps = state => {
    return {
        regHost: state.regHost
    }
}

// dispatch actions
const mapDispatchToProps = dispatch => {
    return {
        postRegDataSaved: data => dispatch(postRegDataSaved(data)),
        newRegDataSaved: data => dispatch(newRegDataSaved(data)),
        postRegHost: (id, data) => dispatch(postRegHost(id, data)),
        newRegHost: (data) => dispatch(registerHost(data))
    }
}


class Registration extends Component {

  constructor (props) {
    super(props);
    this.state = {
        step: 1
    }
  }

  saveAndNext(data) {
      if (this.props.regHost.type==="newReg") {
          this.props.newRegDataSaved(data)
      }
      if (this.props.regHost.type==="postReg") {
          this.props.postRegDataSaved(data)
      }
      this.setState({
          step: this.state.step+1
      })
  }

  previous() {
      this.setState({
          step: this.state.step-1
      })
  }

  saveAndGo(data) {
      if (this.props.regHost.type==="newReg") {
          this.props.newRegDataSaved(data)
          this.props.newRegHost(this.props.regHost.newRegData)
      }
      if (this.props.regHost.type==="postReg") {
          this.props.postRegDataSaved(data)
          this.props.postRegHost(this.props.regHost.postRegHostId, this.props.regHost.postRegData)
      }
  }

  render() {
      console.log("render steps " + this.state.step)
      switch (this.state.step) {
          case 1:
              return (
                  <div>
                    <RegPosition ref={(me)=>{this.regPosRef=me}}/>
                    <button type="button" className="btn btn-primary" onClick={()=>{this.saveAndNext(this.regPosRef.getWrappedInstance().getInput())}}>下一步</button>
                  </div>
              )
          case 2:
              return (
                  <div>
                    <RegOs ref={(me)=>{this.regOsRef=me}}/>
                    <button type="button" className="btn btn-secondary" onClick={()=>{this.previous()}}>上一步</button>
                    <button type="button" className="btn btn-primary" onClick={()=>{this.saveAndNext(this.regOsRef.getWrappedInstance().getInput())}}>下一步</button>
                  </div>
              )
          case 3:
              return (
                  <div>
                    <RegCpu ref={(me)=>{this.regCpuRef=me}}/>
                    <button type="button" className="btn btn-secondary" onClick={()=>{this.previous()}}>上一步</button>
                    <button type="button" className="btn btn-primary" onClick={()=>{this.saveAndNext(this.regCpuRef.getWrappedInstance().getInput())}}>下一步</button>
                  </div>
              )
          case 4:
              return (
                  <div>
                    <RegMem ref={(me)=>{this.regMemRef=me}}/>
                    <button type="button" className="btn btn-secondary" onClick={()=>{this.previous()}}>上一步</button>
                    <button type="button" className="btn btn-primary" onClick={()=>{this.saveAndNext(this.regMemRef.getWrappedInstance().getInput())}}>下一步</button>
                  </div>
              )
          case 5:
              return (
                  <div>
                    <RegDisks ref={(me)=>{this.regDisksRef=me}}/>
                    <button type="button" className="btn btn-secondary" onClick={()=>{this.previous()}}>上一步</button>
                    <button type="button" className="btn btn-primary" onClick={()=>{this.saveAndNext(this.regDisksRef.getWrappedInstance().getInput())}}>下一步</button>
                  </div>
              )
          case 6:
              return (
                  <div>
                    <RegNet ref={(me)=>{this.regNetRef=me}}/>
                      <button type="button" className="btn btn-secondary" onClick={()=>{this.previous()}}>上一步</button>
                      <button type="button" className="btn btn-primary" onClick={()=>{this.saveAndGo(this.regNetRef.getWrappedInstance().getInput())}}>走你</button>
                  </div>
              )
      }
  }

}

Registration.defaultProps = {
    btnColor: "secondary",
    btnText: "注册",
    method: "json",
}

// https://facebook.github.io/react/docs/typechecking-with-proptypes.html
// for more prop types
Registration.propTypes = {
    btnColor: PropTypes.string,
    btnText: PropTypes.string,
    method: PropTypes.string,
}

// export default connect(
//     mapStateToProps,
//     mapDispatchToProps
// ) (Registration)

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(Registration)
