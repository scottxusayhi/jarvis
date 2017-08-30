import React, { Component } from 'react'
import { connect } from 'react-redux'

import { InputGroup, InputGroupAddon, Input } from 'reactstrap';

import {
  updateRegHost, fetchHostDetail
} from '../../states/actions'


// subscribe state
const mapStateToProps = state => {
  return {
    hostDetail: state.hostDetail
  }
}

// dispatch actions
const mapDispatchToProps = dispatch => {
    return {
        updateRegHost: (id, data) => {
            dispatch(updateRegHost(id, data))
        },
        fetchHostDetail: (id) => {
            dispatch(fetchHostDetail(id))
        }
    }
}

class Comments extends Component {

  constructor (props) {
    super(props);
    this.state = {
        comments: ""
    }
  }


  componentWillReceiveProps(nextProps) {
    this.setState({
        comments: nextProps.hostDetail.data.comments
    })
  }


    render() {
    return (
      <div>
            <Input
                type="textarea"
                name="inputCommentsName"
                id="inputComments"
                rows="20"
                ref={(me)=> {this.inputComments = me}}
                key={this.inputComments}
                onChange={(e)=>this.setState({comments: e.target.value})}
                value={this.state.comments}
            />
          <button type="button" className="btn btn-secondary" onClick={()=>this.setState({comments: this.props.hostDetail.data.comments})}>重置</button>
          <button type="button" className="btn btn-primary" onClick={()=>this.props.updateRegHost(this.props.hostDetail.data.systemId, {comments: this.state.comments})}>保存</button>
      </div>
    )
  }



}

export default connect(
    mapStateToProps,
    mapDispatchToProps,
    null,
    {withRef: true}
) (Comments)
