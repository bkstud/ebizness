const BACKEND_URL="http://localhost:1323"

function useApi() {

    async function fetchProducts () {
        const response = await fetch(BACKEND_URL + "/products").then(
            response => response.json())
        return response
    }

    async function updateProducts(data) { 
        const response = await fetch(BACKEND_URL + "/basket", {
                                    method: "POST",
                                    headers: {'Content-Type': 'application/json'},
                                    body: JSON.stringify(data)}).then(
            response => response.json())
        
        return response
    }
    
    return {fetchProducts, updateProducts}
}
export default useApi;