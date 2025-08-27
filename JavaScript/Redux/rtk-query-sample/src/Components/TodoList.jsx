import { useState } from "react";
import { useGetTodosQuery, useAddTodoMutation } from "../features/apiSlice";

function TodoList() {
  return (
    <div>
      <ShowTodo />
    </div>
  );
}

export default TodoList;

const ShowTodo = () => {
  const [newTodo, setNewTodo] = useState("");
  const { data: todos,isLoading,isFetching } = useGetTodosQuery();
  const [addTodo] = useAddTodoMutation();



  if (isLoading) return <div>Loading...</div>;
if (isFetching) return <div>Fetching...</div>;

  const handleAdd = async (e) => {
    e.preventDefault();
    if (!newTodo.trim()) return;

    try {
      await addTodo({
        userId: 233,
        title: newTodo,
        completed: false,
      }).unwrap();
      setNewTodo("");
    } catch (err) {
      console.error("Add Todo failed", err);
    }
  };

  return (
    <div className="min-h-screen bg-gray-100 p-6">
      {/* Form */}
      <div className="flex justify-center">
        <form
          onSubmit={handleAdd}
          className="flex items-center gap-4 bg-white shadow-md p-4 rounded-lg mt-20 w-full max-w-md"
        >
          <input
            className="flex-1 px-4 py-2 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-400"
            type="text"
            placeholder="Add a new todo..."
            value={newTodo}
            onChange={(e) => setNewTodo(e.target.value)}
          />
          <button
            className="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600 transition"
            type="submit"
          >
            Add
          </button>
        </form>
      </div>

      {/* Todo List */}
      <div className="mt-10 flex flex-col items-center space-y-4">
        {todos?.map((todo, index) => (
          <div
            key={todo.id}
            className="w-full max-w-md bg-white shadow-md rounded-lg p-4 flex justify-between items-center"
          >
            <div className="text-gray-800 font-medium">
              <span className="text-gray-500 mr-2">{index + 1}</span>
              {todo.title}
            </div>
            <div className="flex gap-2">
              <button className="bg-red-500 text-white px-3 py-1 rounded-lg hover:bg-red-600 transition">
                Delete
              </button>
              <button className="bg-yellow-500 text-white px-3 py-1 rounded-lg hover:bg-yellow-600 transition">
                Edit
              </button>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

const EditTodo = () => {
  return (
    <div className="fixed inset-0 backdrop-blur-sm bg-opacity-40 flex items-center justify-center z-50">
      {/* Panel */}
      <div className="bg-white rounded-lg shadow-lg p-6 w-full max-w-md">
        <h2 className="text-xl font-semibold mb-4 text-gray-800">Edit Todo</h2>

        {/* Optional: Editable input */}
        <input
          type="text"
          className="w-full px-4 py-2 border rounded-md mb-4 focus:outline-none focus:ring-2 focus:ring-blue-400"
          placeholder="Edit todo..."
        />

        {/* Checkbox */}
        <div className="flex items-center mb-4">
          <input type="checkbox" className="mr-2" />
          <label className="text-gray-700">Mark as completed</label>
        </div>

        {/* Buttons */}
        <div className="flex justify-end space-x-3">
          <button className="px-4 py-2 bg-gray-300 text-gray-800 rounded-md hover:bg-gray-400 transition">
            Cancel
          </button>
          <button className="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 transition">
            Save
          </button>
        </div>
      </div>
    </div>
  );
};
