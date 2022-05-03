import React, { useContext } from 'react';
import ShopContext from "../contexts/ShopContext";
import useApi from '../api/products'

function Checkout(props) {
  const {products, useBasket} = useContext(ShopContext);
  const basket = useBasket.basket
  const result = basket.reduce((total, currentValue) => total = total + currentValue.Price, 0);
  const api = useApi()

  function buyProducts() {
    if(!basket)
      return
    basket.forEach(product => {
      product.State = "bought"
    });
    useBasket.setBasket([])
    api.updateProducts(products)
  }

  return (
    <div>
      <div className="checkout">
        <b>Checkout:</b>
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
      <button onClick={buyProducts}>Buy products</button>
    </div>
  );
}

export default Checkout;