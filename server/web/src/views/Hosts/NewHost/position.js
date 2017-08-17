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
  }




  render() {
    return (
        <div>
            <p className="h7">1/10 位置信息</p>
            <div className="form-group row">
              <label htmlFor="example-text-input" className="col-3 col-form-label">数据中心</label>
              <div className="col-9">
                <input
                    className="form-control"
                    type="text"
                    placeholder="search..."
                    defaultValue={this.props.data.datacenter && this.props.data.datacenter}
                    ref={(input)=>this.input = input}
                    // name="datacenter"
                    // key={`datacenter:${this.input}`}
                    key={this.input}
                    id="example-text-input"
                />
              </div>
            </div>
            <div className="form-group row">
              <label htmlFor="example-text-input" className="col-3 col-form-label">机架</label>
              <div className="col-9">
                  <input
                    className="form-control"
                    type="text"
                    placeholder="search..."
                    defaultValue={this.props.data.rack && this.props.data.rack}
                    ref={(input)=>this.input = input}
                    // name="datacenter"
                    // key={`datacenter:${this.input}`}
                    key={this.input}
                    id="example-text-input"
                />
              </div>
            </div>
            <div className="form-group row">
              <label htmlFor="example-text-input" className="col-3 col-form-label">槽位</label>
              <div className="col-9">
                  <input
                    className="form-control"
                    type="text"
                    placeholder="search..."
                    defaultValue={this.props.data.slot && this.props.data.slot}
                    ref={(input)=>this.input = input}
                    // name="datacenter"
                    // key={`datacenter:${this.input}`}
                    key={this.input}
                    id="example-text-input"
                />
              </div>
            </div>
            <div className="form-group row">
              <label htmlFor="example-text-input" className="col-3 col-form-label">拥有人</label>
              <div className="col-9">
                  <input
                    className="form-control"
                    type="text"
                    placeholder="search..."
                    defaultValue={this.props.data.owner && this.props.data.owner}
                    ref={(input)=>this.input = input}
                    // name="datacenter"
                    // key={`datacenter:${this.input}`}
                    key={this.input}
                    id="example-text-input"
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
    mapDispatchToProps
) (RegPosition)
