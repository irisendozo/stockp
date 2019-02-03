import React, { Component } from 'react';
import NumericLabel from 'react-pretty-numbers';

export class MoneyLabel extends Component {
  render() {
    const moneyFormatParams = {
      currency: true,
      currencyIndicator: 'USD',
      shortFormat: true,
      precision: 2,
      justification: 'L',
    };

    return (
      <NumericLabel params={moneyFormatParams}>
        {this.props.amount || 0}
      </NumericLabel>
    );
  }
}

export default MoneyLabel;
