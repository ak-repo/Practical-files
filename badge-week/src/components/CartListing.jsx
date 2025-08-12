import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { fetchCart, removeCart } from "../reduc/cartSlice";
function CartListing() {
  const { cartList } = useSelector((state) => state.cart);
  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(fetchCart());
  }, [dispatch]);

  return (
    <div>
      <h1>Cart</h1>
      {cartList &&
        cartList?.map((item, index) => (
          <div key={index}>
            <p>{item?.name}</p>
            <button onClick={() => dispatch(removeCart(item?.id))}>
              Remove
            </button>
          </div>
        ))}
    </div>
  );
}

export default CartListing;
