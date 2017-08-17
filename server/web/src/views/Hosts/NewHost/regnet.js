import React, { Component } from 'react'
import PropTypes from 'prop-types';
import { connect } from 'react-redux'
import { Button, Modal, ModalHeader, ModalBody, ModalFooter } from 'reactstrap';

import {
    registerHost
} from '../../../states/actions'

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


class RegNet extends Component {

  constructor (props) {
    super(props);
  }




  render() {
    return (
        <div>
            <p className="h7">1/10 网络信息</p>
            <div className="form-group row">
              <label htmlFor="example-text-input" className="col-3 col-form-label">IP</label>
              <div className="col-9">
                <input className="form-control" type="text" value="Artisanal kale" id="example-text-input"/>
              </div>
            </div>
        </div>

    );
  }

}

RegNet.defaultProps = {
    btnColor: "secondary",
    btnText: "注册",
    method: "json",
}

// https://facebook.github.io/react/docs/typechecking-with-proptypes.html
// for more prop types
RegNet.propTypes = {
    btnColor: PropTypes.string,
    btnText: PropTypes.string,
    method: PropTypes.string,
}

export default connect(
    mapStateToProps,
    mapDispatchToProps
) (RegNet)
