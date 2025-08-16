import React, { useReducer } from "react";

// reduce function
const reducerFun = (state, action) => {
  switch (action.type) {
    case "increment":
      return state + 1;
    case "decrement":
      return state - 1;
    default:
      return state;
  }
};

function Counter() {
  const [count, dispatch] = useReducer(reducerFun, 0);
  return (
    <div>
      <h2>{count}</h2>
      <button
        className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:ring-2 focus:ring-blue-400 focus:outline-none"
        onClick={() => dispatch({ type: "increment" })}
      >
        +
      </button>
      <button
        className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:ring-2 focus:ring-blue-400 focus:outline-none"
        onClick={() => dispatch({ type: "decrement" })}
      >
        -
      </button>
    </div>
  );
}

export default Counter;
