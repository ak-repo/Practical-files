import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import axios from "axios";

const PRODUCTAPI = "http://localhost:3000/products";

//fetching products

export const fetchingProducts = createAsyncThunk(
  "fetching/producst",
  async () => {
    try {
      const { data } = await axios.get(PRODUCTAPI);

      return data;
    } catch (err) {
      console.log("error while fetching product", err.message);
    }
  }
);




const productSlice = createSlice({
  name: "products",
  initialState: {
    productList: [],
    filteredList: [],
    loading: false,
    error: null,
  },
  reducers: {
    searchProducts: (state, action) => {
      state.filteredList = state.productList.filter((item) =>
        item?.name.toLowerCase().includes(action.payload?.toLowerCase())
      );
    },
  },

  extraReducers: (builder) => {
    builder
      .addCase(fetchingProducts.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(fetchingProducts.fulfilled, (state, action) => {
        state.loading = false;
        state.productList = action.payload;
        state.filteredList = action.payload;
      })
      .addCase(fetchingProducts.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      });
  },
});

export const { searchProducts } = productSlice.actions;
export default productSlice.reducer;
