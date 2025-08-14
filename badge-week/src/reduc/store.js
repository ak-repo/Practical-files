import { configureStore } from "@reduxjs/toolkit";
import productSlice from "./productSlice";
import cartSlice from "./cartSlice";
import todoSlice from "./todoSlice";

const store = configureStore({
  reducer: {
    products: productSlice,
    cart: cartSlice,
    todo: todoSlice,
  },
});

export default store;
