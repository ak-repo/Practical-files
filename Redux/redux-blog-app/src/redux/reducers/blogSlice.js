import { createSlice } from "@reduxjs/toolkit";





//blog slice

export const blogSlice = createSlice({
  name: "blogs",
  initialState: {
    blogList: [],
  },
  reducers: {
    addBlog: (state, action) => {
      state.blogList = [...state.blogList, action.payload];
    },
    deleteBlog: (state, action) => {
      state.blogList = state.blogList?.filter(
        (item) => item?.id !== action.payload
      );
    },
    editBlog: (state, action) => {
      state.blogList = state.blogList?.map((item) =>
        item?.id === action.dispatch ? action.dispatch : item
      );
    },
  },
});

export const blogs = blogSlice.reducer;
export const { addBlog, deleteBlog, editBlog } = blogSlice.actions;
