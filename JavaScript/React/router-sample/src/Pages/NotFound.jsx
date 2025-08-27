import { Link } from "react-router-dom";

function NotFound() {
  return (
    <div className="page">
      <h1>Page not found 404</h1>
      <p>
        <Link to="/">Home</Link>
      </p>
    </div>
  );
}

export default NotFound;
