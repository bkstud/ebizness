import React, { useEffect, useState } from "react";
import useApi from '../api/products'
import useBasket from '../hooks/useBasket';

export const ShopContext = React.createContext([]);

export const ShopContextProvider = (props) => {
    const [products, setProducts] = useState([]);
    const {fetchProducts, _} = useApi()
    useEffect(() => {
      fetchProducts()
        .then((products) => {
          setProducts(products)
        })
    }, []);
  
    return (
      <ShopContext.Provider value={{"products": products, "useBasket": useBasket()}}>
        {props.children}
      </ShopContext.Provider>
    );
};
  
export default ShopContext;