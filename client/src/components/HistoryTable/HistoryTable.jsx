import React, { Component } from 'react';
import { Table } from 'react-bootstrap';

import Card from '../Card/Card.jsx';

export class HistoryTable extends Component {
  render() {
    const tableHeaders = ['Code', 'Name', 'Units', 'Price($)', 'Purchase Date'];

    return (
      <Card
        title="Purchase History"
        category="Displays all purchase history"
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

export default HistoryTable;
