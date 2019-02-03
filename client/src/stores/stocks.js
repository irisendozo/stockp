//@flow
import { action, computed, decorate, flow, observable } from 'mobx';

import * as API from '../api';

type FeedbackMessageModel = {
  success?: string,
  error?: string,
};

type Stock = {
  Symbol?: string,
  Name?: string,
  Price: number,
  Count?: number,
  PurchaseDate?: string,
};

type SearchedStock = {
  Symbol?: string,
  Name?: string,
  Price?: string,
};

class StocksStore {
  purchaseHistory: Array<Stock> = [];
  ownedStocks: Array<Stock> = [];
  searchedStock: SearchedStock = {};
  buyStockFeedback: FeedbackMessageModel = {};
  sellStockFeedback: FeedbackMessageModel = {};

  fetchPurchaseHistory = flow(function*() {
    try {
      const { data: apiResponse } = yield API.fetchPurchaseHistory();
      this.purchaseHistory = apiResponse;
    } catch (error) {}
  });

  fetchOwnedStocks = flow(function*() {
    try {
      const { data: apiResponse } = yield API.fetchOwnedStocks();
      this.ownedStocks = apiResponse.OwnedStocks;
    } catch (error) {}
  });

  searchStock = flow(function*(filter) {
    try {
      const { data: apiResponse } = yield API.searchStocks(filter);
      this.searchedStock = apiResponse;
    } catch (error) {}
  });

  buyStock = flow(function*(symbol, quantity) {
    this.buyStockFeedback = {};

    try {
      const { data: apiResponse } = yield API.buyStock(symbol, quantity);
      this.ownedStocks = apiResponse.OwnedStocks;
      this.buyStockFeedback.success = 'You have successfully bought the stock';
    } catch (error) {
      this.buyStockFeedback.error = 'You have failed to buy the stock';
    }
  });

  sellStock = flow(function*(symbol, quantity) {
    this.sellStockFeedback = {};

    try {
      const { data: apiResponse } = yield API.sellStock(symbol, quantity);
      this.ownedStocks = apiResponse.OwnedStocks;
      this.sellStockFeedback.success = 'You have successfully sold the stock';
    } catch (error) {
      this.sellStockFeedback.error = 'You have failed to sell the stock';
    }
  });

  get netAssetPrice() {
    let netAsset = 0;

    this.ownedStocks.forEach(ownedStock => {
      netAsset += netAsset + ownedStock.Count * ownedStock.Price;
    });

    return netAsset;
  }

  get ownedStocksMap() {
    let stocksMap = [];

    this.ownedStocks.forEach((ownedStock, index) => {
      stocksMap[index] = [];
      for (const key of Object.keys(ownedStock)) {
        if (key === 'Price') {
          stocksMap[index].push('$' + ownedStock.Price);
        } else if (key !== 'PurchaseDate') {
          stocksMap[index].push(ownedStock[key]);
        }
      }
    });

    return stocksMap;
  }

   get purchaseHistoryMap() {
    let stocksMap = [];

    this.purchaseHistory.forEach((ownedStock, index) => {
      stocksMap[index] = [];
      for (const key of Object.keys(ownedStock)) {
        if (key === 'Price') {
          stocksMap[index].push('$' + ownedStock.Price);
        } else if (key === 'PurchaseDate') {
          const dateObj = new Date(ownedStock.PurchaseDate).toLocaleString()
          stocksMap[index].push(dateObj);
        } else {
          stocksMap[index].push(ownedStock[key]);
        }
      }
    });

    return stocksMap;
  }
}

export default decorate(StocksStore, {
  purchaseHistory: observable,
  ownedStocks: observable,
  ownedStocksMap: computed,
  purchaseHistoryMap: computed,
  netAssetPrice: computed,
  searchedStock: observable,
  buyStockFeedback: observable,
  sellStockFeedback: observable,
  searchStock: action,
  buyStock: action,
  sellStock: action,
});
