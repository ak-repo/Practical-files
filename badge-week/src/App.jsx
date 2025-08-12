import { Provider } from "react-redux";

import React from "react";
import store from "./reduc/store";
import ProductListing from "./components/ProductListing";
import CartListing from "./components/CartListing";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import StopWatch from "./components/StopWatch";
import Navbar from "./components/Navbar";

function App() {
  return (
    <Provider store={store}>
      <BrowserRouter>
        <Navbar />

        <Routes>
          <Route path="/" element={<ProductListing />} />

          <Route path="/cart" element={<CartListing />} />
          <Route path="/stopWatch" element={<StopWatch />} />
        </Routes>
      </BrowserRouter>
    </Provider>
  );
}

export default App;
