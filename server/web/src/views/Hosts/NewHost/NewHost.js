import React, { Component } from 'react'
import PropTypes from 'prop-types';
import { connect } from 'react-redux'
import { Button, Modal, ModalHeader, ModalBody, ModalFooter } from 'reactstrap';

import NewHostTab from "./NewHostTab"

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

  render() {
    console.log(this.props)
    return (
      <div>
        <Button color={this.props.btnColor} onClick={this.toggle}><i className="fa fa-plus"></i>&nbsp; {this.props.btnText}</Button>
        <Modal isOpen={this.state.modal} toggle={this.toggle} className={this.props.className}>
          <ModalHeader toggle={this.toggle}>注册新主机</ModalHeader>
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
