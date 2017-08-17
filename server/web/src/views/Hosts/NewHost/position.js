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


class RegPosition extends Component {

  constructor (props) {
    super(props);
  }




  render() {
    return (
        <div>
            <p className="h7">1/10 位置信息</p>
            <div className="form-group row">
              <label htmlFor="example-text-input" className="col-3 col-form-label">数据中心</label>
              <div className="col-9">
                <input className="form-control" type="text" value="Artisanal kale" id="example-text-input"/>
              </div>
            </div>
            <div className="form-group row">
              <label htmlFor="example-text-input" className="col-3 col-form-label">机架</label>
              <div className="col-9">
                <input className="form-control" type="text" value="Artisanal kale" id="example-text-input"/>
              </div>
            </div>
            <div className="form-group row">
              <label htmlFor="example-text-input" className="col-3 col-form-label">槽位</label>
              <div className="col-9">
                <input className="form-control" type="text" value="Artisanal kale" id="example-text-input"/>
              </div>
            </div>
            <div className="form-group row">
              <label htmlFor="example-text-input" className="col-3 col-form-label">拥有人</label>
              <div className="col-9">
                <input className="form-control" type="text" value="Artisanal kale" id="example-text-input"/>
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
    mapDispatchToProps
) (RegPosition)
