# stockp-client

This is a stock portfolio client that can buy + sell stocks using real-time market data.

## Technical Stack

- React 16 + React Bootstrap
- Flow
- Mobx for state management

## Development & Testing

### Requirements

- npm > v6.5.0

### Run application

Change the API address in `config/index.js`. This can either be `http://localhost:30878` if you
are running stockp-api locally or https://peaceful-castle-96036.herokuapp.com if you are connecting
to the deployed version in Heroku.

```sh
npm install
npm run start
```

Then check your browser for `http://localhost:3000`.

### Run tests

```sh
npm run test
```

## Relevant Project Structure

- `layouts/MainPage` contains the overall page for the application
- `stores/*.js` contains each of the stores maintaining the state shown in MainPage
- `api/*.js` contains the route + request to the API

## Known bugs & Improvements

- No environment for setting up different API URLs for simplicity
- No snapshot tests of the individual components: https://jestjs.io/docs/en/snapshot-testing
- `/assets` contain unneeded bootstrap elements
- Properly cascading the errors from the API back to the UI i.e. api limits
- On buying/selling stocks, need to click on the `Search` button in order to execute an order.
  Ideally this should be an autocomplete search box doing async calls to the API.
