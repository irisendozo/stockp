import React, { Component } from 'react';
import { Table } from 'react-bootstrap';

import Card from '../Card/Card.jsx';

export class HoldingsTable extends Component {
  render() {
    const tableHeaders = ['Code', 'Name', 'Units', 'Price($)'];

    return (
      <Card
        title="My Holdings + Cash"
        category="Displays all holdings in your portfolio, your cash component and any other assets and liabilities you have entered"
        ctTableFullWidth
        ctTableResponsive
        content={
          <Table striped hover>
            <thead>
              <tr>
                {tableHeaders.map((prop, key) => {
                  return <th key={key}>{prop}</th>;
                })}
              </tr>
            </thead>
            <tbody>
              {this.props.contents.map((prop, key) => {
                return (
                  <tr key={key}>
                    {prop.map((prop, key) => {
                      return <td key={key}>{prop}</td>;
                    })}
                  </tr>
                );
              })}
            </tbody>
          </Table>
        }
      />
    );
  }
}

export default HoldingsTable;
