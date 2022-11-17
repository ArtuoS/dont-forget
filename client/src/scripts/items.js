// Instances
var itemRepository = new ItemRepository();

let btnNew = document.getElementById("new-item");
btnNew.addEventListener("click", function () {
  let item = {
      Name: "Bonito",
      Description: "Descrição do item",
  }

  itemRepository.create(item);
});