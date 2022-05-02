import React, { useContext } from 'react';
import ShopContext from "../contexts/ShopContext";

function Basket(props) {
  const {products, useBasket} = useContext(ShopContext);
  const {basket, addProduct, removeProduct} = useBasket
  return (
    <div>
      <div className="basket">
        Basket:
        <table>
          <thead>
          <tr>
            <th>Name</th>
            <th>Price</th>
            <th>Remove from basket</th>
          </tr>
          </thead>
          <tbody>
            {basket.map((p, id) => 
              (<tr key={id}>
                <th>{p.Name}</th>
                <th>{p.Price}</th>
                <th><button onClick={() => removeProduct(p)}>Remove</button></th>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}

export default Basket;