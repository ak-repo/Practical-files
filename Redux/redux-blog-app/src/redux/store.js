import { configureStore } from "@reduxjs/toolkit";
import { blogsQuery } from "./rtkBlogsQuery";

//store

const store = configureStore({
  reducer: {
    [blogsQuery.reducerPath]: blogsQuery.reducer,
  },

  middleware: (getDefalutMiddleware) =>
    getDefalutMiddleware().concat(blogsQuery.middleware),
});
export default store;
