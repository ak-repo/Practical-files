import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const blogsQuery = createApi({
  reducerPath: "blogsApi",
  baseQuery: fetchBaseQuery({
    baseUrl: "http://localhost:3000",
  }),
  endpoints: (builder) => ({
    getBlogs: builder.query({
      query: () => "/blogs",
    }),
    addBlogs: builder.mutation({
      query: (newBlog) => ({
        url: "/blogs",
        method: "POST",
        body: newBlog,
      }),
    }),
  }),
});

//apis
export const { useGetBlogsQuery, useAddBlogsMutation } = blogsQuery;
