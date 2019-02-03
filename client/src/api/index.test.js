import axios from 'axios';

import {
  fetchOwnedStocks,
  buyStock,
  sellStock,
  fetchBalance,
  fetchPurchaseHistory,
  addBalance,
  searchStocks,
  withdrawBalance,
} from './index';
import config from '../config';

jest.mock('axios', () => ({
  request: jest.fn(),
}));

describe('api', () => {
  afterEach(() => {
    jest.clearAllMocks();
  });

  describe('stocks', () => {
    it('should call history stocks API', async () => {
      await fetchPurchaseHistory();

      expect(axios.request).toBeCalledWith({
        method: 'get',
        url: `${config.apiUrl}/stocks/history/me`,
      });
    });

    it('should call fetch own stocks API', async () => {
      await fetchOwnedStocks();

      expect(axios.request).toBeCalledWith({
        method: 'get',
        url: `${config.apiUrl}/stocks/me`,
      });
    });

    it('should call search stocks API', async () => {
      await searchStocks('hello');

      expect(axios.request).toBeCalledWith({
        method: 'get',
        url: `${config.apiUrl}/stocks/search/hello`,
      });
    });

    it('should call buy stocks with symbol + amount', async () => {
      await buyStock('SOME_SYMBOL', 1);

      expect(axios.request).toBeCalledWith({
        method: 'post',
        url: `${config.apiUrl}/stocks/buy/SOME_SYMBOL/1`,
      });
    });

    it('should call sell stocks with symbol + amount', async () => {
      await sellStock('SOME_SYMBOL', 1);

      expect(axios.request).toBeCalledWith({
        method: 'post',
        url: `${config.apiUrl}/stocks/sell/SOME_SYMBOL/1`,
      });
    });
  });

  describe('balance', () => {
    it('should call fetch own balance API', async () => {
      await fetchBalance();

      expect(axios.request).toBeCalledWith({
        method: 'get',
        url: `${config.apiUrl}/balance/me`,
      });
    });

    it('should call buy stocks with symbol + amount', async () => {
      await addBalance(1.2);

      expect(axios.request).toBeCalledWith({
        method: 'post',
        url: `${config.apiUrl}/balance/add/1.2`,
      });
    });

    it('should call sell stocks with symbol + amount', async () => {
      await withdrawBalance(1.3);

      expect(axios.request).toBeCalledWith({
        method: 'post',
        url: `${config.apiUrl}/balance/withdraw/1.3`,
      });
    });
  });
});
