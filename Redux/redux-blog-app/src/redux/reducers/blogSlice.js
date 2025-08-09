import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import axios from "axios";

const blogURL = "http://localhost:3000/blogs";

//fetching
export const fetchBlogs = createAsyncThunk("blogs/fetchBlogs", async () => {
  try {
    const { data } = await axios.get(blogURL);
    return data;
  } catch (error) {
    throw new Error("Error while fetching blogs", error.message);
  }
});

//adding blog into server
export const addBlogsintoServer = createAsyncThunk(
  "blogs/addBlogintoServer",
  async (newBlog, { dispatch }) => {
    await axios.post(`${blogURL}`, newBlog);
    dispatch(fetchBlogs());
  }
);

//blog slice

export const blogSlice = createSlice({
  name: "blogs",
  initialState: {
    blogList: [],
    loading: false,
    error: null,
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
  extraReducers: (builder) => {
    builder
      .addCase(fetchBlogs.pending, (state) => {
        state.error = null;
        state.loading = true;
      })
      .addCase(fetchBlogs.fulfilled, (state, action) => {
        state.blogList = action.payload;
        state.loading = false;
      })
      .addCase(fetchBlogs.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      });
  },
});

export const blogs = blogSlice.reducer;
export const { addBlog, deleteBlog, editBlog } = blogSlice.actions;
