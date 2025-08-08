import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import axios from "axios";

//async

export const fetchTodo = createAsyncThunk("todos/fatchTodos", async () => {
  try {
    const { data } = await axios.get("https://dummyjson.com/todos/random/10");
    return data;
  } catch (error) {
    console.log("error in fetching", error.message);
    throw new Error("Error..................");
  }
});

// reducer
const todoSlice = createSlice({
  name: "todos",
  initialState: {
    todoList: [],
    loading: false,
    error: null,
  },
  reducers: {
    addTodo: (state, action) => {
      state.todoList = [...state.todoList, action.payload];
    },
    deleteTodo: (state, action) => {
      state.todoList = state.todoList.filter(
        (item) => item.id !== action.payload.id
      );
    },
    editTodo: (state, action) => {
      state.todoList = state.todoList.map((item) =>
        item.id === action.payload.id ? { ...action.payload } : item
      );
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchTodo.pending, (state) => {
        state.loading = true;
      })
      .addCase(fetchTodo.fulfilled, (state, action) => {
        state.loading = false;
        state.todoList = action.payload;
      })
      .addCase(fetchTodo.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error.message;
      });
  },
});

export const { addTodo, editTodo, deleteTodo } = todoSlice.actions;
export const todos = todoSlice.reducer;
