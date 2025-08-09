import { useSelector, useDispatch } from "react-redux";
import { addBlog, deleteBlog, editBlog } from "../redux/reducers/blogSlice";
import { useEffect, useState } from "react";
import { fetchBlogs } from "../redux/reducers/blogSlice";
function BlogCreation() {
  const { blogList } = useSelector((state) => state.blogs);
  const [newBlog, setNewBlog] = useState({
    title: "",
    body: "",
  });
  const [selectedBlog, setSelection] = useState(null);
  const [showEditModel, setEditModel] = useState(false);
  const dispatch = useDispatch();

  //fething
  useEffect(() => {
    console.log("use");

    dispatch(fetchBlogs());
  }, []);

  //filling
  const handleChange = (e) => {
    setNewBlog({ ...newBlog, [e.target.name]: e.target.value });
  };

  //add
  const handleAddtoblogs = (e) => {
    e.preventDefault();
    dispatch(
      addBlog({
        ...newBlog,
        createdDate: new Date().toISOString(),
        id: Date.now(),
        userId: "",
      })
    );
    setNewBlog({ title: "", body: "" });
  };

  //edit

  const handleEditBlog = (blog) => {
    setSelection(blog);
    setEditModel(true);
  };

  return (
    <div className="max-w-6xl mx-auto px-4 py-8">
      {/* Form Section */}
      <h1 className="text-center font-bold text-4xl">Redux Blog</h1>
      <div className="bg-white rounded-lg shadow-md p-6 mb-8">
        <form onSubmit={(e) => handleAddtoblogs(e)} className="space-y-4">
          <h2 className="text-2xl font-bold text-gray-800 mb-4"></h2>

          <div>
            <input
              name="title"
              value={newBlog?.title}
              onChange={(e) => handleChange(e)}
              placeholder="Enter blog title"
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-emerald-500 focus:border-emering-emerald-500 outline-none transition"
            />
          </div>

          <div>
            <textarea
              name="body"
              value={newBlog?.body}
              onChange={(e) => handleChange(e)}
              rows={8}
              placeholder="Write your content here..."
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-emerald-500 focus:border-emering-emerald-500 outline-none transition"
            ></textarea>
          </div>

          <button
            type="submit"
            className="px-6 py-2 bg-emerald-700 text-white font-medium rounded-lg hover:bg-emerald-900 focus:outline-none focus:ring-2 focus:ring-emerald-500 focus:ring-offset-2 transition"
          >
            Submit
          </button>
          <button
            onClick={() => setNewBlog({ title: "", body: "" })}
            type="button"
            className="ml-4 px-6 py-2 bg-gray-500 text-white font-medium rounded-lg hover:bg-gray-600 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2 transition"
          >
            Cancel
          </button>
        </form>
      </div>

      {/* Blog List Section */}
      <div className="space-y-6">
        <h1 className="text-3xl font-bold text-gray-800 mb-6">Latest Posts</h1>

        {blogList &&
          blogList.map((blog, index) => (
            <div
              key={blog?.id}
              className="bg-white rounded-lg shadow-md p-6 hover:shadow-lg transition"
            >
              <div className="flex items-start">
                <span className="bg-blue-100 text-emerald-700 font-bold px-3 py-1 rounded-full mr-4">
                  {index + 1}
                </span>
                <div className="flex-1">
                  <h4 className="text-xl font-semibold text-gray-800 mb-2">
                    {blog?.title}
                  </h4>
                  <p className="text-gray-600 mb-3">{blog?.body}</p>
                  <small className="text-gray-500 text-sm">
                    Posted on {blog?.createdDate}
                  </small>
                </div>
                <div className="flex space-x-2">
                  <button
                    className="p-2 text-emebg-emerald-700 hover:text-emerald-700 hover:bg-blue-50 rounded-full transition"
                    title="Edit"
                    onClick={() => handleEditBlog(blog)}
                  >
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      className="h-5 w-5"
                      viewBox="0 0 20 20"
                      fill="currentColor"
                    >
                      <path d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z" />
                    </svg>
                  </button>
                  <button
                    onClick={() => dispatch(deleteBlog(blog?.id))}
                    className="p-2 text-red-600 hover:text-red-800 hover:bg-red-50 rounded-full transition"
                    title="Delete"
                  >
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      className="h-5 w-5"
                      viewBox="0 0 20 20"
                      fill="currentColor"
                    >
                      <path
                        fillRule="evenodd"
                        d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z"
                        clipRule="evenodd"
                      />
                    </svg>
                  </button>
                </div>
              </div>
            </div>
          ))}

        {showEditModel && (
          <BlogEditModal
            blog={selectedBlog}
            dispatch={dispatch}
            setEditModel={setEditModel}
          />
        )}
      </div>
    </div>
  );
}

export default BlogCreation;

const BlogEditModal = ({ blog, dispatch, setEditModel }) => {
  const [editedBlog, setChangeBlog] = useState(blog);

  //change
  const handleChange = (e) => {
    setChangeBlog({ ...editedBlog, [e.target.name]: e.target.value });
  };

  //submissin
  const handleSubmit = () => {
    dispatch(
      editBlog({
        ...editedBlog,
        lastEditied: new Date().toISOString(),
      })
    );
  };

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center p-4">
      {/* Blurred Background */}
      <div className="absolute inset-0  bg-opacity-50 backdrop-blur-sm" />

      {/* Edit Form */}
      <div className="relative w-full max-w-2xl bg-white rounded-xl shadow-xl z-10">
        <div className="p-6">
          <h2 className="text-2xl font-bold text-gray-800 mb-4">Edit Post</h2>

          <form onSubmit={handleSubmit} className="space-y-4">
            <div>
              <input
                name="title"
                value={editedBlog.title}
                onChange={(e) => handleChange(e)}
                className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-emerald-500 focus:border-emering-emerald-500 outline-none transition"
                required
              />
            </div>

            <div>
              <textarea
                name="body"
                value={editedBlog.body}
                onChange={(e) => handleChange(e)}
                rows={8}
                className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-emerald-500 focus:border-emering-emerald-500 outline-none transition"
                required
              />
            </div>

            <div className="flex justify-end space-x-3 pt-4">
              <button
                type="button"
                onClick={() => setEditModel(false)}
                className="px-5 py-2 bg-gray-200 text-gray-800 rounded-lg hover:bg-gray-300 transition"
              >
                Cancel
              </button>
              <button
                type="submit"
                className="px-5 py-2 bg-emerald-700 text-white rounded-lg hover:bg-emerald-900 transition"
              >
                Save Changes
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};
