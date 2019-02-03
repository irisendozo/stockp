import balance from './balance';
import { withdrawBalance, addBalance, fetchBalance } from '../api';

jest.mock('../api', () => ({
  fetchBalance: jest.fn(),
  withdrawBalance: jest.fn(),
  addBalance: jest.fn(),
}));

describe('BalanceStore', () => {
  let store;
  beforeEach(() => {
    store = new balance();
  });
  afterEach(() => {
    jest.clearAllMocks();
  });

  it('should have default values', () => {
    expect(store.balance).toEqual(0);
    expect(store.withdrawBalanceFeedback).toMatchObject({});
    expect(store.addBalanceFeedback).toMatchObject({});
  });

  describe('fetchBalance()', () => {
    it('should fill up balance from API', async () => {
      fetchBalance.mockReturnValue(
        new Promise((resolve, reject) =>
          resolve({
            data: {
              Balance: 123,
            },
          })
        )
      );

      await store.fetchBalance();
      expect(fetchBalance).toHaveBeenCalled();
      expect(store.balance).toBe(123);
    });
  });

  describe('withdrawBalance()', () => {
    it('should fill up balance from API and call withdraw balance API', async () => {
      withdrawBalance.mockReturnValue(
        new Promise((resolve, reject) =>
          resolve({
            data: {
              Balance: 123,
            },
          })
        )
      );

      await store.withdrawBalance(456);
      expect(withdrawBalance).toHaveBeenCalledWith(456);
      expect(store.balance).toBe(123);
      expect(store.withdrawBalanceFeedback.success).toBeDefined();
    });

    it('should fill up withdraw balance message with error if failed', async () => {
      const failedRequest = {
        request: { status: '502' },
        message: 'This is a mock error',
      };

      withdrawBalance.mockReturnValue(
        new Promise((resolve, reject) => reject(failedRequest))
      );

      await store.withdrawBalance(456);
      expect(withdrawBalance).toHaveBeenCalledWith(456);
      expect(store.withdrawBalanceFeedback.error).toBeDefined();
    });
  });

  describe('addBalance()', () => {
    it('should fill up balance from API and call add balance API', async () => {
      addBalance.mockReturnValue(
        new Promise((resolve, reject) =>
          resolve({
            data: {
              Balance: 123,
            },
          })
        )
      );

      await store.addBalance(456);
      expect(addBalance).toHaveBeenCalledWith(456);
      expect(store.balance).toBe(123);
      expect(store.addBalanceFeedback.success).toBeDefined();
    });

    it('should fill up add balance message with error if failed', async () => {
      const failedRequest = {
        request: { status: '502' },
        message: 'This is a mock error',
      };

      addBalance.mockReturnValue(
        new Promise((resolve, reject) => reject(failedRequest))
      );

      await store.addBalance(456);
      expect(addBalance).toHaveBeenCalledWith(456);
      expect(store.addBalanceFeedback.error).toBeDefined();
    });
  });
});
