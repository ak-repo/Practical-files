import { combineReducers } from "redux";
import * as type from "../types";

const initialState = {
  usersList: [],
  loading: false,
  error: null,
};

//user reducer
const users = (state = initialState, action) => {
  switch (action.type) {
    case type.GET_USERS:
      return {
        ...state,
        users: action.payload,
      };
    case type.GET_USERS_REQUESTED:
      return {
        ...state,
        loading: true,
      };
    case type.GET_USERS_SUCCESS:
      return {
        ...state,
        loading: false,
        usersList: action.usersList,
      };
    case type.GET_USERS_FAILED:
      return {
        ...state,
        loading: false,
        error: action.message,
      };

    default:
      return state;
  }
};

//action creators

export const getUsers = (users) => {
  return {
    type: type.GET_USERS_REQUESTED,
    payload: users,
  };
};

const rootReducer = combineReducers({
  users,
});

export default rootReducer;
