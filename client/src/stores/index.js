import BalanceStore from './balance';
import StocksStore from './stocks';

const balanceStore = new BalanceStore();
const stocksStore = new StocksStore();

export default {
  balanceStore,
  stocksStore,
};
