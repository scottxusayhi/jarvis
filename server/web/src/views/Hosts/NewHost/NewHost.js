import React, { Component } from 'react'
import { connect } from 'react-redux'
import { Button, Modal, ModalHeader, ModalBody, ModalFooter } from 'reactstrap';
import NewHostTab from "../../../components/NewHostTab/NewHostTab";

import {
    registerHost
} from '../../../states/actions'

// subscribe
const mapStateToProps = state => {
    return {
        modal: state.success,
    }
}

// dispatch actions
const mapDispatchToProps = dispatch => {
    return {
        registerHost: (payload) => {
          dispatch(registerHost(payload))
        }
    }
}


class NewHostButton extends Component {

  constructor (props) {
    super(props);
    this.state = {
      modal: false
    };

    this.toggle = this.toggle.bind(this);
  }

  toggle() {
    this.setState({
      modal: !this.state.modal
    });
  }

  getInput() {
    return "{}"
  }

  render() {
    console.log('rendering new host');
    console.log(this.props)
    return (
      <div>
        <Button color="secondary" onClick={this.toggle}><i className="fa fa-plus"></i>&nbsp; 创建</Button>
        <Modal isOpen={this.state.modal} toggle={this.toggle} className={this.props.className}>
          <ModalHeader toggle={this.toggle}>注册新主机</ModalHeader>
          <ModalBody>
            <NewHostTab/>
          </ModalBody>
          <ModalFooter>
            <Button color="secondary" onClick={this.toggle}>取消</Button>
            <Button color="primary" onClick={() => this.props.registerHost(this.getInput())}>注册</Button>{' '}
          </ModalFooter>
        </Modal>
      </div>
    );
  }

}

export default connect(
    mapStateToProps,
    mapDispatchToProps
) (NewHostButton)
