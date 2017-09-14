import React, {Component} from 'react'
import { connect } from 'react-redux'
import {Modal, Button, Input, Tag} from 'antd';

import {
    updateRegHost,
    fetchHostDetail
} from '../../states/actions'


// subscribe state
const mapStateToProps = state => {
    return {
        state: state
    }
}

// dispatch actions
const mapDispatchToProps = dispatch => {
    return {
        updateRegHost: (id, data) => {
            dispatch(updateRegHost(id, data))
        }
    }
}

class CommentEditor extends React.Component {
    state = {
        loading: false,
        visible: false,
        comments: "",
    }
    showModal = () => {
        this.setState({
            visible: true,
        });
    }
    handleOk = () => {
        // this.setState({loading: true});
        // setTimeout(() => {
        //     this.setState({loading: false, visible: false});
        // }, 3000);
        this.props.updateRegHost(this.props.host, {comments: this.state.comments})
        this.setState({
            visible: false,
        })

    }
    handleCancel = () => {
        this.setState({visible: false});
    }
    handleReset = () => {
        this.setState({
            comments: this.props.comments
        })
    }

    componentWillReceiveProps(nextProps) {
        this.setState({
            comments: nextProps.comments
        })
    }

    render() {
        const {visible, loading} = this.state;
        return (
            <div>
                <Tag color="cyan" value={this.props.children} onClick={this.showModal}>{this.props.children}</Tag>
                <Modal
                    visible={visible}
                    title="备注"
                    onOk={this.handleOk}
                    onCancel={this.handleCancel}
                    footer={[
                        <Button key="reset" size="large" onClick={this.handleReset}>重置</Button>,
                        <Button key="back" size="large" onClick={this.handleCancel}>取消</Button>,
                        <Button key="submit" type="primary" size="large" loading={loading} onClick={this.handleOk}>
                            保存
                        </Button>,
                    ]}
                >
            <Input
                type="textarea"
                name="inputCommentsName"
                id="inputComments"
                rows="20"
                onChange={(e)=>this.setState({comments: e.target.value})}
                value={this.state.comments}
            />
                </Modal>
            </div>
        );
    }
}

export default connect(
    mapStateToProps,
    mapDispatchToProps,
) (CommentEditor)

// export default CommentEditor