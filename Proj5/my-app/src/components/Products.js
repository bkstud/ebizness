import React, { useContext } from 'react';

// import useBasket from '../hooks/useBasket';
import ShopContext from "../contexts/ShopContext";

function Products(props) {
  const {products, useBasket} = useContext(ShopContext);
  const {basket, addProduct, removeProduct} = useBasket
  // const { basket, addProduct, removeProduct } = useBasket();
  return (
    <div>
      <div className="products">
        Products:
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