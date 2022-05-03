
import {useState, useEffect} from "react";
import useApi from '../api/products'


function useBasket(props) {
  const updateProducts = useApi().updateProducts

  const [basket, setBasket] = useState([]);
  
  useEffect(() => {
    var initialBasket = props.products.filter(p => p.State === "basket")
    setBasket(initialBasket)
  }, [props.products]);

  function addProduct(product) {
    if(product.State === "basket" || product.State === "bought")
      return
      
      product.State = "basket"

      setBasket([
      ...basket,
      product
    ])
    updateProducts([product])
  }

  function removeProduct(product, new_state="available") {
    product.State = new_state
    console.log('remove:', product)
    const filteredProducts = basket.filter(p => product.Id !== p.Id)
    setBasket([...filteredProducts])
    updateProducts([product])
  }

  return {
    basket,
    addProduct,
    removeProduct,
    setBasket
  }
}

export default useBasket;