import React, { Component } from 'react'
import PropTypes from 'prop-types';
import { connect } from 'react-redux'

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


class RegOs extends Component {

  constructor (props) {
    super(props);
  }


  getInput() {
      return {
          osExpected: {
              type: this.inputOsType.value,
              arch: this.inputOsArch.value,
              hostname: this.inputOsHostname.value,
          }
      }
  }


  render() {
    return (
        <div>
            <p className="h7">1/10 OS信息</p>
            <div className="form-group row">
              <label htmlFor="example-text-input" className="col-3 col-form-label">类型</label>
              <div className="col-9">
                  <input
                    className="form-control"
                    type="text"
                    placeholder="linux|windows..."
                    defaultValue={this.props.data.osExpected && this.props.data.osExpected.type}
                    ref={(me)=> {this.inputOsType = me}}
                    key={this.inputOsType}
                    id="example-text-input"
                    onChange={()=>{console.log(this.inputOsType.value)}}
                />
              </div>
            </div>
            <div className="form-group row">
              <label htmlFor="example-text-input" className="col-3 col-form-label">架构</label>
              <div className="col-9">
                  <input
                    className="form-control"
                    type="text"
                    placeholder="amd64..."
                    defaultValue={this.props.data.osExpected && this.props.data.osExpected.arch}
                    ref={(me)=> {this.inputOsArch = me}}
                    key={this.inputOsArch}
                    id="example-text-input"
                    onChange={()=>{console.log(this.inputOsArch.value)}}
                />
              </div>
            </div>
            <div className="form-group row">
              <label htmlFor="example-text-input" className="col-3 col-form-label">Hostname</label>
              <div className="col-9">
                  <input
                    className="form-control"
                    type="text"
                    placeholder="hostname..."
                    defaultValue={this.props.data.osExpected && this.props.data.osExpected.hostname}
                    ref={(me)=> {this.inputOsHostname = me}}
                    key={this.inputOsHostname}
                    id="example-text-input"
                    onChange={()=>{console.log(this.inputOsHostname.value)}}
                />
              </div>
            </div>

        </div>

    );
  }

}

RegOs.defaultProps = {
    btnColor: "secondary",
    btnText: "注册",
    method: "json",
}

// https://facebook.github.io/react/docs/typechecking-with-proptypes.html
// for more prop types
RegOs.propTypes = {
    btnColor: PropTypes.string,
    btnText: PropTypes.string,
    method: PropTypes.string,
}

export default connect(
    mapStateToProps,
    mapDispatchToProps,
    null,
    {withRef: true}
) (RegOs)
