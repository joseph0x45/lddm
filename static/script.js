document.addEventListener('alpine:init', () => {
  Alpine.store('products', Alpine.reactive([]));
  Alpine.store('current_order', Alpine.reactive({}))

  Alpine.effect(() => {
    localStorage.setItem("products", JSON.stringify(Alpine.store('products')));
    localStorage.setItem("current_order", JSON.stringify(Alpine.store('current_order')));
  });
});

function product_card(el) {
  const product = {
    name: el.dataset.name,
    price: el.dataset.price,
    variant: parseFloat(el.dataset.variant),
  }
  return {
    product,

  }
}
