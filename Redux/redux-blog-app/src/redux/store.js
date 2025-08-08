import { configureStore } from "@reduxjs/toolkit";
import { blogs } from "./reducers/blogSlice";

const store = configureStore({
  reducer: {
    blogs,
  },
});

export default store;
