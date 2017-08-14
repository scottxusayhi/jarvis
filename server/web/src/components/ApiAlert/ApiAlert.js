import React, { Component } from 'react';
import { Alert } from 'reactstrap';


class ApiAlert extends Component {
  render() {
    return (
        <div>
            <Alert color="success">
                <strong>Well done!</strong> You successfully read this important alert message.
            </Alert>
        </div>
    )
  }
}

export default ApiAlert;