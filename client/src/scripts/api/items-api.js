const ITEMS_API = "http://localhost:8081";

class ItemRepository {
  getAll = async () => {
    return await fetch(`${ITEMS_API}/items`)
      .then((response) => response.json())
      .then((data) => {
        return data;
      });
  };

  getByGuid = async (guid) => {
    return await fetch(`${ITEMS_API}/items/${guid}`)
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
      });
  };

  create = async (item) => {
    return await fetch(`${ITEMS_API}/items`, {
      method: "POST",
      body: JSON.stringify(item),
    }).then((response) => response.json());
  }
}
