import React, { Component } from 'react';
import { Pagination, PaginationItem, PaginationLink, Button } from 'reactstrap';
import { connect } from 'react-redux'
import {
    fetchHosts,
} from '../../states/actions'
// import { Select } from "react-select"
import 'react-select/dist/react-select.css';
var Select = require('react-select')

// subscribe
const mapStateToProps = state => {
    return {
        pageInfo: state.hosts.data.pageInfo,
    }
}

// dispatch actions
const mapDispatchToProps = dispatch => {
    return {
        onPageChange: newPage => {
            console.log("new page clicked " + newPage)
            var filter = {
                page: newPage
            }
            dispatch(fetchHosts(filter))
        }
    }
}


class Pager extends Component {

    constructor(props) {
        super(props);
    }

    options = [
  { value: 10, label: '10' },
  { value: 20, label: '20' },
        { value: 50, label: '50' },
        { value: 100, label: '100' },
];
    
    previous() {
       if (this.props.pageInfo && this.props.pageInfo.page==1) {
           return <PaginationItem disabled>
          <PaginationLink previous/>
        </PaginationItem>
       } else {
           return <PaginationItem>
          <PaginationLink previous/>
        </PaginationItem>
       }
    }

    pages() {
       var list = [];
       if (this.props.pageInfo) {
           var i
           for (i=1; i<=this.props.pageInfo.totalPage; i++) {
               list.push(i)
           }
       }
       return list
    }

    next() {
       if (this.props.pageInfo && this.props.pageInfo.page==this.props.pageInfo.totalPage) {
           return <PaginationItem disabled>
          <PaginationLink next/>
        </PaginationItem>
       } else {
           return <PaginationItem>
          <PaginationLink next/>
        </PaginationItem>
       }
    }



    render() {
    return (
        <div className="row">
        总计：{this.props.pageInfo && this.props.pageInfo.totalSize}

          <Pagination size="sm">

              {this.previous()}

              {
                  this.pages().map(page => {
                      {
                       if (page==this.props.pageInfo.page) {
                           return <PaginationItem active>
                                    <PaginationLink>
                                        {page}
                                    </PaginationLink>
                                  </PaginationItem>
                       } else {
                           return <PaginationItem>
                                    <PaginationLink onClick={() => {this.props.onPageChange(page)}}>
                                        {page}
                                    </PaginationLink>
                                  </PaginationItem>
                       }

                      }
                  })
              }

              {this.next()}


      </Pagination>
        </div>
    )
  }
}

// export default connect(
//     mapStateToProps,
//     mapDispatchToProps
// )(Pager)

export default Pager

