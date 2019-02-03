import React, { Component } from 'react';
import { inject, observer } from 'mobx-react';
import { Grid, Row, Col } from 'react-bootstrap';

import Header from '../../components/Header/Header';
import Footer from '../../components/Footer/Footer';
import HistoryTable from '../../components/HistoryTable/HistoryTable.jsx';
import HoldingsTable from '../../components/HoldingsTable/HoldingsTable.jsx';
import SimpleForm from '../../components/SimpleForm/SimpleForm.jsx';
import PlaceOrderForm from '../../components/PlaceOrderForm/PlaceOrderForm.jsx';
import HeroStats from '../../components/HeroStats/HeroStats.jsx';

class MainPage extends Component {
  componentDidMount() {
    const { balanceStore, stocksStore } = this.props;

    balanceStore.fetchBalance();
    stocksStore.fetchOwnedStocks();
    stocksStore.fetchPurchaseHistory();
  }

  render() {
    const {
      balance,
      withdrawBalanceFeedback,
      addBalanceFeedback,
    } = this.props.balanceStore;
    const {
      ownedStocksMap,
      searchedStock,
      netAssetPrice,
      purchaseHistoryMap,
    } = this.props.stocksStore;

    return (
      <div className="wrapper">
        <div id="main-panel" className="main-panel" ref="mainPanel">
          <Header {...this.props} />
          <div className="content">
            <Grid fluid>
              <HeroStats balance={balance} netAsset={netAssetPrice} />
              <Row>
                <Col md={12}>
                  <HistoryTable contents={purchaseHistoryMap} />
                </Col>
              </Row>
              <Row>
                <Col md={6}>
                  <HoldingsTable contents={ownedStocksMap} />
                </Col>
                <Col md={6}>
                  <PlaceOrderForm
                    searchFilter={filter =>
                      this.props.stocksStore.searchStock(filter)
                    }
                    searchedStock={searchedStock}
                    buy={(symbol, quantity) =>
                      this.props.stocksStore.buyStock(symbol, quantity)
                    }
                    sell={(symbol, quantity) =>
                      this.props.stocksStore.sellStock(symbol, quantity)
                    }
                  />
                </Col>
              </Row>
              <Row>
                <Col md={6} />
                <Col md={3}>
                  <SimpleForm
                    title="Withdraw Balance"
                    status="danger"
                    submitForm={amount =>
                      this.props.balanceStore.withdrawBalance(amount)
                    }
                    feedbackMessage={
                      withdrawBalanceFeedback.error ||
                      withdrawBalanceFeedback.success
                    }
                  />
                </Col>
                <Col md={3}>
                  <SimpleForm
                    title="Add Balance"
                    status="success"
                    submitForm={amount =>
                      this.props.balanceStore.addBalance(amount)
                    }
                    feedbackMessage={
                      addBalanceFeedback.error || addBalanceFeedback.success
                    }
                  />
                </Col>
              </Row>
            </Grid>
          </div>
          <Footer />
        </div>
      </div>
    );
  }
}

export default inject(({ balanceStore, stocksStore }) => ({
  balanceStore,
  stocksStore,
}))(observer(MainPage));
