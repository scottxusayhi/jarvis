import React, { Component } from 'react'
import { connect } from 'react-redux'

import { Tag, Input, Tooltip, Button } from 'antd';


import {
    attachTagToHost,
    removeTagFromHost,
} from '../../../../states/actions'

// subscribe
const mapStateToProps = state => {
    return {
        tags: state.hostDetail.data.tags,
        hostId: state.hostDetail.data.systemId,
    }
}

// dispatch actions
const mapDispatchToProps = dispatch => {
    return {
        attachTagToHost: (hostId, tag) => {
            dispatch(attachTagToHost(hostId, tag))
        },
        removeTagFromHost: (hostId, tag) => {
            dispatch(removeTagFromHost(hostId, tag))
        }
    }
}

class HostTags extends Component {
  state = {
    tags: [],
    inputVisible: false,
    inputValue: '',
  };

  handleClose = (removedTag) => {
    const tags = this.state.tags.filter(tag => tag !== removedTag);
    console.log(tags);
    this.setState({ tags });
    this.props.removeTagFromHost(this.props.hostId, [removedTag])
  }

  showInput = () => {
    this.setState({ inputVisible: true }, () => this.input.focus());
  }

  handleInputChange = (e) => {
    this.setState({ inputValue: e.target.value });
  }

  handleInputConfirm = () => {
    const state = this.state;
    const inputValue = state.inputValue;
    let tags = state.tags;
    if (inputValue && tags.indexOf(inputValue) === -1) {
      tags = [...tags, inputValue];
    }
    console.log(tags);
    this.setState({
      tags,
      inputVisible: false,
      inputValue: '',
    });
    this.props.attachTagToHost(this.props.hostId, [inputValue])
  }

  saveInputRef = input => this.input = input

  componentWillReceiveProps(nextProps) {
    this.setState({
        tags: nextProps.tags
    })
  }

  render() {
    const { tags, inputVisible, inputValue } = this.state;
    return (
      <div>
        {tags.map((tag, index) => {
          const isLongTag = tag.length > 20;
          const tagElem = (
            <Tag key={tag} closable={true} color="blue" afterClose={() => this.handleClose(tag)}>
              {isLongTag ? `${tag.slice(0, 20)}...` : tag}
            </Tag>
          );
          return isLongTag ? <Tooltip title={tag}>{tagElem}</Tooltip> : tagElem;
        })}
        {inputVisible && (
          <Input
            ref={this.saveInputRef}
            type="text"
            size="small"
            style={{ width: 78 }}
            value={inputValue}
            onChange={this.handleInputChange}
            onBlur={this.handleInputConfirm}
            onPressEnter={this.handleInputConfirm}
          />
        )}
        {!inputVisible && <Button size="small" type="dashed" onClick={this.showInput}>+ New Tag</Button>}
      </div>
    );
  }


}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(HostTags)
