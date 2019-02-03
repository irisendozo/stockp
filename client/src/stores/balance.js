//@flow
import { action, decorate, flow, observable } from 'mobx';

import * as API from '../api';

type FeedbackMessageModel = {
  success?: string,
  error?: string,
};

class BalanceStore {
  balance: number = 0;
  withdrawBalanceFeedback: FeedbackMessageModel = {};
  addBalanceFeedback: FeedbackMessageModel = {};

  fetchBalance = flow(function*() {
    try {
      const { data: apiResponse } = yield API.fetchBalance();
      this.balance = apiResponse.Balance;
    } catch (error) {}
  });

  withdrawBalance = flow(function*(amount) {
    this.withdrawBalanceFeedback = {};

    try {
      const { data: apiResponse } = yield API.withdrawBalance(amount);
      this.balance = apiResponse.Balance;
      this.withdrawBalanceFeedback.success =
        'You have successfully withdrawn from your balance';
    } catch (error) {
      this.withdrawBalanceFeedback.error =
        'You have failed to withdraw from your balance';
    }
  });

  addBalance = flow(function*(amount) {
    this.addBalanceFeedback = {};

    try {
      const { data: apiResponse } = yield API.addBalance(amount);
      this.balance = apiResponse.Balance;
      this.addBalanceFeedback.success =
        'You have successfully added to your balance';
    } catch (error) {
      this.addBalanceFeedback.error = 'You have failed to add to your balance';
    }
  });
}

export default decorate(BalanceStore, {
  balance: observable,
  withdrawBalanceFeedback: observable,
  addBalanceFeedback: observable,
  withdrawBalance: action,
  addBalance: action,
});
