import { useEffect, useRef } from "react";
import { useDispatch, useSelector } from "react-redux";
import { fetchingProducts, searchProducts } from "../reduc/productSlice";
import { addToCart } from "../reduc/cartSlice";
function ProductListing() {
  const dispatch = useDispatch();
  const { filteredList } = useSelector((state) => state.products);
  const { cartList } = useSelector((state) => state.cart);

  // const [search, setSearch] = useState(null);
  const search = useRef();

  useEffect(() => {
    dispatch(fetchingProducts());
  }, []);

  const handleAddToCart = (item) => {
    if (cartList.some((cart) => cart?.id === item?.id)) {
      console.log("already in cart");
    } else {
      dispatch(addToCart(item));
    }
  };

  return (
    <div>
      <h1>Products</h1>
      <input
        onChange={() => dispatch(searchProducts(search?.current?.value))}
        type="text"
        placeholder="Search.."
        ref={search}
      />

      {filteredList?.length > 0 ? (
        filteredList?.map((item, index) => (
          <div key={item?.id}>
            <h2>
              <span>{index + 1}-</span>
              {item?.name}
            </h2>
            <button onClick={() => handleAddToCart(item)}>Cart</button>
          </div>
        ))
      ) : (
        <p>No products found</p>
      )}
    </div>
  );
}

export default ProductListing;
