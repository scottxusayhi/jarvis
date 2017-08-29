import React, { Component } from 'react'
import PropTypes from 'prop-types';
import { connect } from 'react-redux'
import { Button, Modal, ModalHeader, ModalBody, ModalFooter } from 'reactstrap';

import NewHostTab from "./NewHostTab"

import {
    newRegStart,
    postRegStart,
    regCancelled,
} from '../../../states/actions'


function makeRegDataFromDetected(detected) {
    return {
        datacenter: detected.datacenter,
        rack: detected.rack,
        slot: detected.slot,
        tags: detected.tags,
        owner: detected.owner,
        osExpected: detected.osDetected,
        cpuExpected: detected.cpuDetected,
        memExpected: detected.memDetected,
        diskExpected: detected.diskDetected,
        networkExpected: detected.networkDetected
    }
}

// subscribe
const mapStateToProps = state => {
    return {
        regHost: state.regHost,
        hostDetail: state.hostDetail,
    }
}

// dispatch actions
const mapDispatchToProps = dispatch => {
    return {
        newRegStart: () => dispatch(newRegStart()),
        postRegStart: (id, initData) => dispatch(postRegStart(id, initData)),
        regCancelled: () => dispatch(regCancelled())
    }
}

class NewHostButton extends Component {

  constructor (props) {
    super(props);
    this.state = {
      modal: false
    };

    this.toggle = this.toggle.bind(this);
    this.regStart = this.regStart.bind(this);
  }

  toggle() {
    this.setState({
      modal: !this.state.modal
    });
  }

  regStart() {
      if (this.props.regType==="newReg") {
          this.props.newRegStart()
      }
      else if (this.props.regType==="postReg") {
          this.props.postRegStart(this.props.hostDetail.data.systemId, makeRegDataFromDetected(this.props.hostDetail.data))
      }
      else {
          console.error("unknown reg type: " + this.props.regType)
      }
  }

  render() {
    console.log(this.props)
    return (
      <div>
        {/*<Button color={this.props.btnColor} onClick={this.toggle}><i className="fa fa-plus"></i>&nbsp; {this.props.btnText}</Button>*/}
        {/*<Modal isOpen={this.state.modal} toggle={this.toggle} className={this.props.className}>*/}
        <Button color={this.props.btnColor} onClick={this.regStart}><i className="fa fa-plus"></i>&nbsp; {this.props.btnText}</Button>
        <Modal isOpen={!this.props.regHost.success} toggle={this.props.regCancelled} className={this.props.className}>
          <ModalHeader toggle={this.props.regCancelled}>注册新主机</ModalHeader>
          <ModalBody>
            <NewHostTab regType={this.props.regType} postRegHostId={this.props.postRegHostId}/>
          </ModalBody>
          {/*<ModalFooter>*/}
            {/*<Button color="secondary" onClick={this.toggle}>取消</Button>*/}
            {/*<Button color="primary" onClick={() => this.props.registerHost(this.getInput())}>注册</Button>{' '}*/}
          {/*</ModalFooter>*/}
        </Modal>
      </div>
    );
  }

}

NewHostButton.defaultProps = {
    btnColor: "secondary",
    btnText: "注册",
    regType: "newReg",
    postRegHostId: 0, // takes effect only when regType=postReg
}

// https://facebook.github.io/react/docs/typechecking-with-proptypes.html
// for more prop types
NewHostButton.propTypes = {
    btnColor: PropTypes.string,
    btnText: PropTypes.string,
    regType: PropTypes.oneOf(["newReg", "postReg"]),
    postRegHostId: PropTypes.number,
}

export default connect(
    mapStateToProps,
    mapDispatchToProps
) (NewHostButton)
