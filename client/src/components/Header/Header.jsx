import React, { Component } from 'react';
import { Navbar } from 'react-bootstrap';

class Header extends Component {
  render() {
    return (
      <Navbar fluid>
        <Navbar.Header>
          <Navbar.Brand>
            <a>Stock Portfolio</a>
          </Navbar.Brand>
        </Navbar.Header>
      </Navbar>
    );
  }
}

export default Header;
