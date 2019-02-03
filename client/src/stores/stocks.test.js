import stocks from './stocks';
import {
  fetchOwnedStocks,
  searchStocks,
  buyStock,
  sellStock,
  fetchPurchaseHistory,
} from '../api';

jest.mock('../api', () => ({
  fetchOwnedStocks: jest.fn(),
  searchStocks: jest.fn(),
  buyStock: jest.fn(),
  sellStock: jest.fn(),
  fetchPurchaseHistory: jest.fn(),
}));

describe('StocksStore', () => {
  let store;
  beforeEach(() => {
    store = new stocks();
  });
  afterEach(() => {
    jest.clearAllMocks();
  });

  describe('fetchPurchaseHistory()', () => {
    it('should fill up history from API', async () => {
      fetchPurchaseHistory.mockReturnValue(
        new Promise((resolve, reject) =>
          resolve({
            data: [
              {
                Symbol: 'SYM',
                Price: 123.1,
              },
            ],
          })
        )
      );

      await store.fetchPurchaseHistory();
      expect(fetchPurchaseHistory).toHaveBeenCalled();
      expect(store.purchaseHistory[0].Symbol).toBe('SYM');
    });
  });

  describe('fetchOwnedStocks()', () => {
    it('should fill up stocks from API', async () => {
      fetchOwnedStocks.mockReturnValue(
        new Promise((resolve, reject) =>
          resolve({
            data: {
              OwnedStocks: [
                {
                  Symbol: 'SYM',
                  Price: 123.1,
                },
              ],
            },
          })
        )
      );

      await store.fetchOwnedStocks();
      expect(fetchOwnedStocks).toHaveBeenCalled();
      expect(store.ownedStocks[0].Symbol).toBe('SYM');
    });
  });

  describe('searchStock()', () => {
    it('should fill up stock search from API', async () => {
      searchStocks.mockReturnValue(
        new Promise((resolve, reject) =>
          resolve({
            data: {
              Symbol: 'SYM',
              Price: 123.1,
            },
          })
        )
      );

      await store.searchStock(456);
      expect(searchStocks).toHaveBeenCalledWith(456);
      expect(store.searchedStock.Symbol).toBe('SYM');
    });
  });

  describe('buyStock()', () => {
    it('should fill up buy stock message and call buy stock API', async () => {
      buyStock.mockReturnValue(
        new Promise((resolve, reject) =>
          resolve({
            data: {
              OwnedStocks: [
                {
                  Symbol: 'SYM',
                  Price: 123.1,
                },
              ],
            },
          })
        )
      );

      await store.buyStock('SYM', 1);
      expect(buyStock).toHaveBeenCalledWith('SYM', 1);
      expect(store.ownedStocks[0].Symbol).toBe('SYM');
      expect(store.buyStockFeedback.success).toBeDefined();
    });

    it('should fill up buy stock message with error if failed', async () => {
      const failedRequest = {
        request: { status: '502' },
        message: 'This is a mock error',
      };

      buyStock.mockReturnValue(
        new Promise((resolve, reject) => reject(failedRequest))
      );

      await store.buyStock('SYM', 1);
      expect(buyStock).toHaveBeenCalledWith('SYM', 1);
      expect(store.buyStockFeedback.error).toBeDefined();
    });
  });

  describe('sellStock()', () => {
    it('should fill up sell stock message and call sell stock API', async () => {
      sellStock.mockReturnValue(
        new Promise((resolve, reject) =>
          resolve({
            data: {
              OwnedStocks: [
                {
                  Symbol: 'SYM',
                  Price: 123.1,
                },
              ],
            },
          })
        )
      );

      await store.sellStock('SYM', 1);
      expect(sellStock).toHaveBeenCalledWith('SYM', 1);
      expect(store.ownedStocks[0].Symbol).toBe('SYM');
      expect(store.sellStockFeedback.success).toBeDefined();
    });

    it('should fill up sell stock message with error if failed', async () => {
      const failedRequest = {
        request: { status: '502' },
        message: 'This is a mock error',
      };

      sellStock.mockReturnValue(
        new Promise((resolve, reject) => reject(failedRequest))
      );

      await store.sellStock('SYM', 1);
      expect(sellStock).toHaveBeenCalledWith('SYM', 1);
      expect(store.sellStockFeedback.error).toBeDefined();
    });
  });
});
