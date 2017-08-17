import React, { Component } from 'react'
import PropTypes from 'prop-types';
import { connect } from 'react-redux'
import { Button, Modal, ModalHeader, ModalBody, ModalFooter } from 'reactstrap';

import {
    registerHost
} from '../../../states/actions'

import RegPosition from "./position"
import RegOs from "./regos"
import RegCpu from "./regcpu"
import RegMem from "./regmem"
import RegDisks from "./regdisks"
import RegNet from "./regnet"

// subscribe
const mapStateToProps = state => {
    return {
    }
}

// dispatch actions
const mapDispatchToProps = dispatch => {
    return {
    }
}


class Registration extends Component {

  constructor (props) {
    super(props);
    this.state = {
        step: 1
    }
  }

  saveAndNext() {
      this.setState({
          step: this.state.step+1
      })
  }

  previous() {
      this.setState({
          step: this.state.step-1
      })
  }

  register() {
      console.log("api call should be here")
  }

  render() {
      console.log("render steps " + this.state.step)
      switch (this.state.step) {
          case 1:
              return (
                  <div>
                    <RegPosition/>
                    <button type="button" className="btn btn-primary" onClick={()=>{this.saveAndNext()}}>下一步</button>
                  </div>
              )
          case 2:
              return (
                  <div>
                    <RegOs/>
                    <button type="button" className="btn btn-secondary" onClick={()=>{this.previous()}}>上一步</button>
                    <button type="button" className="btn btn-primary" onClick={()=>{this.saveAndNext()}}>下一步</button>
                  </div>
              )
          case 3:
              return (
                  <div>
                    <RegCpu/>
                    <button type="button" className="btn btn-secondary" onClick={()=>{this.previous()}}>上一步</button>
                    <button type="button" className="btn btn-primary" onClick={()=>{this.saveAndNext()}}>下一步</button>
                  </div>
              )
          case 4:
              return (
                  <div>
                    <RegMem/>
                    <button type="button" className="btn btn-secondary" onClick={()=>{this.previous()}}>上一步</button>
                    <button type="button" className="btn btn-primary" onClick={()=>{this.saveAndNext()}}>下一步</button>
                  </div>
              )
          case 5:
              return (
                  <div>
                    <RegDisks/>
                    <button type="button" className="btn btn-secondary" onClick={()=>{this.previous()}}>上一步</button>
                    <button type="button" className="btn btn-primary" onClick={()=>{this.saveAndNext()}}>下一步</button>
                  </div>
              )
          case 6:
              return (
                  <div>
                    <RegNet/>
                      <button type="button" className="btn btn-secondary" onClick={()=>{this.previous()}}>上一步</button>
                      <button type="button" className="btn btn-primary" onClick={()=>{this.register()}}>走你</button>
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

export default Registration
