import React, {Component} from 'react'
import {connect} from 'react-redux'

// import {InputGroup, InputGroupAddon, Input} from 'reactstrap';

import {Input} from 'antd'

import {
    updateRegHost,
    fetchHostDetail
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

    constructor(props) {
        super(props);
        this.state = {
            comments: ""
        }
    }

    componentDidMount() {
        console.log("Comments will did mount")
        this.handleReset()
    }

    componentWillReceiveProps(nextProp) {
        console.log("Comments will receive props:", nextProp)
        this.setState({
            comments: nextProp.hostDetail.data.comments
        })
    }

    handleReset() {
        this.setState({comments: this.props.hostDetail.data.comments})
        // this.setState({comments: this.props.comments})
    }

    handleSave() {
        this.props.updateRegHost(this.props.hostDetail.data.systemId, {comments: this.state.comments})
        // this.props.updateRegHost(this.props.hostId, {comments: this.state.comments})
    }

    render() {
        console.log("Comments rending, state=", this.state)
        return (
            <div>
                <Input
                    type="textarea"
                    name="inputCommentsName"
                    id="inputComments"
                    rows="20"
                    onChange={(e) => this.setState({comments: e.target.value})}
                    value={this.state.comments}
                />
                <button type="button" className="btn btn-secondary"
                        onClick={() => {this.handleReset()}}>重置
                </button>
                <button type="button" className="btn btn-primary"
                        onClick={() => {this.handleSave()}}>
                    保存
                </button>
            </div>
        )
    }


}

export default connect(
    mapStateToProps,
    mapDispatchToProps,
    // null,
    // {withRef: true}
)(Comments)
