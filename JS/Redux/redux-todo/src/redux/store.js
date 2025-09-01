import { configureStore } from "@reduxjs/toolkit";

import { todos } from "./reducers/todosReducer";


const store = configureStore({
  reducer: { todos: todos },
});


export default store;