import React, { useContext } from 'react';

import ShopContext from "../contexts/ShopContext";
import useBasket from '../hooks/useBasket';

function Products(props) {
  const {products, useBasket} = useContext(ShopContext);
  const addProduct = useBasket.addProduct
  return (
    <div>
      <div className="products">
        <b>Products:</b>
        <table>
          <thead>
          <tr>
            <th>Name</th>
            <th>Price</th>
            <th>State</th>
            <th>Add to basket</th>
            
          </tr>
          </thead>
          <tbody>
            {products.map((p) => 
              (<tr key={p.Id}>
                <th>{p.Name}</th>
                <th>{p.Price}</th>
                <th>{p.State}</th>
                <th><button onClick={() => addProduct(p)}>Add</button></th>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>

  );
}

export default Products;