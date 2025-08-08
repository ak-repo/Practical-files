import { useContext, useEffect, useState } from "react";
import {
  addTodo,
  deleteTodo,
  editTodo,
  fetchTodo,
} from "../redux/reducers/todosReducer";
import { useDispatch, useSelector } from "react-redux";

function TodoList() {
  return (
    <div>
      <ShowTodo />
    </div>
  );
}

export default TodoList;

const ShowTodo = () => {
  const dispatch = useDispatch();
  const [input, setInput] = useState("");
  const [showEdit, setShowEdit] = useState(false);
  const [editTodo, setEditTodo] = useState(null);
  const { todoList, loading } = useSelector((state) => state.todos);
  console.log(todoList);

  useEffect(() => {
    dispatch(fetchTodo());
  }, [dispatch]);
  if (loading) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="w-16 h-16 border-4 border-dashed rounded-full animate-spin border-blue-500"></div>
      </div>
    );
  }

  // handle the add to TodoList

  const submit = (e) => {
    e.preventDefault();
    if (input.length > 1) {
      dispatch(
        addTodo({
          id: new Date().getTime(),
          todo: input,
          completed: false,
        })
      );
      setInput("");
    }
  };

  // delete
  const handleDelete = (id) => {
    dispatch(deleteTodo({ id }));
  };

  //edit
  const handleEdit = (todo) => {
    setEditTodo(todo);
    setShowEdit(true);
  };
  return (
    <div className="min-h-screen bg-gray-100 p-6">
      {/* Form */}
      <div className="flex justify-center">
        <form
          className="flex items-center gap-4 bg-white shadow-md p-4 rounded-lg mt-20 w-full max-w-md"
          onSubmit={(e) => submit(e)}
        >
          <input
            className="flex-1 px-4 py-2 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-400"
            type="text"
            placeholder="Add a new todo..."
            value={input}
            onChange={(e) => setInput(e.target.value)}
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
        {Array.isArray(todoList) &&
          todoList.map((todo, index) => (
            <div
              key={todo.id}
              className="w-full max-w-md bg-white shadow-md rounded-lg p-4 flex justify-between items-center"
            >
              <div className="text-gray-800 font-medium">
                <span className="text-gray-500 mr-2">{index + 1}.</span>
                {todo?.todo}
              </div>
              <div className="flex gap-2">
                <button
                  onClick={() => handleDelete(todo?.id)}
                  className="bg-red-500 text-white px-3 py-1 rounded-lg hover:bg-red-600 transition"
                >
                  Delete
                </button>
                {!todo?.completed && (
                  <button
                    onClick={() => handleEdit(todo)}
                    className="bg-yellow-500 text-white px-3 py-1 rounded-lg hover:bg-yellow-600 transition"
                  >
                    Edit
                  </button>
                )}
              </div>
            </div>
          ))}
      </div>

      {/* Edit Form */}
      <div className="mt-8 flex justify-center">
        {showEdit && (
          <div className="w-full max-w-md">
            <EditTodo todo={editTodo} setShowEdit={setShowEdit} />
          </div>
        )}
      </div>
    </div>
  );
};

//show

const EditTodo = ({ todo, setShowEdit }) => {
  const dispatch = useDispatch();
  const [currentTodo, setCurrentTodo] = useState(todo?.todo || "nothing");
  const [completed, setCompleted] = useState(todo?.completed || false);

  //save
  const handleSave = () => {
    setShowEdit(false);
    console.log(completed, "com[");
    const updatedTodo = { ...todo, todo: currentTodo, completed: completed };
    dispatch(editTodo(updatedTodo));
  };
  return (
    <div className="fixed inset-0 backdrop-blur-sm bg-opacity-40 flex items-center justify-center z-50">
      {/* Panel */}
      <div className="bg-white rounded-lg shadow-lg p-6 w-full max-w-md">
        <h2 className="text-xl font-semibold mb-4 text-gray-800">Edit Todo</h2>

        {/* Optional: Editable input */}
        <input
          value={currentTodo}
          onChange={(e) => setCurrentTodo(e.target.value)}
          type="text"
          className="w-full px-4 py-2 border rounded-md mb-4 focus:outline-none focus:ring-2 focus:ring-blue-400"
          placeholder="Edit todo..."
        />

        {/* Checkbox */}
        <div className="flex items-center mb-4">
          <input
            type="checkbox"
            checked={completed}
            onChange={(e) => setCompleted(e.target.checked)}
            className="mr-2"
          />
          <label className="text-gray-700">Mark as completed</label>
        </div>

        {/* Buttons */}
        <div className="flex justify-end space-x-3">
          <button
            onClick={() => setShowEdit(false)}
            className="px-4 py-2 bg-gray-300 text-gray-800 rounded-md hover:bg-gray-400 transition"
          >
            Cancel
          </button>
          <button
            onClick={handleSave}
            className="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 transition"
          >
            Save
          </button>
        </div>
      </div>
    </div>
  );
};



const useMine = useContext()