import React, { Component } from 'react';
import { Grid } from 'react-bootstrap';

class Footer extends Component {
  render() {
    return (
      <footer className="footer">
        <Grid fluid>
          <p className="copyright pull-right">
            Iris Grace Endozo, January 2019
          </p>
        </Grid>
      </footer>
    );
  }
}

export default Footer;
