import React, {Component} from 'react'

class UselessComponent extends Component {
    state = {a : ""}

    componentDidMount() {
        console.log("UselessComponent will did mount")
    }

    componentWillReceiveProps(nextProp) {
        console.log("UselessComponent will receive props:", nextProp)
        this.setState({
            a: nextProp.a
        })
    }

    render() {
        console.log("Comments rending, state=", this.state)
        return (
            <div>
                useless
            </div>
        )
    }
}

export default UselessComponent
