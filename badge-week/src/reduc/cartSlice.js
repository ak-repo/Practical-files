import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import axios from "axios";

const CARTAPI = "http://localhost:3000/cart";

//fetch cart

export const fetchCart = createAsyncThunk("cart/fetch", async () => {
  try {
    const { data } = await axios.get(CARTAPI);
    return data;
  } catch (error) {
    console.log("error while fetching data", error.message);
    return error.message;
  }
});

//add to cart

export const addToCart = createAsyncThunk("cart/addToCart", async (cart) => {
  try {
    const { data } = await axios.post(CARTAPI, cart);
    return data;
  } catch (error) {
    console.log("error while add to cart");
    return new Error(error.message);
  }
});

export const removeCart = createAsyncThunk("cart/removeCart", async (id) => {
  try {
    const { data } = await axios.delete(`${CARTAPI}/${id}`);
    return data?.id;
  } catch (error) {
    console.log("error while deleting", error.message);
    return new Error(error.message);
  }
});
const cartSlice = createSlice({
  name: "cart",
  initialState: {
    cartList: [],
  },
  reducers: {
    // addTocart: (state, action) => {
    //   state.cartList = [...state.cartList, action.payload];
    // },
    // removeCart: (state, action) => {
    //   state.cartList = state.cartList.filter(
    //     (item) => item?.id !== action.payload
    //   );
    // },
  },

  extraReducers: (builder) => {
    builder
      .addCase(fetchCart.fulfilled, (state, action) => {
        state.cartList = action.payload;
      })
      .addCase(addToCart.fulfilled, (state, action) => {
        state.cartList.push(action.payload);
      })
      .addCase(removeCart.fulfilled, (state, action) => {
        state.cartList = state.cartList.filter(
          (item) => item?.id !== action.payload
        );
      });
  },
});

export default cartSlice.reducer;

// export const { addTocart, removeCart } = cartSlice.actions;
