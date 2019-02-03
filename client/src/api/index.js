import axios from 'axios';

import config from '../config';

export const fetchPurchaseHistory = async () =>
  await axios.request({
    method: 'get',
    url: `${config.apiUrl}/stocks/history/me`,
  });

export const fetchOwnedStocks = async () =>
  await axios.request({
    method: 'get',
    url: `${config.apiUrl}/stocks/me`,
  });

export const searchStocks = async filter =>
  await axios.request({
    method: 'get',
    url: `${config.apiUrl}/stocks/search/${filter}`,
  });

export const buyStock = async (symbol, quantity) =>
  await axios.request({
    method: 'post',
    url: `${config.apiUrl}/stocks/buy/${symbol}/${quantity}`,
  });

export const sellStock = async (symbol, quantity) =>
  await axios.request({
    method: 'post',
    url: `${config.apiUrl}/stocks/sell/${symbol}/${quantity}`,
  });

export const fetchBalance = async () =>
  await axios.request({
    method: 'get',
    url: `${config.apiUrl}/balance/me`,
  });

export const addBalance = async amount =>
  await axios.request({
    method: 'post',
    url: `${config.apiUrl}/balance/add/${amount}`,
  });

export const withdrawBalance = async amount =>
  await axios.request({
    method: 'post',
    url: `${config.apiUrl}/balance/withdraw/${amount}`,
  });
