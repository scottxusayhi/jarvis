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


class RegMem extends Component {

  constructor (props) {
    super(props);
  }

  getInput() {
      return {
        memExpected: {
            total: Number(this.inputMemTotal.value)
        }
      }
  }


  render() {
    return (
        <div>
            <p className="h7">1/10 内存信息</p>
            <div className="form-group row">
              <label htmlFor="example-text-input" className="col-3 col-form-label">总容量</label>
              <div className="col-9">
                  <input
                    className="form-control"
                    type="text"
                    placeholder="in GB..."
                    defaultValue={this.props.data.memExpected && this.props.data.memExpected.total}
                    ref={(me)=> {this.inputMemTotal = me}}
                    key={this.inputMemTotal}
                    id="example-text-input"
                    onChange={()=>{console.log(this.inputMemTotal.value)}}
                />
              </div>
            </div>
        </div>

    );
  }

}

RegMem.defaultProps = {
    btnColor: "secondary",
    btnText: "注册",
    method: "json",
}

// https://facebook.github.io/react/docs/typechecking-with-proptypes.html
// for more prop types
RegMem.propTypes = {
    btnColor: PropTypes.string,
    btnText: PropTypes.string,
    method: PropTypes.string,
}

export default connect(
    mapStateToProps,
    mapDispatchToProps,
    null,
    {withRef: true}    
) (RegMem)
