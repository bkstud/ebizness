import React, { useEffect, useState } from "react";
import useApi from '../api/products'
import useBasket from '../hooks/useBasket';

export const ShopContext = React.createContext([]);

export const ShopContextProvider = (props) => {
    const [products, setProducts] = useState([]);
    const {fetchProducts, _} = useApi()
    const UseBasket = useBasket({products, setProducts})
    useEffect(() => {
      fetchProducts()
        .then((products) => {
          setProducts(products)
        })
      }, []);
  
    return (
      <ShopContext.Provider value={{"products": products, "useBasket": UseBasket}}>
        {props.children}
      </ShopContext.Provider>
    );
};
  
export default ShopContext;