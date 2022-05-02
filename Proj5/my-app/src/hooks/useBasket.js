
import {useState} from "react";
import useApi from '../api/products'


function useBasket() {
  const {_, updateProducts} = useApi()

  const [basket, setBasket] = useState([]);

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
    const filteredProducts = basket.filter(p => product.Id !== p.Id)
    setBasket([...filteredProducts])
    updateProducts([product])
  }

  return {
    basket,
    addProduct,
    removeProduct,
  }
}

export default useBasket;