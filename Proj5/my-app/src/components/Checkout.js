import React, { useContext } from 'react';
import ShopContext from "../contexts/ShopContext";

function Checkout(props) {
  const {products, useBasket} = useContext(ShopContext);
  const {basket, addProduct, removeProduct} = useBasket
  const result = basket.reduce((total, currentValue) => total = total + currentValue.Price, 0);
  return (
    <div>
      <div className="checkout">
        Checkout:
        <table>
          <thead>
          <tr>
            <th>Name</th>
            <th>Price</th>
          </tr>
          </thead>
          <tbody>
            {basket.map((p, id) => 
              (<tr key={id}>
                <th>{p.Name}</th>
                <th>{p.Price}</th>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
      <p>Sum: {result}</p>
      <button onClick={()=>{}}>Buy products</button>
    </div>
  );
}

export default Checkout;