import React, { useContext, useEffect} from 'react';
import ShopContext from "../contexts/ShopContext";

function Basket(props) {
  const {products, useBasket} = useContext(ShopContext);
  const basket = useBasket.basket
  
  return (
    <div>
      <div className="basket">
        <b>Basket:</b>
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
                <th><button onClick={() => useBasket.removeProduct(p)}>Remove</button></th>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}

export default Basket;